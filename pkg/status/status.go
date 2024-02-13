package status

import (
	"fmt"
	// "encoding/json"
	// "io"
	// "io/fs"
	// "log"
	// "net/http"
	// "os"
	// "path/filepath"
	// "time"
)

func Status() {
	fmt.Println("Partner!")
}

// type NoEmbedInfo struct {
// 	Title           string `json:"title"`
// 	AuthorURL       string `json:"author_url"`
// 	AuthorName      string `json:"author_name"`
// 	ThumbnailURL    string `json:"thumbnail_url"`
// 	HTML            string `json:"html"`
// 	Height          int    `json:"height"`
// 	ProviderURL     string `json:"provider_url"`
// 	Type            string `json:"type"`
// 	ProviderName    string `json:"provider_name"`
// 	ThumbnailWidth  int    `json:"thumbnail_width"`
// 	Width           int    `json:"width"`
// 	URL             string `json:"url"`
// 	Version         string `json:"version"`
// 	ThumbnailHeight int    `json:"thumbnail_height"`
// 	Error           string `json:"error"`
// }

// func GetNoEmbedInfo(video_id string) NoEmbedInfo {
// 	resp, err := http.Get("https://noembed.com/embed?url=https://www.youtube.com/watch?v=" + video_id)
// 	if err != nil {
// 		fmt.Println("No response from request")
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body) // response body is []byte

// 	var info NoEmbedInfo
// 	err = json.Unmarshal(body, &info)
// 	if err != nil {
// 		fmt.Println("An error occured: %v", err)
// 		os.Exit(1)
// 	}
// 	return info
// }

// func timer(name string) func() {
// 	start := time.Now()
// 	return func() {
// 		fmt.Printf("%s took %v\n", name, time.Since(start))
// 	}
// }

// func walk(s string, d fs.DirEntry, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	if !d.IsDir() { // if file
// 		println(s)
// 		println(d.Name())
// 	}
// 	return nil
// }

// func main() {
// 	defer timer("main")()

// 	// var wg sync.WaitGroup
// 	// // private := "98TyXNPdMRo"
// 	// normal := "mmYv3Haj6uc"
// 	// video_ids := []string{
// 	// 	normal,
// 	// }
// 	// fmt.Print(len(video_ids))
// 	// for _, element := range video_ids {
// 	// 	wg.Add(1)
// 	// 	go func() {
// 	// 		defer wg.Done()
// 	// 		GetNoEmbedInfo(element)
// 	// 	}()
// 	// 	wg.Wait()
// 	// }
// 	// // fmt.Printf("%+v\n", data)
// 	// fmt.Println(data.Title)

// 	dirname, err := os.UserHomeDir()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var myslice []string
// 	filepath.WalkDir(dirname, walk())

// 	// file, err := os.Open(dirname)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// defer file.Close()

// 	// names, err := file.Readdirnames(0)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// for _, v := range names {
// 	// 	fmt.Println(v)
// 	// }
// }
