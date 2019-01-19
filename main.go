package main

import (
	"encoding/json"
	"fmt"
	"github.com/Mimerel/go-logger-client"
	"github.com/op/go-logging"
	"go-zway-last-values/package"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)



var log = logging.MustGetLogger("default")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`,
)


func main() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.NOTICE, "")
	logging.SetBackend(backendLeveled, backendFormatter)

	config := _package.ReadConfiguration()
	Port := config.Port
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		urlParams := strings.Split(urlPath, "/")
		if len(urlParams) > 1  {
			collectMetricValuesAndSend(w, r, urlParams, &config)
		} else {
			w.WriteHeader(500)
		}
	})
	http.ListenAndServe(":" + Port, nil)
}

func collectMetricValuesAndSend(w http.ResponseWriter, r *http.Request, urlParams []string, config *_package.Configuration) {
	var results []_package.Result
	yamlFile, err := ioutil.ReadFile(config.Local)
	if err != nil {
		logs.Error(config.Elasticsearch.Url, config.Host, fmt.Sprintf("Unable to read file from local storeg %s ", err))
		w.WriteHeader(500)
	}

	var data []_package.StructuredData

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		logs.Error(config.Elasticsearch.Url, config.Host, fmt.Sprintf("Unable to yaml to structure %s ", err))

	}
	for key, value := range urlParams {
		if key != 0 {
			result, err := requestMetric(data, value)
			if err != nil {
				logs.Error(config.Elasticsearch.Url, config.Host, fmt.Sprintf("Unable to find value for metric %s %s", value, err))
			} else {
				results = append(results, result)
			}
		}
	}
	js, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func requestMetric(data []_package.StructuredData, metricName string) (_package.Result, error) {
	for _, value := range data {
		if value.Metric == metricName {
			metricValue, err := strconv.ParseFloat(value.Value,  64)
			if err != nil {
				return _package.Result{} , fmt.Errorf("Unable to convert metric value to float")
			}
			return _package.Result{metricName, metricValue}, nil
		}
	}
	return _package.Result{} , fmt.Errorf("Unable to find metric in storage_file")
}
