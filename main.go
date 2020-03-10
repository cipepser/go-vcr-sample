package main

import (
	"fmt"
	"github.com/cipepser/go-vcr-sample/qiita"
)

func main() {
	c := qiita.NewClient()
	u, err := c.FetchUser("cipepser")
	if err != nil {
		panic(err)
	}
	fmt.Println(*u)
}
