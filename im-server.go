package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"webSocketDemo/config"
)

func main() {

	var (
		configDir = ".\\config\\config.json"
		content   []byte
		err       error
		Config    *config.Config
	)
	path := flag.String("f", configDir, "user im json")
	flag.Parse()
	//log.Fatalln(*path)

	if content, err = ioutil.ReadFile(*path); err != nil {
		log.Fatalf("Error %v\n", err)
	}
	Config = &config.Config{}
	//Config = new(config.Config)
	err = json.Unmarshal(content, Config)
	if err != nil {
		log.Fatalf("Error %v\n", err)
	}
	log.Println(string(content), Config.Path, Config.Port)
}
