package main

import (
	"log"
	"os"

	_ "github.com/heqingbao/goinaction-code/chapter2/matchers"
	"github.com/heqingbao/goinaction-code/chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
