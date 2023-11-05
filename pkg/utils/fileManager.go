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

func GetFile(fileName string) *File {
	if !Init() {
		return nil
	}
	return files[fileName]
}

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

func AddFile(file *File) {
	files[file.Name] = file
}

func Init() bool {
	if initilized {
		return true
	}

	initilized = true
	resp, err := http.Get("https://api.github.com/users/" + OWNER + "/repos")

	check(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	check(err)

	var repos []map[string]interface{}
	err = json.Unmarshal(body, &repos)
	check(err)

	fileUrl := ""
	for _, repo := range repos {
		name := repo["name"].(string)
		if strings.ToLower(name) == ".gitgitgo" {
			fileUrl = "https://raw.githubusercontent.com/" + OWNER + "/.gitgitgo/" + repo["default_branch"].(string) + "/files.json"
			break
		}
	}
	if fileUrl == "" {
		logger.AddError("No .gitgitgo repo found for owner '" + OWNER + "'")
		return false
	}

	resp, err = http.Get(fileUrl)
	check(err)

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	check(err)

	var content map[string]interface{}
	err = json.Unmarshal(body, &content)
	check(err)

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

		AddFile(fileObj)

	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
