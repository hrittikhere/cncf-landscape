package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
	// "github.com/hrittikhere/cncf-landscape/platforms"
	"github.com/mmcdole/gofeed"
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
				if !(item.RepoUrl == "") {
				rssFeed := item.RepoUrl+"/releases.atom"
				parser(rssFeed)
				}
			}

		}
	}

}


func parser(feedLink string){
	
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL(feedLink)

	for _, item := range feed.Items {

		NowTime := time.Now()
		ParsedNowTime := time.Unix(NowTime.Unix(), 0)

		PublishedTime := item.PublishedParsed
		ParsedPublishedTime := time.Unix(PublishedTime.Unix(), 0)

		if ParsedNowTime.Sub(ParsedPublishedTime).Hours() < 400 {
			// 0 */4 * * *
			// PostTitle := item.Title
			// PostLink := item.Link
			// PostDescription := item.Description
			// PostPublished := item.Published
			// Categories := feed.Categories
			

			// fmt.Printf("%s \n %s %s \n %s %s \n", PostTitle, PostLink, PostDescription, PostPublished, Categories)
			fmt.Println(item.Link)
			fmt.Println("====================================================")
			// Tweet := fmt.Sprintf("%s was published by %s 🎉🎉🎉 \n %s ", PostTitle, PostLink)

			// TweeetId, _ := cmd.PublishToTwitter(Tweet)

			// fmt.Println(TweeetId, "Posted")

		}

	}

}