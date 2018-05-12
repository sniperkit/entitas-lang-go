Entitas-Lang-Go is currently a work in progress.
Currently implemented:
- Target
- Namespace
- Context

additional syntax added:
- KeyValue for context, component etc. ```context Game(default, key:value), GameState, Input```

Entitas-Lang Source:

```
target entitas_csharp

namespace my.game.test


context Game(default, key:value), GameState, Input

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
```
package main

import (
	"fmt"
	"os"
	"strings"

	elang "github.com/SirMetathyst/entitas-lang-go"
)

func main() {
	// Open File
	filename := os.Args[1] // First arg
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}

	f, err := elang.Parse(file)
	if err != nil {
		panic(err)
	}

	if f.TargetDecl != nil {
		fmt.Println(f.TargetDecl.Target)
	}

	if f.NamespaceDecl != nil {
		fmt.Println(f.NamespaceDecl.Namespace)
	}

	if f.ContextDecl != nil {
        // Print context name and context parameter key value pairs.
		for _, c := range f.ContextDecl.Context {
			fmt.Print("Context: " + c.ContextName)
			fmt.Print("(")
			for k, v := range c.ContextParameter {
				if k != "" {
					if v != "" {
						fmt.Printf("[%s:%s]", k, v)
					} else {
						fmt.Printf("[%s]", k)
					}
				}
			}
			fmt.Println(")")
		}
	}
}
```