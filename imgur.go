package imgur

import (
	"fmt"
	"http"
	"io"
	"json"
	"os"
	"path"
	"url"
)

type ImageInfo struct {
	Downs  int
	Title  string
	Rating float64
	Views  int
	Ups    int
}

func (i *ImageInfo) UpdateRating() {
	i.Rating = float64(i.Ups) / float64(i.Ups+i.Downs) * 100
}

func Load(r io.Reader) (*ImageInfo, os.Error) {
	var result struct {
		Gallery struct {
			Image ImageInfo
		}
	}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&result); err != nil {
		return nil, err
	}

	result.Gallery.Image.UpdateRating()
	return &result.Gallery.Image, nil
}

func ParseUrl(incoming string) (hash string, err os.Error) {
	// Parse URL
	u, err := url.Parse(incoming)
	if err != nil {
		return "", err
	}

	if u.Scheme != "http" {
		return "", os.NewError("Incorrect Scheme")
	}

	if u.Host != "i.imgur.com" {
		return "", os.NewError("Incorrect Host")
	}

	hash = getHash(u)
	if hash == "" {
		err = os.NewError("Unable to find hash")
	}

	return
}

func getHash(u *url.URL) string {
	file := path.Base(u.Path)
	return file[:len(file)-len(path.Ext(file))]
}

func HashInfo(hash string) (*ImageInfo, os.Error) {
	path := fmt.Sprintf("http://imgur.com/gallery/%s.json", hash)
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return Load(resp.Body)
}
