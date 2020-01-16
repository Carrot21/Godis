package main

import (
	"Godis/cache-benchmark/cacheClient"
	"flag"
	"fmt"
	"os"
)

var(
	s string
	p string
	op string
	k string
	v string
)
func usage() {
	_, err := fmt.Fprintf(os.Stderr, `Godis client
Usage: client [-s server] [-p port] [-op operation] [-k key] [-v value]

Options:
 -s  指定Godis服务器地址
 -p  指定连接端口
 -op 指定操作类型，可以是get/set/del
 -k  指定键
 -v  指定值
`)
	if err != nil {
		panic(err.Error())
	}
}
func init() {
	flag.StringVar(&s,"s","localhost","服务器地址")
	flag.StringVar(&p,"p","2333","端口号")
	flag.StringVar(&op,"op","get","操作选项，可以是get/set/del")
	flag.StringVar(&k,"k","","key")
	flag.StringVar(&v,"v","","value")
	flag.Usage = usage
	flag.Parse()
}

func main()  {
	c := cacheClient.New("tcp", s,p)
	cmd := &cacheClient.Cmd{op,k,v,nil}
	c.Run(cmd)
	if cmd.Error != nil {
		fmt.Println("error:", cmd.Error)
	} else {
		fmt.Println(cmd.Value)
	}
}
