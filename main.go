package main

import (
	"log"
	"os"

	"github.com/ocalvet/blockchain_concepts/blockchain"

	"github.com/joho/godotenv"
	"github.com/ocalvet/blockchain_concepts/database"
	"github.com/ocalvet/blockchain_concepts/server"
)

// // Blockchain holds the blockchain in memory
// var Blockchain []block.Block

func main() {
	// Load environment
	godotenv.Load()

	// Get Working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Load database
	db, err := database.New(dir + string(os.PathSeparator) + "data")
	if err != nil {
		log.Fatal(err)
	}

	// If not genesisBlock create it
	// TODO - go func() {
	// 	t := time.Now()
	// 	genesisBlock := block.Block{0, t.String(), 0, "", ""}
	// 	spew.Dump(genesisBlock)
	// 	Blockchain = append(Blockchain, genesisBlock)
	// }()

	// Load blockchain
	blkChain := blockchain.New(db)

	// Start server
	log.Fatal(server.Run(blkChain))
}
