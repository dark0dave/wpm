package dropbox

import (
	"io"
	u "net/url"
	"os"

	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/schollz/progressbar/v3"
)

var (
	Token string = os.Getenv("DROPBOX_TOKEN")
)

type Dependency struct {
	*manifest.Dependency
	config *dropbox.Config
}

func New(name, version string, url u.URL) *Dependency {
	return &Dependency{
		Dependency: &manifest.Dependency{
			Name:     name,
			Url:      url,
			Version:  version,
			Protocol: manifest.DropBox,
		},
		config: &dropbox.Config{
			Token: Token,
		},
	}
}

func (d *Dependency) Download(folderPath string) error {
	link := files.NewDownloadArg(d.Url.String())
	dbx := files.New(*d.config)
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
