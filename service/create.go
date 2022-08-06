package service

import (
	"errors"
	"io/ioutil"
	"log"
)

func GenerateProject(project string) {
	createApiaHandle.setProjectName(project)
	createApiaHandle.run()
}

var createApiaHandle = createProject{}

type createProject struct {
	project   string
	basecache string
}

func (c *createProject) setProjectName(p string) {
	c.project = p
}

func (c *createProject) judgeDirIsEmpty() bool {
	cur, err := ioutil.ReadDir(".")
	if err != nil {
		return false
	}

	if len(cur) > 0 {
		return false
	}

	return true
}

func (c *createProject) run() error {
	if !c.judgeDirIsEmpty() {
		log.Println("目录非空")
		return errors.New("current workdir is not empty, please use empty dir")
	}

	return nil
}
