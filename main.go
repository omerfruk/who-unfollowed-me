package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type followersResponse struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

func main() {
	var userName string
	var githubApiUrl = "https://api.github.com/users/"
	fmt.Print("Please enter your name: ")
	fmt.Scanf("%s", &userName)
	fmt.Println("Hello", userName)
	var response []followersResponse

	// For Followers List
	followers := make(map[string]string)
	for i := 0; ; i++ {
		resp, err := http.Get(githubApiUrl + userName + "/followers?page=" + strconv.Itoa(i) + "&per_page=100")
		if err != nil {
			fmt.Println(err)
			fmt.Println("Sorry the program will close.")
			return
		}
		if resp.StatusCode != 200 {
			fmt.Println("Status Code:", resp.StatusCode)
			fmt.Println("Sorry the program will close.")
			return
		}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Sorry the program will close.")
			return
		}
		if len(response) == 0 {
			break
		}
		for _, follower := range response {
			followers[follower.Login] = follower.HtmlUrl
		}
	}
	// For Following List
	following := make(map[string]string)
	for i := 0; ; i++ {
		resp, err := http.Get(githubApiUrl + userName + "/following?page=" + strconv.Itoa(i) + "&per_page=100")
		if err != nil {
			fmt.Println(err)
			fmt.Println("Sorry the program will close.")
			return
		}
		if resp.StatusCode != 200 {
			fmt.Println("Status Code:", resp.StatusCode)
			fmt.Println("Sorry the program will close.")
			return
		}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Sorry the program will close.")
			return
		}
		if len(response) == 0 {
			break
		}
		for _, follower := range response {
			following[follower.Login] = follower.HtmlUrl
		}
	}
	// For Save to File
	fmt.Print("Do you save the result to a file? (y/N): ")
	var save string
	fmt.Scanf("%s", &save)
	if save == "y" {
		file, err := os.Create("result.txt")
		if err != nil {
			fmt.Println(err)
			fmt.Println("Sorry the program will close.")
			return
		}
		defer file.Close()
		writeString := "Followers:"
		for key, value := range followers {
			writeString += "\n" + key + " - " + value
		}
		writeString += "\n\nFollowing:"
		for key, value := range following {
			writeString += "\n" + key + " - " + value
		}
		file.WriteString(writeString)
	}

	// For who is not following you back
	fmt.Println("Who is not following you back?")
	writeString := "<!DOCTYPE html>\n<html>\n<body>\n\n<h1>They are don't following you back</h1>"
	for name, url := range following {
		if _, ok := followers[name]; !ok {
			fmt.Println(name, "is not following you back. User Name: "+name+" URL:", url)
			writeString += "\n" + "<a href='" + url + "' target='_blank'>" + name + "</a> </br>"
		}
	}
	writeString += "\n\n</body>\n</html>"
	file, err := os.Create("unfollowers.html")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Sorry the program will close.")
		return
	}
	defer file.Close()
	file.WriteString(writeString)
}
