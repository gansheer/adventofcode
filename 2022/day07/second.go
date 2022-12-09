package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type Data struct {
	currentDir string
	dirSizes   map[string]int64
}

const totalsize = int64(70000000)
const minfreesize = int64(30000000)

func main() {

	fileName := "input1.txt"

	// open the file
	file, err := os.Open(fileName)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	// read line by line
	fmt.Println(doTheThing(fileScanner))

	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()

}

func doTheThing(fileScanner *bufio.Scanner) string {
	data := Data{currentDir: "", dirSizes: make(map[string]int64)}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("********", line)
		switch string(line[0]) {
		case "$":
			exploitCommand(line, &data)
		case "d":
			exploitDir(line, &data)
		default:
			exploitFile(line, &data)
		}
		print(data)
	}

	fmt.Println("---> nothing")

	return strconv.FormatInt(computeResult(&data), 10)
}

func exploitCommand(line string, data *Data) {
	fmt.Println("----> exploitCommand")
	fmt.Println(line)
	details := strings.Split(line, " ")
	switch string(details[1]) {
	case "ls":
		fmt.Println("ls")
	case "cd":
		fmt.Println("cd", details[2])
		executeCd(line, data)
	default:
		fmt.Println("unknown ERROR ERRROR ERRROR", details[1])
	}
}

func executeCd(line string, data *Data) {
	details := strings.Split(line, " ")
	switch string(details[2]) {
	case "..":
		fmt.Println("go to parent")
		gotToParentDir(data)
	case "/":
		fmt.Println("go to root")
		data.currentDir = "/"
	default:
		fmt.Println("go to child", details[2])
		data.currentDir = data.currentDir + details[2] + "/"
		// add current dir if not exists
		//if _, ok := data.dirSizes[data.currentDir]; ok {
		//nothinh to do something here
		//} else {
		//data.dirSizes[data.currentDir] = 0
		//}
	}
}

func gotToParentDir(data *Data) {
	dirs := strings.Split(data.currentDir, "/")
	//fmt.Println("BEFORE", data.currentDir)
	data.currentDir = strings.Join(dirs[:len(dirs)-2], "/") + "/"
	//fmt.Println("AFTER", data.currentDir)
}

func exploitDir(line string, data *Data) {
	fmt.Println("----> exploitDir")
	details := strings.Split(line, " ")
	dirPath := data.currentDir + details[1] + "/"
	if _, ok := data.dirSizes[dirPath]; ok {
		//nothinh to do something here
	} else {
		data.dirSizes[dirPath] = 0
	}

	fmt.Println(line)
}

func exploitFile(line string, data *Data) {
	fmt.Println("----> exploitFile")
	details := strings.Split(line, " ")
	filesize, _ := strconv.Atoi(details[0])
	data.dirSizes[data.currentDir] = data.dirSizes[data.currentDir] + int64(filesize)
	fmt.Println(line)
}

func computeResult(data *Data) int64 {
	fmt.Println("----> computeResult")
	fullDirs := make(map[string]int64)

	for k, v := range data.dirSizes {
		dirNode := k
		// if not "/"
		if len(dirNode) > 1 {
			dirNode = path.Dir(dirNode)
		}
		fmt.Println(dirNode, v)

		if _, ok := fullDirs[dirNode]; ok {
			fullDirs[dirNode] = fullDirs[dirNode] + v
		} else {
			fullDirs[dirNode] = v
		}

		for dirNode != "/" {
			dirNode = path.Dir(dirNode)
			fmt.Println(dirNode)
			if _, ok := fullDirs[dirNode]; ok {
				fullDirs[dirNode] = fullDirs[dirNode] + v
			} else {
				fullDirs[dirNode] = v
			}
		}

	}

	fmt.Println("fulldirs")
	usedsize := fullDirs["/"]
	unusedsize := totalsize - fullDirs["/"]
	targetunusedsize := minfreesize
	targetfreenecessary := targetunusedsize - unusedsize

	fmt.Println("Used size :", usedsize)
	fmt.Println("Unused size :", unusedsize)
	fmt.Println("Target unused size :", targetunusedsize)
	fmt.Println("To free size :", targetfreenecessary)

	dirToDelete := "/"
	dirToDeleteSize := fullDirs["/"]
	fmt.Println(dirToDelete, dirToDeleteSize)
	for dir, dirSize := range fullDirs {
		fmt.Println(dir, dirSize)

		if dirSize > targetfreenecessary && dirSize < dirToDeleteSize {
			dirToDelete = dir
			dirToDeleteSize = dirSize

		}

		fmt.Println(dirToDelete, dirToDeleteSize)
	}

	return dirToDeleteSize
}

func print(data Data) {
	fmt.Println("------------")
	for k, v := range data.dirSizes {
		fmt.Println(k, v)
	}
	fmt.Println("---")
	fmt.Println(data.currentDir)
	fmt.Println("------------")
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
