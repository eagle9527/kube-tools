package pkg

type CommandOptions interface {
	Exe(parameter []string)
}

type CmdFactory struct{}

func (fac *CmdFactory) CreateCommandOption(kind string) CommandOptions {
	var cmdop CommandOptions

	if kind == "ls" {
		cmdop = new(CmdLs)
	} else if kind == "cat" {
		cmdop = new(CmdCat)
	} else if kind == "mkdir" {
		cmdop = new(CmdMkdir)
	} else if kind == "rm" {
		cmdop = new(CmdRm)
	} else if kind == "touch" {
		cmdop = new(CmdTouch)
	} else if kind == "mv" {
		cmdop = new(CmdMv)
	} else if kind == "tar" {
		cmdop = new(CmdTar)
	}

	return cmdop
}
