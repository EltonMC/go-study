package util

import (
	"github.com/silentsokolov/go-vimeo/vimeo"
	"golang.org/x/oauth2"
)

func GetVideoDownloadLink(token string, id int) string {
	video, _, _ := getClientVimeo(token).Videos.Get(id)
	downloadList := video.Download
	bestQuality := downloadList[len(downloadList)-1]

	return bestQuality.Link
}

func getClientVimeo(token string) *vimeo.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return vimeo.NewClient(tc, nil)
}
