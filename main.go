package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Subcategories
type Subcategories struct {
	Name        string      `yaml:"name"`
	Items       []Items     `yaml:"items"`
	Subcategory interface{} `yaml:"subcategory"`
}

// Items
type Items struct {
	Item        interface{} `yaml:"item"`
	Name        string      `yaml:"name"`
	HomepageUrl string      `yaml:"homepage_url"`
	RepoUrl     string      `yaml:"repo_url"`
	Logo        string      `yaml:"logo"`
	Twitter     string      `yaml:"twitter"`
	Crunchbase  string      `yaml:"crunchbase"`
}

// LandscapeSchema
type LandscapeSchema struct {
	Landscape []Landscape `yaml:"landscape"`
}

// Landscape
type Landscape struct {
	Subcategories []Subcategories `yaml:"subcategories"`
	Category      interface{}     `yaml:"category"`
	Name          string          `yaml:"name"`
}

func main() {

	landscape, err := ioutil.ReadFile("landscape.yml")
	if err != nil {
		fmt.Println("Error reading config file: ", err)
	}

	var landscapeconfig LandscapeSchema

	err = yaml.Unmarshal(landscape, &landscapeconfig)
	if err != nil {
		fmt.Println("Error parsing config file: ", err)
	}

	for _, feed := range landscapeconfig.Landscape {
		for _, subcategory := range feed.Subcategories {
			for _, item := range subcategory.Items {
				RepoUrl := item.RepoUrl
				if (RepoUrl == "") {
					fmt.Println( item.Name)
				}
			}

		}
	}

}
