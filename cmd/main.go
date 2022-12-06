package main

// PORT Can be taken from environment variable by creating a config package using viper.
const PORT = 9000

func main() {
	initRouter(PORT)
}
