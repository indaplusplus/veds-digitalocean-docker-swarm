package main

import (
    "log"
    "encoding/xml"
    "net/http"
    "io/ioutil"
)

type RSS struct {
    URL string
}

type RSSRoot struct {
    XMLName xml.Name `xml:"rss"`
    Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
    Items []RSSItem `xml:"item"`
}

type RSSItem struct {
    GUID string `xml:"guid"`
    URL string `xml:"url"`
    PubDate string `xml:"pubDate"`
    Title string `xml:"title"`
    Description string `xml:"description"`
    Category string `xml:"category"`
}

func (rss *RSS) Read() []byte {
    request, err := http.Get(rss.URL)
    if err != nil {
        log.Fatal(err)
    }
    defer request.Body.Close()
    text,_ := ioutil.ReadAll(request.Body)
    return text
}

func ParseRSS(rss_text []byte) RSSRoot {
    var q RSSRoot
    xml.Unmarshal(rss_text, &q)
    return q
}
