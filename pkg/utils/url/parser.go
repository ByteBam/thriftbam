package url

import (
	"fmt"
	"strings"
)

const IDLPath = "idl"

type FilePath struct {
	Owner       string
	Repo        string
	AccessToken string
	Path        string
	Ref         string
}

func NewFilePath(
	storeLink string,
	AccessToken string,
	Ref string,
) *FilePath {
	_, str, ok := strings.Cut(storeLink, "https://gitee.com/")
	if !ok {
		return nil
	}
	strs := strings.Split(str, "/")
	owner := strs[0]
	repo, _, ok := strings.Cut(strs[1], ".git")
	if !ok {
		return nil
	}
	return &FilePath{
		Owner:       owner,
		Repo:        repo,
		AccessToken: AccessToken,
		Path:        IDLPath,
		Ref:         Ref,
	}
}

func (f *FilePath) Parser() string {
	// 拼接url
	return fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s?access_token=%s&ref=%s", f.Owner, f.Repo, f.Path, f.AccessToken, f.Ref)
}
