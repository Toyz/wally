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
)

func SingleImage(apiKey, id string) (wallpaper Wallpaper, err error) {
	var data Single
	err = getJson(fmt.Sprintf("https://wallhaven.cc/api/v1/w/%s", id), apiKey, &data)
	if err != nil {
		log.Printf("failed to decode json: %s", err)
		err = errors.New("Rate limit has been hit")
		return
	}
	if data.Error != "" {
		err = errors.New("image does not exist")
		return
	}

	wallpaper = data.Data
	return
}

func RandomImage(apiKey, category, purity, resolution string) (wallpaper []Wallpaper, err error) {
	var data Multi
	err = getJson(fmt.Sprintf("https://wallhaven.cc/api/v1/search?sorting=random&categories=%s&purity=%s&seed=%s&resolutions=%s", category, purity, rand.String(6), resolution), apiKey, &data)
	if err != nil {
		log.Printf("failed to decode json: %s", err)
		err = errors.New("Rate limit has been hit")
		return
	}

	if data.Error != "" {
		err = errors.New("image does not exist")
		return
	}

	wallpaper = data.Data
	return
}

func getJson(url, apiKey string, target interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil
	}

	if apiKey != "" {
		req.Header.Add("X-API-Key", apiKey)
	}

	r, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}