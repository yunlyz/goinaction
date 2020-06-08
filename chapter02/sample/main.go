package main

import (
	"log"
	"os"

	"github.com/yunlyz/goinaction/chapter02/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
