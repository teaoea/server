package tools

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Mkdir 创建文章目录,如果目录已存在,中止创建
func Mkdir(filepath string) {
	path := fmt.Sprintf("./%s/%d/%s/%d",
		filepath, time.Now().Year(), time.Now().Month(), time.Now().Day())
	m := os.MkdirAll(path, os.ModePerm)
	if m != nil {
		return
	}
}

// WriteMd 把文章内容保存为md文件
// path: 文件路径
// content: 内容
// return 返回文件路径,数据库存储文件路径
func WriteMd(path, content string) string {
	Mkdir(path)
	// 文章保存路径,文章内容保存为md文件
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
