/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package pkg

import (
	"fmt"
	"k8s.io/klog/v2"
	"os"
)

type CmdMv struct{}

func (mv *CmdMv) Exe(parameter []string) {
	if err := os.Rename(parameter[0], parameter[1]); err != nil {
		klog.Errorln(fmt.Sprintf("Rename File Or Dir fail, %s", err))
	}
}

type MvFactory struct {
	AbstractFactory
}

func (mv *MvFactory) Option() CommandOptions {
	return new(CmdMv)
}
