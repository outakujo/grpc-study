package main

type SysInfo struct {
	Id   string
	Os   string
	Arch string
}

func (SysInfo) TableName() string {
	return "sys_info"
}

type ScriptResult struct {
	Id     string
	Code   int
	Stdout string
}

func (ScriptResult) TableName() string {
	return "script_result"
}
