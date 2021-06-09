package pkg

import "golang.org/x/xerrors"

func NewRepositoryError(err error) error {
	return xerrors.Errorf("failed to repository: %w", err)
}

func NewValidationError() error {
	return xerrors.New("failed to validation")
}
