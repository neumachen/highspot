package main

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/magicalbanana/highspot/mixtape"
)

// I opted to `panic` here instead of logging since this application will not be
// a streaming one. But ideally, we can improve this application's error
// handling by logging the error instead and making them even more verbose.
func main() {
	// load file mixtape
	mxt, err := mixtape.LoadMixtapeFromFile("./testdata/mixtape.json")
	if err != nil {
		panic(err.Error())
	}
	// load file change
	changes, err := mixtape.LoadChangesFromFile("./testdata/changes.json")
	if err != nil {
		panic(err.Error())
	}
	// process changes
	err = mxt.ApplyChanges(changes)
	if err != nil {
		panic(err.Error())
	}

	// output changes
	b, err := json.Marshal(mxt)
	if err != nil {
		panic(err.Error())
	}
	// extra step here so that we print the JSON content in a formatted
	// manner
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, b, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	f, err := os.Create("./testdata/output.json")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	_, err = f.Write(prettyJSON.Bytes())
	if err != nil {
		panic(err.Error())
	}
}
