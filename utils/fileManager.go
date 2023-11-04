package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var initilized = false
var files = make(map[string]*File)

func GetFile(fileName string, gName string) *File {
	if !Init(gName) {
		return nil
	}
	return files[fileName]
}

func GetAllFiles(gName string) []*File {
	if !Init(gName) {
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

func Init(s string) bool {
	if initilized {
		return true
	}

	initilized = true
	resp, err := http.Get("https://api.github.com/users/" + s + "/repos")

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
			fileUrl = "https://raw.githubusercontent.com/" + s + "/.gitgitgo/" + repo["default_branch"].(string) + "/files.json"
			break
		}
	}
	if fileUrl == "" {
		fmt.Println("No .gitgitgo repo found")
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
