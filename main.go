package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/nlpodyssey/safetensors"
	"golang.org/x/term"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func waitForEnter() {
        // Ugly mess of code to display a prompt and wait for Enter keypress
        // without displaying user input if other keys are pressed
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	check(err)
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stderr}
	terminal := term.NewTerminal(screen, "")
	_, _ = terminal.ReadPassword("Press 'Enter' to exit...")
}

func main() {
	// Check if argument is provided
	if len(os.Args) == 1 {
		panic(errors.New("no argument provided"))
	}

	// Open file provided as argument
	f, err := os.Open(os.Args[1])
	check(err)
	defer f.Close()

	// Read header size
	var l uint64
	err = binary.Read(f, binary.LittleEndian, &l)
	check(err)

	// Read header
	header := make([]byte, l)
	_, err = f.ReadAt(header, 8)
	check(err)

	// Get metadata from header
	var metadata safetensors.Metadata
	err = json.Unmarshal(header, &metadata)
	check(err)

	// Get tensors infos
	tensors := metadata.Tensors()

	// Get all tensors names
	layers := []string{}
	for name, _ := range tensors {
		layers = append(layers, name)
	}

	// Sort tensors names
	sort.Strings(layers)

	// Print names
	for _, name := range layers {
		fmt.Println(name)
	}

	// Wait for Enter keypress
	waitForEnter()
}
