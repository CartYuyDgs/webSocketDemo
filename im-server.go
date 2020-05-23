package main

import (
	"container/list"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"webSocketDemo/config"
	"webSocketDemo/handler"
)

var G_Config *config.Config
var configDir = ".\\config\\config.json"

func main() {

	var err error
	var imHandler *handler.ImHandler

	if G_Config, err = getConfigInfo(); err != nil {
		log.Fatalln("Error Get Config Info failed")
		return
	}

	m := make(map[string]*list.List)

	imHandler = handler.NewImHandler(m, G_Config.LinkNum)
	http.HandleFunc(G_Config.Path, imHandler.Handler)

	http.ListenAndServe(":"+strconv.Itoa(G_Config.Port), nil)
}

func getConfigInfo() (*config.Config, error) {
	var (
		configDir = ".\\config\\config.json"
		content   []byte
		err       error
		Config    *config.Config
	)
	Config = &config.Config{}
	path := flag.String("f", configDir, "user im json")
	flag.Parse()
	//log.Fatalln(*path)

	if content, err = ioutil.ReadFile(*path); err != nil {
		log.Fatalf("Error %v\n", err)
		return Config, err
	}

	//Config = new(config.Config)
	err = json.Unmarshal(content, Config)
	if err != nil {
		log.Fatalf("Error %v\n", err)
		return Config, err
	}
	return Config, nil
}
