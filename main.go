package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	//loads the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		//Creates the Genesis block, the first block, which can't be generated from the generateBlock function
		//Since there's no previous block
		t := time.Now()
		genesisBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()
	log.Fatal(run())

}