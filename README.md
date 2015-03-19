# Dfa-go
>Dfa-go is a Deterministic Finite Automation algorithm for golang.This program is for filter some text if it is a lllegal words.

### Example
```go
package main

import (
    "fmt"

    . "common"
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


``` 

### Output
```txt
Original： 我的微信请联系我吧 ，Result： true
Original： 我的微信号123456789请联系我吧 ，Result： true
Original： 加我微信1235445 ，Result： true
Original： 加我wechat ，Result： true
Original： 我的微 信 123565656fdsaf，联系我吧 ，Result： true
Original： 画的不错！这位同学 ，Result： false
--------Filter And Replace words---------
Original： 我的微信请联系我吧 ，Result： 我的**请联系我吧
Original： 我的微信号123456789请联系我吧 ，Result： 我的************请联系我吧
Original： 加我微信1235445 ，Result： 加我*********
Original： 加我wechat ，Result： 加我******
Original： 我的微 信 123565656fdsaf，联系我吧 ，Result： 我的******************，联系我吧
Original： 画的不错！这位同学 ，Result： 画的不错！这位同学
```

### Lllegal Words
>微.信.

>微 信

>微　信
　
>微.信.号R[\d\w]+R

>微信

>微信号R[\d\w]+R

>微信R[\d\w]+R

>微 信 R[\d\w]+R

>wechat