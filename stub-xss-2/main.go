package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stderr, "<script>alert(1);</script>\n")
	fmt.Fprintf(os.Stdout, "<script>alert(1);</script>\n")
	for {
		time.Sleep(time.Second)
	}
}
