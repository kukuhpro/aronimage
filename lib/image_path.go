package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ImagePath struct {
	PathFolder string
	ListOfPath map[string]Path
}

func (i *ImagePath) LoopKeys(closure func(key string)) {
	for key, _ := range i.ListOfPath {
		closure(key)
	}
}

func (i *ImagePath) GetPath(key string) Path {
	return i.ListOfPath[key]
}

func (i *ImagePath) readFile(fileName string) {
	raw, err := ioutil.ReadFile(i.PathFolder + "/" + fileName)

	if err != nil {
		log.Fatalln(err)
	}
	var data Path
	json.Unmarshal(raw, &data)

	i.ListOfPath[data.Name] = data
}

func (i *ImagePath) handleFiles() {
	files, err := ioutil.ReadDir(i.PathFolder)

	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		i.readFile(f.Name())
	}
}

func NewImagePath() *ImagePath {
	var imagePath ImagePath

	imagePath.ListOfPath = map[string]Path{}
	imagePath.PathFolder = "./path"
	imagePath.handleFiles()

	return &imagePath
}
