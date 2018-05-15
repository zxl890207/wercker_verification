package main

import (
	"net/http"
	"time"
	"fmt"
	"log"
)
func TestRespond(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"第一次更新")
}
func main()  {
	routerlist := http.NewServeMux()
	addr := fmt.Sprintf(":%d",7986)
	s := &http.Server{
		Addr: addr,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	routerlist.HandleFunc("/test/", TestRespond)
	s.Handler = routerlist
	log.Fatal(s.ListenAndServe())
}
