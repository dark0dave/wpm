package util

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
)

func getFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(filepath.Clean(dir), func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel := file
		if dir != "." {
			rel = file[len(dir)+1:]
		}
		files = append(files, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func hash512(files []string) (string, error) {
	h := sha512.New()
	files = append([]string(nil), files...)
	slices.Sort(files)
	for _, file := range files {
		r, err := os.Open(file)
		if err != nil {
			return "", err
		}
		defer r.Close()
		hf := sha512.New()
		if _, err = io.Copy(hf, r); err != nil {
			return "", err
		}
		fmt.Fprintf(h, "%x  %s\n", hf.Sum(nil), file)
	}
	return "sha512-" + base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func CheckSum(dir string) (string, error) {
	files, err := getFiles(dir)
	if err != nil {
		return "", err
	}
	return hash512(files)
}
