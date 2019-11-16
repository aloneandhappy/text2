package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
	str := bufio.NewReader(os.Stdin)
	fmt.Println("输入时间戳：")
	maintimestamp, _ := str.ReadString('\n')
	s_fields := strings.Fields(maintimestamp)
	for _, v := range s_fields {
		timestamp,_ :=strconv.Atoi(v)
		fmt.Println(time.Unix(int64(timestamp),0))
	}
}