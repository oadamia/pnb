package mock

import (
	"pnb/service/store"
	"time"
)

func ListSource() []store.Source {
	return []store.Source{
		{
			Name:        "Source1",
			Url:         "http://source1.com",
			CreatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			ID:          "1",
			Description: "This is Source1",
			Category:    "general",
			Language:    "en",
			Country:     "US",
		},
		{
			Name:        "Source2",
			Url:         "http://source2.com",
			CreatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			ID:          "2",
			Description: "This is Source2",
			Category:    "sport",
			Language:    "ge",
			Country:     "GE",
		},
		{
			Name:        "Source3",
			Url:         "http://source3.com",
			CreatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			ID:          "3",
			Description: "This is Source3",
			Category:    "business",
			Language:    "fr",
			Country:     "FR",
		},
	}

}

func GetSource() store.Source {
	return store.Source{
		Name:        "Source4",
		Url:         "http://source4.com",
		CreatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		ID:          "4",
		Description: "This is Source4",
		Category:    "technology",
		Language:    "en",
		Country:     "US",
	}
}

func CreateSourceParams() store.CreateSourceParams {
	return store.CreateSourceParams{
		ID:          "source4",
		Name:        "Source4",
		Url:         "http://source4.com",
		Description: "This is Source4",
		Category:    "technology",
		Language:    "en",
		Country:     "US",
	}
}

func UpdateSourceParams() store.UpdateSourceParams {
	return store.UpdateSourceParams{
		ID:          "source4",
		Name:        "Source4 Updated",
		Url:         "http://source4updated.com",
		Description: "This is Source4 Updated",
		Category:    "technology",
		Language:    "en",
		Country:     "US",
	}
}
