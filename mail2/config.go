package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	pathLib "github.com/hilaily/kit/path"
)

/*
	读取配置文件
	配置文件默认在 ~/.config/mail2/config.toml
*/

var (
	configFileName = "mail2.toml"
	configDir      = ""
)

type configStruct struct {
	SmtpHost string `toml:"smtp_host"`
	SmtpPort int    `toml:"smtp_port"`
	UseSsl   bool   `toml:"use_ssl"`
	Mail     string `toml:"mail"`
	Password string `toml:"password"`
}

func init() {
	home, err := pathLib.GetHome()
	if err != nil {
		fmt.Printf("get home directory failed: %s\n", err)
		os.Exit(1)
	}
	configDir = home + "/.config/mail2"
}

// 初始化配置文件
func initConfigFile() {
	exist, err := configIsExist()
	if err != nil {
		fmt.Println("check config file failed: ", err)
		return
	}
	if exist {
		fmt.Println("config file is exist, please edit it and run")
		return
	} else {
		genConfigFile()
	}
}

// 检查配置文件是否存在
func configIsExist() (bool, error) {
	_, err := os.Stat(configDir + "/" + configFileName)
	if err != nil && os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func genConfigFile() {
	home, _ := pathLib.GetHome()
	filePath := home + "/.config/mail2"
	data := `stmp_host=
stmp_port=
use_ssl=
mail=
password=`
	fmt.Println("config is not exist, create it")
	os.MkdirAll(filePath, 0711)
	f, err := os.OpenFile(filePath+"/"+configFileName, os.O_WRONLY|os.O_CREATE, 0650)
	defer f.Close()
	if err != nil {
		fmt.Println("create config file failed: ", err)
		return
	}
	_, err = f.WriteString(data)
	if err != nil {
		fmt.Println("create config file failed: ", err)
	}
	fmt.Println("create config file succ")
}

func loadConfig(configPath string) *configStruct {
	c := &configStruct{}
	if configPath == "" {
		configPath = configDir + "/" + configFileName
	}
	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		fmt.Println("decode config file err: ", err)
		os.Exit(1)
	}
	return c
}
