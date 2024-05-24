package verifier

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Verifier(prx string) {

	// proxy
	proxyUrl, err := url.Parse(prx)
	if err != nil {
		log.Println(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 6,
	}

	t0 := time.Now()

	// req, err := http.NewRequest("GET", "http://checkip.amazonaws.com", nil)
	// req, err := http.NewRequest("GET", "http://www.httpbin.org/ip", nil)
	req, err := http.NewRequest("GET", "http://httpforever.com/", nil)
	if err != nil {
		log.Println(err)
	}

	// Fazer a solicitação HTTP com o cliente personalizado
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	// Adicionar headers à solicitação
	// req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	// req.Header.Add("Accept-Encoding", "gzip, deflate")
	// req.Header.Add("Accept-Language", "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7")
	// req.Header.Add("Cache-Control", "max-age=0")
	// req.Header.Add("Connection", "keep-alive")
	// req.Header.Add("Host", "httpforever.com")
	// req.Header.Add("If-Modified-Since", "Wed, 22 Mar 2023 14:54:48 GMT")
	// req.Header.Add("If-None-Match", `W/"641b16b8-1404"`)
	// req.Header.Add("Upgrade-Insecure-Requests", "1")
	// req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	t1 := time.Now()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s", body)
	fmt.Println(res.Status, res.Proto, t1.Sub(t0), proxyUrl)

}
