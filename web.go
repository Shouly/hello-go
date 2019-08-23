package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	for k, v := range r.Form {
		fmt.Println(k, " : ", strings.Join(v, ""))
	}
	_, _ = fmt.Fprint(w, "Hello World!")
}

func main() {

	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("listen server error!", err)
	}

}
