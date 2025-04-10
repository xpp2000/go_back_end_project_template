package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string
var iJsonData map[string]any

const stJsonFileName = "dir.json"

// step1. read json
func TestGenerateDir01(t *testing.T) {
	fmt.Println("testset")
	loadJson()
	parseMap(iJsonData, "")
}

func loadJson() {
	stSeparator = string(filepath.Separator)
	stWorkDir, _ := os.Getwd()
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]

	fmt.Println("WorkDir is: " + stWorkDir)
	fmt.Println("RootDir is: " + stRootDir)

	gnJsonBytes, _ := os.ReadFile(stWorkDir + stSeparator + stJsonFileName)
	err := json.Unmarshal(gnJsonBytes, &iJsonData)
	if err != nil {
		panic("Load Json Data Error: " + err.Error())
	}

}
func parseMap(mapData map[string]any, stParentDir string) {
	for k, v := range mapData {
		switch v.(type) {
		case string:
			{
				path, _ := v.(string)
				if path == "" {
					continue
				}
				if stParentDir != "" {
					path = stParentDir + stSeparator + path
					if k == "text" {
						stParentDir = path
					}
				} else {
					stParentDir = path
				}
				createDir(path)

			}
		case []any:
			{
				parseArray(v.([]any), stParentDir)
			}
		}
	}
}
func parseArray(giJsonData []any, stParentDir string) {
	for _, v := range giJsonData {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, stParentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}
	fmt.Println("creating path :" + path)
	err := os.MkdirAll(stRootDir+stSeparator+path, fs.ModePerm)
	if err != nil {
		panic("Create Dir Error: " + err.Error())
	}
}
