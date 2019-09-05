package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	// 1. 索引字符串可访问单个字节，而不是字符
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x,", sample[i])
	}
	// 输出 bd b2 3d bc 20 e2 8c 98

	// 2. ％和x之间放置一个空格
	fmt.Println()
	fmt.Printf("% X\n", sample)
	// 输出 bd b2 3d bc 20 e2 8c 98

	// 3. 转义字符串中任何不可打印的字节序列
	fmt.Printf("%q\n", sample)
	// 输出 "\xbd\xb2=\xbc ⌘"

	// 4. 以正确格式化的UTF-8的Unicode值输出
	fmt.Printf("%+q\n", sample)
	// 输出 "\xbd\xb2=\xbc \u2318"

	// 5. 字符`原意
	fmt.Println("\xe2\x8c\x98", `\xe2\x8c\x98`)
}
