package mocks

type HackernewsMockFetcher struct {
	StoriesID []int
}

func (hsm *HackernewsMockFetcher) GetStories() ([]int, error) {
	return hsm.StoriesID, nil
}
