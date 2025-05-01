package main

import (
	"debug/elf"
	"fmt"
	"log"
	"strings"
)

const path = "http_server"

func main() {
	elfFile, err := elf.Open(path)
	if err != nil {
		log.Fatalf("error while opening ELF File %s: %+s", path, err.Error())
	}

	// Extract the symbol table
	symbolTable, err := elfFile.Symbols()
	if err != nil {
		log.Fatalf("could not extract symbol table %+s", err.Error())
	}

	funcSymbols := []elf.Symbol{}
	s := "net/http"
	for _, symbol := range symbolTable {
		if strings.HasPrefix(symbol.Name, s) {
			if symbol.Info == 18 {
				fmt.Printf("Func Name: %s Func Type: %d \n", symbol.Name, symbol.Info)
				funcSymbols = append(funcSymbols, symbol)
			}

		}
	}

	// print the func Symbols
	// for _, value := range funcSymbols {
	// 	fmt.Printf("Func Name: %s \n", value.Name)
	// }
}
