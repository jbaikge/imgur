package imgur

import (
	"testing"
)

type sample struct {
	Url string
	Valid bool
	Hash string
	Title string
}

var sampleTests = []sample {
	sample{
		"http://imgur.com/gallery/jZv4f",
		true,
		"jZv4f",
		"So at the Lego store you can make your own lego guys. I made Ron Swanson",
	},
	sample{
		"http://i.imgur.com/jZv4f.jpg",
		true,
		"jZv4f",
		"So at the Lego store you can make your own lego guys. I made Ron Swanson",
	},
	sample{
		"http://imgur.com/jZv4f",
		true,
		"jZv4f",
		"So at the Lego store you can make your own lego guys. I made Ron Swanson",
	},
}

func TestValidUrl(t *testing.T) {
	for _, s := range sampleTests {
		valid, hash := ValidUrl(s.Url)
		if valid != s.Valid {
			t.Errorf("Invalid URL: %s", s.Url)
		}
		if hash != s.Hash {
			t.Errorf("Hashes do not match: %s != %s", hash, s.Hash)
		}
	}
}

func TestHashInfo(t *testing.T) {
	for _, s := range sampleTests {
		valid, hash := ValidUrl(s.Url)
		if valid {
			info, err := HashInfo(hash)
			if err != nil {
				t.Errorf("Error getting info: %s", err.String())
			}
			if info.Title != s.Title {
				t.Errorf("Invalid title, expected '%s', got '%s'", s.Title, info.Title)
			}
		}
	}
}
