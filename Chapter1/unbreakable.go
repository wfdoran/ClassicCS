package main

import (
	"crypto/rand"
	"fmt"
)

func random_key(length int) []uint8 {
	rv := make([]uint8, length)
	n, err := rand.Read(rv)
	if err != nil {
		panic(err)
	}
	if n != length {
		panic(n)
	}
	return rv
}

func encrypt(orig string) ([]uint8, []uint8) {
	length := len(orig)
	key := random_key(length)

	cipher := make([]uint8, length)

	for i := 0; i < length; i++ {
		cipher[i] = orig[i] ^ key[i]
	}

	return cipher, key
}

func decrypt(cipher []uint8, key []uint8) string {
	rv := ""
	length := len(cipher)

	for i := 0; i < length; i++ {
		rv += string(cipher[i] ^ key[i])
	}
	return rv
}

func main() {
	s := "Hello World!"

	cipher, key := encrypt(s)

	for _, c := range cipher {
		fmt.Printf("%02x ", c)
	}
	fmt.Println()
	for _, c := range key {
		fmt.Printf("%02x ", c)
	}
	fmt.Println()
	t := decrypt(cipher, key)
	fmt.Println(t)
}
