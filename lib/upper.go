package lib

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const url = "https://gitee.com/api/v5/repos/%s/%s/contents/%s"

func fileName(path string) string {
	_, name := filepath.Split(path)
	return name
}

func Upload(user, repo, token, path string, files []string) {
	var succ []string
	for i := range files {
		// fmt.Println(files[i])
		name := fileName(files[i])
		log.Printf("uploading: %s", files[i])
		targetparh := path + "/" + name
		targeturl := fmt.Sprintf(url, user, repo, targetparh)
		data, err := ioutil.ReadFile(files[i])
		if err != nil {
			panic(err)
		}
		d := new(Req)
		d.Access_token = token
		d.Branch = "master"
		d.Message = "upload"
		d.Content = base64.StdEncoding.EncodeToString(data)
		jsondata, err := json.Marshal(d)
		if err != nil {
			panic(err)
		}
		r, err := http.Post(targeturl, "application/json;charset=UTF-8", bytes.NewReader(jsondata))
		if err != nil {
			panic(err)
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		rsp := new(Resp)
		if r.StatusCode < 200 || r.StatusCode >= 400 {
			fmt.Printf("%s\n%s\n", r.Status, string(b))
			os.Exit(0)
		}
		json.Unmarshal(b, rsp)
		succ = append(succ, rsp.Content.DownloadURL)
	}
	fmt.Println("Upload Success:")
	for i := range succ {
		fmt.Println(succ[i])
	}

}
