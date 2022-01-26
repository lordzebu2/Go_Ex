package main

import (
	"fmt"
	"log"
	"net/http"
)

func wep(){
	http.HandleFunc("/", index)  //사이트 접속 시 index이름의 함수를 실행
	log.Fatal(http.ListenAndServe(":8090", nil)) //http.ListenAndServe로 웹서버 시작, 8080포트 외부에서 들어오는 http요청을 지속적으로 수신.
								//포트넘버, nil= Go에서 null이다. 
	//http.ListenAndServe는 항상 오류값을 반환한다. 평소에는 실행을 유지하다가 오류 발생때만 값을 반환함. 
	//따라서 이를 처리하게 log.Fatal로 감싸줬다. log.Fatal은 실행시 오류 내용을 메시지로 출력하고 프로그램을 종료함

}



func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "안녕하세요! %s에 오신 것을 환영합니다!", r.URL.Path[1:])
	//http.ResponseWriter는 http서버의 응답과 관련된 모든 데이터를 담고있음.
	//이 객체 변수를 대상으로 출력(메시지)을 하면 방문자인 클라이언트로 데이터(문장) 전송.
	//http.Request는 클라이언트의 http 요청관련정보.
}

