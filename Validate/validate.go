package validate

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	claspName            = "clasp"
	claspJsonFileName    = ".clasp.json"
	claspallJsonFileName = ".claspall.json"
)

func Validate() {
	if !hasClasp(claspName) {
		log.Fatalln("Could not find the 'clasp' command.")
	}
	fmt.Println("'clasp' found!")

	hasJsonFile(claspJsonFileName)
	hasJsonFile(claspallJsonFileName)

}

func hasClasp(cmdName string) bool {
	_, err := exec.LookPath(cmdName)
	return err == nil
}

func hasJsonFile(fileName string) {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		log.Fatalf("%q file does not exist\n***TERMINATE***", fileName)
	}
	fmt.Printf("%q file exists\n", fileName)

}
