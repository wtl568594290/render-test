package main

import (
	"fmt"
	"os"
)

type Storage map[string]string

const (
	ImgDir = "images"
)

var StorageMap = make(Storage)

func init() {
	// 删除images目录下的所有文件
	_ = os.RemoveAll(ImgDir)
	// 重新创建images目录
	_ = os.Mkdir(ImgDir, 0755)

}
func (s Storage) add(name string, path string) {
	StorageMap[name] = path
}

func (s Storage) get(name string) (string, error) {
	if path, ok := StorageMap[name]; ok {
		return path, nil
	}
	return "", fmt.Errorf("img not found")
}

func (s Storage) has(name string) bool {
	if _, ok := StorageMap[name]; ok {
		return true
	}
	return false
}

func (s Storage) sync(dir []FileInfoShortT) []ImgBrief {
	// docKeys 收集新 docList 中所有图片 key(hash+name),并找出 map 中缺失的
	docKeys := make(map[string]struct{})
	addList := make([]ImgBrief, 0)
	for _, item := range dir {
		for _, img := range item.ImgBaseNames {
			key := img.GetHashName()
			docKeys[key] = struct{}{}
			if _, ok := s[key]; !ok {
				addList = append(addList, img)
			}
		}
	}
	// 删除 map 中存在、但新 docList 中已不存在的图片
	for key := range s {
		if _, ok := docKeys[key]; !ok {
			// 删除图片文件
			_ = os.Remove(s[key])
			delete(s, key)
		}
	}
	return addList
}
