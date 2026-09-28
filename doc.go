package main

type ImgBrief struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
	Hash   string `json:"hash"`
	Ext    string `json:"ext"`
}
type FileInfoShortT struct {
	Name         string     `json:"name"`
	ImgBaseNames []ImgBrief `json:"imgs"`
}

var DocList []FileInfoShortT

func (i *ImgBrief) GetHashName() string {
	return i.Hash + i.Name
}
