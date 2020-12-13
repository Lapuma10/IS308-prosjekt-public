package downloader

import (
	"io"
	"log"
	"net/http"
	"os"
)

//Code from:
//https://golangcode.com/download-a-file-from-a-url/

func Download(fileUrl string) string {

	fileName := "File.jsonl"
	dir := "/var/downloader/downloads/" + fileName
	err := DownloadFile(dir, fileUrl)
	if err != nil {
		panic(err)
	}
	log.Println("Downloaded: " + fileName)
	return dir
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//This renames file by the downloaded name
	//path.Base(resp.Request.URL.String())
	// filepath := path.Base(resp.Request.URL.String())

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
