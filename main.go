package main

import (
	"brasileirao/internal/times"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	timesLista := []times.Time{
		{Nome: "Botafogo", Pontos: 50, Jogos: 25, Vitorias: 15, Empates: 5, Derrotas: 5},
		{Nome: "Fortaleza", Pontos: 48, Jogos: 24, Vitorias: 14, Empates: 6, Derrotas: 4},
		{Nome: "Palmeiras", Pontos: 47, Jogos: 25, Vitorias: 14, Empates: 5, Derrotas: 6},
		{Nome: "Flamengo", Pontos: 44, Jogos: 24, Vitorias: 13, Empates: 5, Derrotas: 6},
		{Nome: "Cruzeiro", Pontos: 41, Jogos: 25, Vitorias: 12, Empates: 5, Derrotas: 8},
		{Nome: "São Paulo", Pontos: 41, Jogos: 25, Vitorias: 12, Empates: 5, Derrotas: 8},
		{Nome: "Bahia", Pontos: 29, Jogos: 25, Vitorias: 11, Empates: 6, Derrotas: 8},
		{Nome: "Vasco da Gama", Pontos: 34, Jogos: 24, Vitorias: 10, Empates: 4, Derrotas: 10},
		{Nome: "Atlético-MG", Pontos: 33, Jogos: 23, Vitorias: 8, Empates: 9, Derrotas: 6},
		{Nome: "Internacional", Pontos: 32, Jogos: 22, Vitorias: 8, Empates: 8, Derrotas: 6},
		{Nome: "Bragantino", Pontos: 30, Jogos: 24, Vitorias: 8, Empates: 6, Derrotas: 10},
		{Nome: "Athletico-PR", Pontos: 29, Jogos: 23, Vitorias: 8, Empates: 5, Derrotas: 10},
		{Nome: "Juventude", Pontos: 29, Jogos: 25, Vitorias: 7, Empates: 8, Derrotas: 10},
		{Nome: "Criciúma", Pontos: 28, Jogos: 24, Vitorias: 7, Empates: 7, Derrotas: 10},
		{Nome: "Grêmio", Pontos: 27, Jogos: 23, Vitorias: 8, Empates: 3, Derrotas: 12},
		{Nome: "Fluminense", Pontos: 27, Jogos: 24, Vitorias: 7, Empates: 6, Derrotas: 11},
		{Nome: "Corinthians", Pontos: 25, Jogos: 25, Vitorias: 5, Empates: 10, Derrotas: 10},
		{Nome: "EC Vitória", Pontos: 22, Jogos: 25, Vitorias: 6, Empates: 4, Derrotas: 15},
		{Nome: "Cuiabá", Pontos: 22, Jogos: 24, Vitorias: 5, Empates: 7, Derrotas: 12},
		{Nome: "Atlético-GO", Pontos: 18, Jogos: 25, Vitorias: 4, Empates: 6, Derrotas: 15},
	}

	times.OrdenarTimes(timesLista)

	times.ExibirTabela(timesLista)

	err := times.ExportarParaCSV(timesLista, "tabela_brasileirao.csv")
	if err != nil {
		fmt.Println("Erro ao exportar para CSV:", err)
	} else {
		fmt.Println("Tabela exportada com sucesso para tabela_brasileirao.csv")
	}

	http.HandleFunc("/tabela", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		var sb strings.Builder
		sb.WriteString(`
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Tabela do Brasileirão</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					text-align: center;
					background-color: #f0f0f0;
				}
				table {
					margin: auto;
					border-collapse: collapse;
					width: 80%;
				}
				th, td {
					padding: 12px;
					text-align: center;
					border: 1px solid #ddd;
				}
				th {
					background-color: #4CAF50;
					color: white;
				}
				tr:nth-child(even) {
					background-color: #f2f2f2;
				}
				tr:hover {
					background-color: #ddd;
				}
				h1 {
					color: #333;
				}
			</style>
		</head>
		<body>
			<h1>Tabela do Brasileirão</h1>
			<table>
				<tr>
					<th>Classificação</th>
					<th>Time</th>
					<th>Pontos</th>
					<th>Jogos</th>
					<th>Vitórias</th>
					<th>Empates</th>
					<th>Derrotas</th>
				</tr>`)

		for i, time := range timesLista {
			sb.WriteString(fmt.Sprintf(
				"<tr><td>%d</td><td>%s</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td><td>%d</td></tr>",
				i+1, time.Nome, time.Pontos, time.Jogos, time.Vitorias, time.Empates, time.Derrotas))
		}

		sb.WriteString(`
			</table>
		</body>
		</html>`)

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sb.String()))
	})

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
