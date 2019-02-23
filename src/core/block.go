package core

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64  // 区块创建的时间戳
	Data          []byte // 区块包含的数据
	PrevBlockHash []byte // 前一个区块的哈希值
	Hash          []byte // 区块自身的哈希值，用于校验数据的有效
	Nonce         int
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	//block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
// 第一个区块，创世纪区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
