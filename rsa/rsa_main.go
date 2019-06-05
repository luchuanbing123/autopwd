package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"io/ioutil"
	"os"
)

func main() {
	CreateKey()
}

func CreateKey() {
	size := 1024
	privateKey, _ := rsa.GenerateKey(rand.Reader, size)
	bf := x509.MarshalPKCS1PrivateKey(privateKey)
	ioutil.WriteFile("../key.autopwd", []byte(bf), os.ModeTemporary)
}

// func GenRsaKey(bits int) error {
// 	s := "hellod大家记得加"
// 	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
// 	if err != nil {
// 		return err
// 	}

// 	d, _ := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, []byte(s))
// 	fmt.Println(d)
// 	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
// 	ddds := []byte(derStream)

// 	ioutil.WriteFile("key.pwd", ddds, os.ModeTemporary)

// 	dadwad, _ := ioutil.ReadFile("key.pwd")
// 	pre, _ := x509.ParsePKCS1PrivateKey(dadwad)

// 	s1, _ := pre.Decrypt(rand.Reader, d, nil)

// 	fmt.Println(s1)
// 	return nil
// }
