package sources

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnArticleHashWhenCreatingArticleTitle1Url1Source1(t *testing.T) {
	article := Article{Title: "title1", Url: "url1", Source: "source1"}

	hash := article.GetHash()

	wantedHash := "dabea3c8c7b4bf4b5f204a5e8fefb2ab"
	assert.Equal(t, wantedHash, hash)
}

func TestShouldReturnArticleHashWhenCreatingArticleTitle2Url2Source2(t *testing.T) {
	article := Article{Title: "title2", Url: "url2", Source: "source2"}

	hash := article.GetHash()

	wantedHash := "a5f87a6ef9621fe3678fa539143e221a"
	assert.Equal(t, wantedHash, hash)
}
