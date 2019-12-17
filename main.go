package main

import (
	"crypto/md5"
	"fmt"
	"github.com/mpetavy/common"
	"io"
	"os"
	"time"
)

func init() {
	common.Init("0.0.0", "2018", "test", "mpetavy", common.APACHE, true, nil, nil, run, 0)
}

func run() error {
	start := time.Now()

	hash := md5.New()

	f, err := os.Open("/home/ransom/aur/intellij-idea-ultimate-edition/ideaIU-2019.3.tar.gz")
	if common.Error(err) {
		return err
	}

	n, err := common.Stream(hash, f)
	if common.Error(err) {
		return err
	}

	common.Error(f.Close())

	fmt.Printf("%d %v %v\n", n, hash.Sum(nil), time.Now().Sub(start))
	hash.Reset()

	buf := make([]byte, 32*1024)
	buf = nil

	f, err = os.Open("/home/ransom/aur/intellij-idea-ultimate-edition/ideaIU-2019.3.tar.gz")
	if common.Error(err) {
		return err
	}

	n, err = io.CopyBuffer(hash, f, buf)
	if common.Error(err) {
		return err
	}

	common.Error(f.Close())

	fmt.Printf("%d %v %v\n", n, hash.Sum(nil), time.Now().Sub(start))

	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
