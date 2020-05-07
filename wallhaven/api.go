package wallhaven

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/toyz/wally/tools/rand"
	"net/http"
	"time"
)

var (
	myClient = &http.Client{Timeout: 10 * time.Second}
	apiKey string
)

func SetAPIKey(key string) {
	apiKey = key
}

func SingleImage(id string) (wallpaper Wallpaper, err error) {
	var data Single
	err = getJson(fmt.Sprintf("https://wallhaven.cc/api/v1/w/%s?apikey=%s", id, apiKey), &data)
	if data.Error != "" {
		err = errors.New("image does not exist")
		return
	}

	wallpaper = data.Data
	return
}

func RandomImage(category, purity, resolution string) (wallpaper []Wallpaper, err error) {
	var data Multi
	err = getJson(fmt.Sprintf("https://wallhaven.cc/api/v1/search?sorting=random&categories=%s&purity=%s&seed=%s&resolutions=%s&apikey=%s", category, purity, rand.String(6), resolution, apiKey), &data)
	if data.Error != "" {
		err = errors.New("image does not exist")
		return
	}

	wallpaper = data.Data
	return
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}