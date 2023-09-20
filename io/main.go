package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func readArchive() {
	file, err := os.Open("teste.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// file size
	stat, err := file.Stat()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stat.Size())

	// read content

	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)

	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bs)
	fmt.Println(str)

}

func ioutilReader() {
	bs, err := ioutil.ReadFile("teste.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

func createFile() {
	file, err := os.Create("teste2.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	file.WriteString("Unable are the loved to die,\nFor love is immortality")
}

func osReadDir() {

	userDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		return
	}

	dir, err := os.Open(userDir + "/Documents")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer dir.Close()

	fileInfos, err := dir.ReadDir(-1)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func walk() {
	userDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		return
	}

	filepath.Walk(userDir+"/Documents", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

func main() {
	walk()
}
