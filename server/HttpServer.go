package server

import (
	"applib/conf"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)



type HttpServer struct {
	MuxRouter *mux.Router
	Conf      conf.ServerConf
	instance  *http.Server
}

func (server *HttpServer) Create(conf conf.ServerConf, mdf mux.MiddlewareFunc) {

	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	server.Conf = conf
	router := mux.NewRouter()

	if mdf!=nil {
		router.Use(mdf)
	}

	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	server.MuxRouter = router

	d1 := strconv.Itoa(conf.WriteTimeout) + "s"
	writeDuration, err := time.ParseDuration(d1)
	if err != nil {
		panic(err)
	}

	d2 := strconv.Itoa(conf.ReadTimeout) + "s"
	readDuration, err := time.ParseDuration(d2)
	if err != nil {
		panic(err)
	}

	server.instance = &http.Server{
		Handler:      server.MuxRouter,
		Addr:         conf.GetUrl(),
		WriteTimeout: writeDuration,
		ReadTimeout:  readDuration,
	}
}

func (server *HttpServer) AddHandler(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.MuxRouter.HandleFunc(path, f)
}

func (server *HttpServer) Listen() {

	log.Printf("Computation Server is listening to " + server.Conf.GetUrl() + "\n")
	fmt.Println("Computation Server is listening to", server.Conf.GetUrl())

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
		"trID", "trType", "requestAccount", "requestUser", "requestApp", "requestSystem", "requestSession",
		"page", "pageSize", "usePaging",
	})
	//headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router := server.MuxRouter
	http.ListenAndServe(server.Conf.GetUrl(), handlers.CORS(headers, methods, origins)(router))	
}

func (server *HttpServer) Stop() {
	//server.instance.Shutdown(context.TODO())
}

