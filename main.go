package main

import (
	"fmt"
	"habitat/cmd"
	"os"
)

func init() {

}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
