package core

import (
	"crypto/md5"
	"fmt"
	"os"
)

const BASE_PATH = "/tmp/protocg"

var WORK_CACHE_PATH string

func init() {
	WORK_CACHE_PATH = cachePathOfWork()
	if err := os.MkdirAll(WORK_CACHE_PATH, os.ModePerm); err != nil {
		panic(err)
	}
}

func workSig() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	sig := md5.Sum([]byte(dir))
	return fmt.Sprintf("%x", sig)[:8]
}

func cachePathOfWork() string {
	return fmt.Sprintf("%s-%s", BASE_PATH, workSig())
}

func cachePathOf(subpath string) string {
	return fmt.Sprintf("%s/%s", WORK_CACHE_PATH, subpath)
}
