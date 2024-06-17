package main

import (
	"log"
	"net/http"
)

func main() {
	// 정적 파일 서빙
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// INFO 페이지
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/info.html")
	})

	// PROJECTS 페이지
	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/projects.html")
	})

	// ARCHIVE 페이지
	http.HandleFunc("/archive", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/archive.html")
	})

	// 기본 페이지 설정 (index.html 이라면, / 로 접근시 index.html 로 이동)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/info.html", http.StatusMovedPermanently)
	})

	log.Println("서버가 포트 8080에서 실행 중입니다...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
