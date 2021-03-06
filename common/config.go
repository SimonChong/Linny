package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
)

type ConfigLinnyMeasure struct {
	Enabled      bool
	AuthRequired bool
}

type ConfigLinny struct {
	ContentRoot string
	Measure     ConfigLinnyMeasure
}

func LoadConfigLinny() (ConfigLinny, error) {
	conf := ConfigLinny{}
	err := conf.Init()
	return conf, err
}

func (c *ConfigLinny) Init() error {

	confStr, err := ioutil.ReadFile("configLinny.json")
	if err != nil {
		log.Println("configLinny.json is missing")
		return err
	}
	err = json.Unmarshal(confStr, c)
	if err != nil {
		log.Println("configLinny.json error:", err)
		return err
	}

	c.ContentRoot, err = filepath.Abs(c.ContentRoot)
	if err != nil {
		log.Println("ContentRoot Error: ", err)
		return err
	}
	// log.Println("ContentRoot: ", c.ContentRoot)
	return nil
}

func (c *ConfigLinny) Save() {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Println("Martial Error", err)
		return
	}
	ioutil.WriteFile("configLinny.json", data, 0777)
}

type ConfigAd struct {
	Id string
}

func LoadConfigAd(cl *ConfigLinny) (ConfigAd, error) {
	conf := ConfigAd{}
	err := conf.Init(cl)
	return conf, err
}

func (c *ConfigAd) Init(cl *ConfigLinny) error {

	fileStr := cl.ContentRoot + "/configAd.json"
	confStr, err := ioutil.ReadFile(fileStr)
	if err != nil {
		log.Println("configAd.json is missing")
		return err
	}
	err = json.Unmarshal(confStr, c)
	if err != nil {
		log.Println("configAd.json error:", err)
		return err
	}

	if c.Id == "" {
		log.Println("configAd.json error: Id field is missing")
		return errors.New("configAd.json error: Id field is missing")
	}
	return nil
}
