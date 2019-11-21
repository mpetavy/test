package main

import (
	"fmt"
	"github.com/mpetavy/common"
)

func init() {
	common.Init("0.0.0", "2018", "test", "mpetavy", common.APACHE, true, nil, nil, run, 0)
}

func run() error {
	phrase := []byte("1234567890123456")
	txt, err := common.EncryptString(phrase, "Hello world!")
	if common.Error(err) {
		return nil
	}

	fmt.Printf("%s\n", txt)

	txt, err = common.EncryptString(phrase, "Hello world!")
	if common.Error(err) {
		return nil
	}

	fmt.Printf("%s\n", txt)

	txt, err = common.DecryptString(phrase, txt)
	if common.Error(err) {
		return nil
	}

	fmt.Printf("%s\n", txt)

	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
