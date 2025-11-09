package port

type FileWatcher interface {
	Watch(folderPath string, handleOnChange func())
	Cancel()
}