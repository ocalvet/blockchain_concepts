package server

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/ocalvet/blockchain_concepts/blockchain"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/ocalvet/blockchain_concepts/block"
)

var Blockchain blockchain.Chain

func makeMuxRouter() http.Handler {
	dir, err := os.Getwd()
	if err != nil {
		panic("Could not create current directory")
	}

	Blockchain = blockchain.New(dir + string(os.PathSeparator) + "data")
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain.Get(), "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// Message is struct to save the http message
type Message struct {
	BPM int
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	blocks := Blockchain.Get()
	newBlock, err := block.NewBlock(blocks[len(blocks)-1], m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if newBlock.IsValid(blocks[len(blocks)-1]) {
		newBlockchain := append(blocks, newBlock)
		Blockchain.Replace(newBlockchain)
		spew.Dump(Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
