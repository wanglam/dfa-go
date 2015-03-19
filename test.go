package main

import (
	"fmt"

	. "yikaobang.cn/app/server/common"
)

func main() {
	arr := []string{
		"我的微信请联系我吧",
		"我的微信号123456789请联系我吧",
		"加我微信1235445",
		"加我wechat",
		"我的微 信 123565656fdsaf，联系我吧",
		"画的不错！这位同学",
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
