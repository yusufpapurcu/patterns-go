package main

import (
	"github.com/yusufpapurcu/patterns-go/dependency_injection/module"
	"github.com/yusufpapurcu/patterns-go/dependency_injection/writer"
)

func main() {
	w := &writer.DisplayWriter{}
	m := module.NewModule(w)

	m.ProcessEmails("aa", "bb")
}
