package main

import (
	"fmt"
	"github.com/cipepser/go-vcr-sample/qiita"
	"net/http"
)

func main() {
	c := qiita.NewClient(http.DefaultClient)
	u, err := c.FetchUser("cipepser")
	if err != nil {
		panic(err)
	}
	fmt.Println(*u)
}
