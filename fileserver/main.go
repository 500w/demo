package main

import (
	"flag"
	"log"
	"net/http"
)

// 快速启动一个文件服务器
// go install
// fileserver [flags]
// -d 指定目录 默认为当前目录
// -p 指定端口 默认为8888
func main() {
	port := flag.String("p", "8888", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
