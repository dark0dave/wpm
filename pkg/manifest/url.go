package manifest

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

const url_folder_name string = "url"

type UrlDependancy struct {
	Name     string `yaml:"name"`
	Location string `yaml:"location"`
	Version  string `yaml:"version"`
}

func (u *UrlDependancy) GetName() string {
	return u.Name
}

func (u *UrlDependancy) Download(folderPath string) (err error) {
	res, err := http.Get(string(u.Location))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	mtype, err := mimetype.DetectReader(res.Body)

	path := filepath.Join(folderPath, url_folder_name)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	file_path := filepath.Join(folderPath, url_folder_name, u.Name+mtype.Extension())
	out, err := os.Create(file_path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}
