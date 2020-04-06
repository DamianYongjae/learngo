package main

import (
	"fmt"

	"github.com/DamianYongjae/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")

	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	_, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
}
