package main

import "os"

func main() {
	err := Main(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}

func Main(args []string) error {
	return nil
}
