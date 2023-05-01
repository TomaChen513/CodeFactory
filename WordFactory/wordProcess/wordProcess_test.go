package wordprocess

import (
	"fmt"
	"strings"
	"testing"

	// "github.com/blevesearch/go-porterstemmer"
	"github.com/kljensen/snowball"
)

func TestWordDemo(t *testing.T) {
	words,_:=DataProcessWord("/Users/toma/Projects/wordFactory/source/paragraph.txt")
	orderWords:=UniqueStore(words)

	fmt.Println(orderWords)
	// for word:=range wordMap{
	// 	fmt.Println(word)
	// }
	// a,b:=OrderMapWord(wordMap)
	// fmt.Println(a)

	// fmt.Println("")

	// fmt.Println(b)

}

func TestStem(t *testing.T){
	fmt.Println(preProcess("experience"))
	fmt.Println(preProcess("experienced"))

}

func preProcess(str string) string {
	if strings.IndexAny(str, ",.!?") == len(str)-1 {
        str = str[:len(str)-1]
    }
	lang := "english"
	word := strings.ToLower(str)
	wordStem,_ :=snowball.Stem(word,lang,false)
	
	return wordStem
}


//  func main() {
// 	word1 := "jumps"
// 	word2 := "jumped"
// 	word3 := "running"
//  	lang := "english" // 选择英语作为处理语言
//  	// 将单词转换为原形
// 	word1 = snowball.Stem(word1, lang, true)
// 	word2 = snowball.Stem(word2, lang, true)
// 	word3 = snowball.Stem(word3, lang, true)
//  	fmt.Println(word1) // 输出 "jump"
// 	fmt.Println(word2) // 输出 "jump"
// 	fmt.Println(word3) // 输出 "run"
// }
