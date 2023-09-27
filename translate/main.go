package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
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
