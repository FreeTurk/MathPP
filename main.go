package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func preinit(line string, linenum int) {
	errortext := color.New(color.FgRed).Add(color.Bold).SprintFunc()
	erroremp := color.New(color.FgRed).Add(color.Underline).SprintFunc()
	keywords := map[string]bool{"iter": true, "var": true, "//": true, "print": true, "chkpol": true, "": true}
	words := strings.Split(line, " ")
	if !keywords[words[0]] {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println(errortext("\nUnknown keyword used at line ", linenum, "\n", ">>>"), erroremp(words[0]), strings.Join(words[1:], " ")+"\n")
		os.Exit(0)
	}
}

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	successtext := color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	// get an argument
	if len(os.Args) != 2 {
		fmt.Println("Please input a file: matt FILENAME.mtt")
		os.Exit(0)
	}
	file := os.Args[1]

	// open the file
	f, err := os.Open(file)
	check(err)

	// read the file line by line
	scanner := bufio.NewScanner(f)
	i := 0
	spinnerd := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	spinnerd.Prefix = "Checking for errors "
	spinnerd.Start()
	time.Sleep(400 * time.Millisecond)
	for scanner.Scan() {
		i += 1
		preinit(scanner.Text(), i)
	}
	spinnerd.Stop()

	fmt.Println(successtext("\nNo errors found!\n"))
	spinnerd = spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	spinnerd.Prefix = "Initializing "
	spinnerd.Start()
	time.Sleep(400 * time.Millisecond)

	spinnerd.Stop()

	// close the file
	err = f.Close()
	check(err)
}
