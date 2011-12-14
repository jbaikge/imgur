package imgur

import (
	"fmt"
	"http"
	"io"
	"io/ioutil"
	"json"
	"os"
	"strings"
	"url"
)


type ImageInfo struct {
	Downs int
	Title string
	Rating float64
	Views int
	Ups int
}

type GalleryInfo struct {
	Image ImageInfo
}

type Message struct {
	Gallery GalleryInfo
}

func Load(r io.Reader) (*ImageInfo, os.Error) {
	result := new(Message)
	jsonBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(jsonBytes, &result)

	// Calculate the rating
	i := result.Gallery.Image
	result.Gallery.Image.Rating = float64(i.Ups) / float64(i.Ups + i.Downs) * 100

	return &result.Gallery.Image, nil
}

func ValidUrl(incoming string) (bool, string) {
	// Force prefix with http:// 
	if !strings.HasPrefix(incoming, "http://") {
		incoming = "http://" + incoming
	}

	// Parse URL
	u, err := url.Parse(incoming)
	if err != nil {
		return false, ""
	}

	// Check host
	path := u.Path
	if u.Host == "i.imgur.com" {
		path = strings.Split(path, ".")[0]
	}
	if strings.HasSuffix(u.Host, "imgur.com") {
		bits := strings.Split(path, "/")
		hash := bits[len(bits) - 1]
		return true, hash
	}
	return false, ""
}

func HashInfo(hash string) (*ImageInfo, os.Error) {
	path := fmt.Sprintf("http://imgur.com/gallery/%s.json", hash)

	response, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return Load(response.Body)
}
