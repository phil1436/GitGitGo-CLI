package utils

import (
	"fmt"
	"os"
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
	text = strings.ReplaceAll(text, "${YEAR}", string(rune(time.Now().Year())))
	text = strings.ReplaceAll(text, "${MONTH}", string(rune(time.Now().Month())))
	text = strings.ReplaceAll(text, "${DAY}", string(rune(time.Now().Day())))
	text = strings.ReplaceAll(text, "${GITHUBNAME}", githubName)
	text = strings.ReplaceAll(text, "${FULLNAME}", fullname)
	text = strings.ReplaceAll(text, "${REPONAME}", repoName)
	return text

}

func (t *File) Save(basePath string, githubName string, fullname string, repoName string) bool {
	text := t.ParseText(githubName, fullname, repoName)

	url := basePath + "/" + t.Name
	_, err := os.Stat(url)

	if err == nil {
		return false
	}
	// create file
	f, err := os.Create(url)
	if err != nil {
		return false
	}
	defer f.Close()
	_, err = f.WriteString(text)
	return err == nil
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
	result := "File: " + t.Name + "\n"
	if onlyNames {
		return result
	}
	result += "Ignore: " + fmt.Sprintf("%v", t.Ignore) + "\n"
	result += "OnInit: " + fmt.Sprintf("%v", t.OnInit) + "\n"
	result += "Text: \n" + t.Text + "\n"
	return result
}
