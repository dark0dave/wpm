package manifest

type Dependency interface {
	Download(folderPath string) error
}
