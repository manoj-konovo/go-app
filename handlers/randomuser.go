package handlers

import (
	"net/http"

	"github.com/integrationninjas/go-app/models"
)

func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	userData := getStaticUserData()
	encodeJSON(w, userData.Results[0]) // Encode and return the first user data
}

func getStaticUserData() models.UserData {
	return models.UserData{
		Results: []models.User{
			{
				Gender: "male",
				Name: models.Name{
					Title: "Mr",
					First: "John",
					Last:  "Doe",
				},
				Location: models.Location{
					Street: struct {
						Number int    `json:"number"`
						Name   string `json:"name"`
					}{
						Number: 123,
						Name:   "Main Street",
					},
					City:     "Springfield",
					State:    "Illinois",
					Country:  "United States",
					Postcode: 62701,
					Coordinates: struct {
						Latitude  string `json:"latitude"`
						Longitude string `json:"longitude"`
					}{
						Latitude:  "39.7817",
						Longitude: "-89.6501",
					},
					Timezone: struct {
						Offset      string `json:"offset"`
						Description string `json:"description"`
					}{
						Offset:      "-6:00",
						Description: "Central Time (US & Canada)",
					},
				},
				Email: "john.doe@example.com",
				Login: models.Login{
					UUID:     "12345678-1234-1234-1234-123456789012",
					Username: "johndoe123",
					Password: "password123",
					Salt:     "salt123",
					Md5:      "md5hash",
					Sha1:     "sha1hash",
					Sha256:   "sha256hash",
				},
				Dob: models.Dob{
					Date: "1990-01-15T08:00:00.000Z",
					Age:  36,
				},
				Registered: models.Registered{
					Date: "2015-03-20T12:30:00.000Z",
					Age:  10,
				},
				Phone: "(555) 123-4567",
				Cell:  "(555) 987-6543",
				ID: models.ID{
					Name:  "SSN",
					Value: "123-45-6789",
				},
				Picture: models.Picture{
					Large:     "https://via.placeholder.com/512x512.jpg",
					Medium:    "https://via.placeholder.com/256x256.jpg",
					Thumbnail: "https://via.placeholder.com/96x96.jpg",
				},
				Nat: "US",
			},
		},
		Info: struct {
			Seed    string `json:"seed"`
			Results int    `json:"results"`
			Page    int    `json:"page"`
			Version string `json:"version"`
		}{
			Seed:    "static-seed",
			Results: 1,
			Page:    1,
			Version: "1.0",
		},
	}
}
