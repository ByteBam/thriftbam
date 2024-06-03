package download

import (
	"context"
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

func File(ctx context.Context, contents []Content, dirPath string) error {
	// Communication between goroutines is implemented with the help of context
	// if err != nil, it breaks
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	// create the file storage path
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	var wg sync.WaitGroup
	var once sync.Once
	var err error
	// Buffer 10, maximum concurrent downloads 10
	results := make(chan int, 10)

	for i, content := range contents {
		// download queue + 1 and goroutine queue - 1
		wg.Add(1)
		results <- i

		go func(content Content) {
			defer func() {
				wg.Done()
				<-results
			}()

			// download file raw
			resp, e := http.Get(content.DownloadURL)
			if e != nil {
				// once.Do to make sure err written only once
				// if err != nil, it breaks
				once.Do(func() {
					err = e
					cancel(e)
				})
				return
			}
			defer resp.Body.Close()

			// creat file if the file not exists
			file, e := os.OpenFile(filepath.Join(dirPath, content.Name), os.O_CREATE|os.O_WRONLY, 0644)
			if e != nil {
				once.Do(func() {
					err = e
					cancel(e)
				})
				return
			}
			defer file.Close()

			// copy the file raw into a file form response body
			if _, err = io.Copy(file, resp.Body); err != nil {
				once.Do(func() {
					err = e
					cancel(e)
				})
				return
			}
		}(content)
	}

	wg.Wait()

	return err
}
