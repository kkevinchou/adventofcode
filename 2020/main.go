package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dirs, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dir := range dirs {
		dirName := dir.Name()
		fmt.Println(dirName)
		if strings.HasPrefix(dirName, "day") {
			continue
		}

		split := strings.Split(dir.Name(), "_")
		day := split[0]
		part := split[1]

		os.MkdirAll(fmt.Sprintf("day%s/part%s", day, part), os.ModeDir)

		files, err := os.ReadDir("./" + dirName)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, file := range files {
			fileName := "./" + dirName + "/" + file.Name()
			readFile, err := os.Open(fileName)
			if err != nil {
				panic(err)
			}

			outFile := fmt.Sprintf("day%s/part%s/%s", day, part, filepath.Base(file.Name()))
			writeFile, err := os.Create(outFile)
			if err != nil {
				panic(err)
			}

			io.Copy(writeFile, readFile)
		}
	}
}
