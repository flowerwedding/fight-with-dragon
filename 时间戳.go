package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var slicee =make([]int64,0,0)
	var sliceee =make([]string,0,0)
	var count int
    for i:=0;;i++ {
    	var s string
		_, _ = fmt.Scanln(&s)
    	if s=="" {break}
		v,_:=strconv.ParseInt(s,10,64)
    	slicee =append(slicee,v)
		t := time.Now().Unix()
		t -= slicee[i]
		defer fmt.Println("   ",time.Unix(t, 0).Format("2006-01-02 15:04:05.000000001 +0800 CST"))
		sliceee =append(sliceee,"input ok!")
		count++
    }
    fmt.Println("input：")
    for i:=0;i<=count-1;i++{
        fmt.Println("   ",slicee[i])
    }
    fmt.Println("   ","result")
    fmt.Println("output:")
    for i:=0;i<=count-1;i++{
        fmt.Println("   ",sliceee[i])
    }
    fmt.Println("   ","the result are :")
}