# Dfa-go
>Dfa-go is a Deterministic Finite Automation algorithm for golang

### Example
```go
package main

import (
    "fmt"

    . "common"
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

``` 

### Output
```text
原字符串： 你的微信号是多少？ ，结果： true
原字符串： 你的微博号是多少？ ，结果： true
原字符串： 你的QQ号是多少？ ，结果： true
原字符串： 画的不错！ ，结果： false
原字符串： 加我微信1235445 ，结果： true
过滤字符串
原字符串： 你的微信号是多少？ ，结果： 你的**号是多少？
原字符串： 你的微博号是多少？ ，结果： 你的**号是多少？
原字符串： 你的QQ号是多少？ ，结果： 你的**号是多少？
原字符串： 画的不错！ ，结果： 画的不错！
原字符串： 加我微信1235445 ，结果： 加我**1235445
```