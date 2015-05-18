package override

import "time"

// A Signal is used to communicate with the Spawn function.
type Signal bool

const (
	// Force instructs Spawn to skip to the next cycle.
	Force Signal = true
	// Stop instructs Spawn to terminate its goroutine.
	Stop Signal = false
)

// Go spawns a goroutine that repeatedly calls fn and then waits for the
// specified duration. It returns a channel that can be used to bypass the wait
// and immediately continue to the next call/wait cycle.
func Go(fn func(), every time.Duration) chan<- struct{} {
	force := make(chan struct{})
	go func() {
		for {
			fn()
			select {
			case <-time.After(every):
			case <-force:
			}
		}
	}()
	return force
}

// Spawn behaves the same as Go, except that it allows the caller to terminate
// the goroutine. To perform a normal override, send Force. To terminate the
// goroutine, send Stop.
func Spawn(fn func(), every time.Duration) chan<- Signal {
	sigChan := make(chan Signal)
	go func() {
		// call/wait loop
		for {
			fn()
			select {
			case <-time.After(every):
			case s := <-sigChan:
				if s == Stop {
					return
				}
				// any other Signal is treated as a Force
			}
		}
	}()
	return sigChan
}
