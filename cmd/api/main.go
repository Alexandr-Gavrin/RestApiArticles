package main

import (
	"ServerAndDB2/internal/app/api"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	// скажем, что наше приложение будет на этапе запуска получать пусть до конфиг файла из "внешнего мира"
	flag.StringVar(&configPath, "path", "configs/api.toml", "Path to config file in .toml format")

}

func main() {
	// В этот момент происходит инициализация переменной configPath
	flag.Parse()
	log.Println("It works!")
	// server instance initialization
	config := api.NewConfig()
	// теперь тут надо попробовать прочитать из .toml/.env, тк там может быть новая инфа
	_, err := toml.DecodeFile(configPath, config) // десеарилизуем содержимое .toml файл
	if err != nil {
		log.Println("Cant find configs file. Using default values:", err)

	}
	server := api.New(config)
	// api start
	log.Fatal(server.Start())
}
