package main

import (
	"encoding/json"
	"math/rand/v2"
	"os"
	"time"
)

func main() {
	//parse json
	var logs []map[string]string
	file, _ := os.Open("dummy.json")
	json.NewDecoder(file).Decode(&logs)
	for {
		time.Sleep(3 * time.Second)

		rnd := rand.Int() % len(logs)

		json.NewEncoder(os.Stdout).Encode(logs[rnd])
	}
}
