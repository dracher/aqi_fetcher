package libs

import (
	"net/http"

	"io/ioutil"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

func fetchURL(url string, ch chan<- []byte) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		ch <- []byte(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		ch <- []byte(err.Error())
		return
	}
	ch <- body
}

// FetchAqiData is
func FetchAqiData(urls []string) []AqiData {
	ch := make(chan []byte)
	res := []AqiData{}
	// lock := &sync.Mutex{}
	for _, url := range urls {
		go fetchURL(url, ch)
	}
	for range urls {
		tmp := AqiData{}
		json.Unmarshal(<-ch, &tmp)
		// lock.Lock()
		res = append(res, tmp)
		// lock.Unlock()
	}
	return res
}
