package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", string(data))
}
