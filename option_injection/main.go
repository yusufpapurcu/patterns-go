package main

func main() {
	m := NewModule(WithTLS(true), WithMaxConnection(10), WithPort(8080))
	m.PrintOption()
}
