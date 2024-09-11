package times

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ExportarParaCSV(times []Time, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Nome", "Pontos", "Jogos", "Vitórias", "Empates", "Derrotas", "SaldoGols"})
	for _, time := range times {
		writer.Write([]string{
			time.Nome,
			fmt.Sprintf("%d", time.Pontos),
			fmt.Sprintf("%d", time.Jogos),
			fmt.Sprintf("%d", time.Vitorias),
			fmt.Sprintf("%d", time.Empates),
			fmt.Sprintf("%d", time.Derrotas),
			fmt.Sprintf("%d", time.SaldoGols),
		})
	}

	return nil
}
