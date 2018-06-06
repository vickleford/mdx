package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type markdownRunner struct {
	tmpfile  string
	markdown io.Reader
}

func (m *markdownRunner) scanForCommands(reader io.Reader) []string {
	var commands []string
	var reapCommands bool

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		if (line == "```" || line == "-->") && reapCommands {
			reapCommands = false
		} else if reapCommands {
			commands = append(commands, line)
		} else if line == "```shell" || line == "<!--mdx" {
			reapCommands = true
		}
	}

	return commands
}

func (m *markdownRunner) writeGuideSteps(contents []string) (func(), error) {
	filename := m.tmpfile

	cleanup := func() {
		err := os.Remove(filename)
		if err != nil {
			fmt.Printf("Could not clean up file %q automatically because %q\n",
				filename, err)
		}
	}

	dest, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Something bad happened when trying to open file %q"+
			" for reading: %q\n", filename, err)
		return cleanup, err
	}

	dest.Write([]byte("#!/usr/bin/env bash\n\n"))
	for _, line := range contents {
		_, err := dest.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Printf("Something bad happened when writing line %q: %q\n",
				line, err)
			return cleanup, err
		}
	}

	return cleanup, nil
}

func (m *markdownRunner) Execute() error {
	var err error
	commands := m.scanForCommands(m.markdown)
	cleanGuideScript, err := m.writeGuideSteps(commands)
	defer cleanGuideScript()
	if err != nil {
		return fmt.Errorf("Could not create temporary file %q to execute guide scripts from: %q", m.tmpfile, err)
	}

	commandsFromMd := exec.Command("/bin/bash", "-x", "-e", m.tmpfile)
	commandsFromMd.Stdout = os.Stdout
	commandsFromMd.Stderr = os.Stderr

	err = commandsFromMd.Run()
	if err != nil {
		return fmt.Errorf("Could not execute temporary file %q created from markdown: %q", m.tmpfile, err)
	}

	return nil
}

func NewMarkdownRunner(md io.Reader) *markdownRunner {
	m := markdownRunner{markdown: md}
	m.tmpfile = "/tmp/farout2" // maybe /tmp is a bad idea because noexec
	return &m
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Printf("Error opening file %q: %q\n", filename, err)
		os.Exit(1)
	}

	runmd := NewMarkdownRunner(file)
	runErr := runmd.Execute()
	if runErr != nil {
		fmt.Printf("Got error: %s", err)
		os.Exit(1)
	}
}
