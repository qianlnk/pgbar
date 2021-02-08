package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/qianlnk/pgbar"
)

type Downloader struct {
	io.Reader
	bar *pgbar.Bar
}

func NewDownloader(resp *http.Response) *Downloader {
	nb := pgbar.NewBar(0, "下载进度", int(resp.ContentLength))
	if resp.ContentLength > 10*1024 {
		nb.SetUnit("B", "kb", 1024*1024)
	}

	if resp.ContentLength > 10*1024*1024 {
		nb.SetUnit("B", "MB", 1024*1024)
	}
	return &Downloader{
		Reader: resp.Body,
		bar:    nb,
	}
}

func (d *Downloader) Read(p []byte) (n int, err error) {
	n, err = d.Reader.Read(p)

	d.bar.Add(n)

	return n, err
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("xget url <filename>")
	}
	url := os.Args[1]

	us := strings.Split(url, "/")
	filename := us[len(us)-1]

	if len(os.Args) == 3 {
		filename = os.Args[2]
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	downloader := NewDownloader(resp)

	io.Copy(file, downloader)

	file.Close()
}
