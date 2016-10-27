package main

import (
	"fmt"
	"io/ioutil"
	"lab/searcher"
)

func main() {
	dat, _ := ioutil.ReadFile("searchfile.txt")
	for index := 0; index < 1; index++ {
		result := searcher.Search(string(dat), "*o*em")
		for _, elem := range result {
			fmt.Printf("%v - %v\n", elem.Position, elem.Result)
		}
	}
}
