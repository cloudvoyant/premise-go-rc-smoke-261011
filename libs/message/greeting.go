package greeting
import basegreeting "github.com/cloudvoyant/premise-go-rc-smoke-261011/libs/greeting"
// Public API -----------------------------------------------------------------
// Hello returns a greeting from the public sibling module.
func Hello(name string) string { return basegreeting.Hello(name) }
