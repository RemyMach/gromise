package process

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func ForInt(n int) (interface{}, error) {
	if n == 2 || n == 1 {
		return 0, errors.New("Error processing number 2")
	}

	fmt.Println("Processing", n)
	time.Sleep(4 * time.Second)

	// Calcul du résultat
	result := n*2 - 1
	return result, nil
}

func ForString(s string) (interface{}, error) {
	// Ici, je retourne simplement la chaîne en majuscule pour l'exemple.
	return strings.ToUpper(s), nil
}