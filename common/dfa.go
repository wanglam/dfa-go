package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	childrens map[string]*Node
	terminate bool
}

type Dfa struct {
	root Node
}

//添加敏感词
//@word敏感词
func (p *Dfa) addWord(word string) {
	node := &p.root
	str := []rune(word)
	lastPos := len(str) - 1
	for i, s := range str {
		sKey := string(s)
		if node.childrens[sKey] == nil {
			node.childrens[sKey] = &Node{map[string]*Node{}, i == lastPos}
		} else {
			if i == lastPos {
				node.childrens[sKey] = &Node{map[string]*Node{}, true}
			}
		}
		node = node.childrens[sKey]
	}
}

// 构建敏感词过滤器
// @filename 敏感词列表
func (p *Dfa) BuildTree(filename string) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	p.root = Node{map[string]*Node{}, false}

	r := bufio.NewReader(f)
	s, e := readln(r)
	p.addWord(s)
	for e == nil {
		s, e = readln(r)
		p.addWord(s)
	}
}

func readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

// 判断是否包含敏感词
// @text 输入文本
func (p *Dfa) IsContain(text string) bool {
	str := []rune(text)
	strLen := len(str)
	for i, _ := range str {
		j := i
		node := &p.root
		for j < strLen {
			currentWord := string(str[j])
			if node.childrens[currentWord] == nil {
				break
			} else {
				if node.childrens[currentWord].terminate {
					return true
				}
				node = node.childrens[currentWord]
			}
			j++
		}
	}
	return false
}

// 过滤关键词，替换为“*”
// @text 输入文本
func (p *Dfa) FilterWords(text, replaceChar string) string {
	str := []rune(text)
	strLen := len(str)
	strs := []string{}
	//是否继续添加
	isAppend := true
	for i := 0; i < strLen; i++ {
		node := &p.root
		for j := i; j < strLen && len(node.childrens) > 0; j++ {
			currentWord := string(str[j])
			if node.childrens[currentWord] == nil {
				break
			} else {
				if node.childrens[currentWord].terminate {
					isAppend = false
					temp := []string{}
					for k := 0; k < j-i+1; k++ {
						temp = append(temp, replaceChar)
					}
					strs = append(strs, strings.Join(temp, ""))
					i = j
				}
				node = node.childrens[currentWord]
			}
		}
		if isAppend {
			strs = append(strs, string(str[i]))
		} else {
			isAppend = true
		}
	}
	return strings.Join(strs, "")
}
