package _package

import (
	"fmt"
	"github.com/Mimerel/go-logger-client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

/**
Reads configuration file
 */
func ReadConfiguration() (Configuration) {
	pathToFile := os.Getenv("HEATING_CONFIGURATION_FILE")
	if _, err := os.Stat("./configuration.yaml"); !os.IsNotExist(err) {
		pathToFile = "./configuration.yaml"
	} else if pathToFile == "" {
		pathToFile = "/home/pi/go/src/go-zway-last-values/configuration.yaml"
	}
	yamlFile, err := ioutil.ReadFile(pathToFile)

	if err != nil {
		fmt.Printf("Unable to read conf file : %+v \n", err)
		panic(err)
	}

	var config Configuration

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Unable to convert file to yaml : %+v \n", err)
		panic(err)
	} else {
		config.Logger = logs.New(config.Elasticsearch.Url, config.Host)
		config.Logger.Info("Configuration Loaded : %+v \n", config)
	}
	return config
}

