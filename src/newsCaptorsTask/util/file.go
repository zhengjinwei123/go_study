package util

import (
	"runtime"
	"strings"
	"path/filepath"
	"os"
	"errors"
)

func GetFilePath(resolvePath string) (string,error){
	if resolvePath == "" {
		panic("GetFilePath params is empty")
	}
	_,file,_,ok := runtime.Caller(1)
	if !ok {
		panic("GetFilePath error")
	}

	resolvePathList := strings.Split(resolvePath,"/")

	var oneDot int = 0
	var twoDot int = 0

	var twoDotList []int = make([]int,20)
	if len(resolvePathList) != 0 {
		for k,v := range resolvePathList {
			if v == "." {
				if k == 0 {
					oneDot ++
				} else {
					panic("invalid path")
				}
			} else if v == ".." {
				twoDot ++
				twoDotList[k] = 1
			} else {
				break
			}
		}
	}


	// 获取调用文件所在目录
	fileDir := filepath.Dir(file)

	if twoDot == 0 {
		if oneDot != 0 && oneDot > 1 {
			panic("invalid path")
		} else {
			for _,v := range resolvePathList {
				if v != "." {
					fileDir += "\\"+v
				}
			}
		}
	} else {
		for _,v := range twoDotList {
			_ = v
			if v == 0 {
				break
			}
			fileDir = filepath.Dir(fileDir)
		}

		for _,v := range resolvePathList {
			if v != ".."  && v != "." {
				fileDir += "\\"+v
			}
		}
	}

	_,err := PathExists(fileDir)
	if err != nil {
		panic(err)
	}

	return fileDir,nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, errors.New(path + " is not a valid file")
	}
	return false, err
}
