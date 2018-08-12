package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ImagePath struct {
	PathFolder string
	ListOfPath map[string]Path
}

/**
 * Loop Keys for functional function
 * @param closure func(key string)
 * @return void
 */
func (i *ImagePath) LoopKeys(closure func(key string)) {
	for key, _ := range i.ListOfPath {
		closure(key)
	}
}

/**
 * Get return path
 * @param key string
 * @return Path
 */
func (i *ImagePath) GetPath(key string) Path {
	return i.ListOfPath[key]
}

/**
 * Process read file from JSON file, to generate path upload image.
 * @param fileName string
 * @return void
 */
func (i *ImagePath) readFile(fileName string) {
	raw, err := ioutil.ReadFile(i.PathFolder + "/" + fileName)

	if err != nil {
		log.Fatalln(err)
	}
	var data Path
	json.Unmarshal(raw, &data)

	i.ListOfPath[data.Name] = data
}

/**
 * Handle to read all files on specific folder.
 * @return void
 */
func (i *ImagePath) handleFiles() {
	files, err := ioutil.ReadDir(i.PathFolder)

	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		i.readFile(f.Name())
	}
}

/**
 * Constructor function for handle image path
 * @return *ImagePath
 */
func NewImagePath() *ImagePath {
	var imagePath ImagePath

	pwd, _ := os.Getwd()

	if strings.Contains(pwd, "/tests") {
		pwd = strings.Replace(pwd, "/tests", "", -1)
	}
	imagePath.ListOfPath = map[string]Path{}
	imagePath.PathFolder = pwd + "/path"
	imagePath.handleFiles()

	return &imagePath
}
