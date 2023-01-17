/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package pkg

import (
	"fmt"
	"k8s.io/klog/v2"
	"os"
	"strings"
)

type CmdMkdir struct{}

func (mkdir *CmdMkdir) IsExist(parameter []string) bool {
	_, err := os.Stat(strings.Join(parameter, ""))
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}

	return true
}

func (mkdir *CmdMkdir) Exe(parameter []string) {
	if !mkdir.IsExist(parameter) {
		if err := os.MkdirAll(strings.Join(parameter, ""), os.ModePerm); err != nil {
			klog.Errorln(fmt.Sprintf("Create Dir fail, %s", err))
		}
	}
}
