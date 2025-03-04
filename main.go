package main

import "fmt"

type Question struct {
	Text 	string
	Options []string
	Answer 	int
}

type GameState struct {
	Name 		string
	Points 		string
	Questions	[]Question
}

func (g *GameState) Init() {
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Println("Escreva o seu nome: ")
}

func main() {
	game := &GameState{}
	game.Init()
}
