package lists

import (
	"bufio"
	"log"
	"net/http"
)

func HandlerProxyscrape() []string {

	prx := getProxyscrape()

	// fmt.Println(prx)

	return prx

}

func getProxyscrape() []string {

	var prx []string

	url := "https://api.proxyscrape.com/v3/free-proxy-list/get?request=displayproxies&protocol=http&proxy_format=protocolipport&format=text&anonymity=Elite&timeout=1618"

	r, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	body := bufio.NewScanner(r.Body)
	body.Split(bufio.ScanLines)

	for body.Scan() {
		prx = append(prx, body.Text())
	}

	return prx

}
