package main

import (
	"log"
	"net/http"
)

func main() {
	// 정적 파일을 서빙
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// API 엔드포인트
	http.HandleFunc("/api/projects", projectsHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// API 핸들러
func projectsHandler(w http.ResponseWriter, r *http.Request) {
	projects := `[{"name": "Project 1", "description": "Description 1"}, {"name": "Project 2", "description": "Description 2"}]`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(projects))
}
