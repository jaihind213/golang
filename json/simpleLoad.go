package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//simple parsing of json file using readAll which reads all content into mem in 1 go, so for large files => huge spike in mem usage.
func ReadAll(input *os.File) {
	var counter int64 = 0
	data, _ := ioutil.ReadAll(input)

	var msgs []Msg

	if err := json.Unmarshal(data, &msgs); err != nil {
		log.Fatal(err.Error())
	}
	for _, _ = range msgs {
		counter++
	}
	fmt.Println("--------")
	fmt.Println("record counter : = " + strconv.FormatInt(counter, 10))
	fmt.Println("--------")
}
