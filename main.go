package main

import "github.com/mpetavy/common"

func init() {
	common.Init("test", "0.0.0", "2018", "test", "mpetavy", common.APACHE, "https://github.com/golang/mpetavy/golang/tresor", true, nil, nil, run, 0)
}

func run() error {
	common.DebugFunc()
	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
