package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/folder", folderHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func folderHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	getdrive := query["drive"][0]
	fmt.Fprintf(w, "drive = %s\n", getdrive)
	var dirve = getdrive + ":/BaiduNetdiskDownload/"
	dirlist, e := ioutil.ReadDir(dirve)
	if e != nil {
		fmt.Fprintf(w, "read dir error = %q\n", e)
	}
	for i, v := range dirlist {
		fmt.Fprintf(w, "%s:/%s\n", getdrive, v.Name())
		fmt.Println(i, "=", v.Name())
	}
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	wg.Add(1)
	go walkDir(w, dirve, &wg, fileSizes)

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	tick = time.Tick(100 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			fmt.Fprintf(w, "%d files %f MB\n", nfiles, nbytes/1024/1024)
		}
	}
	fmt.Fprintf(w, "%d files %f GB\n", nfiles, nbytes /1024/1024/1024)
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("read dir %s error\n", dir)
		return nil
	}
	return entries
}

func walkDir(w http.ResponseWriter, dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			fmt.Fprintf(w, "%s\n", subDir)
			go walkDir(w, subDir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}

	}
}
