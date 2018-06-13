package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	exibeIntroducao()

	for {
		exibeComando()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(5 * time.Minute)
	}

	fmt.Println("")
}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func exibeIntroducao() {
	nome := "Rafael"
	versao := 1.1
	fmt.Println("Olá ", nome)
	fmt.Println("Versão: ", versao)
}

func exibeComando() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
