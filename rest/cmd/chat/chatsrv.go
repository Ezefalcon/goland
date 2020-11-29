package main

import config "config"

func main() {
	config.LoadConfig
	cfg, err := config.LoadConfig("asd")
}
