package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"phil1436/GitGitGo-CLI/logger"
	"strings"
	"time"
)

type File struct {
	Name   string
	Ignore bool
	OnInit bool
	Text   string
}

func NewFile(name string, text string, ignore bool, onInit bool) *File {

	return &File{
		Name:   name,
		Ignore: ignore,
		OnInit: onInit,
		Text:   text,
	}
}

func NewEmptyFile(name string) *File {
	return &File{
		Name:   name,
		Ignore: false,
		OnInit: false,
		Text:   "",
	}
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
	result += "Text: \n" + t.Text + "\n"
	return result
}
