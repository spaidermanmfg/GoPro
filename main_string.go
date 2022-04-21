package main
import (
	"strings"
	"fmt"
	"bytes"
)
//字符串类型

func mainc() {
	//定义字符串
	var a0 string = "Hello, World!"
	var a1 = "Hello, World!"
	a2 := "Hello, World!"
	//打印字符串
	fmt.Printf("%T %v\n", a0, a0)
	fmt.Printf("%T %v\n", a1, a1)
	fmt.Printf("%T %v\n", a2, a2)

	//字符串拼接
	var b0 = "hello, "
	var b1 = "world!"
	var b2 = b0 + b1
	fmt.Printf("%T %v\n", b2, b2)

	//格式化字符串
	c0 := "hello"
	c1 := 18
	c2 := fmt.Sprintf("name=%s, age=%d", c0, c1)
	fmt.Printf("%T %v\n", c2, c2)

	//使用Join方法连接字符串
	d0 := "hello"
	d1 := "world"
	d2 := strings.Join([]string{d0, d1}, "-")
	fmt.Printf("%T %s\n", d2, d2)

	//使用bytes.buffer方法连接字符串
	var e0 = "hello"
	var e1 = "world"
	var buffer = bytes.Buffer{}//var buffer bytes.Buffer
	buffer.WriteString(e0)
	buffer.WriteString(",")
	buffer.WriteString(e1)
	fmt.Printf("%T %s\n", buffer.String(), buffer.String())

	//字符串切片
	var f0 = "Hello,World!"
	var f1 = f0[0:5]
	var f2 = f0[6:11]
	var f3 = 2
	var f4 = 5
	fmt.Printf("%T %s\n", f1, f1)
	fmt.Printf("%T %s\n", f2, f2)
	fmt.Printf("%v %c\n", f0[f3], f0[f3])//取第二位字符
	fmt.Printf("%s %s\n", f0[f3:f4], f0[f3:f4])//取第二位到第五位字符
	fmt.Printf("%s\n", f0[f3:])//从第二位取到最后一位字符
	//字符串函数
	//字符串长度
	fmt.Printf("%T %d\n", len(f0), len(f0))
	//分割字符串
	fmt.Printf("%T %v\n", strings.Split(f0, ","), strings.Split(f0, ","))
	//是否包含某个字符串，返回true或者false
	fmt.Printf("%T %v\n", strings.Contains(f0, "ll"), strings.Contains(f0, "ll"))
	//将字符串转换为小写
	fmt.Printf("%T %v\n", strings.ToLower(f0), strings.ToLower(f0))
	//将字符串转换为大写
	fmt.Printf("%T %v\n", strings.ToUpper(f0), strings.ToUpper(f0))
	//判断字符串是否以某个字符串开头，返回true或者false
	fmt.Printf("%T %v\n", strings.HasPrefix(f0, "H"), strings.HasPrefix(f0, "H"))
	//判断字符串是否以某个字符串结尾，返回true或者false
	fmt.Printf("%T %v\n", strings.HasSuffix(f0, "d"), strings.HasSuffix(f0, "d"))
	//查找字符串，返回第一个匹配的位置，如果没有则返回-1
	fmt.Printf("%T %v\n", strings.Index(f0, "o"), strings.Index(f0, "o"))
	//查找字符串，返回最后一个匹配的位置，如果没有则返回-1
	fmt.Printf("%T %v\n", strings.LastIndex(f0, "o"), strings.LastIndex(f0, "o"))
	//替换字符串
	fmt.Printf("%T %v\n", strings.Replace(f0, "o", "K", -1), strings.Replace(f0, "o", "K", -1))
	

	
	


}