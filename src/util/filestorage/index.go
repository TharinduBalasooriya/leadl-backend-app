package filestorageHandler

/*
	Interface for file upload
*/

type FileStorage interface {
	AddFiles() error
}

type File interface {
	GetContent(fileId string) error
}
