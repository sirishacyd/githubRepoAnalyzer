package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Release struct {
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
}

func getLatestRelease(repo string) (*Release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "gittoken")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Error fetching data for %s. StatusCode: %d, Response: %s", repo, resp.StatusCode, string(bodyBytes))
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

func getNextVersion(tag string) string {
	parts := strings.Split(tag, ".")
	if len(parts) < 3 {
		return tag // Return unchanged if not following x.y.z format
	}
	patchVersion, err := strconv.Atoi(parts[2])
	if err != nil {
		return tag // Return unchanged if error occurs
	}
	parts[2] = strconv.Itoa(patchVersion + 1)
	return strings.Join(parts, ".")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program_name repo1 repo2 ...")
		return
	}

	repos := os.Args[1:]

	for _, repo := range repos {
		release, err := getLatestRelease(repo)
		if err != nil {
			fmt.Printf("Error fetching release for %s: %s\n", repo, err)
			continue
		}

		if release.TagName == "" {
			fmt.Printf("No release found for %s\n", repo)
			continue
		}

		yamlContent := fmt.Sprintf(`---
Repository: %s
LatestReleaseVersion: %s
Changes: |
%s
PredictedNextVersion: %s
`, repo, release.TagName, strings.ReplaceAll(release.Body, "\n", "\n  "), getNextVersion(release.TagName))

		filename := strings.ReplaceAll(repo, "/", "_") + ".yaml"
		err = ioutil.WriteFile(filename, []byte(yamlContent), 0644)
		if err != nil {
			fmt.Printf("Error writing to file %s: %s\n", filename, err)
			continue
		}

		fmt.Println("Generated:", filename)
	}
}
