package wallhaven

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/toyz/wally/tools/rand"
)

var (
	myClient = &http.Client{Timeout: 10 * time.Second}
	ProxyURL = "https://wallhaven.cc/api/v1/"
)

func SetProxyURL(key *string) {
	ProxyURL = *key
}

func SingleImage(id string) (wallpaper Wallpaper, err error) {
	var data Single
	err = getJson(fmt.Sprintf("w/%s", id), &data)
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

func RandomImage(category, purity, resolution string) (wallpaper []Wallpaper, err error) {
	var data Multi
	err = getJson(fmt.Sprintf("search?sorting=random&categories=%s&purity=%s&seed=%s&resolutions=%s", category, purity, rand.String(6), resolution), &data)
	if err != nil {
		log.Printf("failed to decode json: %s", err)
		err = errors.New("Rate limit has been hit")
		return
	}

	if data.Error != "" {
		err = fmt.Errorf("Error: %s", data.Error)
		return
	}

	wallpaper = data.Data
	return
}

func getJson(url string, target interface{}) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", ProxyURL, url), nil)
	if err != nil {
		return nil
	}

	r, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
