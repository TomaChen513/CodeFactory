package wordprocess

import (
	"fmt"
	"testing"
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
