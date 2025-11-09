package port

type FileFinder interface {
	Find(folderPath string) ([]string, error)
}