package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func GetLinks() map[string]string {
	data, err := ioutil.ReadFile("links.yml")
	if err != nil {
		log.Panic("Error panic")
	}

	links := make(map[string]string)
	yaml.Unmarshal(data, links)

	return links
}