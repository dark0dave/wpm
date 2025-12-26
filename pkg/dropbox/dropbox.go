package dropbox

import (
	"io"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/schollz/progressbar/v3"
)

type Dependency struct {
	Name    string `yaml:"name"`
	Url     string `yaml:"url"`
	Version string `yaml:"version"`
}

func (d *Dependency) Download(config dropbox.Config, folderPath string) error {
	link := files.NewDownloadArg(d.Url)
	dbx := files.New(config)
	meta, contents, err := dbx.Download(link)
	if err != nil {
		return err
	}
	defer contents.Close()

	f, err := os.Create(folderPath)
	if err != nil {
		return err
	}
	defer f.Close()

	bar := progressbar.DefaultBytes(
		int64(meta.Size),
		"downloading",
	)

	if _, err := io.Copy(io.MultiWriter(f, bar), contents); err != nil {
		return err
	}

	return nil
}
