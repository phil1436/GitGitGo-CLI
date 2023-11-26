package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"strings"
	"time"
)

type File struct {
	Name        string
	Ignore      bool
	OnInit      bool
	Text        string
	Description string
	Keywords    []string
}

func NewFile(name string, text string, ignore bool, onInit bool, description string, keywords string) *File {

	return &File{
		Name:        name,
		Ignore:      ignore,
		OnInit:      onInit,
		Text:        text,
		Description: description,
		Keywords:    strings.Split(keywords, ","),
	}
}

func NewEmptyFile(name string) *File {
	return &File{
		Name:        name,
		Ignore:      false,
		OnInit:      false,
		Text:        "",
		Description: "",
		Keywords:    []string{},
	}
}

func (t *File) SetKeywords(keyword string) {
	t.Keywords = strings.Split(keyword, ",")
}

func (t *File) ParseText(githubName string, fullname string, repoName string) string {
	text := t.Text
	text = strings.ReplaceAll(text, "${YEAR}", fmt.Sprintf("%d", time.Now().Year()))
	text = strings.ReplaceAll(text, "${MONTH}", fmt.Sprintf("%d", time.Now().Month()))
	text = strings.ReplaceAll(text, "${DAY}", fmt.Sprintf("%d", time.Now().Day()))
	text = strings.ReplaceAll(text, "${GITHUBNAME}", githubName)
	text = strings.ReplaceAll(text, "${FULLNAME}", fullname)
	text = strings.ReplaceAll(text, "${REPONAME}", repoName)
	return text

}

func (t *File) ContainsKeyword(keyword string) bool {
	for _, key := range t.Keywords {
		if strings.EqualFold(key, keyword) {
			return true
		}
	}
	return false
}

func (t *File) ContainsKeywordArr(keywords []string) bool {
	for _, key := range keywords {
		if t.ContainsKeyword(key) {
			return true
		}
	}
	return false
}

func (t *File) Save(basePath string, githubName string, fullname string, repoName string, as string, force bool, dryrun bool) bool {
	name := t.Name
	if as != "" {
		name = as
	}
	text := t.ParseText(githubName, fullname, repoName)

	url := filepath.Join(basePath, name)
	_, err := os.Stat(url)

	if err == nil && !force {
		logger.Log("File '" + name + "' already exists...")
		return true
	}
	if dryrun {
		logger.Log("File '" + name + "' created!")
		return true
	}

	// create all directories
	err = os.MkdirAll(filepath.Dir(url), os.ModePerm)
	if err != nil {
		logger.AddErrObj("Error while creating directories", err)
		return false
	}
	// create file
	f, err := os.Create(url)
	if err != nil {
		logger.AddErrObj("Error while creating file", err)
		return false
	}
	defer f.Close()
	_, err = f.WriteString(text)
	if err != nil {
		logger.AddErrObj("Error while writing to file", err)
		return false
	}
	logger.Log("File '" + name + "' created!")
	return true
}

func (t *File) ImportTextArr(arr []string) {
	if arr == nil {
		return
	}
	newText := ""
	for _, line := range arr {
		newText += line + "\n"
	}
	t.Text = newText
}
func (t *File) ImportTextArrI(arr []interface{}) {
	if arr == nil {
		return
	}
	newText := ""
	for _, line := range arr {
		newText += line.(string) + "\n"
	}
	t.Text = newText
}

func (t *File) ToString(onlyNames bool) string {
	if onlyNames {
		return t.Name
	}
	result := "File: " + t.Name + "\n"
	result += "Ignore: " + fmt.Sprintf("%v", t.Ignore) + "\n"
	result += "OnInit: " + fmt.Sprintf("%v", t.OnInit) + "\n"
	if t.Description != "" {
		result += "Description: " + t.Description + "\n"
	}
	if len(t.Keywords) > 0 {
		result += "Keywords: "
		for _, keyword := range t.Keywords {
			result += keyword + ", "
		}
		result += "\n"
	}
	result += "Text: \n" + t.Text + "\n"
	return result
}
