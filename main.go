package main

import (
	"bytes"
	"fmt"
	"github.com/mpetavy/common"
)

func init() {
	common.Init("0.0.0", "2018", "test", "mpetavy", common.APACHE, true, nil, nil, run, 0)
}

func test(ba *bytes.Buffer) {
	ba.Reset()
	ba.WriteString("modified")
}

func run() error {
	fmt.Printf("title: %s\n", common.Title())

	ba := []byte("default")

	bb := bytes.NewBuffer(ba)

	test(bb)

	fmt.Printf("%s\n", bb.String())

	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
