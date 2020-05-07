package wallhaven

import "testing"

const (
	imageID = "2kroqg"
)

func TestSingleImage(t *testing.T) {
	wallpaper, err := SingleImage(imageID, "")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if wallpaper.ID  != imageID {
		t.Error("image id didn't match" + imageID)
		t.Fail()
	}
}
