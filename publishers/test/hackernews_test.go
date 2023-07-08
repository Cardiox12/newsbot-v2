package test

import (
	"github.com/stretchr/testify/assert"
	"publishers/internal/sources"
	"publishers/internal/sources/hackernews"
	"testing"
)

func TestShouldRetrieve2ArticlesFromHackernewsSource(t *testing.T) {
	hnSource := hackernews.NewHackernewsSource()

	articles, err := hnSource.GetArticles(2)

	assert.Nil(t, err)
	assert.Len(t, articles, 2)

	for _, article := range articles {
		assertArticleValid(t, hackernews.SourceName, article)
	}
}

func TestShouldRetrieve4ArticlesFromHackernewsSource(t *testing.T) {
	hnSource := hackernews.NewHackernewsSource()

	articles, err := hnSource.GetArticles(3)

	assert.Nil(t, err)
	assert.Len(t, articles, 3)

	for _, article := range articles {
		assertArticleValid(t, hackernews.SourceName, article)
	}
}

func assertArticleValid(t *testing.T, expectedSourceName string, givenArticle sources.Article) {
	assert.NotEmpty(t, givenArticle.Title)
	assert.NotEmpty(t, givenArticle.Url)
	assert.Equal(t, expectedSourceName, givenArticle.Source)
}
