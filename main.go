package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"
)

var (
	appVersion = "dev"
	buildTime  = ""
)

type metadataObj struct {
	groupID     string
	artifactID  string
	version     string
	timestamp   string
	buildNumber string
}

func main() {
	versionPtr := flag.Bool("v", false, "show version")
	repoPathPtr := flag.String("r", "~/.m2/repository", "maven repo path")
	flag.Parse()
	if *versionPtr {
		fmt.Println("version", appVersion)
		fmt.Println("build time", buildTime)
		os.Exit(0)
	}
	err := handleTargetDirs(*repoPathPtr)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}

func handleTargetDirs(repoPath string) error {
	if strings.HasPrefix(repoPath, "~") {
		u, err := user.Current()
		if err != nil {
			return err
		}
		repoPath = path.Join(u.HomeDir, repoPath[1:])
	}
	count := 0
	err := handleDir(repoPath, &count)
	if err != nil {
		return err
	}
	fmt.Println("==============================================")
	fmt.Printf("Total %d entries cleaned\n", count)
	return nil
}

func handleDir(pDir string, count *int) error {
	files, err := ioutil.ReadDir(pDir)
	if err != nil {
		return err
	}
	metadataFile := ""
	filenames := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			handleDir(path.Join(pDir, f.Name()), count)
		} else {
			if strings.HasPrefix(f.Name(), "maven-metadata") && strings.HasSuffix(f.Name(), ".xml") {
				metadataFile = f.Name()
			} else {
				filenames = append(filenames, f.Name())
			}
		}
	}
	if metadataFile == "" {
		return nil
	}
	mo, err := readMetadataFile(path.Join(pDir, metadataFile))
	if err != nil {
		return err
	}
	if mo.artifactID != "" && strings.HasSuffix(mo.version, "-SNAPSHOT") && mo.timestamp != "" && mo.buildNumber != "" {
		prefixWithoutSnapshot := mo.artifactID + "-" + mo.version[:len(mo.version)-9]
		prefixStr := mo.artifactID + "-" + mo.version
		latestName := prefixWithoutSnapshot + "-" + mo.timestamp + "-" + mo.buildNumber
		historyCount := 0
		for _, ff := range filenames {
			if strings.HasPrefix(ff, prefixWithoutSnapshot) &&
				!strings.HasPrefix(ff, latestName) &&
				!strings.HasPrefix(ff, prefixStr) {
				if strings.HasSuffix(ff, ".jar") {
					historyCount++
				}
				filePath := path.Join(pDir, ff)
				err := os.Remove(filePath)
				if err != nil {
					fmt.Println("Could not remove file:", filePath)
				}
			}
		}
		if historyCount > 0 {
			fmt.Printf("DELETE %s:%s:%s [%d history items]\n", mo.groupID, mo.artifactID, mo.version, historyCount)
			*count++
		}
	}
	return nil
}

func readMetadataFile(filePath string) (metadataObj, error) {
	mo := metadataObj{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return mo, err
	}
	var t xml.Token
	var name string
	decoder := xml.NewDecoder(bytes.NewBuffer(content))
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name = token.Name.Local
		// 处理元素结束（标签）
		case xml.EndElement:
			name = ""
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content := string([]byte(token))
			switch name {
			case "groupId":
				mo.groupID = content
			case "artifactId":
				mo.artifactID = content
			case "version":
				mo.version = content
			case "timestamp":
				mo.timestamp = content
			case "buildNumber":
				mo.buildNumber = content
			}
		}
	}
	return mo, nil
}
