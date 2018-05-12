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
```go
package main

import (
	"fmt"
	"os"
	"strings"

	elang "github.com/SirMetathyst/entitas-lang-go"
)

func main() {
	// Open file.
	file, err = os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

    // Parse source.
	project, err := elang.Parse(file)
	if err != nil {
		panic(err)
	}

    // Print target and namespace.
    // target entitas_csharp
    // namespace my.game
	fmt.Println(project.TargetDecl.Target)
	fmt.Println(project.NamespaceDecl.Namespace)

    // Print context names and their parameters. e.g. Game(default, myParam, etc, key:value)
	for _, context := range project.ContextDecl.Context {
		fmt.Print("Context: " + context.ContextName)
		fmt.Print("(")
		for k, v := range context.ContextParameter {
			if v != "" {
				fmt.Printf("[%s:%s]", k, v)
			} else {
				fmt.Printf("[%s]", k)
			}
		}
		fmt.Println(")")
	}
}
```