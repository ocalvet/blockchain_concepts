package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ocalvet/blockchain_concepts/server"
)

// // Blockchain holds the blockchain in memory
// var Blockchain []block.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// TODO - go func() {
	// 	t := time.Now()
	// 	genesisBlock := block.Block{0, t.String(), 0, "", ""}
	// 	spew.Dump(genesisBlock)
	// 	Blockchain = append(Blockchain, genesisBlock)
	// }()
	log.Fatal(server.Run())
}
