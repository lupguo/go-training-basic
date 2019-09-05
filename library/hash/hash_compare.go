// 对比文件是否存在差异
package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}
func main() {
	h1, err := getHash("/data/github.com/tkstorm/go-example/testdata/temp/file_a.txt")
	if err != nil {
		log.Fatalln(err)
	}
	h2, err := getHash("/data/github.com/tkstorm/go-example/testdata/temp/file_b.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(h1, h2, h1 == h2)
}
