package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

//Created for config hold
//env-default:"production"
type Config struct{
	Env string `Yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `Yaml:"storage_path" env-required:"true"`
	HttpServer HTTPServer `Yaml:"http_server"`
}

type HTTPServer struct{
	Address string
}

//Execute File
func MustLoad() *Config{
	var configPath string 
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags:=flag.String("config","","path to config file")
		flag.Parse()

		configPath = *flags
	}

	if configPath == ""{
		log.Fatal("Config path is not set")
	}

	//check error in file
	if _,error := os.Stat(configPath); os.IsNotExist(error){
		log.Fatalf("config does not exist: %s", configPath)
	}
	var cfg Config
	error:=cleanenv.ReadConfig(configPath,&cfg)
	if error != nil {
		log.Fatalf("Can not read config file: %s",error.Error())
	}


	return &cfg

}


