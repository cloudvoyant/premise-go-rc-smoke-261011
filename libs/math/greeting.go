package greeting

import local "github.com/cloudvoyant/premise-go-rc-smoke-261011/v4/libs/local"

// Public API -----------------------------------------------------------------
// Hello delegates to another package in the shared workspace module.
func Hello(name string) string { return local.Hello(name) }
