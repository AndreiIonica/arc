package template

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// PLATFORM INDEPENTENT way of copying stuff

// Copy CONTENTS of src directly into dst
func copyFolder(src, dst string) error {
	// TODO better error handling
	var err error
	var content []os.FileInfo
	var srcInfo os.FileInfo

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	if content, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range content {
		srcFilePath := path.Join(src, fd.Name())
		destFilePath := path.Join(dst, fd.Name())

		if fd.IsDir() {
			// TODO: Recursive, might wanna check performance for deeply nested directories
			if err = copyFolder(srcFilePath, destFilePath); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = copyFile(srcFilePath, destFilePath); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

// copyFile copies a single file from src to dst
func copyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
