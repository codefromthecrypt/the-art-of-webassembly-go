package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tetratelabs/wazero/wasm"
	"github.com/tetratelabs/wazero/wasm/wazeroir"
)

func main() {
	value1, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid args %v: %v", os.Args[1], err)
		os.Exit(1)
	}

	value2, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid args %v: %v", os.Args[2], err)
		os.Exit(1)
	}

	buf, err := os.ReadFile("addint.wasm")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read wasm file: %v", err)
		os.Exit(1)
	}

	mod, err := wasm.DecodeModule(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode module: %v", err)
		os.Exit(1)
	}

	store := wasm.NewStore(wazeroir.NewEngine())

	err = store.Instantiate(mod, "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to instantiate: %v", err)
		os.Exit(1)
	}

	results, _, err := store.CallFunction("", "AddInt", value1, value2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to call function: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%d + %d = %d\n", value1, value2, results[0])
}
