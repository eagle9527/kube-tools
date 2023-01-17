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

type CmdRm struct{}

func (rm *CmdRm) Exe(parameter []string) {
	if err := os.RemoveAll(strings.Join(parameter, "")); err != nil {
		klog.Errorln(fmt.Sprintf("Remove file Or dir fail, %s", err))
	}
}

type RmdirFactory struct {
	AbstractFactory
}

func (mv *RmdirFactory) Option() CommandOptions {
	return new(CmdRm)
}
