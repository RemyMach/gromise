package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	Value int32
	Error error
}

// func longRunningTask(ctx context.Context, num float32) <-chan Result {
// 	r := make(chan Result)

// 	go func() {
// 		defer close(r)

// 		// Simule une charge de travail.
// 		durRand := time.Duration(rand.Int63n(10000)) * time.Millisecond
// 		fmt.Println("Durée aléatoire:", durRand, "pour", num)
// 		time.Sleep(durRand)

// 		select {
// 		case <-ctx.Done():
// 			return // Retourne immédiatement si le contexte est annulé
// 		default:
// 			if num < 0.5 { // Simule une erreur 50% du temps
// 				r <- Result{Error: errors.New("une erreur est survenue")}
// 				return
// 			}

// 			r <- Result{Value: rand.Int31n(100)}
// 		}
// 	}()

// 	return r
// }

func longRunningTask(ctx context.Context, num float32) <-chan Result {
	r := make(chan Result)

	go func() {
		defer close(r)

		// Simule une charge de travail.
		durRand := time.Duration(rand.Int63n(10000)) * time.Millisecond
		fmt.Println("Durée aléatoire:", durRand, "pour", num)

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		// Utilise un ticker pour faire des pauses régulières
		// et vérifie l'état du contexte à chaque pause.

		for remaining := durRand; remaining > 0; {
			select {
			case <-ctx.Done():
				// Si le contexte est annulé, termine immédiatement.
				fmt.Println("Context cancelled for", num)
				return
			case <-ticker.C:
				// Sinon, continue à attendre.
				remaining -= 1 * time.Second
			}
		}

		// Une fois que l'attente est terminée, vérifie à nouveau l'état du contexte.
		select {
		case <-ctx.Done():
			return
		default:
			if num < 0.5 { // Simule une erreur 50% du temps
				r <- Result{Error: errors.New("une erreur est survenue")}
				return
			}

			r <- Result{Value: rand.Int31n(100)}
		}
	}()

	return r
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Assurez-vous d'annuler pour libérer des ressources même si tout se passe bien

	aCh, bCh, cCh := longRunningTask(ctx, 2), longRunningTask(ctx, 0.4), longRunningTask(ctx, 0.2)
	a, b, c := <-aCh, <-bCh, <-cCh

	if a.Error != nil || b.Error != nil || c.Error != nil {
		fmt.Println("Une erreur est survenue dans une ou plusieurs goroutines")
		cancel() // Annule le contexte pour stopper les autres goroutines
		return
	}

	fmt.Println(a.Value, b.Value, c.Value)
}
