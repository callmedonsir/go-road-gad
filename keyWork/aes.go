package main

import (
	//"os"
	"fmt"
	"os"
	"crypto/aes"
	"crypto/cipher"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
func main(){

	//需要去加密的字符串
	plaintext := []byte("My name is Dong")
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	//aes的加密字符串
	key_text := "My name is Dong , im 22 year old"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}

	fmt.Println(len(key_text))
	fmt.Println(len(commonIV))
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	//解密字符串
	cfbdec :=cipher.NewCFBDecrypter(c,commonIV)
	dec_ciphertext := make([]byte,len(ciphertext))
	cfbdec.XORKeyStream(dec_ciphertext,ciphertext)
	fmt.Printf("%x => %s\n",ciphertext,dec_ciphertext)
}


