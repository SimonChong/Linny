package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	ContentRoot string
}

func NewConfig() Config {
	conf := Config{}
	conf.Init()
	return conf
}

func (c *Config) Init() {

	confStr, err := ioutil.ReadFile("linnyConfig.json")
	if err != nil {
		fmt.Println("linnyConfig.json is missing")
		return
	}
	err = json.Unmarshal(confStr, c)
	if err != nil {
		fmt.Println("linnyConfig.json error:", err)
		return
	}

	c.ContentRoot, err = filepath.Abs(c.ContentRoot)
	if err != nil {
		fmt.Println("ContentRoot Error: ", err)
		return
	}
	fmt.Println("ContentRoot: ", c.ContentRoot)
}
