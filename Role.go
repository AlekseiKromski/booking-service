package main

type Role struct {
	Name    string
	Actions map[string]func()
}
