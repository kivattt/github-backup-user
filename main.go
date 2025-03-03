package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	jsonFile, err := os.Open("repos.json")
	if err != nil {
		log.Fatal("Failed to open repos.json, err:", err)
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Failed to read from repos.json, err:", err)
	}

	var repositories []interface{}
	err = json.Unmarshal(jsonData, &repositories)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON data from repos.json, err:", err)
	}

	for _, repository := range repositories {
		var repoData map[string]interface{}
		repoData = repository.(map[string]interface{})
		repoName := repoData["name"].(string)
		repoFullName := repoData["full_name"].(string)
		if !strings.HasPrefix(repoFullName, "kivattt/")	 {
			continue
		}

		fmt.Println("Cloning", repoFullName, "...")

		cmd := exec.Command("gh", "repo", "clone", repoFullName, "output_folder/" + repoName)
		err := cmd.Run()
		if err != nil {
			log.Fatal("ERROR:", err)
		}
	}
}
