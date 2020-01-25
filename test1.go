package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	drives := getDrives()
	files := []string{}
	SMALL := []string{}
	BIG := []string{}
	for _, d := range drives {
		FindFileFromExtension([]string{".jpg", ".gif"}, d, &files, &SMALL, &BIG)
	}
	for _, F := range files {
		fmt.Printf("Path : %v\n", F)
	}
	fmt.Printf("Files in one month : %v\n", len(files))
	fmt.Printf("Small file : %v\n", len(SMALL))
	fmt.Printf("big file : %v\n", len(BIG))
	fmt.Println("Run time : ", time.Since(start))
}

func getDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d := string(drive) + ":/"
			r = append(r, d)
			f.Close()
		}
	}
	return
}

//FindFileFromExtension ddd
func FindFileFromExtension(extension []string, dir string, files *[]string, SMALL *[]string, BIG *[]string) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range fs {
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {
					if (time.Now().Unix() - f.ModTime().Unix()) < 2591999 {
						*files = append(*files, dir+"/"+f.Name())
						if f.Size() < 1000000 {
							*SMALL = append(*files, dir+"/"+f.Name())
						}
						if f.Size() > 1000000 {
							*BIG = append(*files, dir+"/"+f.Name())
						}
					}
				}
			}

			if f.IsDir() {
				path := dir + "/" + f.Name()
				FindFileFromExtension(extension, path, files, SMALL, BIG)
			}
		}
	}
}
