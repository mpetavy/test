package main

import (
	"fmt"
	"github.com/mpetavy/common"
)

func init() {
	common.Init("0.0.0", "2018", "test", "mpetavy", common.APACHE, true, nil, nil, run, 0)
}

func run() error {
	fmt.Printf("title: %s\n", common.Title())

	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
