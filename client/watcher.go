package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/zsxawerdu/pgrok/service"
)

type Watcher struct {
	pipeName string
	filePath string
	writer   RequestWriter
}

func (w *Watcher) Watch() {
	go w.listen()
}

func (w *Watcher) listen() {
	b := make([]byte, 4096)

OpenFile:
	for {
		// os.Open will block and wait for the mkfifo pipe to open
		f, err := os.Open(w.filePath)
		if err != nil {
			log.Println(err)
		}

		for {
			n, err := f.Read(b)
			switch err {
			case nil:
				wr := service.WriteReq{
					PipeName: w.pipeName,
					Data:     b[:n],
				}
				w.writer.WriteRequest(wr)
			case io.EOF:
				f.Close()
				time.Sleep(250 * time.Millisecond)
				continue OpenFile
			default:
				f.Close()
				log.Fatal(err)
				return
			}
		}
	}
}
