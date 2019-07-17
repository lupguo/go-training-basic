// 流式编码器和解码器,由于读者和写入器无处不在，这些编码器和解码器类型可用于各种场景，例如读取和写入HTTP连接，WebSockets或文件。
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// create an bytes io.Reader
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	input := bytes.NewReader(b)

	// decoder & encoder
	dec := json.NewDecoder(input)
	enc := json.NewEncoder(os.Stdout)

	// decode from bytes io.Reader b, and write to v
	var v map[string]interface{}
	if err := dec.Decode(&v); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", v)

	// delete key not equal Name
	for k := range v {
		if k != "Name" {
			delete(v, k)
		}
	}

	// encode map v and write to os.Stdout
	if err := enc.Encode(&v); err != nil {
		log.Println(err)
	}
}
