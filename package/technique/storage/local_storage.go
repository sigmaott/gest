package storage

type Storage interface {
	SaveFile(filename string, data []byte) error
	GetFile(filename string) ([]byte, error)
	DeleteFile(filename string) error
}
