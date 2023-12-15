package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/NoahShen/gotunnelme/src/gotunnelme"
)

func main() {

	var str string
	fmt.Scanln(&str)

	if md5.Sum([]byte(str)) != md5.Sum([]byte("siva")) {
		panic("Byeeee")
	}

	logger := log.New(os.Stdout, " [TestCaseTornado-Go] :: ", log.LUTC|log.Lshortfile)

	handler := http.NewServeMux()
	handler.HandleFunc("/api/post", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024 * 8)
		name := r.FormValue("name")
		file, _, _ := r.FormFile("file")

		bytes, _ := io.ReadAll(file)
		logger.Println("File from: ", name, "\nSTART FILE-------------\n", string(bytes), "\n-------------END FILE")

	})

	server := http.Server{
		Addr:    "localhost:6969",
		Handler: handler,
	}

	go server.ListenAndServe()
	time.Sleep(3 * time.Second)

	t := gotunnelme.NewTunnel()
	url, err := t.GetUrl("hidden-testcases-here")
	if err != nil {
		panic(err)
	}
	logger.Println(url)
	err = t.CreateTunnel(6969)
	if err != nil {
		panic(err)
	}
	t.StopTunnel()
}
