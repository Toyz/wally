package wallhaven

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/toyz/wally/tools/rand"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func SingleImage(id string) (wallpaper Wallpaper, err error) {
	var data Single
	err = getJson(fmt.Sprintf("https://wallhaven.cc/api/v1/w/%s", id), &data)
	if data.Error != "" {
		err = errors.New("image does not exist")
		return
	}

	wallpaper = data.Data
	return
}

func RandomImage(category string, resolution string) (wallpaper []Wallpaper, err error) {
	var data Multi
	err = getJson(fmt.Sprintf("https://wallhaven.cc/api/v1/search?sorting=random&categories=%s&purity=110&seed=%s&resolutions=%s", category, rand.String(6), resolution), &data)
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