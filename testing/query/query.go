package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ReleaseInfo ...
type ReleaseInfo struct {
	ID      uint   `json:"id"`
	TagName string `json:"tag_name"`
}

// ReleaseInfoer ...
type ReleaseInfoer interface {
	GetLatestReleaseTag(string) (string, error)
}

// GithubReleaseInfo ...
type GithubReleaseInfo struct{}

// GetLatestReleaseTag function to actually query the Github API for the release information
func (gh GithubReleaseInfo) GetLatestReleaseTag(repo string) (string, error) {
	api := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	res, err := http.Get(api)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	releases := []ReleaseInfo{}
	if err := json.Unmarshal(body, &releases); err != nil {
		return "", err
	}

	return releases[0].TagName, nil
}

// Function to get the message to display to the end user
func getReleaseTagMessage(ri ReleaseInfoer, repo string) (string, error) {
	tag, err := ri.GetLatestReleaseTag(repo)
	if err != nil {
		return "", fmt.Errorf("Error querying Github API: %s", err)
	}
	return fmt.Sprintf("The latest release is %q", tag), nil
}

func main() {
	gh := GithubReleaseInfo{}
	msg, err := getReleaseTagMessage(gh, "docker/machine")
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
	}

	fmt.Println(msg)
}

/*
  Example of the bad version
*/

//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// )
//
// // ReleaseInfo information of the release
// type ReleaseInfo struct {
// 	ID      uint   `json:"id"`
// 	TagName string `json:"tag_name"`
// }
//
// // Function to actually query the Github API for the release information
// func getLatestReleaseTag(repo string) (string, error) {
// 	api := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
// 	res, err := http.Get(api)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer res.Body.Close()
//
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return "", nil
// 	}
//
// 	releases := []ReleaseInfo{}
// 	if err := json.Unmarshal(body, &releases); err != nil {
// 		return "", err
// 	}
//
// 	return releases[0].TagName, nil
// }
//
// // Function to get the message to display to the end user
// func getReleaseTagMessage(repo string) (string, error) {
// 	tag, err := getLatestReleaseTag(repo)
// 	if err != nil {
// 		return "", fmt.Errorf("Error querying Github API: %s", err)
// 	}
// 	return fmt.Sprintf("The latest release is %q", tag), nil
// }
//
// func main() {
// 	msg, err := getReleaseTagMessage("docker/machine")
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, msg)
// 	}
//
// 	fmt.Println(msg)
// }
