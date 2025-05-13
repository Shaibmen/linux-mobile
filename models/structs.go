package models

type BashFile struct {
	NameField string `json:"nameField"`
	TextField string `json:"textField"`
}

type ErrorLog struct {
	err error
}

type Resource struct {
	Memory string
	Disk   []string
	CPU    []string
}

type Network struct {
	Netstat []string
	Ssi     []string
}

type Process struct {
	Process []string
}

type ProcessGrep struct {
	Prefix string `json:"prefix"`
}
