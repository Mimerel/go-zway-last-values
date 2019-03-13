package _package

import "github.com/Mimerel/go-logger-client"

type Configuration struct {
	Elasticsearch Elasticsearch `yaml:"elasticSearch,omitempty"`
	Host string `yaml:"host,omitempty"`
	Port string `yaml:"port,omitempty"`
	Local string `yaml:"local,omitempty"`
	Logger logs.LogParams
}

type Elasticsearch struct {
	Url string `yaml:"url,omitempty"`
}

type StructuredData struct {
	Metric string
	Labels map[string]string
	Timestamp string
	Timestamp2 string
	Value string
}