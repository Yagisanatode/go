package updateall

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

var (
	claspName            = "clasp"
	claspJsonFileName    = ".clasp.json"
	claspallJsonFileName = ".claspall.json"
)

func UpdateAllProjects() {

	files := getClaspAllJsonData()

	updateClaspId(files)

}

type Files struct {
	Files  []File `json:"files"`
	CoreId string
}

type File struct {
	Title string `json:"title"`
	Id    string `json:"id"`
}

func getClaspAllJsonData() Files {
	// Open our jsonFile
	jsonFile, err := os.Open(claspallJsonFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully Opened " + claspallJsonFileName)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read opened jsonFile
	byteValue, _ := io.ReadAll(jsonFile)

	// initialise Files array
	var files Files

	// Un marshal byteAarray containing the jsonFile's contents
	// into 'Files' which is defined in the struct above.
	json.Unmarshal(byteValue, &files)

	for i := 0; i < len(files.Files); i++ {
		fmt.Println("Title: " + files.Files[i].Title)
		fmt.Println("ID: " + files.Files[i].Id)
	}

	return files

}

type ClaspConfig struct {
	ScriptId string `json:"scriptId"`
	RootDir  string `json:"rootDir"`
}

func updateClaspId(files Files) {

	// Read the existing JSON file
	file, err := os.ReadFile(claspJsonFileName)
	if err != nil {
		log.Panicln("Error reading file: ", err)
	}

	// Unmarshal the clasp JSON file in Go data struct
	var config ClaspConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Panicln("Error unmarshalling JSON: ", err)
	}

	// Store the original scriptID
	originalId := config.ScriptId

	// Iterate over the claps location arrays.
	for _, f := range files.Files {
		// Update the scriptId
		config.ScriptId = f.Id
		writeToClaspJson(config)
		runCmdClaspPush()
	}

	config.ScriptId = originalId
	writeToClaspJson(config)
	runCmdClaspPush()

}

func writeToClaspJson(config ClaspConfig) {
	// Marshal the data back to JSON
	updatedData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Panicln("Error marshalling JSON: ", err)
	}

	// Write the updated JSON data to the file.
	err = os.WriteFile(claspJsonFileName, updatedData, 0644)
	if err != nil {
		log.Panicln("Error writing to file: ", err)
	}

	log.Println("File updated successfuly")
}

func runCmdClaspPush() {
	cmd := exec.Command(claspName, "push")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
