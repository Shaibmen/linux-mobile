package models

type BashFile struct {
	NameField string `json:"nameField"`
	TextField string `json:"textField"`
}

type ErrorLog struct {
	err error
}

type Resource struct {
	Memory []string
	Disk   []string
	CPU    []string
}
