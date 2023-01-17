/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package pkg

import (
	"fmt"
	"io/ioutil"
	"k8s.io/klog/v2"
	"os"
	"strings"
)

type CmdCat struct{}

func (cat *CmdCat) IsDir(parameter []string) bool {
	fileOrDir, err := os.Stat(parameter[0])
	if err != nil {
		klog.Errorln(fmt.Sprintf("Get fileOrDir message fail, %s", err))
		return false
	}

	return fileOrDir.IsDir()
}

func (cat *CmdCat) Exe(parameter []string) {
	if !cat.IsDir(parameter) {
		file, err := ioutil.ReadFile(strings.Join(parameter, ""))
		if err != nil {
			klog.Errorln(fmt.Sprintf("Read file fail, %s", err))
		}

		fmt.Println(string(file))
	}
}

type CatFactory struct {
	AbstractFactory
}

func (cat *CatFactory) Option() CommandOptions {
	return new(CmdCat)
}
