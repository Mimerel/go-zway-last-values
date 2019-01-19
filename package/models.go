package main

type Configuration struct {
	Token string `yaml:"token,omitempty"`
	Elasticsearch Elasticsearch `yaml:"elasticSearch,omitempty"`
	Host string `yaml:"host,omitempty"`
	Port string `yaml:"port,omitempty"`
	Local string `yaml:"local,omitempty"`
}

type Elasticsearch struct {
	Url string `yaml:"url,omitempty"`
}

type Result struct {
	Name string
	Value float64
}

type StructuredData struct {
	Metric string
	Labels map[string]string
	Timestamp string
	Timestamp2 string
	Value string
}