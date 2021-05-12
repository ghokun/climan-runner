package main

import (
	"log"

	"github.com/ghokun/climan-runner/pkg/tool"
)

func main() {
	err := tool.GenerateTools()
	if err != nil {
		log.Fatal(err)
	}
}
