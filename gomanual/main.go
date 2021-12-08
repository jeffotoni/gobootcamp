package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

var (
	HTTP_PORT = "0.0.0.0:8181"
	HOST      = "http://localhost:" + HTTP_PORT + "/manual"
)

// go:embed manual
// var res embed.FS
func main() {
	// fs := http.FileServer(http.FS(res))
	// http.Handle("/manual/", http.StripPrefix("/", fs))

	http.Handle("/", http.StripPrefix("/",
		http.FileServer(http.Dir("./manual"))))

	err := openBrowser(HOST)
	if err != nil {
		log.Println(err)
		return
	}
	println("Run Server:", HOST)
	http.ListenAndServe(HTTP_PORT, nil)
}

func openBrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	return exec.Command(cmd, append(args, url)...).Start()
}
