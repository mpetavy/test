package main

import (
	"fmt"
	"github.com/mpetavy/common"
)

func init() {
	common.Init("0.0.0", "2018", "test", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, true, nil, nil, run, 0)
}

func run() error {
	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
