package main

import (
	"fmt"
	"math/rand"
	"strings"
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main(){
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)
	result := name + RandString(12)
	fmt.Println(result)
	for i := 0; i<len(result); i+=3 {
		result = strings.Replace(result, result[i:i+1], "X", 1)
	}
	fmt.Println(result)
}

/* Problem
1. your first name + 12 bit unique character string
2. Replace every 3rd character to X from index O
 */