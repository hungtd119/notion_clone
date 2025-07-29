package main

import (
	"log"
	"notion/src/cmd"
)

func main() {
	rootCmd := cmd.Root()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
