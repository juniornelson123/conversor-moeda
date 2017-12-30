package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/juniornelson123/conversor-moeda/converter"
)

func main() {
	clear()
	ui()

}

func ui() {
	scanner3 := bufio.NewScanner(os.Stdin)
	var action string

	for action != "2" {

		fmt.Println("*****************Conversor de Moedas*******************")
		fmt.Println("1 - Converter moeda")
		fmt.Println("2 - Sair")
		fmt.Println("Selecione uma ação: ")

		scanner3.Scan()
		action = scanner3.Text()
		switch {
		case action == "1":
			convertCoin()
			break
		case action == "2":
			fmt.Println("Saindo...")
			os.Exit(200)
			break

		default:
			fmt.Println("Valor invalido")
			main()
			break
		}

	}
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func convertCoin() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner2 := bufio.NewScanner(os.Stdin)

	var option string
	var coin string
	var value float64

	for option != "q" {

		//Enter type coin - euro/libra/real/dolar
		fmt.Println("Entre com a moeda(euro, libra, real, dolar): ")

		scanner.Scan()
		option = scanner.Text()

		//Enter value for coin
		fmt.Println("Entre com o valor(ex: 15.23, 1.00): ")
		scanner1.Scan()

		value, _ = strconv.ParseFloat(scanner1.Text(), 64)

		//Enter type for convert coin
		fmt.Println("Entre com a moeda para conversão(euro, libra, real, dolar): ")
		scanner2.Scan()

		coin = scanner2.Text()

		c := converter.Coin{option, value}

		if option == "" || coin == "" || value == 0 {

			fmt.Println("### Preencha todos os campos ##\n")
			convertCoin()
		} else {

			convertValue, err := c.ConvertCoin(coin)

			if err != nil {
				fmt.Println("Erro ao tentar converter", err)
			} else {
				fmt.Printf("%.2f %s(s) são %.2f %s(s)\n\n\n", value, option, convertValue, coin)
			}

			restart()
		}

	}
}

func restart() {
	scanner3 := bufio.NewScanner(os.Stdin)
	var action string

	fmt.Println("Deseja Realizar mais alguma ação ?")
	fmt.Println("1 - Menu Principal")
	fmt.Println("2 - Sair")

	scanner3.Scan()
	action = scanner3.Text()
	switch {
	case action == "1":
		main()
		break
	case action == "2":
		fmt.Println("Saindo...")
		os.Exit(200)
		break

	default:
		fmt.Println("Valor invalido")
		restart()
		break
	}

}
