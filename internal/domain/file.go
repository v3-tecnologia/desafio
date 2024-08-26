package domain

import "errors"

type File struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
	Ext     string `json:"ext"`
}

type FileDto struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
	Ext     string `json:"ext"`
}

func NewFile(f *FileDto) *File {
	return &File{
		Name:    f.Name,
		Content: f.Content,
		Ext:     f.Ext,
	}
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) GetContent() []byte {
	return f.Content
}

func (f *File) GetExt() string {
	return f.Ext
}

func (f *File) Validate() (bool, error) {
	if f.Name == "" {
		return false, ErrEmptyFile
	}

	if f.Content == nil {
		return false, ErrEmptyContentFile
	}

	if f.Ext == "" {
		return false, ErrEmptyExtFile
	}

	return true, nil
}

var (
	ErrEmptyFile        = errors.New("name is empty")
	ErrEmptyContentFile = errors.New("content is empty")
	ErrEmptyExtFile     = errors.New("ext is empty")
)
