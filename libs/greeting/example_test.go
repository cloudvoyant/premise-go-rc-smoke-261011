package greeting_test

import (
	greeting "github.com/cloudvoyant/premise-go-rc-smoke-261011/libs/greeting"
	"fmt"
)

// Public API -----------------------------------------------------------------
func ExampleHello() {
	fmt.Println(greeting.Hello("Ada"))
	// Output: Hello, Ada!
}
