package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := flag.String("dir","/home/","directory to search")
	files := flag.String("files",".py,.go", "suffixes for files to be searched")
	search := flag.String("search","import requests", "phrase to search in files")
	flag.Parse()
	filesslice := strings.Split(*files,",")
	log.Printf("Searching directory %s and subdirectories, where files have suffixes %s and contain phrase: %s\n", *dir, *files, *search)
	filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error  {
		if info.IsDir(){
			return nil
		}else {
				for _, suffix := range filesslice{
					if strings.HasSuffix(path, suffix){
						if Check(path, *search){
							fmt.Println(path)
						}
					}
			}
		}
		return nil
	})
	log.Println("Done")
}

func Check(path string, search string) bool {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
		return false
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), search){
			return true
		}// the line
	}
	return false
}
