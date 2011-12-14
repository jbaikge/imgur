package imgur

import (
	"testing"
)

type sample struct {
	Url string
	Hash string
	Title string
}

var sampleTests = []sample {
	sample{
		"http://imgur.com/gallery/jZv4f",
		"jZv4f",
		"So at the Lego store you can make your own lego guys. I made Ron Swanson",
	},
	sample{
		"http://i.imgur.com/jZv4f.jpg",
		"jZv4f",
		"So at the Lego store you can make your own lego guys. I made Ron Swanson",
	},
	sample{
		"http://imgur.com/jZv4f",
		"jZv4f",
		"So at the Lego store you can make your own lego guys. I made Ron Swanson",
	},
}

func TestParseUrl(t *testing.T) {
	for _, s := range sampleTests {
		hash, err := ParseUrl(s.Url)
		if err != nil {
			t.Errorf("Invalid URL: %s", s.Url)
		}
		if hash != s.Hash {
			t.Errorf("Hashes do not match: %s != %s", hash, s.Hash)
		}
	}
}

func TestHashInfo(t *testing.T) {
	for _, s := range sampleTests {
		hash, err := ParseUrl(s.Url)
		if err != nil {
			t.Errorf("Invalid URL: %s", s.Url)
			continue
		}
		info, hErr := HashInfo(hash)
		if hErr != nil {
			t.Errorf("Error getting info: %s", hErr.String())
		}
		if info.Title != s.Title {
			t.Errorf("Invalid title, expected '%s', got '%s'", s.Title, info.Title)
		}
	}
}
