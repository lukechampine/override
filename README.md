# override #

`override` provides two functions, `Go` and `Spawn`, which can be used to implement an override pattern.

### Example usage ###

```go
package main

import "time"
import "github.com/lukechampine/override"

func foo() { println("hi!") }

func main() {
	ch := override.Spawn(foo, time.Second) // will print "hi!" once per second

	ch <- override.Force // "hi!" is printed immediately
	// Force resets the sleep timer, so the next "hi" will be printed one
	// second after this send.

	ch <- override.Stop // not available when using Go
}
```

Full documentation can be found on [godoc](http://godoc.org/github.com/lukechampine/override)
