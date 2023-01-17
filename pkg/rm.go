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