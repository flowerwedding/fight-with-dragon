package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var arrayy [10]byte
	var arrayx [6]byte
	var mingzi string
	var xingqi string
	var xuehao string
	_, _ = fmt.Scanln(&xuehao)
	resp, err := http.Get("http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + xuehao)
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
		for i := 0; i <= 4; i++ {
			arrayx[i] = arrayx[i+1]
		}
		arrayx[5] = v
		if string(arrayx[:]) == "星期" {
			xingqi = "星期" + string(body[u+2])
			break
		}
	}
	for u, v := range body {
		for i := 0; i <= 8; i++ {
			arrayy[i] = arrayy[i+1]
		}
		arrayy[9] = v
		if string(arrayy[:]) == xuehao {
			for m, n := range body[u:] {
				if string(n) == " " {
					mingzi = string(body[u+1 : m+u])
					break
				}
			}
			break
		}
	}
	var arrayz [len(" align='center'>理论")]byte
	var arrayj [len(" align='center'>理论")]byte
	mapp:=make(map[int]string)
	var arrayk [4]byte
	var arrayl [3]byte
	var kecheng string
	var shijian string
	var didian string
	fmt.Println("input:")
	fmt.Println("   " + xuehao)
	fmt.Println("output:")
	fmt.Println("   " + mingzi + "同学，你今日的课表为：")
	for u, v := range body {
		for i := 0; i <= len(" align='center'>理论")-2; i++ {
			arrayz[i] = arrayz[i+1]
		}
		arrayz[len(" align='center'>理论")-1] = v
		if string(arrayz[:]) == " align='center'>理论" || string(arrayz[:]) == " align='center'>实验" { //一个课程的循环开始
			for m, n := range body[u:] {
				for i := 0; i <= 4; i++ {
					arrayx[i] = arrayx[i+1]
				}
				arrayx[5] = n
				var xingqi2 string
				if string(arrayx[:]) == "星期" {
					xingqi2 = "星期" + string(body[u+m+1])
				}
				if xingqi2 == xingqi { //确定这门课是今天
					for d, c := range body[u:] { //课程名
						if string(c) == "-" {
							for a, b := range body[u+d:] {
								if string(b) == "<" {
									kecheng = string(body[d+u+1 : d+u+a])
									break
								}
							}
							break
						}
					}
					for i, o := range body[u+m:] { //时间名
						for i := 0; i <= 1; i++ {
							arrayl[i] = arrayl[i+1]
						}
						arrayl[2] = o
						if string(arrayl[:]) == "第" {
							for a, b := range body[u+m+i:] {
								for i := 0; i <= 1; i++ {
									arrayl[i] = arrayl[i+1]
								}
								arrayl[2] = b
								if string(arrayl[:]) == "节" {
									shijian = string(body[u+m+i+1 : u+m+i+a-2])
									break
								}
							}
							break
						}
					}
					for i, o := range body[u+m+12:] { //地点名
						for i := 0; i <= 2; i++ {
							arrayk[i] = arrayk[i+1]
						}
						arrayk[3] = o
						if string(arrayk[:]) == "<td>" {
							for a, b := range body[u+m+12+i:] {
								if string(b) == "<" {
									didian = string(body[u+m+12+i+1 : u+m+12+i+a])
									break
								}
							}
							break
						}
					}
					v := "   " + kecheng + "," + zhuanhuan(shijian) + "," + didian + "教室"
					x, _ :=strconv.Atoi(shijian[0:1])
					mapp[x]=v
				}
				for i := 0; i <= len(" align='center'>理论")-2; i++ { //一个课程的循环结束
					arrayj[i] = arrayj[i+1]
				}
				arrayj[len(" align='center'>理论")-1] = n
				if string(arrayj[:]) == " align='center'>理论" || string(arrayj[:]) == " align='center'>实验" {
					break
				}
			}
		}
	}
	for x:=1;x<=11;x+=2{
		if mapp[x]!=""{
			fmt.Println(mapp[x])
		}
	}
}
func zhuanhuan(shijian string)string {
	var dm [2]int
	var ans string
	j:=0
	for i:=1;i<=len(shijian);i++{
		if shijian[i-1:i]!="-" {
			a, _ := strconv.Atoi(shijian[i-1 : i])
			dm[j]=dm[j]*10+a
		}else{j++}
	}
	strr:="一二三四五六七八九十十一十二"
	for k:=dm[0];k<=dm[1];k++{
		ans+=strr[(k-1)*3:k*3]
	}
	ans+="节"
	p:=dm[1]-dm[0]
	if p!=1{ans+=strr[(p)*3:(p+1)*3]+"连上"}
	return ans
}