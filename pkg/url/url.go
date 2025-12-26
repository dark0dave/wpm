package url

import (
	"io"
	"net/http"
	u "net/url"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

const URLFOLDERNAME string = "url"

type Dependency struct {
	Name    string `yaml:"name"`
	Url     u.URL  `yaml:"url"`
	Version string `yaml:"version"`
}

func New(name, version string, url u.URL) *Dependency {
	return &Dependency{
		Name:    name,
		Url:     url,
		Version: version,
	}
}

func (u *Dependency) Download(folderPath string) (err error) {
	res, err := http.Get(u.Url.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	mtype, err := mimetype.DetectReader(res.Body)

	path := filepath.Join(folderPath, URLFOLDERNAME)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	file_path := filepath.Join(folderPath, URLFOLDERNAME, u.Name+mtype.Extension())
	out, err := os.Create(file_path)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, res.Body); err != nil {
		return err
	}

	return nil
}
