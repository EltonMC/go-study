package main

import (
	"fmt"
	"github.com/silentsokolov/go-vimeo/vimeo"
	"golang.org/x/oauth2"
)

func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := vimeo.NewClient(tc, nil)

	video, _, _ := client.Videos.Get(687283583)

	fmt.Println(video.Name)

	//for video := range videos {
	//	fmt.Println(video)
	//}

}
