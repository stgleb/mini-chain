package minichain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	PrevBlockHash []byte        `json:"prev-block-hash"`
	BlockHash     []byte        `json:"block-hash"`
	Transactions  []Transaction `json:"transactions"`
}

func NewBlock(prevBlockHash []byte, transactions []Transaction) *Block {
	var txHashes [][]byte
	var txHash [32]byte

	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Transactions:  transactions,
	}

	GetLogger().Debugf("Create new block with timestamp %d prev-block-hash %s tx count",
		block.Timestamp, string(block.PrevBlockHash), len(block.Transactions))

	for _, tx := range transactions {
		txHashes = append(txHashes, tx.Id)
	}

	timestampBytes := []byte(strconv.FormatInt(block.Timestamp, 10))
	txHashes = append(txHashes, timestampBytes)

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	block.BlockHash = txHash[:]

	return block
}
