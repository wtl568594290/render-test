package main

type ImgBrief struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}
type FileInfoShortT struct {
	Name         string     `json:"name"`
	ImgBaseNames []ImgBrief `json:"imgs"`
}

var DocList []FileInfoShortT
