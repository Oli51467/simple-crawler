package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code: ", resp.StatusCode)
		return
	} else {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		processAll(all)
	}
}

func processAll(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com.zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		fmt.Printf("City: %s, Url: %s\n", match[2], match[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}
