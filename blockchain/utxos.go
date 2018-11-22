package blockchain

import (
	"encoding/json"
)

type UtxoEntry struct {
	amount      uint64
	script      []byte
	blockHeight int

	coinBase bool
}

type utxoEntryJSON struct {
	amount      uint64 `json:"amount"`
	script      []byte `json:"script"`
	blockHeight int    `json:"blockHeight"`

	coinBase bool `json:"coinbase"`
}

func (entry *UtxoEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(utxoEntryJSON{
		amount:      entry.amount,
		script:      entry.script,
		blockHeight: entry.blockHeight,
		coinBase:    entry.coinBase,
	})
}

func (entry *UtxoEntry) UnmarshalJSON(b []byte) error {
	tmp := &utxoEntryJSON{}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	entry.amount = tmp.amount
	entry.script = tmp.script
	entry.blockHeight = tmp.blockHeight
	entry.coinBase = tmp.coinBase

	return nil
}

func newUtxoEntry() *UtxoEntry {
	return &UtxoEntry{}
}

func (entry *UtxoEntry) Amount() uint64 {
	return entry.amount
}

func (entry *UtxoEntry) Script() []byte {
	return entry.script
}

func (entry *UtxoEntry) BlockHeight() int {
	return entry.blockHeight
}

func (entry *UtxoEntry) IsCoinBase() bool {
	return entry.coinBase
}

type Utxos struct {
	entries map[Outpoint]*UtxoEntry
}

func newUtxos() *Utxos {
	return &Utxos{
		entries: make(map[Outpoint]*UtxoEntry),
	}
}

func (utxos *Utxos) AddEntry(outpoint Outpoint, entry *UtxoEntry) {
	utxos.entries[outpoint] = entry
}

func (utxos *Utxos) AddEntryWithTx(outpoint Outpoint, tx *Tx, blockHeight int) {
	output := tx.Outputs[outpoint.Index]

	utxos.AddEntry(outpoint, &UtxoEntry{
		amount:      output.Value,
		script:      output.Script,
		blockHeight: blockHeight,
		coinBase:    tx.IsCoinBase(),
	})
}

func (utxos *Utxos) FindEntry(outpoint Outpoint) *UtxoEntry {
	return utxos.entries[outpoint]
}