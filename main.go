package main

import (
	"bufio"
	"flag"
	"os"
	"regexp"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func replacr(text string, toReplace string, color string, ignoreCase bool) string {
	if toReplace == "" {
		return text
	}
	if ignoreCase {
		toReplace = "(?i)" + toReplace
	}
	//panic(toReplace)
	r := regexp.MustCompile(toReplace)
	res := r.ReplaceAllFunc([]byte(text), func(s []byte) []byte {
		str := string(s)
		return []byte(color + str + Reset)
	})
	return string(res)
}

func main() {
	ignoreCaseFlag := flag.Bool("i", false, "ignore case")
	filterFlag := flag.Bool("filter", false, "do not filter the output")
	greenFlag := flag.String("green", "", "green text")
	redFlag := flag.String("red", "", "red text")
	yellowFlag := flag.String("yellow", "", "yellow text")
	blueFlag := flag.String("blue", "", "blue text")
	magentaFlag := flag.String("magenta", "", "magenta text")
	cyanFlag := flag.String("cyan", "", "cyan text")
	grayFlag := flag.String("gray", "", "gray text")
	whiteFlag := flag.String("white", "", "white text")

	flag.Parse()

	//fmt.Println("green", *greenFlag, "red", *redFlag)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sourceText := scanner.Text()
		text := replacr(sourceText, *greenFlag, Green, *ignoreCaseFlag)
		text = replacr(text, *redFlag, Red, *ignoreCaseFlag)
		text = replacr(text, *yellowFlag, Yellow, *ignoreCaseFlag)
		text = replacr(text, *blueFlag, Blue, *ignoreCaseFlag)
		text = replacr(text, *magentaFlag, Magenta, *ignoreCaseFlag)
		text = replacr(text, *cyanFlag, Cyan, *ignoreCaseFlag)
		text = replacr(text, *grayFlag, Gray, *ignoreCaseFlag)
		text = replacr(text, *whiteFlag, White, *ignoreCaseFlag)
		if *filterFlag && text == sourceText {
			continue
		}
		os.Stdout.WriteString(text + "\n")
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
