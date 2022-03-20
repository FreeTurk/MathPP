package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	error := color.New(color.FgRed, color.Bold)
	file, err := os.Open(os.Args[1])
	check(err)
	defer file.Close()
	out, err := os.Create("OUTPUT.go")
	check(err)

	out.WriteString("package main\nimport (\n\"fmt\"\n)\nfunc main() {\n")
	defer out.Close()
	scanner := bufio.NewScanner(file)
	line := 1
	for scanner.Scan() {
		text := scanner.Text()
		keyword := strings.SplitN(text, " ", 2)
		if scanner.Text() == "" {
			out.WriteString("\n")
		} else if keyword[0] == "chkpol" {
			data := strings.Split(keyword[1], " = ")
			out.WriteString("var " + data[0] + " bool\n" + "if " + data[1] + " % 2 == 0 {\n" + data[0] + " = true\n} else {\n" + data[0] + " = false\n}")
		} else if keyword[0] == "iter" {
			varname := strings.Split(keyword[1], " = ")
			data := strings.Split(varname[1], ", ")
			out.WriteString("var " + varname[0] + " []int\n")
			out.WriteString("for i := " + data[3] + "; i <= " + data[4] + "; i++ {\n" + "for n := " + data[0] + "; n <= " + data[1] + "; n++ {\nresult := n " + data[2] + " i\n" + varname[0] + " = append(" + varname[0] + ", result)}}\n")
		} else if keyword[0] == "var" {
			data := strings.Split(keyword[1], " = ")
			out.WriteString(data[0] + " := " + data[1] + "\n")
		} else if keyword[0] == "print" {
			data := keyword[1]
			out.WriteString("fmt.Println(" + data + ")\n")
		} else if keyword[0] == "//" {
			out.WriteString(scanner.Text() + "\n")
		} else {
			error.Println("An error occured at line " + strconv.Itoa(line) + ". You probably used an undefined keyword.\n" + ">>> " + scanner.Text())
			os.Remove("OUTPUT.go")
			out.Close()
			break
		}
		line++
	}

	out.WriteString("}")
	out.Close()
}
