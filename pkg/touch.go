/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package pkg

import (
	"fmt"
	"k8s.io/klog/v2"
	"os"
)

type CmdTouch struct{}

func (tc *CmdTouch) Exe(parameter []string) {
	fmt.Println(parameter)
	file, err := os.OpenFile(parameter[0], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		klog.Errorln(fmt.Sprintf("OpenFile fail, %s", err))
	}

	if _, err := file.Write([]byte(parameter[1])); err != nil {
		klog.Errorln(fmt.Sprintf("File Write fail, %s", err))
	}

	file.WriteString("\n")

}
