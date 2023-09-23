package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"crypto/sha1"
)

func crc32Cryp() {
	h := crc32.NewIEEE()
	h.Write([]byte("teste"))
	i := crc32.NewIEEE()
	i.Write([]byte("testes"))
	j := crc32.NewIEEE()
	j.Write([]byte("teste"))
	fmt.Println(h.Sum32(), "|", i.Sum32(), "|", j.Sum32())
}

func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer f.Close()

	h := crc32.NewIEEE()

	_, err = io.Copy(h, f)

	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func fileHashRun() {
	h1, err := getHash("teste1.txt")

	if err != nil {
		return
	}

	h2, err := getHash("teste2.txt")

	if err != nil {
		return
	}

	h3, err := getHash("teste3.txt")

	if err != nil {
		return
	}

	result := fmt.Sprintf(
		"h1 == h2: %t | h1 == h3: %t | h2 == h3: %t", h1 == h2, h1 == h3, h2 == h3,
	)

	fmt.Println(result)
}

func sha1Crypt() {
	h := sha1.New()
	h.Write([]byte("test"))
	fmt.Println(h.Sum([]byte{}))
	s := sha1.New()
	s.Write([]byte("teste"))
	fmt.Println(s.Sum([]byte{}))
	f := sha1.New()
	f.Write([]byte("test"))
	fmt.Println(f.Sum([]byte{}))
}

func main() {
	sha1Crypt()
}
