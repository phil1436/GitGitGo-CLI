package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"runtime"
	"strings"
)

const VERSION = "0.1.0"
const PRE_RELEASE = true

const LATEST_VERSION_URL = "https://api.github.com/repos/phil1436/gitgitgo-cli/releases/latest"

const BINARY_DOWNLOAD_URL_LINUX = "https://github.com/phil1436/GitGitGo-CLI/releases/download/$$version$$/gitgitgo"
const BINARY_DOWNLOAD_URL_WINDOWS = "https://github.com/phil1436/GitGitGo-CLI/releases/download/$$version$$/gitgitgo.exe"

func Update(attValue []interface{}, fs *cmdtool.FlagSet) bool {
	logger.Log("*** GitGitGo UPDATE ***")
	logger.Log(fs.GetStateString())

	homeDir := HOME
	if homeDir == "" {
		homeDir = fs.GetValue("home").(string)
	}
	if homeDir == "" {
		logger.AddError("Cannot find home directory of GitGitGo-CLI run again with -home <path> flag")
		return false
	}

	latestVersion := GetLatestVersion()

	logger.Log("Update to " + latestVersion + "...")
	logger.Log("Home directory: " + homeDir)
	// get os
	myOS := ""
	if runtime.GOOS == "windows" {
		myOS = "windows"
	} else if runtime.GOOS == "linux" {
		myOS = "linux"
	} else if runtime.GOOS == "macos" {
		myOS = "macos"
	} else {
		logger.AddError("Unsupported OS: " + runtime.GOOS)
		return false
	}
	logger.Log("OS: " + myOS)
	logger.Log("Downloading binary...")
	url := BINARY_DOWNLOAD_URL_LINUX
	if myOS == "windows" {
		url = BINARY_DOWNLOAD_URL_WINDOWS
	}

	url = strings.ReplaceAll(url, "$$version$$", latestVersion)

	msg, err := downloadFile(url, homeDir)
	if err != nil {
		logger.AddErrObj("Failed to download binary", err)
		return false
	}
	logger.Log(msg)
	logger.Log("Successfully updated GitGitGo-CLI")
	logger.Log("Run 'gitgitgo version' to check the version")
	logger.Log("Closing...")
	os.Exit(0)

	return true
}

func downloadFile(url string, outputDir string) (string, error) {

	// Extract the file name from the URL
	fileName := filepath.Base(url)

	// Create the output file
	outputPath := filepath.Join(outputDir, fileName)
	file, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Make a GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Check if the request was successful (status code 200)
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download file, status code: %d", response.StatusCode)
	}

	// Copy the binary data to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}
	return "File downloaded and saved to: " + outputPath, nil
}

func CheckForUpdate() {
	latestVersion := GetLatestVersion()
	if latestVersion == "" {
		return
	}
	if latestVersion[0] == 'v' {
		latestVersion = latestVersion[1:]
	}

	if strings.Compare(VERSION, latestVersion) >= 0 {
		return
	}

	logger.Log("New release of GitGitGo-CLI available: " + VERSION + " -> " + latestVersion)
	logger.Log("Run 'gitgitgo update' to update")
}

func GetLatestVersion() string {
	resp, err := http.Get(LATEST_VERSION_URL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var meta map[string]interface{}
	err = json.Unmarshal(body, &meta)

	if err != nil {
		return ""
	}

	version := meta["tag_name"].(string)
	return version
}
