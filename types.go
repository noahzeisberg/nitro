package main

type NitroPackage struct {
}

type Command struct {
	Name        string
	Description string
	Args        Args
	Run         func([]string)
}

type Args struct {
	Count int
	Get   []string
}

type Arguments []string
