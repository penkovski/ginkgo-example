package service

import (
	"context"
	"errors"
)

//go:generate counterfeiter . Users
//go:generate counterfeiter . Storage
//go:generate counterfeiter . Download

var (
	ErrUnauthorized = errors.New("unauthorized request")
	ErrFileNotFound = errors.New("file not found")
	ErrInternal     = errors.New("internal error")
	ErrFileDownload = errors.New("error downloading file")
	ErrFileSave     = errors.New("error saving file")
)

type Users interface {
	Get(ctx context.Context, token string) (*User, error)
}

type Storage interface {
	File(ctx context.Context, filename string, userGUID string) (*File, error)
	SaveFile(ctx context.Context, file *File, userGUID string) error
}

type Download interface {
	Download(ctx context.Context, filename string, userGUID string) (*File, error)
}

type User struct {
	GUID  string
	Email string
	Name  string
}

type File struct {
	Name string `json:name`
	Size int    `json:size`
}

type Service struct {
	users    Users
	storage  Storage
	download Download
}

func New(users Users, storage Storage, download Download) *Service {
	return &Service{
		users:    users,
		storage:  storage,
		download: download,
	}
}

func (s *Service) File(ctx context.Context, token string, filename string) (*File, error) {
	user, err := s.users.Get(ctx, token)
	if user == nil || err != nil {
		return nil, ErrUnauthorized
	}

	file, err := s.storage.File(ctx, filename, user.GUID)
	if err != nil {
		if err != ErrFileNotFound {
			return nil, ErrInternal
		}

		file, err = s.download.Download(ctx, filename, user.GUID)
		if err != nil {
			return nil, ErrFileDownload
		}

		err = s.storage.SaveFile(ctx, file, user.GUID)
		if err != nil {
			return nil, ErrFileSave
		}
	}

	return file, nil
}
