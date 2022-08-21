package u2

import (
	"errors"
	"log"
	"os"
	"path"
)

func Madir(workdir, dirname string) error {
	err := os.MkdirAll(path.Join(".", workdir, dirname), 0777)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			// 已存在文件夹不用关心
			return nil
		}
		log.Println("initWorkdir other err:", err)
		return err
	}
	return nil
}
