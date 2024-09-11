package times

import (
	"fmt"
	"sort"
)

func OrdenarTimes(times []Time) {
	sort.Slice(times, func(i, j int) bool {
		return times[i].Pontos > times[j].Pontos
	})
}

func ExibirTabela(times []Time) {
	fmt.Println("Classificação | Time | Pontos | Jogos | Vitórias | Empates | Derrotas")
	for i, time := range times {
		fmt.Printf("%d | %s | %d | %d | %d | %d | %d\n", i+1, time.Nome, time.Pontos, time.Jogos, time.Vitorias, time.Empates, time.Derrotas)
	}
}
