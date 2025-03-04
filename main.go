package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

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
	reader := bufio.NewReader(os.Stdin)

	/* ReadString recebe o caracterer limite de leitura. Ou seja, quando identificar o \n para a leitura */
	/* ReadString retorna um "name" ou erro */
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = name

	fmt.Printf("Vamos ao jogo %s", name)
}

func (g *GameState) ProcessCSV() {
	file, err := os.Open("questions.csv")

	if err != nil {
		panic("Erro ao ler arquivo")
	}

	/* defer é uma função que executa só após todo o resto da função ser executada */
	/* defer é a última execução, independente da ordem */
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Error ao ler CSV")
	}

	for index, record := range records {
		fmt.Println(index, record)
		if index > 0 {
			question := Question{
				Text: 		record[0],
				Options: 	record[1:5],
				Answer: 	toInt(record[5]),
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func main() {
	game := &GameState{}
	go game.ProcessCSV()
	game.Init()

	fmt.Println(game.Questions)
}
