package main

import (
	"fmt"
	"github.com/jaihind213/golang/json"
	"github.com/jaihind213/golang/mem"
	"os"
)
func main() {
	fmt.Println("generate data using generate_data.sh it generate ~900mB in ~600 seconds.")
	go mem.PrintUsage()
	input, _ := os.Open("/tmp/large.json")
	//json.ReadAll(input)
	json.StreamedParsing(input)
}
