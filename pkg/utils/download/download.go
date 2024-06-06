package download

import (
	"context"
	"encoding/base64"
	"github.com/bytedance/sonic"
	"io"
	http "net/http"
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

type UrlContents struct {
	Type        string `json:"type"`
	Encoding    string `json:"encoding"`
	Size        int    `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Content     string `json:"content"`
	Sha         string `json:"sha"`
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	DownloadUrl string `json:"download_url"`
	Links       struct {
		Self string `json:"self"`
		Html string `json:"html"`
	} `json:"_links"`
}

type IDL struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Content string `json:"content"`
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
	defer close(results)

	for i, content := range contents {
		// download queue + 1 and goroutine queue - 1
		wg.Add(1)
		results <- i

		go func(content Content) {
			defer func() {
				<-results
				wg.Done()
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

func FileContent(ctx context.Context, contents *[]Content) (*[]IDL, error) {
	var idls []IDL
	for _, content := range *contents {
		response, err := http.Get(content.URL)
		if err != nil {
			return nil, err
		}
		var urlBody UrlContents
		body, err := io.ReadAll(response.Body)
		if err = sonic.Unmarshal(body, &urlBody); err != nil {
			return nil, err
		}
		decodeString, err := base64.StdEncoding.DecodeString(urlBody.Content)
		if err != nil {
			return nil, err
		}
		idls = append(idls, IDL{
			Name:    urlBody.Name,
			Path:    urlBody.Path,
			Content: string(decodeString),
		})
	}
	return &idls, nil
}
