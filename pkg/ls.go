package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/klog/v2"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	Mode    string
	ModTime time.Time
	IsDir   bool
}

type CmdLs struct{}

func (ls *CmdLs) Exe(parameter []string) {
	files, err := ioutil.ReadDir(parameter[0])
	if err != nil {
		klog.Errorln(fmt.Sprintf("Read Dir Message fail, %s", err))
	}

	var fileDirMsg []FileInfo
	for _, file := range files {
		fileDirMsg = append(fileDirMsg, FileInfo{
			Name:    file.Name(),
			Size:    file.Size(),
			Mode:    file.Mode().String(),
			ModTime: file.ModTime(),
			IsDir:   file.IsDir(),
		})
	}

	jsonFileDirs, err := json.Marshal(fileDirMsg)
	if err != nil {
		klog.Errorln(fmt.Sprintf("fileDirMsg to Json fail, %s", err))
	}

	fmt.Println(string(jsonFileDirs))
	return
}
