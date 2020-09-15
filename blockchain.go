package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int    //Position in Blockchain
	Timestamp string //data that was written
	BPM       int    //pulse rate (Beats per minute)
	Hash      string //Identifier (Sha256)
	PrevHash  string //previous hash (Sha256)

}

var Blockchain []Block

//How does it chain
/*
+----------+        +----------+        +----------+
|  BLOCK 1 |        |  BLOCK 2 |        |  BLOCK 3 |
|          |        |          |        |          |
+----------+        +----------+        +----------+
|Index     |        |Index     |        |Index     |
+----------+        +----------+        +----------+
|Timestamp |        |Timestamp |        |Timestamp |
+----------+        +----------+        +----------+
|Bpm       |        |Bpm       |        |Bpm       |
+----------+        +----------+        +----------+
|Hash      +---+    |Hash      +---+    |Hash      |
+----------+   |    +----------+   |    +----------+
|PrevHash  |   +--->+PrevHash  |   +--->+PrevHash  |
+----------+        +----------+        +----------+

the hash and prevhash ensure that the Blockchain is in the right order, making a chain.

*/

// Why do we use hashing
/*
The main purposes of hashing are:
- calculating an hash from a Value is easy, but calculating a Value from an Hash is almost impossible, we use it to prevent replication
- we use it to maintain the order in the chain, like an Identification
- Either alot of characters or a few amount will be turned into the same amount of characters after hashed, 64 characters (idempotency)

*/

func calculateHash(block Block) string {
	//The order of this concatenation can be any
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash

	//hashing
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
	//

}

func generateBlock(previousBlock Block, BPM int) (Block, error) {

	//Pretty Straightforward

	//The index is one more than the previous block
	//The Timestamp of the block is the current time
	//Previous Hash is the hash of the previous block, like shown before
	//The Hash is generated from the calculateHash, more on that above the calculateHash function

	var nextBlock Block
	t := time.Now()

	nextBlock.Index = previousBlock.Index + 1
	nextBlock.Timestamp = t.String()
	nextBlock.BPM = BPM
	nextBlock.PrevHash = previousBlock.Hash
	nextBlock.Hash = calculateHash(nextBlock)

	return nextBlock, nil
}

func isBlockValid(nextBlock, previousBlock Block) bool {

	//check if the new block is just an increment
	if previousBlock.Index+1 != nextBlock.Index {
		return false
	}

	//check if the hash chain was broken
	if previousBlock.Hash != nextBlock.PrevHash {
		return false
	}

	//check if the hash really matches
	if calculateHash(nextBlock) != nextBlock.Hash {
		return false
	}

	return true

}

//if we want to know which, between two chains/nodes, which one we use, we will consider that the longest one (with more data/more updates) is the real chain
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}

}

