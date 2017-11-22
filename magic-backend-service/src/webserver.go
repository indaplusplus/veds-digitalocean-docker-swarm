package main

import (
    "github.com/gorilla/mux"
    "encoding/json"
    "net/http"
    "log"
)

type JSONResponce struct {
    URL string
    PubDate string
    Title string
    Description string
}

func RunWebserver() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/{site}/{request}", HandleRequest)
    log.Fatal(http.ListenAndServe(":7331", router))
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
    var URL string
    vars := mux.Vars(r)
    switch vars["site"] {
        case "dn":
            URL = "https://www.dn.se/nyheter/m/rss/"
        case "svd":
            URL = "https://www.svd.se/?service=rss"
        default:
            return
    }

    rss := &RSS{URL: URL}
    rss_root := ParseRSS(rss.Read())
    switch vars["request"] {
        case "recent":
            json.NewEncoder(w).Encode(rss_root.Channel.Items[0])
        case "all":
            json.NewEncoder(w).Encode(rss_root.Channel.Items)
    }
}
