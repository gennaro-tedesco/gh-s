package cmd

import (
	"log"
	"net/url"

	gh "github.com/cli/go-gh"
)

type repoInfo struct {
	Name        string
	Description string
	URL         string
	Stars       float64
}

func getRepos(query url.Values) []repoInfo {
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	var apiResults map[string]interface{}
	err = client.Get("search/repositories?"+query.Encode(), &apiResults)
	if err != nil {
		log.Fatal(err)
	}

	itemsResults := apiResults["items"].([]interface{})

	var repos []repoInfo
	for _, item := range itemsResults {
		repos = append(repos, repoInfo{
			Name:        item.(map[string]interface{})["full_name"].(string),
			Description: checkNil(item, "description"),
			URL:         item.(map[string]interface{})["html_url"].(string),
			Stars:       item.(map[string]interface{})["stargazers_count"].(float64),
		})
	}
	return repos
}
