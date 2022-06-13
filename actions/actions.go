package actions

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"

	"github.com/sourcegraph/run"
)

type Map map[string]any

func render(tmpl string, data any) (string, error) {
	var buf bytes.Buffer
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}
	err = t.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}
	return buf.String(), nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func envFilename() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("getting current user: %w", err)
	}

	envFilename := filepath.Join(u.HomeDir, "niflheim.env")
	return envFilename, nil
}

// func homeDir() (string, error) {
// 	u, err := user.Current()
// 	if err != nil {
// 		return "", fmt.Errorf("getting current user: %w", err)
// 	}
// 	return u.HomeDir, nil
// }

func stdData() (Map, error) {
	env := os.Environ()
	m := Map{}
	for _, v := range env {
		b, a, found := strings.Cut(v, "=")
		if !found {
			continue
		}
		m[b] = a
	}
	return m, nil
}

func LoadEnv() error {
	exists := func(path string) bool {
		_, err := os.Stat(path)
		return !errors.Is(err, os.ErrNotExist)
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Errorf("error getting current directory: %v", err))
		os.Exit(2)
	}
	if exists(filepath.Join(pwd, "niflheim.env")) {
		if err := godotenv.Load(filepath.Join(pwd, "niflheim.env")); err != nil {
			return fmt.Errorf("error loading %s/niflheim.env file: %w", pwd, err)
		}
	}
	if exists(filepath.Join(os.Getenv("HOME"), "niflheim.env")) {
		if err := godotenv.Load(filepath.Join(os.Getenv("HOME"), "niflheim.env")); err != nil {
			return fmt.Errorf("error loading %s/niflheim.env file: %w", os.Getenv("HOME"), err)
		}
	}
	return nil
}

func runCmd(script string) error {
	ctx := context.Background()

	// Easily stream all output back to standard out
	return run.Cmd(ctx, script).Run().Stream(os.Stdout)
}

func getLatestGithubRelease(user, repo string) (string, string) {
	u := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", user, repo)
	client := resty.New()

	releases := GitHubReleases{}
	resp, err := client.R().
		EnableTrace().
		SetResult(&releases).
		Get(u)

	if err != nil {
		fmt.Println(fmt.Errorf("error getting latest release: %v", err))
		os.Exit(2)
	}

	if resp.StatusCode() != 200 {
		fmt.Println(fmt.Errorf("error getting latest release: %v", resp.StatusCode()))
		os.Exit(2)
	}
	for _, v := range releases.Assets {
		if v.Name == "UnixServer.tar.gz" {
			return v.BrowserDownloadURL, v.Name
		}
	}
	return "XXX", "XXX"
}

type GitHubReleases struct {
	Assets []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}
