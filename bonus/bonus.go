package bonus

import (
	"errors"
	"fmt"
)

func Destino(paths [][2]string) (string, error) {
	cityOutgoing := make(map[string]bool)

	for _, path := range paths {
		fmt.Println(cityOutgoing)
		cityOutgoing[path[0]] = true
	}

	for _, path := range paths {
		if !cityOutgoing[path[1]] {
			return path[1], nil
		}
	}
	return "", errors.New("Not implemented yet")
}
