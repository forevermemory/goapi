package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
)

/**
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * @param srcPath 需要拷贝的文件夹路径: .../projects
 * @param destPath 拷贝到的位置: .../projects/.cache
 */
func CopyDir(src string, dest string) error {
	// 检测原目录正确性
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !srcInfo.IsDir() {
		return errors.New("srcPath不是一个正确的目录！")
	}

	// 目标文件夹是否存在
	destInfo, err := os.Stat(dest)
	if err != nil {
		os.Mkdir(dest, os.ModePerm)
	}
	if !destInfo.IsDir() {
		return errors.New("destInfo不是一个正确的目录！")
	}

	// 递归拷贝文件夹和文件
	fmap(src, dest)

	return nil
}

func fmap(src string, dest string) error {
	// 遍历原文件夹内部所有item
	items, _ := ioutil.ReadDir(src)
	for _, item := range items {
		// 忽略.cache
		if item.Name() == ".cache" {
			continue
		}

		// 文件
		if !item.IsDir() {
			cpoyFile2(path.Join(src, item.Name()), path.Join(dest, item.Name()))
			continue
		}

		// 目录
		os.Mkdir(path.Join(dest, item.Name()), os.ModePerm)
		// 递归
		fmap(path.Join(src, item.Name()), path.Join(dest, item.Name()))
	}

	return nil
}

func cpoyFile2(src, dest string) error {
	// open src readonly
	srcFp, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFp.Close()

	// create dest
	dstFp, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dstFp.Close()
	_, err = io.Copy(dstFp, srcFp)
	return err
}
