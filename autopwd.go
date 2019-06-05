package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/LuChuanBing/autopwd/asset"
)

func main() {
	keyBuffer, _ := asset.Asset("key.autopwd")
	privateKey, _ := x509.ParsePKCS1PrivateKey(keyBuffer)

	rsaHelper := RSAHelper{*privateKey}

	s := "daidhwoaihdw好的好的哈"
	d := rsaHelper.Encrypt(s)
	s1 := rsaHelper.Decrypt(d)
	fmt.Println(s1)

	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9876", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
}

type RSAHelper struct {
	privateKey rsa.PrivateKey
}

func (helper *RSAHelper) Encrypt(in string) []byte {
	publicKey := helper.privateKey.PublicKey
	out, _ := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, []byte(in))
	return out
}

func (helper *RSAHelper) Decrypt(in []byte) string {
	privateKey := helper.privateKey
	out, _ := rsa.DecryptPKCS1v15(rand.Reader, &privateKey, in)
	return string(out)
}
