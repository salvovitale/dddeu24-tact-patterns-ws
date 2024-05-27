package infra_repository

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/salvovitale/dddeu24-tact-patterns-ws/internal/domain"
)

const (
	svcUrl     = "https://ddd-in-language.aardling.eu/api/users"
	token      = "emk7srgDuZ"
	workshopID = "ImplementingTacticalPatternsDDDEU24"
)

/*
	{
		"id": "Bald Eagle",
		"type": "private",
		"address": "Point Dume 111",
		"city": "Pineville",
		"card_id": "123"
	}
*/
type UserInRepo struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Address string `json:"address"`
	City    string `json:"city"`
	CardID  string `json:"card_id"`
}

type UserDto struct {
	ID   string
	City domain.City
}

type UserRepository struct{}

func (r *UserRepository) Get(id string) (*UserDto, error) {
	client := &http.Client{
		Timeout: time.Duration(5) * time.Second,
	}

	// Create a new request
	req, err := http.NewRequest("GET", svcUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add custom headers
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-auth-token", token)
	req.Header.Add("x-workshop-id", workshopID)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	var users []UserInRepo

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	var user *UserDto
	for _, u := range users {
		if u.ID == id {
			slog.Info("user found", slog.Any("user", u))
			city, err := domain.ParseCity(u.City)
			if err != nil {
				return nil, err
			}
			user = &UserDto{
				ID:   u.ID,
				City: city,
			}
			break
		}
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
