package bonus

import (
	"errors"
	"strconv"
)

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	for linha := range caminhos {
		for _, elemento := range caminhos[linha] {
			strconv.Atoi(elemento)
		}
		sort
	}

	return "", errors.New("not implemented yet")
}
