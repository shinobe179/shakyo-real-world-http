// Use nameless function in http.HandleFunc
// Start listening http with http.ListenAndServe()
//
// cf: https://www.yoheim.net/blog.php?q=20170403

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	dump, err := httputil.DumpRequest(r, true)
//	if err != nil {
//		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
//		return
//	}
//	fmt.Println(string(dump))
//	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
//}

func main() {
	//var httpServer http.Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		fmt.Println(string(dump))
		fmt.Fprintf(w, "<html><body>hello</body></html>\n")
	})
	log.Println("start http listening :18888")
//	httpServer.Addr = ":18888"
//	log.Println(httpServer.ListenAndServe())
	http.ListenAndServe(":18888", nil)
}
