package main

import (
	"fmt"
	"github.com/hrittikhere/cncf-landscape/platforms"
	"github.com/mmcdole/gofeed"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
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
	// fmt.Println("Starting Landscape Loop", len(landscapeconfig.Landscape))

	for _, category := range landscapeconfig.Landscape {
		// fmt.Println("Currently inspecting category ", index)
		// fmt.Println("Starting subcategory Loop", len(category.Subcategories))

		for _, subcategory := range category.Subcategories {
			// fmt.Println("Currently inspecting subcategory ", index)
			// fmt.Println("Starting items Loop", len(subcategory.Items))

			for _, item := range subcategory.Items {
				// fmt.Println("Currently inspecting item ", index)

				if !(item.RepoUrl == "") {
					rssFeed := item.RepoUrl + "/releases.atom"

					parser(rssFeed, item.Name)
				}
			}

		}
	}

}

func parser(feedLink string, itemName string) {

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(feedLink)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(" number of feed items: ", len(feed.Items))

	for _, item := range feed.Items {
		// fmt.Println("Currently Inspecting Feed Item", item.Link)
		NowTime := time.Now()
		ParsedNowTime := time.Unix(NowTime.Unix(), 0)

		PublishedTime := item.PublishedParsed
		ParsedPublishedTime := time.Unix(PublishedTime.Unix(), 0)

		if ParsedNowTime.Sub(ParsedPublishedTime).Hours() < 6 {
			// 0 */4 * * *
			PostTitle := item.Title
			PostLink := item.Link

			Tweet := fmt.Sprintf("%s is released by %s 🎉🎉🎉 \n %s ", PostTitle, itemName, PostLink)

			TweeetId, _ := cmd.PublishToTwitter(Tweet)

			fmt.Println(TweeetId, "Posted")

		}

	}

}
