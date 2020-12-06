package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("pb.bin.log")
	parts := strings.Split(string(content), " ")
	f, err := os.Create("./data.json")
	check(err)
	defer f.Close()

	for _, part := range parts {
		// fmt.Println(part)
		bitNum, err := strconv.ParseUint(part, 2, 8)
		_, err = f.WriteString(string(rune(bitNum)))
		check(err)
	}

	f.Sync()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
