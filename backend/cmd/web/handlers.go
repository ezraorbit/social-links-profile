package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/ezraorbit/social-links-profile/internal/githubapi"
	"github.com/ezraorbit/social-links-profile/internal/models"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/index.tmpl",
	}

	query := "followers:>100"

	// Call the function to search for GitHub users
	user, err := githubapi.SearchUsers(query)
	if err != nil {
		log.Fatalf("Error fetching GitHub users: %v", err)
	}

	socialAccounts, err := githubapi.GetSocialAccounts(user.Login)
	if err != nil {
		http.Error(w, "Error fetching social accounts", http.StatusInternalServerError)
		return
	}

	userData := models.UserProfile{
		Name:           user.Name,
		Location:       user.Location,
		AvatarURL:      user.AvatarURL,
		Bio:            user.Bio,
		GithubURL:      user.GithubURL,
		SocialAccounts: socialAccounts,
	}

	// Print the users
	fmt.Printf("User: %s, Profile: %s, Social Media: %s\n", user.Name, user.Location, user.SocialAccounts)

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}

	err = ts.ExecuteTemplate(w, "base", userData)
	if err != nil {
		log.Fatal(err)
	}

}
