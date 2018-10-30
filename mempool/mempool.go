package mempool

import (
	"github.com/EnsicoinDevs/ensicoin-go/blockchain"
	"sync"
)

type Mempool struct {
	txs     map[string]*blockchain.Tx
	orphans map[string]*blockchain.Tx
	mutex   sync.RWMutex
}

func NewMempool() *Mempool {
	return &Mempool{
		txs: make(map[string]*blockchain.Tx),
	}
}

func (mempool *Mempool) addTx(tx *blockchain.Tx) {
	mempool.mutex.Lock()
	mempool.txs[tx.Hash] = tx
	mempool.mutex.Unlock()
}

func (mempool *Mempool) findTxByHash(hash string) *blockchain.Tx {
	mempool.mutex.RLock()
	defer mempool.mutex.RUnlock()

	tx := mempool.txs[hash]
	if tx == nil {
		tx = mempool.orphans[hash]
	}

	return tx
}

func (mempool *Mempool) ProcessTx(tx *blockchain.Tx) bool {
	if !tx.IsSane() {
		return false
	}

	if tx.IsCoinBase() {
		return false
	}

	if mempool.findTxByHash(tx.Hash) != nil {
		return false
	}

	mempool.addTx(tx)

	return true
}

func (mempool *Mempool) FindTxByHash(hash string) *blockchain.Tx {
	return mempool.findTxByHash(hash)
}

func (mempool *Mempool) RemoveTxByHash(hash string) {
	mempool.mutex.Lock()
	delete(mempool.txs, hash)
	mempool.mutex.Unlock()
}
