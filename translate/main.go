package main

import (
	"fmt"
	"os"

	ydfanyi "github.com/hnmaonanbei/go-youdao-fanyi"
)

func main() {
	keys := make(map[int]bool)

	for k, _ := range os.Args {
		keys[k] = true
	}

	if _, ok := keys[1]; !ok {
		fmt.Println("USEAGE: ")
		fmt.Println("\t- $ yd 词/句")
		fmt.Println("\t- $ yd hello  /* 试一下 */")
		return
	}
	opts := ydfanyi.NewOptions("", "", "")
	// opts.proxy = ""
	//opts.From = ydfanyi.EN
	//opts.To = ydfanyi.ZH
	input := os.Args[1]
	res, _ := ydfanyi.Do(input, opts)
	// 翻译结果
	fmt.Println(res.String())
	// 详细信息
	fmt.Println(res.SmartResult.Entries)
	// out

	//	你好
	//	[ int. 喂；哈罗，你好，您好（表示问候， 惊奇或唤起注意时的用语）
	//	n. （Hello）（法）埃洛（人名）
	//]

}
