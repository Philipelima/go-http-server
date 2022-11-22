package server

import (
	"fmt"
	"net/http"

	"github.com/philipelima/http-server/types"
)

type ServerSettings struct {
	port string
	ip   string
}

type Server struct {
	directory string
	settings  ServerSettings
}

func (s *Server) SetSettings(ip string, port string) {
	s.settings = ServerSettings{ip: ip, port: port}
}

func (s *Server) SetDirectory(directory string) {
	s.directory = directory
}

func (s *Server) Run() {

	s.start()

	fileDir := http.Dir(s.directory)
	fileHandler := http.FileServer(fileDir)

	http.Handle("/", fileHandler)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println("Error")
	}
}

func (s *Server) Get(pattern string, function types.Callback) {
	http.HandleFunc(pattern, function)
}

func (s *Server) start() {

	fmt.Printf("\nStarting server on %s:%s...\n", s.settings.ip, s.settings.port)
	fmt.Println("Directory: ", s.directory)
	fmt.Println("\n ")

}
