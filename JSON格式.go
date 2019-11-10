package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var arrayy [10]byte
	var mingzi string
	xingxi:=make(map[string]int)
	for xuehao:=2019210001;xuehao<=2019215203;xuehao++ {
		resp, err := http.Get("http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + strconv.Itoa(xuehao))
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		for u, v := range body {
			for i := 0; i <= 8; i++ {
				arrayy[i] = arrayy[i+1]
			}
			arrayy[9] = v
			if string(arrayy[:]) == strconv.Itoa(xuehao) {
				for m, n := range body[u:] {
					if string(n) == " " {
						mingzi = string(body[u+1 : m+u])
						break
					}
				}
				break
			}
		}
		xingxi[mingzi]++
		fmt.Println(mingzi,xingxi[mingzi])
	}
	var maxnCount int
	var maxnCount2 int
	var maxnName string
	var maxnName2 string
	for m,n:=range xingxi {
		if m=="" {continue}
		if n>=maxnCount{maxnCount2=maxnCount;maxnName2=maxnName;maxnCount=n;maxnName=m}
	}
	shuchu(maxnName,maxnCount)
	if maxnCount==maxnCount2{shuchu(maxnName2,maxnCount2)}
}

type Student struct {
	Name    string
	Count   int
}

func shuchu(maxnName string,maxnCount int){
	st := &Student {
		maxnName,
		maxnCount,
	}
	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding faild")
	} else {
		fmt.Println(string(b))
	}
}