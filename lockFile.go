package main

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
)

var LockFileName string

func setLockFileName(name string, url string) {
	h := sha1.New()
	h.Write([]byte(name + url))

	LockFileName = hex.EncodeToString(h.Sum(nil)) + ".lock"
}

func createLockFile() error {
	lockFile, err := os.Create(LockFileName)
	if err != nil {
		return err
	}

	lockFile.Close()
	return nil
}

func isLockFileExist() bool {
	if _, err := os.Stat(LockFileName); err == nil {
		return true
	}

	return false
}

func removeLockFile() error {
	if err := os.Remove(LockFileName); err != nil {
		return err
	}

	return nil
}
