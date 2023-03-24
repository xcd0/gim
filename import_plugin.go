package gim

import "plugin"

func import_plugin() {
	p, err := plugin.Open("plugin/plugin.so")
	// error handling
	f, err := p.Lookup("F")
	// error handling
	f.(func())() // prints "Hello, world"
}
