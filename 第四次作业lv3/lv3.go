package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func Getxm(num int)string {
	url := []string{"http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh="}
	newurl := append(url, strconv.Itoa(num))
	Url := newurl[0] + newurl[1]
	resp, _ := http.Get(Url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result :=regexp.MustCompile(">>.*?  ")
	result1:=regexp.MustCompile("[\u4e00-\u9fa5].*[\u4e00-\u9fa5]")
	n:=result.FindString(string(body))
	xinming := result1.FindString(n)
	return xinming
}

func Tonji(newMap map[string]int)([]string,int) {
	var allcount []int
	var maxcount int
	for _, value := range newMap {
		allcount = append(allcount, value)
	}
	maxcount = allcount[0]
	for i:=0; i<len(allcount);i++ {
		if maxcount < allcount[i] {
			maxcount =allcount[i]
		}
	}
	var maxvalue []string
	for key, value := range newMap {
		if value == maxcount {
			maxvalue = append(maxvalue, key)
		}
	}
	return maxvalue,maxcount
}

func main()  {
	newMap := make(map[string]int)
	var num int
	  for x:=2016;x<=2019;x++ {
		for y:=210000;y<=229000;y++{
			num =x*1000000+y
		    go Getxm(num)
		    x := Getxm(num)
		    if x != "" {
		    v, ok := newMap[x]
		    if ok != false {
			      newMap[x] = v + 1
		    } else {
			      newMap[x] = 1
		    }
		    }
	    }
	}
	go Tonji(newMap)
	fmt.Println(Tonji(newMap))
	/*var allcount []int
	var maxcount int
	for _, value := range newMap {
		allcount = append(allcount, value)
	}
	maxcount = allcount[0]
	for i:=0; i<len(allcount);i++ {
		if maxcount < allcount[i] {
			maxcount =allcount[i]
		}
	}
	var maxvalue []string
	for key, value := range newMap {
		if value == maxcount {
			maxvalue = append(maxvalue, key)
		}
	}*/
}