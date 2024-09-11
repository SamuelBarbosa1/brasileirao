package times

import "fmt"

type Time struct {
	Nome      string
	Pontos    int
	Jogos     int
	Vitorias  int
	Empates   int
	Derrotas  int
	SaldoGols int
	Historico []Jogo
}

func (t *Time) AtualizarEstatisticas(jogo Jogo) {
	t.SaldoGols += (jogo.GolsCasa - jogo.GolsFora)

	// Exemplo para calcular aproveitamento
	if t.Jogos > 0 {
		aproveitamento := (float64(t.Pontos) / float64(t.Jogos*3)) * 100
		fmt.Printf("Aproveitamento do %s: %.2f%%\n", t.Nome, aproveitamento)
	}
}
