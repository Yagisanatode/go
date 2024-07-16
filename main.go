package main

import (
	"fmt"

	updateall "github.com/Yagisanatode/go-claspall/UpdateAll"
	validate "github.com/Yagisanatode/go-claspall/Validate"
)

func main() {
	fmt.Println("ready to go")

	validate.Validate()

	updateall.UpdateAllProjects()

}
