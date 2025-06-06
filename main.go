package main

import (
	"fmt"

	"github.com/netflix/weep/pkg/util"
)

func main() {
	fmt.Println("Running OpenLink with default command (open)...")
	err := util.OpenLink("https://example.com")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("OpenLink executed.")
}
