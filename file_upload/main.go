package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

//save용 핸들러
func saveHandler(w http.ResponseWriter, r *http.Request){
	//MultipartReder를 이용해서 받은 파일을 읽는다.
	reader, err := r.MultipartReader()

	//에러가 발생하면 던진다.
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//for로 복수 파일이 있는 경우에 모든 파일이 끝날때까지 읽는다.
	for{
		part, err := reader.NextPart()
		if err == io.EOF{
			break
		}

		//파일명이 없는 경우는 스킵
		if part.FileName() == ""{
			continue			
		}

		//uploadFile 디렉토리에 받았던 파일 명으로 파일을 만든다.
		uploadFile, err := os.Create("/upload" + part.FileName())
		
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			uploadFile.Close()
			//os자원이라 파일을 만들고나면 닫아줘야함
			redirectToErrorPage(w,r)
			return
		}
		//만든 파일에 읽었던 파일 내용을 모두 복사
		_, err = io.Copy(uploadFile, part)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			uploadFile.Close()
			redirectToErrorPage(w,r)
			return
		}
	}

	//upload 페이지에 리다이렉트
	http.Redirect(w, r, "/file.html", http.StatusFound)
}

// upload용 핸들러
func uploadHandler(w http.ResponseWriter, r *http.Request){
	var templatefile = template.Must(template.ParseFiles("file.html"))
	templatefile.Execute(w, "file.html")
}

//errorPage용 핸들러
func errorPageHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "%s", "<p>에러남<p>")
}

//error가 발생했을 때에 에러 페이지로 이동한다.
func redirectToErrorPage(w http.ResponseWriter, r *http.Request){
	http.Redirect(w,r,"/errorPage", http.StatusFound)
}

func main(){
	//핸들러 등록
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/errorPage", errorPageHandler)
	//서버시작
	http.ListenAndServe(":8080", nil)
}
