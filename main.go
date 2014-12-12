package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
		fmt.Printf("%s\t%+q\t%s\n\n", f.Name(), f.Name(), hex(f.Name()))
	}

}

func hex(in string) (out string) {

	for i := 0; i < len(in); i++ {
		out = out + fmt.Sprintf("% X", in[i])
	}

	return
}
