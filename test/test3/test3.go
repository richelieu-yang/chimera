package main

import (
	"fmt"
	"log"
	"net/http"

	"crypto/tls"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've reached a secure server!")
}

func main() {
	// 服务器监听地址和端口
	addr := ":443"

	// 加载SSL证书和私钥文件
	certFile := "/path/to/your/cert.pem" // SSL证书路径
	keyFile := "/path/to/your/key.pem"   // 私钥路径

	// 解析并加载证书
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal("Failed to load certificate: ", err)
	}

	// 配置TLS配置结构体
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// 可以添加其他TLS配置项，例如支持的协议版本、加密套件等
	}

	// 创建一个基于TLS配置的HTTP服务器
	server := &http.Server{
		Addr: addr,
		//Handler:   http.HandlerFunc(hello),
		Handler:   http.HandlerFunc(hello),
		TLSConfig: config,
	}

	// 启动HTTPS服务
	log.Printf("Starting HTTPS server at %s", addr)
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
