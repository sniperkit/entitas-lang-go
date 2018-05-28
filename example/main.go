package main

import (
	"fmt"
	"os"

	. "github.com/SirMetathyst/entitas-lang-go"
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
	if len(os.Args) <= 1 {
		return
	}

	//RunScanner()

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	parser := NewParser(file)

	/* Will be built-in eventually with default parser but can be overridden */
	parser.HandleContextDecl(func(p *Project, c *ContextDecl) error {
		if c.GetContextWithParameter("default") == nil {
			return fmt.Errorf("default context not defined!")
		}
		return nil
	})

	/* Will be built-in eventually with default parser but can be overridden */
	parser.HandleContext(func(p *Project, cd *ContextDecl, c *Context) error {
		if cd.GetContextWithName(c.Name) != nil {
			return fmt.Errorf("context with name '%s' is already defined!", c.Name)
		}
		return nil
	})

	project, err := parser.Parse()
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
