package wordprocess

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 按序提取所有字符
// Todo:去除特殊标点符号
func DataProcessWord(path string) ([]string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	res := make([]string, 0)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Single data process
	for scanner.Scan() {
		lineStr := scanner.Text()
		if lineStr != "" {
			arr := strings.Fields(lineStr)
			res = append(res, arr...)
		}
	}
	return res, nil
}

// 根据字符出现顺序唯一保存
// 采用map[string]int
func UniqueStore(words []string) []string {
	uniqueMap := make(map[string]int)
	index := 0

	for _, word := range words {
		// 假如未出现
		if _, ok := uniqueMap[word]; !ok {
			uniqueMap[word] = index
			index++
		}
	}

	// 用数组保存map中的key和value，并排序
	orderedWords,_:=orderMapWord(uniqueMap)
	return orderedWords
}

func orderMapWord(uniqueMap map[string]int) ([]string,[]int) {
	wordNums:=len(uniqueMap)
	words:=make([]string,0)
	sequences:=make([]int,0)

	for word,sequence:=range uniqueMap{
		words = append(words, word)
		sequences=append(sequences, sequence)
	}

	for i := 0; i < wordNums; i++ {
		for j := 0; j < wordNums-1-i; j++ {
			if sequences[j] > sequences[j+1] {
				sequences[j], sequences[j+1] = sequences[j+1], sequences[j]
				words[j], words[j+1] = words[j+1], words[j]
			}
		}
	}
	return words,sequences
}