package blockchain

import (
	"github.com/EnsicoinDevs/ensicoincoin/consensus"
	"github.com/EnsicoinDevs/ensicoincoin/network"
	"github.com/EnsicoinDevs/ensicoincoin/utils"
	bolt "github.com/etcd-io/bbolt"
	"github.com/pkg/errors"
	"math/big"
)

type Blockchain struct {
	db *bolt.DB

	GenesisBlock *Block

	index *blockIndex
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{
		GenesisBlock: &genesisBlock,
	}

	blockchain.index = newBlockIndex(blockchain)

	return blockchain
}

var (
	blocksBucket    = []byte("blocks")
	statsBucket     = []byte("stats")
	followingBucket = []byte("following")
	utxosBucket     = []byte("utxos")
)

func (blockchain *Blockchain) Load() error {
	var err error
	blockchain.db, err = bolt.Open("blockchain.db", 0600, nil)
	if err != nil {
		return errors.Wrap(err, "error opening the blockchain database")
	}

	shouldStoreGenesisBlock := false

	blockchain.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(blocksBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the blocks bucket")
		}

		genesisBlockInDb := b.Get(genesisBlock.Hash()[:])
		if genesisBlockInDb == nil {
			shouldStoreGenesisBlock = true
		}

		b, err = tx.CreateBucketIfNotExists(statsBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the stats bucket")
		}

		longestChainInDb := b.Get([]byte("longestChain"))
		if longestChainInDb == nil {
			b.Put([]byte("longestChain"), genesisBlock.Hash()[:])
		}

		_, err = tx.CreateBucketIfNotExists(followingBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the following bucket")
		}

		_, err = tx.CreateBucketIfNotExists(utxosBucket)
		if err != nil {
			return errors.Wrap(err, "error creating the utxos bucket")
		}

		return nil
	})

	if shouldStoreGenesisBlock {
		blockchain.StoreBlock(&genesisBlock)
	}

	return nil
}

func (blockchain *Blockchain) FetchUtxos(tx *Tx) (*Utxos, []*network.Outpoint, error) {
	utxos := newUtxos()
	var missings []*network.Outpoint

	err := blockchain.db.View(func(btx *bolt.Tx) error {
		b := btx.Bucket(utxosBucket)

		for _, input := range tx.Msg.Inputs {
			outpoint := input.PreviousOutput
			outpointBytes, err := outpoint.MarshalBinary()
			if err != nil {
				return errors.Wrap(err, "error marshaling an outpoint")
			}

			utxoBytes := b.Get(outpointBytes)
			if utxoBytes == nil {
				missings = append(missings, outpoint)
				continue
			}

			var utxo *UtxoEntry
			err = utxo.UnmarshalBinary(utxoBytes)
			if err != nil {
				return err
			}

			utxos.AddEntry(outpoint, utxo)
		}

		return nil
	})

	return utxos, missings, err
}

func (blockchain *Blockchain) StoreBlock(block *Block) error {
	err := blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(blocksBucket)

		blockBytes, err := block.MarshalBinary()
		if err != nil {
			return err
		}

		b.Put(block.Hash()[:], blockBytes)

		return nil
	})

	return err
}

func (blockchain *Blockchain) FindBlockByHash(hash *utils.Hash) (*Block, error) {
	var block *Block

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(blocksBucket)

		blockBytes := b.Get(hash[:])
		if blockBytes == nil {
			return nil
		}

		block.UnmarshalBinary(blockBytes)

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error finding a block")
	}

	return block, nil
}

func (blockchain *Blockchain) FindLongestChain() (*Block, error) {
	var longestChainHash *utils.Hash

	err := blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(statsBucket)
		longestChainHash = utils.NewHash(b.Get([]byte("longestChain")))

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error finding the longest chain hash")
	}

	block, err := blockchain.FindBlockByHash(longestChainHash)
	if err != nil {
		return nil, errors.Wrap(err, "error finding the longest chain")
	}

	return block, nil
}

func (blockchain *Blockchain) FindBlockHashesStartingAt(hash *utils.Hash) ([]*utils.Hash, error) {
	var hashes []*utils.Hash

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(followingBucket)

		c := b.Cursor()

		current := hash

		for k, v := c.Seek(hash[:]); k != nil; k, v = c.Seek(current[:]) {
			hash := utils.NewHash(v)
			hashes = append(hashes, hash)

			current = hash
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(nil, "error finding all the hashes")
	}

	return hashes, nil
}

