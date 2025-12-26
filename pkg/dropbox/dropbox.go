package dropbox

import (
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

type Dependency struct {
	Name    string `yaml:"name"`
	Url     string `yaml:"url"`
	Version string `yaml:"version"`
}

func (d *Dependency) Download(folderPath string) error {
	c := files.NewDownloadArg(d.Url)
	dbx := files.New(dropbox.Config{})
	_, contents, err := dbx.Download(c)
	if err != nil {
		return err
	}
	defer contents.Close()

	f, err := os.Create(folderPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// progressbar := &ioprogress.Reader{
	// 	Reader: contents,
	// 	DrawFunc: ioprogress.DrawTerminalf(os.Stderr, func(progress, total int64) string {
	// 		return fmt.Sprintf("Downloading %s/%s",
	// 			uint64(progress), uint64(total))
	// 	}),
	// 	Size: int64(res.Size),
	// }

	// if _, err = io.Copy(f, progressbar); err != nil {
	// 	return err
	// }

	return nil
}
