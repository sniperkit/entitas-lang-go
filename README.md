Entitas-Lang-Go is currently a work in progress.
Currently implemented:
- Target
- Namespace
- Context
- Alias

TODO:
- Component (in progress)
- System

additional syntax added:
- KeyValue for context, component etc. ```context Game(default, key:value), GameState, Input```

Entitas-Lang Source:

```
target entitas_csharp

namespace my.game.test


context Game(default), GameState, Input

alias intList : "System.Collections.Generic.List<int>" stringList : "System.Collections.Generic.List<string>"
alias blueprints : "Entitas.Unity.Blueprints.Blueprints"
alias int : "int"
alias string : "string"
alias go : "UnityEngine.GameObject"

comp Blueprints (unique)
    value : blueprints

comp Destroyed

comp GameBoard(unique, event:global)
    columns : int
    rows : int

comp GameBoardElement
comp Movable

comp Position 
    x (entityindex:multiple) : int
    y : int

comp Score(unique) in GameState
    value : int

comp BurstModeComponent(unique) in Input

comp InputComponent in Input
    x : int
    y : int

comp Interactive

comp Asset
    name : string

comp View
    gameObject : go



sys FallSystem
    trigger:
        removed(GameBoardElement)
        noFilter
    access:
        _context : Game

sys FillSystem
    trigger:
        removed(GameBoardElement)
        noFilter
    access:
        _context : Game 

sys GameBoardSystem(init)
    trigger:
        added(GameBoard)
        filter allOf(GameBoard)
    access:
        _gameBoardElements : allOf(GameBoardElement, Position)
        _context : Game

sys ScoreSystem (init)
    trigger:
        removed(GameBoardElement)
        noFilter
    access:
        gameState : GameState

sys EmitInputSystem (cleanup)
    access:
        _context : Input
        _inputs : allOf(InputComponent)

sys ProcessInputSystem
    trigger:
        added(InputComponent)
        filter allOf(InputComponent)
    access:
        game: Game

sys AddViewSystem
    trigger:
        added(Asset)
        filter allOf(Asset) noneOf(View)
    access:
        _context : Game

sys AnimatePositionSystem
    trigger:
        added(Position)
        filter allOf(View, Position)
    access:
        _context : Game

sys RemoveViewSystem
    trigger:
        removed(Asset) added(Destroyed)
        filter allOf(View)

sys SetViewPositionSystem
    trigger:
        added(View)
        filter allOf(View, Position)

sys DestroySystem
    trigger:
        added(Destroyed)
        filter allOf(Destroyed)
    access:
        _context : Game
```

Entitas-Lang-Go Example Usage
```go
package main

import (
	"fmt"
	"os"
	"strings"

	elang "github.com/SirMetathyst/entitas-lang-go"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

    parser := elang.NewParser(file)
    
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
```