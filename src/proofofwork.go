package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

)

var maxNonce = math.MaxInt64
const targetBits = 16

type Proofofwork struct {
	block *Block
	target *big.Int
}

//builds and returns a pow
func NewProofofwork(b *Block) *Proofofwork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &Proofofwork{b, target}
	return pow

}

//Prepares the data to hash
func (pow *Proofofwork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},

	)

	return data
}

//Run() performs the proof-of-work
func (pow *Proofofwork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	//Loop prepares data, hashes data with SHA-256, Conv hash to big integer
	// and compare integer with the target
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]


}

func (pow *Proofofwork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}




