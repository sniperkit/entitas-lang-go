package main

import (
	"fmt"
	"os"

	EntitasLang "github.com/SirMetathyst/entitas-lang-go"
)

/*
func RunScanner() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := NewScanner(file)
	tok := ILLEGAL
	for tok != EOF {
		tok, _ = scanner.Scan()
		fmt.Print("{" + strings.ToUpper(tok.String()) + "}")
	}
	file.Close()
	fmt.Printf("\n\n")
}*/

func main() {
	//RunScanner()

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	project, err := EntitasLang.Parse(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(project.TargetDecl.Target)
	fmt.Println(project.NamespaceDecl.Namespace)

	fmt.Printf("ContextDecl: %s\n\n", project.ContextDecl)

	for _, aliasDecl := range project.AliasDecl {
		fmt.Printf("AliasDecl: %s\n", aliasDecl)
	}
	fmt.Println()

	for _, componentDecl := range project.ComponentDecl {
		fmt.Print("Component: " + componentDecl.Name)
		fmt.Printf("(%s)", componentDecl.Parameter)

		if len(componentDecl.Context) > 0 {
			fmt.Print(" in ")
			for _, c := range componentDecl.Context {
				fmt.Printf("[%s]", c)
			}
		}
		fmt.Println()
	}
}
