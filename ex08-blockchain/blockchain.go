package main

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PrevHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	block.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}}
	block.SetHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) *Block {
	newBlock := NewBlock(data, []byte{})
	bc.blocks = append(bc.blocks, newBlock)
	return newBlock
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewRow(db *sql.DB, bc *Block) {
	_, err := db.Exec("INSERT INTO BLOCKCHAIN (timestamp, data, hash, prevHash) VALUES ($1, $2, $3, $4)", bc.Timestamp, bc.Data, bc.Hash, bc.PrevHash)
	if err != nil {
		fmt.Println("Fail to create new row")
		panic(err)
	}
}

func NewTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS blockchain (id INTEGER PRIMARY KEY, timestamp INTEGER, data TEXT, hash TEXT, prevHash TEXT)")
	if err != nil {
		fmt.Println("creating of new table failed!")
		panic(err)
	}
	rows, _ := db.Query("SELECT data FROM blockchain WHERE id = 1")
	defer rows.Close()
	if !rows.Next() {
		bc := NewGenesisBlock()
		NewRow(db, bc)
	}

}

func ListPrint(db *sql.DB) {
	var id int

	rows, err := db.Query("select * from blockchain")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		p := Block{}
		err := rows.Scan(&id, &p.Timestamp, &p.Data, &p.Hash, &p.PrevHash)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("ID: %d\n", id)
		fmt.Printf("Prev Hash: %x\n", p.PrevHash)
		fmt.Printf("Data: %s\n", p.Data)
		fmt.Printf("Hash: %x\n", p.Hash)
		fmt.Println()
	}
}

func FindPrevHash(db *sql.DB) []byte {
	var id int

	rows, _ := db.Query("SELECT * FROM blockchain WHERE ID=(SELECT MAX(id) FROM blockchain)")
	defer rows.Close()
	rows.Next()
	p := Block{}
	e := rows.Scan(&id, &p.Timestamp, &p.Data, &p.Hash, &p.PrevHash)
	if e != nil {
		fmt.Println("It's first block!")
	}
	return p.Hash
}

func main() {
	db, err := sql.Open("sqlite3", "blockchain.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	NewTable(db)

	bc := NewBlockchain()
	if len(os.Args) < 2 {
		fmt.Println("Please, add parametr!")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please, add some data to block!")
			os.Exit(1)
		}
		pHash := FindPrevHash(db)
		block := bc.AddBlock(os.Args[2])
		block.PrevHash = pHash
		NewRow(db, block)
		os.Exit(0)
	case "list":
		ListPrint(db)
		os.Exit(0)
	default:
		fmt.Println("Please, use parametrs!")
		os.Exit(1)
	}
}
