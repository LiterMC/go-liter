
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type HttpStatusError struct{
	Code int
}

var _ error = (*HttpStatusError)(nil)

func (e *HttpStatusError)Error()(string){
	return fmt.Sprintf("%d %s", e.Code, http.StatusText(e.Code))
}

func GetOrUseCached(url string, filename string)(r io.ReadCloser, err error){
	res, err := http.DefaultClient.Get(url)
	if err == nil {
		if res.StatusCode == http.StatusOK {
			defer res.Body.Close()
			buf := bytes.NewBuffer(nil)
			_, err = io.Copy(buf, res.Body)
			if err == nil {
				os.MkdirAll(filepath.Dir(filename), 0755)
				os.WriteFile(filename, buf.Bytes(), 0644)
				return io.NopCloser(buf), nil
			}
		}else{
			res.Body.Close()
			err = &HttpStatusError{res.StatusCode}
		}
	}
	fd, er := os.Open(filename)
	if er != nil {
		return
	}
	return fd, nil
}
