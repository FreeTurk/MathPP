package main

import (
	"bufio"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	check(err)
	defer file.Close()
	out, err := os.Create("OUTPUT.go")
	check(err)

	out.WriteString("package main\nimport (\n\"fmt\"\n)\nfunc main() {\n")

	defer out.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			out.WriteString("\n")
		} else if scanner.Text()[0:6] == "chkpol" {
			data := strings.Split(scanner.Text()[6:], " = ")
			out.WriteString("var" + data[0] + " bool\n" + "if " + data[1] + " % 2 == 0 {\n" + data[0] + " = true\n} else {\n" + data[0] + " = false\n}")
		} else if scanner.Text()[0:4] == "iter" {
			varname := strings.Split(scanner.Text()[4:], " = ")
			data := strings.Split(varname[1], ", ")
			out.WriteString("var" + varname[0] + " []int\n")
			out.WriteString("for i := " + data[3] + "; i <= " + data[4] + "; i++ {\n" + "for n := " + data[0] + "; n <= " + data[1] + "; n++ {\nresult := n " + data[2] + " i\n" + varname[0] + " = append(" + varname[0] + ", result)}}\n")
		} else if scanner.Text()[0:3] == "var" {
			data := strings.Split(scanner.Text()[4:], " = ")
			out.WriteString(data[0] + ":=" + data[1] + "\n")
		} else if scanner.Text()[0:5] == "print" {
			data := scanner.Text()[6:]
			out.WriteString("fmt.Println(" + data + ")\n")
		} else if scanner.Text()[0:2] == "//" {
			out.WriteString(scanner.Text() + "\n")
		}
	}

	out.WriteString("}")
	out.Close()
}
