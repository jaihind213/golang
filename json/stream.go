package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Msg struct {
	Name, Text string
	//  sample json file which we wish to parse.
	//  [
	//		{"Name": "Ed", "Text": "Knock knock."},
	//		{"Name": "Sam", "Text": "Go fmt who?"}
	//	]
}

//parsing of json file as a stream of data => less memory footprint, as we read record by record
func StreamedParsing(input *os.File)  {
	var count int64 = 0
	dec := json.NewDecoder(input)

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Msg
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		count++
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
	fmt.Println("--------")
	fmt.Println("count : = " + strconv.FormatInt(count, 10))
	fmt.Println("--------")
}