package main

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	buf := make([]byte, n)

	nRead, err := r.Read(buf)
	if err == io.EOF {
	} else if err != nil {
		return err
	}

	if nRead > 0 {
		_, err := w.Write(buf[:nRead])
		if err != nil {
			return err
		}
	}

	return nil
}
