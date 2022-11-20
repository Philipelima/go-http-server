package server

import (
	"fmt"
	"net/http"
)

type ServerSettings struct {
	port string
	ip   string
}

type Server struct {
	directory string
	settings  ServerSettings
}

type Callback func(http.ResponseWriter, *http.Request)

func (s *Server) SetSettings(ip string, port string) {
	s.settings = ServerSettings{ip: ip, port: port}
}

func (s *Server) SetDirectory(directory string) {
	s.directory = directory
}

func (s *Server) Run() {
	s.start()

	fileHandler := http.FileServer(http.Dir(s.directory))

	http.Handle("/", fileHandler)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println("Error")
	}
}

func (s *Server) Get(pattern string, function Callback) {
	http.HandleFunc(pattern, function)
}


func (s *Server) start() {

	fmt.Printf("\n Starting server on %s:%s...\n", s.settings.ip, s.settings.port)
	fmt.Println(" Directory: ", s.directory)
	fmt.Println("\n ")

}
