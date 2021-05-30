package tools

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Mkdir
/// create a directory, if the directory already exists, abort the creation
func Mkdir(filepath string) {
	path := fmt.Sprintf("./%s/%d/%s/%d",
		filepath, time.Now().Year(), time.Now().Month(), time.Now().Day())
	m := os.MkdirAll(path, os.ModePerm)
	if m != nil {
		return
	}
}

// WriteMd
/// save the content of the article as a file,
/// and return the file path, the database save file path
func WriteMd(path, content string) string {
	Mkdir(path)
	filepath := fmt.Sprintf("%s/%d/%s/%d/%d.md",
		path, time.Now().Year(), time.Now().Month(), time.Now().Day(), NewId())
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	write := bufio.NewWriter(file)
	_, _ = write.WriteString(content)
	_ = write.Flush()
	return filepath
}
