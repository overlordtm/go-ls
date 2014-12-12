package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir, _ = os.Getwd()
	}

	fi, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range fi {

		stat, serr := os.Stat(filepath.Join(dir, f.Name()))
		fmt.Printf("%s\t%+q\t%s\t%v\n\n", f.Name(), f.Name(), hex(f.Name()), serr == nil)
	}

}

func hex(in string) (out string) {

	for i := 0; i < len(in); i++ {
		out = out + fmt.Sprintf("% X", in[i])
	}

	return
}
