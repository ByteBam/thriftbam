package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type Links struct {
	Self string `json:"self"`
	HTML string `json:"html"`
}

type Content struct {
	Type        string `json:"type"`
	Size        int    `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	SHA         string `json:"sha"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	DownloadURL string `json:"download_url"`
	Links       Links  `json:"_links"`
}

func File(contents []Content, dirPath string) error {
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(contents))

	for _, content := range contents {
		wg.Add(1)
		go func(content Content) {
			defer wg.Done()

			// 下载文件逻辑
			resp, err := http.Get(content.DownloadURL)
			if err != nil {
				errCh <- err
				return
			}
			defer resp.Body.Close()

			file, err := os.OpenFile(filepath.Join(dirPath, content.Name), os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				errCh <- err
				return
			}
			defer file.Close()

			if _, err := io.Copy(file, resp.Body); err != nil {
				errCh <- err
				return
			}
		}(content)
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("encountered errors: %v", errs)
	}

	return nil
}
