package main

import (
	dictsearch "berpa/wordFactory/dictSearch"
	wordprocess "berpa/wordFactory/wordProcess"
	"io"
	"os"
	"sort"
	"strings"
	// "time"

	"github.com/blevesearch/go-porterstemmer"

	"fmt"
	"io/ioutil"
	"regexp"

	"net/http"
	"net/url"
)



func main() {
	words, _ := wordprocess.DataProcessWord("./source/paragraph.txt")
	orderWords := wordprocess.UniqueStore(words)
	fmt.Println(orderWords)
	for i := 0; i < len(orderWords); i++ {
		dictsearch.Query(orderWords[i])
	}
	for _, word := range words {
		// queryColor(word)
		writeHtml(word, queryColor(preProcess(word)))
	}
	writeHtmlByWord(wordClassification(orderWords))
}

type Item struct {
	Color string `json:"color"`
}

func queryColor(w string) string {
	// URL encoded query params
	params := url.Values{}
	params.Set("word", w)
	query := params.Encode()
	// Compose the URL
	url := fmt.Sprintf("http://sj.tes-sys.com/words/query?%s", query)
	// Create a GET request with headers that mimic a browser
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return "red"
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	// Make the request and print the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "red"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "red"
	}
	// var color Item
	// json.Unmarshal(body,&color)
	// fmt.Println(string(body))
	return parseColorJson(string(body))
}

func parseColorJson(str string) string {
	re := regexp.MustCompile(`"([^"]+)"`)

	matches := re.FindAllStringSubmatch(str, -1)
	//  for _, match := range matches {
	//     fmt.Println(match[1])
	// }
	if len(matches) >= 1 {
		if len(matches[1]) >= 1 {
			if matches[1][1] == "black" || matches[1][1] == "blue" {
				return matches[1][1]
			}
		}
	}
	return "red"
}

func writeHtml(str string, color string) {
	// 这里重新判断字符串的合理性

	// <font color="red">第二段</font>
	newStr := "<span style=\"color: " + color + "; float: left;\">" + str + "</span>"
	f, err := os.OpenFile("test.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 写入字符串
	if str == "xxxx" {
		if _, err := io.WriteString(f, "<br><br>"); err != nil {
			panic(err)
		}
	} else {
		if _, err := io.WriteString(f, newStr); err != nil {
			panic(err)
		}
	}
	if _, err := io.WriteString(f, "<span style=\"float: left;\">&nbsp;</span>"); err != nil {
		panic(err)
	}
}

func preProcess(str string) string {
	if strings.IndexAny(str, ",.!?") == len(str)-1 {
        str = str[:len(str)-1]
    }
	word := strings.ToLower(str)
	wordStem := porterstemmer.Stem([]rune(word))
	return string(wordStem)
}

//  func main() {
// 	word1 := "jumped"
// 	word2 := "jump"
// 	// 将单词都转换为小写字母
// 	word1 = strings.ToLower(word1)
// 	word2 = strings.ToLower(word2)
// 	// 将单词转换为它们的词干
// 	word1Stem := porterstemmer.Stem(word1)
// 	word2Stem := porterstemmer.Stem(word2)
// 	// 判断两个单词的词干是否相同
// 	if word1Stem == word2Stem {
// 		fmt.Printf("%s 和 %s 是相同的单词\n", word1, word2)
// 	} else {
// 		fmt.Printf("%s 和 %s 不是相同的单词\n", word1, word2)
// 	}
// }

func wordClassification(uniqueWords []string) ([]string, []string, []string) {
	color := make([]string, 0)
	for i := 0; i < len(uniqueWords); i++ {
		color = append(color, queryColor(uniqueWords[i]))
	}

	blackWords := make([]string, 0)
	blueWords := make([]string, 0)
	redWords := make([]string, 0)
	for i := 0; i < len(uniqueWords); i++ {
		if color[i] == "red" {
			redWords = append(redWords, uniqueWords[i])
		} else if color[i] == "blue" {
			blueWords = append(blueWords, uniqueWords[i])
		} else {
			blackWords = append(blackWords, uniqueWords[i])
		}
	}
	sort.Strings(redWords)
	sort.Strings(blueWords)
	sort.Strings(blackWords)
	return blackWords, blueWords, redWords
}

func writeHtmlByWord(a, b, c []string) {
	f, err := os.OpenFile("test.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 写入字符串
	for i := 0; i < len(a); i++ {
		str := a[0] + "<br>"
		if _, err := io.WriteString(f, str); err != nil {
			panic(err)
		}
	}
	if _, err := io.WriteString(f, "<br>"); err != nil {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		str := b[0] + "\n"
		if _, err := io.WriteString(f, str); err != nil {
			panic(err)
		}
	}
	if _, err := io.WriteString(f, "<br>"); err != nil {
		panic(err)
	}
	for i := 0; i < len(c); i++ {
		str := c[0] + "<br>"
		if _, err := io.WriteString(f, str); err != nil {
			panic(err)
		}
	}

}