func (blockchain *Blockchain) CalcNextBlockDifficulty(block *blockIndexEntry, nextBlock *Block) (uint32, error) {
	if nextBlock.Msg.Header.Height < consensus.BLOCKS_PER_RETARGET {
		return block.height, nil
	}

	lastRetargetBlockEntry, err := blockchain.index.findAncestor(nextBlock, nextBlock.Msg.Header.Height-consensus.BLOCKS_PER_RETARGET)
	if err != nil {
		return 0, err
	}

	realizedTimespan := uint64(nextBlock.Msg.Header.Timestamp.Unix()) - uint64(lastRetargetBlockEntry.timestamp.Unix())
	if realizedTimespan < consensus.MIN_RETARGET_TIMESPAN {
		realizedTimespan = consensus.MIN_RETARGET_TIMESPAN
	} else if realizedTimespan > consensus.MAX_RETARGET_TIMESPAN {
		realizedTimespan = consensus.MAX_RETARGET_TIMESPAN
	}

	targetBefore := BitsToBig(block.bits)
	target := new(big.Int).Mul(targetBefore, big.NewInt(int64(realizedTimespan)))
	target.Div(target, big.NewInt(consensus.BLOCKS_MEAN_TIMESPAN*consensus.BLOCKS_PER_RETARGET))

	return BigToBits(target), nil
}

func (blockchain *Blockchain) ValidateBlock(block *Block) (bool, error) {
	parentBlock, err := blockchain.index.findBlock(block.Msg.Header.HashPrevBlock)
	if parentBlock.height+1 != block.Msg.Header.Height {
		return false, nil
	}

	nextBits, err := blockchain.CalcNextBlockDifficulty(parentBlock, block)
	if err != nil {
		return false, err
	}
	if block.Msg.Header.Bits != nextBits {
		return false, nil
	}

	var totalFees uint64

	for _, tx := range block.Txs {
		utxos, notfound, err := blockchain.FetchUtxos(tx)
		if err != nil {
			return false, err
		}
		if len(notfound) != 0 {
			return false, nil
		}

		valid, fees, err := blockchain.ValidateTx(tx, block, utxos)
		if err == nil {
			return false, nil
		}
		if !valid {
			return false, nil
		}

		totalFees += fees
	}

	var totalCoinbaseAmount uint64

	for _, output := range block.Txs[0].Msg.Outputs {
		totalCoinbaseAmount += output.Value
	}

	if totalCoinbaseAmount != totalFees+block.CalcBlockSubsidy() {
		return false, nil
	}

	return true, nil
}

func (blockchain *Blockchain) ValidateTx(tx *Tx, block *Block, utxos *Utxos) (bool, uint64, error) {
	var inputsSum uint64

	for _, input := range tx.Msg.Inputs {
		utxo := utxos.FindEntry(input.PreviousOutput)

		if utxo.IsCoinBase() {
			if block.Msg.Header.Height-utxo.blockHeight < consensus.COINBASE_MATURITY {
				return false, 0, nil
			}
		}

		inputsSum += utxo.Amount()
	}

	var outputsSum uint64

	for _, output := range tx.Msg.Outputs {
		outputsSum += output.Value
	}

	if inputsSum <= outputsSum {
		return false, 0, nil
	}

	return false, inputsSum - outputsSum, nil
}

func (blockchain *Blockchain) ProcessBlock(block *Block) (bool, error) {
	entry, err := blockchain.index.findBlock(block.Hash())
	if err != nil {
		return false, err
	}
	if entry != nil {
		return false, nil
	}

	// TODO: check if the block exist as an orphan

	if !block.IsSane() {
		return false, nil
	}

	entry, err = blockchain.index.findBlock(block.Msg.Header.HashPrevBlock)
	if err != nil {
		return false, err
	}
	if entry == nil {
		// TODO: add the block to the orphans pool
		return false, nil
	}

	valid, err := blockchain.ValidateBlock(block)
	if err != nil {
		return false, err
	}
	if !valid {
		return false, nil
	}

	// TODO: process orphans

	return true, nil
}
