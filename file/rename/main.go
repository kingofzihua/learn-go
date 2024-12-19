package main

import (
	"fmt"
	hashids "github.com/speps/go-hashids"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Defiens alphabet.
const (
	Alphabet62   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Alphabet36   = "abcdefghijklmnopqrstuvwxyz1234567890"
	HASH_ID_SALT = "xiangzhi"
)

func main() {
	// 检查是否传递了目录参数
	if len(os.Args) < 2 {
		log.Fatalf("请提供要重命名文件的目录")
	}
	dir := os.Args[1]

	// 确保提供的路径是一个目录
	dirInfo, err := os.Stat(dir)
	if err != nil {
		log.Fatalf("无法访问目录 %s: %v", dir, err)
	}
	if !dirInfo.IsDir() {
		log.Fatalf("提供的路径 %s 不是一个目录", dir)
	}

	// 读取目录下的文件和文件夹
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("读取目录内容失败: %v", err)
	}

	// 遍历目录中的文件
	for _, file := range files {
		// 检查是否为文件而不是文件夹
		if file.IsDir() {
			continue
		}

		var err error

		// 获取文件名和扩展名
		ext := filepath.Ext(file.Name())
		baseName := strings.TrimSuffix(file.Name(), ext)

		idx, err := strconv.Atoi(baseName)
		if err != nil {
			fmt.Printf("文件名[%s]不能为您转化为数字，已为您跳过改文件: %s \n", baseName, err)
			continue
		}

		// 检查文件是否满足格式n.mp4
		suffix := GetInstanceID(idx, "")
		newName := fmt.Sprintf("%s-%s%s", baseName, suffix, ext)
		oldPath := filepath.Join(dir, file.Name())
		newPath := filepath.Join(dir, newName)

		err = os.Rename(oldPath, newPath)
		if err != nil {
			log.Printf("重命名文件 %s 失败: %v\n", file.Name(), err)
		} else {
			fmt.Printf("已将文件 %s 重命名为 %s\n", file.Name(), newName)
		}
	}
}

// GetInstanceID returns id format like: secret-2v69o5
func GetInstanceID(uid int, prefix string) string {
	hd := hashids.NewData()
	hd.Alphabet = Alphabet36
	hd.MinLength = 6
	hd.Salt = HASH_ID_SALT

	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}

	i, err := h.Encode([]int{int(uid)})
	if err != nil {
		panic(err)
	}

	return prefix + Reverse(i)
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
