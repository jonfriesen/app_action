package main

import (
	"os"
	"testing"
)

func TestReadFileFrom(t *testing.T) {
	//Test to check if read is working correctly
	//For this I will read test1 file and verify the output
	os.Remove("_test")
	testFileInput := `"[
	{
	  "name": "frontend",
	  "repository": "registry.digitalocean.com/<my-registry>/<my-image>",
	  "tag": "latest"
	},
	{
	  "name": "landing",
	  "repository": "registry.digitalocean.com/<my-registry>/<my-image>",
	  "tag": "test1"
	},
	{
	  "name": "api",
	  "repository": "registry.digitalocean.com/<my-registry>/<my-image>",
	  "tag": "test2"
	}
  ]"`
	testFile := []byte(testFileInput)
	file, err := os.Create("test")
	if err != nil {
		t.Error("Not able to create a file", err)
	}
	_, err = file.Write(testFile)
	if err != nil {
		t.Error("Not able to write on a file", err)
	}
	jsonFile, err := readFileFrom("test")
	if err != nil {
		t.Error("Unable to read file", err)
	}

	if string(jsonFile) != testFileInput {
		t.Error("mismatched file: ", testFileInput)
	}
	os.Remove("_test")

}
