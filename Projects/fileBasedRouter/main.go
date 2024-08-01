package main

import (
	"fmt"
	"os"
)

func main() {
	l, _ := getList("pages/")
	for _, route := range l {
		fmt.Println(route)
	}
}

func getList(root string) ([]string, error) {
	filelist := make([]string, 0)
	getFilePaths(root, &filelist)
	if len(filelist) == 0 {
		fmt.Println("The directory is empty or does not exits")
		return nil, fmt.Errorf("Dir empty OR !exits")
	}
	return filelist, nil
}

func getFilePaths(dir string, arr *[]string) {
	dirContent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, thing := range dirContent {
		if thing.IsDir() {
			getFilePaths(dir+thing.Name(), arr)
		}
		*arr = append(*arr, dir+thing.Name())
	}
}
