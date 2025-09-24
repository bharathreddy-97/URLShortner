package utils

import (
	"math/rand"
	"regexp"

	persistence "github.com/bharathreddy-97/URLShortner/Persistence"
)

func ReturnHost(url string) string {
	r := regexp.MustCompile(`/`)

	return r.Split(url, -1)[2]
}

func CreateShortURL(url string, dataMap *persistence.SMap) string {
	if shortURLFound := dataMap.Get(url); shortURLFound != "" {
		return shortURLFound
	}

	randomShortURL := randomCode(5)

	//TODO: should check if there is a duplicate short code....
	dataMap.Set(url, randomShortURL)

	return randomShortURL
}

func randomCode(n int) string {
	output := ""
	for i := 0; i < n; i++ {
		output = output + string(allowedLetters[rand.Intn(len(allowedLetters))])
	}

	return output
}

func GetMetrics(dataMap *persistence.SMap) map[string]int64 {
	hostLinksMap := make(map[string]int64)
	return hostLinksMap
}
