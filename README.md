# go-vcr-sample

[【Golang\+VCR】外部APIとの通信を保存してテストに使用する話 \- yyh\-gl's Tech Blog](https://yyh-gl.github.io/tech-blog/blog/golang-vcr/)を触ってみる。


```text
❯ go test -count=1 ./...
?   	github.com/cipepser/go-vcr-sample	[no test files]
ok  	github.com/cipepser/go-vcr-sample/qiita	0.015s
```

上記テストを初回実行した場合、`fixtures`ディレクトリが自動的に作られ、レスポンスがymlで保存される。

## References
- [【Golang\+VCR】外部APIとの通信を保存してテストに使用する話 \- yyh\-gl's Tech Blog](https://yyh-gl.github.io/tech-blog/blog/golang-vcr/)
- [rpcreplay \- GoDoc](https://godoc.org/cloud.google.com/go/rpcreplay)
- [yyh\-gl/go\-vcr\-sample: go\-vcr 使ってみた](https://github.com/yyh-gl/go-vcr-sample)