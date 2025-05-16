package models

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

type Folder struct {
	Files []string
}

type ProcessGrep struct {
	Prefix string `json:"prefix"`
}

type PID struct {
	ID string `json:"id"`
}

type Dir struct {
	Path  string `json:"path"`
	Force bool   `json:"flag"`
}

type BashFile struct {
	NameField string `json:"nameField"`
	TextField string `json:"textField"`
}
