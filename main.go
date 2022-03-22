package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hash := sha256.New()
	hash.Write([]byte("sand"))
	fmt.Printf("%x\n", hash.Sum(nil))

	hash1 := sha256.New()
	hash1.Write([]byte("sand"))
	fmt.Printf("%x\n", hash1.Sum(nil))

	fmt.Println(hash == hash1)
	fmt.Println(hash)
	fmt.Println(hash1)
}
