package hackernews

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"publishers/internal/sources/mocks"
	"testing"
)

func TestShouldReturnInstanceWithHackernewsSourceWhenCreated(t *testing.T) {
	hnSource := NewHackernewsSource()

	assert.Equal(t, SourceName, hnSource.Source)
}

func TestStoryUrlShould(t *testing.T) {
	tests := []struct {
		given int
		want  string
	}{
		{0, "https://hacker-news.firebaseio.com/v0/item/0.json"},
		{2, "https://hacker-news.firebaseio.com/v0/item/2.json"},
		{42, "https://hacker-news.firebaseio.com/v0/item/42.json"},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("ReturnUrlWithId%dWhenGivenId%d", test.given, test.given)

		t.Run(testName, func(t *testing.T) {
			hnSource := NewHackernewsSource()

			url := hnSource.getStoryUrl(test.given)

			assert.Equal(t, test.want, url)
		})
	}
}

func TestGettingTopStories(t *testing.T) {
	tests := []struct {
		n            int
		givenStories []int
		want         []int
	}{
		{n: 3, givenStories: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3}},
		{n: 5, givenStories: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("ShouldReturn%dTopStoriesWhenGetting%dTopStories", test.n, test.n)

		t.Run(testName, func(t *testing.T) {
			hnSource := NewHackernewsSource()
			hnMockFetcher := &mocks.HackernewsMockFetcher{StoriesID: test.givenStories}

			topStories, _ := hnSource.getTopStories(hnMockFetcher, test.n)

			assert.Equal(t, test.want, topStories)
		})
	}
}
