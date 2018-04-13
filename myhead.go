package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	startOpt = flag.Int("s", 10, "number of start line is -s=%d")
	endOpt = flag.Int("e", 11, "number of end line is -n=%d")
	diffOpt = flag.Bool("d", false, "")
)

func readFile(path string) ([]string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var contents []string
	for scanner.Scan() {
		content := scanner.Text()
		contents = append(contents, content)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return contents, nil
}

func call(path string, contents []string, start int, end int) {
	if len(contents) < end {
		fmt.Printf("Sorry, but the number of lines in %s is less than %d", path, end)
	} else {
		lines := contents[start: end]
		fmt.Println("<<<<<<" + path + ">>>>>>>")
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}


func main() {
	flag.Parse()

	filePath := flag.Arg(0)
	otherFilePath := flag.Arg(1)
	fileContent, err := readFile(filePath)
	if err != nil {
		panic("")
	}
	call(filePath, fileContent, *startOpt, *endOpt)

	if *diffOpt == true {
		otherFile, err := readFile(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		call(otherFilePath, otherFile, *startOpt, *endOpt)
	}

}
