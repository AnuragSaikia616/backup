package main

import (
	"fmt"
	"os"
)

func main() {
	root := "/home/anurag/x/"
	arr, err := ListDir(root)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("List of files in dir: ", arr)
}

func ListDir(root string) ([]string, error) {
	filearr := make([]string, 0)
	NameFilesInDir(root, &filearr)
	if len(filearr) == 0 {
		return filearr, fmt.Errorf("ERROR: The dir is empty or does not exist")
	}
	return filearr, nil

}

func NameFilesInDir(dir string, arr *[]string) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			NameFilesInDir(dir+file.Name(), arr)
		}
		*arr = append(*arr, dir+file.Name())
	}
}
