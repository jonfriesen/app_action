package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//used for parsing json object of changed repo
type UpdatedRepo struct {
	Name       string
	Repository string
	Tag        string
}

//reads the file from fileLocation
func readFileFrom(fileLocation string) ([]byte, error) {
	jsonFile, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal("Error in opening the file", err)
		return []byte{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error in reading from file: ", err)
		return []byte{}, err
	}
	return byteValue, err
}

//reads the file and return json object of type UpdatedRepo
func getAllRepo(location string) ([]UpdatedRepo, error) {
	byteValue, err := readFileFrom(location)
	if err != nil {
		log.Fatal("Error in reading from file: ", err)
		return nil, err
	}
	var allRepos []UpdatedRepo
	err = json.Unmarshal(byteValue, &allRepos)
	if err != nil {
		log.Fatal("Error in parsing json data from file: ", err)
		return nil, err
	}
	return allRepos, nil

}
func checkForDockerHub(allFiles []UpdatedRepo, key int) {
	cmd := exec.Command("sh", "-c", `yq eval '.*[]| select(.name == "`+allFiles[key].Name+`").image.registry_type' _temp.yaml`)
	val, err := cmd.Output()
	if err != nil {
		log.Fatal("Error in checking docr path file: ", err)
		os.Exit(1)
	}
	if string(val) != "DOCR" {
		cmd = exec.Command("sh", "-c", `cat _temp.yaml | yq eval 'del(.*[]| select("`+allFiles[key].Name+`").`+string(val)+`' _temp.yaml-| sponge _temp.yaml`)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal("Error in updating docr path file: ", err)
			os.Exit(1)
		}
	}
}
func checkForGit(allFiles []UpdatedRepo, key int) {
	cmd := exec.Command("sh", "-c", `cat _temp.yaml | yq eval 'del(.*[]| select(.name = "`+allFiles[key].Name+`").git*' _temp.yaml-| sponge _temp.yaml`)
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("Error in updating docr path file: ", err)
		os.Exit(1)
	}
}

func execCommand(allFiles []UpdatedRepo) error {
	for key, _ := range allFiles {
		checkForDockerHub(allFiles, key)
		checkForGit(allFiles, key)
		cmd := exec.Command("sh", "-c", `cat _temp.yaml |yq eval '(.*[]| select(.name == "`+allFiles[key].Name+`").image.repository) |=  "`+allFiles[key].Repository+
			`" |`+`(.*[]| select(.name == "`+allFiles[key].Name+`").image.registry_type) |= "DOCR" |(.*[]|select(.name == "`+allFiles[key].Name+`").image.tag) |=  "`+allFiles[key].Tag+`"' -| sponge _temp.yaml`)
		_, err := cmd.Output()
		if err != nil {
			log.Fatal("Error in checking docr path file: ", err)
			os.Exit(1)
		}

	}
	return nil

}

func main() {
	//import and return json object of changed repo
	allFiles, err := getAllRepo("test1")
	if err != nil {
		fmt.Println("Error in Retrieving json data: ", err)
		os.Exit(1)
	}
	err = execCommand(allFiles)
	if err != nil {
		log.Fatal("Error in retrieving data from")
	}

}
