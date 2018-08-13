package main

import (
	"log"
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"net/http"
	"encoding/json"
	"io"
)

type Block struct {
	Index int64  // 区块编号
	Timestamp int64 // 区块时间戳
	PrevBlockHash string // 上一个区块哈希值
	Hash string // 当前区块哈希值
	Data  string // 区块数据
}

func calculateHash(b Block) string{
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([] byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

func GenerateNewBlock(prevBlock Block,data string) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = prevBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block{
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	preBlock = GenerateNewBlock(preBlock,"Genesis Block")
	return preBlock
}

//--------------------------------------------
type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
}

func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}

func (bc *Blockchain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0{
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index:%d \n", block.Index)
		fmt.Printf("PrevHash:%s \n", block.PrevBlockHash)
		fmt.Printf("CurrHash:%s \n", block.Hash)
		fmt.Printf("CurrData:%s \n", block.Data)
		fmt.Printf("CurTimestamp:%d \n", block.Timestamp)
		fmt.Printf("\n")
	}
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}

	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

var blockChain *Blockchain

func blockChainGetHandler(w http.ResponseWriter, r *http.Request){
	byte , err := json.Marshal(blockChain)
	if err != nil{
		http.Error(w,error.Error(err),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(byte))
}
func blockChainWriteHandler(w http.ResponseWriter, r *http.Request){
	blockData := r.URL.Query().Get("data")
	blockChain.SendData(blockData)
	blockChainGetHandler(w,r)
}

func run(){
	http.HandleFunc("/blockchain/get", blockChainGetHandler)
	http.HandleFunc("/blockchain/write", blockChainWriteHandler)
	http.ListenAndServe(":8080",nil)
}

func main() {
	//bc := NewBlockchain()
	//bc.SendData("111")
	//bc.SendData("222")
	//bc.Print()
	blockChain = NewBlockchain()
	run()
}

