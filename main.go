package main

import (
	"fmt"
	"github.com/mpetavy/common"
	"os"
	"path/filepath"
	"time"
)

func init() {
	common.Init(false, "0.0.0", "", "2018", "test", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, nil, run, 0)
}

func readDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	common.DebugError(f.Close())
	if err != nil {
		return nil, err
	}
	return list, nil
}

func walk1(path string) error {
	fmt.Printf("%s\n", path)

	fis, err := readDir(path)
	if common.Error(err) {
		return err
	}

	for _, fi := range fis {
		if fi.IsDir() {
			err = walk1(filepath.Join(path, fi.Name()))
			if common.Error(err) {
				return err
			}
		}
	}

	return nil
}

func run() error {
	start := time.Now()
	defer func() {
		fmt.Printf("%v\n", time.Since(start))
	}()

	root := "d:\\Dicom-Files.big"

	return walk1(root)
}

func main() {
	defer common.Done()

	common.Run(nil)
}
