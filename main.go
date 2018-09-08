package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ocalvet/blockchain_concepts/block"
	"github.com/ocalvet/blockchain_concepts/blockchain"

	"github.com/joho/godotenv"
	"github.com/ocalvet/blockchain_concepts/database"
	"github.com/ocalvet/blockchain_concepts/server"
)

var bcServer chan []block.Block
var blkChain blockchain.Blockchain

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
	blkChain = blockchain.New(db)

	// Start web server
	log.Fatal(server.Run(blkChain))

	// Channel to handle the blockchain
	bcServer = make(chan []block.Block)

	// start TCP and serve TCP server
	tcpPort := os.Getenv("NETWORK_PORT")
	log.Printf("PORT %S", tcpPort)
	server, err := net.Listen("tcp", ":"+tcpPort)
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
	log.Println("Handling connection")
	defer conn.Close()
	io.WriteString(conn, "Enter a new BPM:")

	scanner := bufio.NewScanner(conn)

	// take in BPM from stdin and add it to blockchain after conducting necessary validation
	go func() {
		for scanner.Scan() {
			bpm, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("%v not a number: %v", scanner.Text(), err)
				continue
			}
			blks := blkChain.Get()
			newBlock, err := block.NewBlock(blks[len(blks)-1], bpm)
			if err != nil {
				log.Println(err)
				continue
			}
			if newBlock.IsValid(blks[len(blks)-1]) {
				newBlockchain := append(blks, newBlock)
				blkChain.Replace(newBlockchain)
			}

			bcServer <- blkChain.Get()
			io.WriteString(conn, "\nEnter a new BPM:")
		}
	}()

	// simulate receiving broadcast
	go func() {
		for {
			time.Sleep(30 * time.Second)
			output, err := json.Marshal(blkChain.Get())
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump(blkChain.Get())
	}
}
