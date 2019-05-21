package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sgonzalez-r7/jtools/j"
)

func main() {
	file := os.Args[1]

	r, err := os.Open(file)
	if err != nil {
		log.Println("Error", err)
		return
	}

	json, err := j.Decode(r)
	if err != nil {
		log.Println("Error", err)
		r.Close()
		return
	}
	r.Close()

	j.Walk(json, func(n j.Node) {
		if !n.IsLeaf() {
			return
		}

		fmt.Println(n.Path.Slice(`1:`), n.Type())

	})


}
