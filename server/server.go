package server

import (
	"encoding/json"
	"fmt"
	stdLibLog "log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/stock-tracker/config"
)

var (
	log = logrus.WithField("component", "server")
)

type Server struct {
	Log        *logrus.Logger
	HttpServer *http.Server
}

type ResponseBody struct {
	Symbol   string `json:"symbol"`
	Day      string `json: "date as text"`
	Open     int    `json: "Open"`
	High     int    `json:"High"`
	Low      int    `json:Low"`
	Close    int    `json:"close"`
	Avgprice int    `json:"average price"`
}

// type ServerConfig struct {
//     Port int
//     BindAddress string
// }

func RunHTTPServer(serverconfig config.ServerConfig) {
	httpServer, err := New(serverconfig)
	if err != nil {
		log.Error(err)
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Error(err)
	}
}

func New(config config.ServerConfig) (*Server, error) {
	s := &Server{
		Log: logrus.New(),
	}

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         fmt.Sprintf("%s:%d", config.BindAddress, config.Port),
		Handler:      s.GetRouter(),
	}
	s.HttpServer = server
	return s, nil
}

func (s *Server) GetRouter() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(NotFound)
	router.GET("/v1/stocker", s.Stocker)
	return router

}

func (s *Server) ListenAndServe() error {
	w := log.Logger.Writer()
	s.HttpServer.ErrorLog = stdLibLog.New(w, "", 0)
	err := s.HttpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.WithError(err).Error("error from server")
		panic(err)
	}
	return err
}

func (s *Server) Stocker(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var log = log.WithField("func", "Stocker")
	enc := json.NewEncoder(w)
	response := struct {
		Symbol   string
		Day      string
		Open     float32
		High     float32
		Low      float32
		Close    float32
		Avgprice float32
	}{
		Symbol:   "MSFT",
		Day:      "2022-10-07",
		Open:     240.9020,
		High:     241.3200,
		Low:      233.1700,
		Close:    234.2400,
		Avgprice: 236.2433,
	}

	err := enc.Encode(response)
	if err != nil {
		log.Error(err)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "The requested route does not exist. \n")
}
