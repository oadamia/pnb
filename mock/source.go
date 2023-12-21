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
