package hackernews

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"publishers/internal/sources"
)

type Source struct {
	Source string
}

type fetcher interface {
	GetStories() ([]int, error)
}

const SourceName = "hackernews"

type hackernewsAPI struct{}

func (api *hackernewsAPI) GetStories() ([]int, error) {
	var stories []int

	response, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")

	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(response.Body).Decode(&stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

func NewHackernewsSource() *Source {
	return &Source{Source: SourceName}
}

func (hs *Source) getStoryUrl(id int) string {
	return fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
}

func (hs *Source) getTopStories(api fetcher, n int) ([]int, error) {
	stories, err := api.GetStories()
	if err != nil {
		return nil, err
	}
	return stories[:n], nil
}

func (hs *Source) getArticle(id int) (sources.Article, error) {
	var story map[string]any
	storyURL := hs.getStoryUrl(id)

	response, err := http.Get(storyURL)
	if err != nil {
		log.Println(err)
		return sources.Article{}, err
	}
	err = json.NewDecoder(response.Body).Decode(&story)
	if err != nil {
		log.Println(err)
		return sources.Article{}, err
	}
	article := sources.Article{
		Title:  story["title"].(string),
		Url:    story["url"].(string),
		Source: SourceName,
	}
	return article, nil
}

func (hs *Source) GetArticles(n int) ([]sources.Article, error) {
	var articles []sources.Article
	api := new(hackernewsAPI)
	topStories, err := hs.getTopStories(api, n)

	if err != nil {
		return nil, err
	}
	for _, storyID := range topStories {
		article, err := hs.getArticle(storyID)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}
