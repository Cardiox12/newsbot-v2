package sources

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Article struct {
	Title  string
	Url    string
	Source string
}

func (a Article) GetHash() string {
	text := fmt.Sprintf("%s+%s+%s", a.Title, a.Url, a.Source)
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
