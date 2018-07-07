package persist

import (
	"testing"
	"joewt.com/joe/learngo/crawler/model"
)

func TestSave(t *testing.T) {


	profiles := model.Profile{
		Age: 123,
		Height: 111,
		Weight: 111,
		Income: "123-123",
	}

	save(profiles)
}
