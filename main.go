package main

import (
	"bufio"
	"encoding/csv"
	"errors"
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
	Points 		int
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

func (g *GameState) ProcessCSV() int {
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
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Question{
				Text: 		record[0],
				Options: 	record[1:5],
				Answer: 	correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}

	return len(records) - 1
}

// toInt pode retornar integer e error
func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, errors.New("não é permitido caracter diferente de número")
	}

	return i, nil
}

func (g *GameState) Run() {
	for index, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Println("Digite uma alternativa")

		var answer int
		var err error

		// for infinito, só sai quando o usuário digita o valor correto.
		// Sai apenas no "break"
		// No "continue" executa o for novamente
		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[:len(read)-1]) // é feito um slice na string

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		if answer == question.Answer {
			fmt.Println("Parabéns você acertou!")
			g.Points += 10
		} else {
			fmt.Println("Você errou!")
			fmt.Println("----------------------")
		}
	}
}

func main() {
	game := &GameState{}
	len := /* go */ game.ProcessCSV()
	game.Init()
	game.Run()

	score := (game.Points * 100) / (len * 10) // calcula porcentagem de acertos

	var msg string = ""

	if score >= 90 {
		msg = "Excelente pontuação!"
	} else if score >= 70 {
		msg = "Boa pontuação!"
	} else {
		msg = "Baixa pontuação."
	}

	fmt.Printf(msg + "Fim de jogo, você fez %d pontos\n", game.Points)
}
