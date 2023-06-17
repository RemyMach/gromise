package gromise

import (
	"context"
	"sync"
)

type ProcessFunc func() (interface{}, error)

func All(funcs []ProcessFunc) ([]interface{}, error) {
	var wg sync.WaitGroup
	results := make([]interface{}, len(funcs))
	errs := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i, fn := range funcs {
		wg.Add(1)
		go func(i int, fn ProcessFunc) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				res, err := fn()
				if err != nil {
					select {
					case errs <- err:
					default:
					}
					cancel() // Annule le contexte, ce qui déclenche ctx.Done()
					return
				}
				results[i] = res
			}
		}(i, fn)
	}

	wg.Wait()
	close(errs)

	if err, ok := <-errs; ok { // Vérifiez s'il y a une erreur dans le canal
		return nil, err
	}


	return results, nil
}
