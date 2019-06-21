package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mattn/go-isatty"
)

func main() {
	if isatty.IsTerminal(os.Stdin.Fd()) {
		os.Stderr.WriteString("Input string and hit ^D when complete\n")
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", string(data))
}
