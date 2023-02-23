package main

import (
	dictsearch "berpa/wordFactory/dictSearch"
	wordprocess "berpa/wordFactory/wordProcess"
)


func main(){
	words,_:=wordprocess.DataProcessWord("/Users/toma/Projects/wordFactory/source/paragraph.txt")
	orderWords:=wordprocess.UniqueStore(words)
	for i := 0; i < len(orderWords); i++ {
		dictsearch.Query(orderWords[i])
	}
}