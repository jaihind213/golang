package json
//stream.go contains code to do a streaming read of a json file, this way the memory footprint is less
//instead of 'ioutil.ReadAll(input) & then json.unmarshall' which reads all data into mem first, which is bad incase of large files.
