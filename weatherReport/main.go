package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type openWeatherApiConfig struct {
	ApiKey string `json:"ApiKey,omitempty"`
}

type weatherData struct {
	Name string `json:"name,omitempty"`
	Main struct {
		Kelvin float64 `json:"temp,omitempty"`
	} `json:"main,omitempty"`
}

func loadApiConfig(filename string) (openWeatherApiConfig, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return openWeatherApiConfig{}, err
	}
	var c openWeatherApiConfig
	json.Unmarshal(bytes, &c)
	if err != nil {
		return openWeatherApiConfig{}, err
	}
	return c, nil
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiconfig")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiConfig.ApiKey)
	if err != nil {
		return weatherData{}, err
	}
	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()

	return d, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weathter/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		fmt.Println(city)
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(":8000", nil)
}
