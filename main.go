package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// var config []Season = []Season{
// 	{1, 2},
// 	{3, 3},
// 	{4, 5},
// }

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("No directory provided!")
		return
	}
	path := args[0] + "/"

	b, err := ioutil.ReadFile(args[0] + ".json")
	if err != nil {
		panic(err)
	}

	var config []Season
	if err := json.Unmarshal(b, &config); err != nil {
		panic(err)
	}
	// check config
	for _, season := range config {
		if season.Beginning > season.End {
			panic("Configuration error!")
		}
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	if len(files) < config[len(config)-1].End {
		fmt.Println(len(files), ":", config[len(config)-1].End)
		panic("Configuration exceeds number of files!")
	}

	for _, f := range files {
		if f.IsDir() {
			panic("Directory CANNOT contain sub directories")
		}
	}

	for i, season := range config {
		seasonPath := path + fmt.Sprintf("Season %d", i+1) + "/"
		os.MkdirAll(seasonPath, 0755)
		subFiles := files[(season.Beginning - 1):season.End]
		for j, f := range subFiles {
			if f.IsDir() {
				panic("Directory CANNOT contain sub directories")
			}
			ext := filepath.Ext(f.Name())
			os.Rename(path+f.Name(), fmt.Sprintf("%sS%.2dE%.2d%s", seasonPath, i+1, j+1, ext))
		}
	}
}

type Season struct {
	Beginning int `json:"beginning"`
	End       int `json:"end"`
}
