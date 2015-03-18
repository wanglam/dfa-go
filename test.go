package main

import (
	"fmt"

	. "yikaobang.cn/app/server/common"
)

func main() {
	arr := []string{
		"你的微信号是多少？",
		"你的微博号是多少？",
		"你的QQ号是多少？",
		"画的不错！",
		"加我微信1235445",
	}

	d := new(Dfa)
	d.BuildTree("words.txt")
	for _, v := range arr {
		fmt.Println("Original：", v, "，Result：", d.IsContain(v))
	}
	fmt.Println("--------Filter And Replace words---------")
	for _, v := range arr {
		fmt.Println("Original：", v, "，Result：", d.FilterWords(v, "*"))
	}
}
