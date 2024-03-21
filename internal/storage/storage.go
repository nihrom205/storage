package storage

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/nihrom205/ozonCurses/05-folders/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(fileName, blob)
	if err != nil {
		return &file.File{}, err
	}

	s.files[newFile.ID] = newFile

	return newFile, err
}

func (s *Storage) GetByID(fileId uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[fileId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("File %v not found", fileId))
	}
	return foundFile, nil
}
