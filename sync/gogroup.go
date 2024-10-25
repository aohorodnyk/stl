package sync

import "sync"

// GoGroup is a very simple wrapper on top of standard sync.WaitGroup.
// There are many mistakes happened during using the standard WaitGroup by forgetting to add `Add`
// or calling `Done` multiple times or in the wrong place, etc.
// Many people use `ErrGroup` from `x`, but it's counter-intuitive and makes things more complicated,
// rather than simplify things.
// GoGroup is called to simplify code and help newbyes in Go to write simple and clear code.
// This struct is useful when multiple goroutins should be ran and waited until all of them will finish execution.
type GoGroup struct {
	wg sync.WaitGroup
}

// Go function just manages `Add` and `Done` logic from the sync.WaitGroup
// and creates a gorouting with provided callback.
func (g *GoGroup) Go(callback func()) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		callback()
	}()
}

// Wait is calling `Wait` function from sync.WaitGroup to wait until all goroutines will finish.
func (g *GoGroup) Wait() {
	g.wg.Wait()
}
