package config

import (
	"encoding/json"
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

type Category struct {
	Name   string   `json:"name"`
	Parent string   `json:"parent"`
	Auto   []string `json:"auto"`
}

type Cfg struct {
	Categories []Category `json:"categories"`
}

const (
	Directory  = ".config/songbird"
	File       = "config.json"
	TokensFile = "tokens.json"
)

func LoadConfig() (Cfg, error) {
	var ret Cfg
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return ret, err
	}
	file, err := os.ReadFile(homeDirectory + "/" + Directory + "/" + File)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("config file not found, creating template")
			return createNewConfig(homeDirectory)
		}
		return ret, err
	}
	log.Println("config file found, loading")
	err = json.Unmarshal(file, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func createNewConfig(homePath string) (Cfg, error) {
	var ret Cfg
	ret.Categories = []Category{}
	err := os.MkdirAll(homePath+"/"+Directory, 0755)
	if err != nil {
		return ret, err
	}
	template, err := json.Marshal(ret)
	if err != nil {
		return ret, err
	}
	err = os.WriteFile(homePath+"/"+Directory+"/"+File, template, 0644)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func LoadTokens() ([]string, error) {
	var ret []string
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return ret, err
	}
	file, err := os.ReadFile(homeDirectory + "/" + Directory + "/" + TokensFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("tokens file not found, creating template")
			return createNewTokensFile(homeDirectory)
		}
		return ret, err
	}
	log.Println("tokens file found, loading")
	err = json.Unmarshal(file, &ret)
	if err != nil {
		return ret, err
	}
	if len(ret) == 0 {
		log.Println("no tokens found, please add a token")
		return ret, errors.New("no tokens found")
	}
	return ret, nil
}

func createNewTokensFile(homePath string) ([]string, error) {
	var ret []string
	err := os.MkdirAll(homePath+"/"+Directory, 0755)
	if err != nil {
		return ret, err
	}
	err = os.WriteFile(homePath+"/"+Directory+"/"+TokensFile, []byte("[]"), 0644)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
