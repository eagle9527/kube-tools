/*
Copyright © 2023 eagle9527 <1032231418@qq.com>
*/
package pkg

type CommandOptions interface {
	Exe(parameter []string)
}

type AbstractFactory interface {
	Option() CommandOptions
}
