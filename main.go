package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func randIndex() string {
	return `<!DOCTYPE html>
	<html lang="cn">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>CCTV</title>
	  <link href="https://vjs.zencdn.net/7.2.3/video-js.css" rel="stylesheet">
	  <script src="https://vjs.zencdn.net/7.2.3/video.js"></script>
	  <style type="text/css">
		body,html {margin: 0; padding: 0; overflow-y: hidden;}
		body,html, .hls-cctv-dimensions, .video-js {width: 100%%; height: 100%%;}
	  </style>
	</head>
	<body>
	  <!-- http://183.207.249.15/PLTV/3/224/3221225530/index.m3u8 -->
	  <video id='hls-cctv'  class="video-js vjs-default-skin" controls>
		<source type="application/x-mpegURL" src="/PLTV/3/224/3221225530/index.m3u8">
	  </video>
	  <script>
		var player = videojs('hls-cctv');
		player.play();
	  </script>
	</body>
	</html>`
}

func index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasPrefix(path, "/PLTV") {
		url, _ := url.Parse("http://183.207.249.15")
		proxy := httputil.NewSingleHostReverseProxy(url)
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host
		proxy.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, randIndex())
	}
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
