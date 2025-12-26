package manifest

const (
	GIT DependencyType = iota
	URL
	DROPBOX
)

type DependencyType int

type Dependency interface {
	Download(folderPath string) error
}
