package githubapi

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/ezraorbit/social-links-profile/internal/models"
)

func SearchUsers(query string) (models.UserProfile, error) {
	// Create a new random generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	searchURL := fmt.Sprintf("https://api.github.com/search/users?q=%s", query)

	resp, err := http.Get(searchURL)
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}

	// Decode the JSON response for the user search
	var searchResult models.UserSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return models.UserProfile{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	// Pick a random user from the result
	if len(searchResult.Items) == 0 {
		return models.UserProfile{}, fmt.Errorf("No users found.")
	}

	randomUser := searchResult.Items[rng.Intn(len(searchResult.Items))]

	// Fetch detailed profile information for the selected user
	profileURL := fmt.Sprintf("https://api.github.com/users/%s", randomUser.Login)

	profileResp, err := http.Get(profileURL)
	if err != nil {
		log.Fatalf("Error fetching user profile: %v", err)
	}
	defer profileResp.Body.Close()

	if profileResp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", profileResp.Status)
	}

	// Decode the JSON response for the user profile
	var userProfile models.UserProfile
	if err := json.NewDecoder(profileResp.Body).Decode(&userProfile); err != nil {
		return models.UserProfile{}, fmt.Errorf("Error decoding JSON: %v", err)
	}

	return userProfile, nil
}

func GetSocialAccounts(username string) ([]map[string]string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/social_accounts", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching social accounts: %v", err)
	}
	defer resp.Body.Close()

	var socialAccounts []map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&socialAccounts); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return socialAccounts, nil
}
