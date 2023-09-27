package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/antchfx/htmlquery"
	ydfanyi "github.com/hnmaonanbei/go-youdao-fanyi"
)

const URL = "https://dict.youdao.com/w/eng/%s/#keyfrom=dict2.index"

func translate(words string) string {
	words = strings.ReplaceAll(words, "/", "／")
	url := fmt.Sprintf(URL, words)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	doc, err := htmlquery.Parse(strings.NewReader(string(content)))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	result := htmlquery.FindOne(doc, "//div[@id='results-contents']//ul")
	// content = []byte(htmlquery.OutputHTML(result, true))
	content = []byte(htmlquery.InnerText(result))
	return string(content)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a word to translate.")
		os.Exit(1)
	}

	result := translate(os.Args[1])
	fmt.Println(result)
}

func main1() {
	keys := make(map[int]bool)
	for k, _ := range os.Args {
		keys[k] = true
	}

	if _, ok := keys[1]; !ok {
		fmt.Println("USEAGE: ")
		fmt.Println("\t- $ yd 词/句 [是否读出来]")
		fmt.Println("\t- $ yd hello 1 /* 试一下 */")
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
	if _, ok := keys[2]; ok {
		cmd := exec.Command("say", input)
		if err := cmd.Run(); err != nil {
			fmt.Println("出错啦：不好意思，我不知道怎么读")
		}
	}

}
