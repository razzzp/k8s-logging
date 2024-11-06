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
		log := logs[rnd]
		log["timestamp"] = time.Now().UTC().Format(time.RFC3339)
		json.NewEncoder(os.Stdout).Encode(logs[rnd])
	}
}
