package main

import (
	"log"
	"net"
	"os"

	"github.com/ocalvet/blockchain_concepts/block"
	"github.com/ocalvet/blockchain_concepts/blockchain"

	"github.com/joho/godotenv"
	"github.com/ocalvet/blockchain_concepts/database"
	"github.com/ocalvet/blockchain_concepts/server"
)

var bcServer chan []block.Block

func main() {
	// Load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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

	// Load blockchain
	blkChain := blockchain.New(db)

	// Start web server
	log.Fatal(server.Run(blkChain))

	// Channel to handle the blockchain
	bcServer = make(chan []block.Block)

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
}
