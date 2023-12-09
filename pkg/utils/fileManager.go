package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"strings"
)

var initilized = false
var files = make(map[string]*File)

// Get a file by name if it does not exist return nil
func GetFile(fileName string) *File {
	if !Init() {
		return nil
	}
	return files[fileName]
}

// Get all files
func GetAllFiles() []*File {
	if !Init() {
		return nil
	}
	var filesArr []*File
	for _, file := range files {
		filesArr = append(filesArr, file)
	}
	return filesArr
}

// Add a file to the file manager
func AddFile(file *File) {
	files[file.Name] = file
}

// Initialize the file manager
func Init() bool {
	if initilized {
		return true
	}

	initilized = true
	resp, err := http.Get("https://api.github.com/users/" + PROVIDER + "/repos")

	Check(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	Check(err)

	var repos []map[string]interface{}
	err = json.Unmarshal(body, &repos)
	Check(err)

	fileUrl := ""
	for _, repo := range repos {
		name := repo["name"].(string)
		if strings.ToLower(name) == ".gitgitgo" {
			fileUrl = "https://raw.githubusercontent.com/" + PROVIDER + "/.gitgitgo/" + repo["default_branch"].(string) + "/files.json"
			break
		}
	}
	if fileUrl == "" {
		logger.AddError("No .gitgitgo repo found for provider '" + PROVIDER + "'")
		return false
	}

	resp, err = http.Get(fileUrl)
	Check(err)

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	Check(err)

	var content map[string]interface{}
	err = json.Unmarshal(body, &content)
	Check(err)

	files := content["files"].([]interface{})

	for _, fileI := range files {
		file := fileI.(map[string]interface{})

		fileObj := NewEmptyFile(file["name"].(string))
		fileObj.ImportTextArrI(file["text"].([]interface{}))
		if file["gitignore"] != nil {
			fileObj.Ignore = file["gitignore"].(bool)
		}
		if file["oninit"] != nil {
			fileObj.OnInit = file["oninit"].(bool)
		}
		if file["description"] != nil {
			fileObj.Description = file["description"].(string)
		}
		if file["keywords"] != nil {
			fileObj.SetKeywords(file["keywords"].(string))
		}

		AddFile(fileObj)

	}
	return true
}

// Reload the file manager
func ReloadFileManager() bool {
	initilized = false
	return Init()
}

// Check if there is an error and log it
func Check(e error) {
	if e != nil {
		logger.AddErrObj("Something went wrong", e)
		initilized = false
	}
}
