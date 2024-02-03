package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// livp 文件本质就是一个zip包，里面包含一个MOV的视频和 HIEC的图片。转成zip包之后，可以直接解压。至于HIEC图片，直接用MacBook的调试即可。
func main() {
	root := "/home/eric/Pictures/新建共享文件夹图片"
	err := filepath.Walk(root, func(path string, info os.FileInfo, inerr error) error {

		if inerr != nil {
			log.Printf("error walking directories, skipping: %v", inerr)
			return filepath.SkipDir
		}
		if info.IsDir() {
			//log.Printf(" --> %s", path)
		}

		if !strings.HasSuffix(path, "livp") {
			return nil
		}

		if info.Mode()&os.ModeType != 0 {
			// skip everything but regular files
			//log.Printf("skip everything but regular files")
			return nil
		}

		destinationFileName := path + ".zip"

		err := renameLivpToZip(path, destinationFileName)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("File renamed successfully!")
		}
		// only consider files with requested extensions
		return nil
	})
	if err != nil {
		log.Printf("error walking %s: %v", root, err)
	}

}

func renameLivpToZip(sourceFileName, destinationFileName string) error {
	// 检查源文件是否存在
	if _, err := os.Stat(sourceFileName); os.IsNotExist(err) {
		return fmt.Errorf("Source file '%s' not found", sourceFileName)
	}

	// 检查目标文件是否存在，避免覆盖现有文件
	if _, err := os.Stat(destinationFileName); err == nil {
		return fmt.Errorf("Destination file '%s' already exists", destinationFileName)
	}

	// 获取源文件的绝对路径
	sourceFilePath, err := filepath.Abs(sourceFileName)
	if err != nil {
		return err
	}

	// 获取目标文件的绝对路径
	destinationFilePath, err := filepath.Abs(destinationFileName)
	if err != nil {
		return err
	}

	// 执行重命名操作
	err = os.Rename(sourceFilePath, destinationFilePath)
	if err != nil {
		return err
	}

	return nil
}
