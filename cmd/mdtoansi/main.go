package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	markdown "github.com/crgimenes/term-markdown"
)

type config struct {
	File string `json:"file"`
}

func main() {
	var err error
	cfg := &config{}

	flag.StringVar(&cfg.File, "f", "-", "file to read")
	flag.Parse()

	f := os.Stdin
	if cfg.File != "-" {
		f, err = os.Open(cfg.File)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
	}

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := markdown.Render(string(b), 80, 6)
	fmt.Println(string(result))
}
