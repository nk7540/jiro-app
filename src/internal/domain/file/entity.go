package file

import "io"

type File struct {
	Path string
	Body io.Reader
}
