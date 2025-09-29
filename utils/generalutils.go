package utils

import (
	"math/rand"
	"regexp"

	persistence "github.com/bharathreddy-97/URLShortner/Persistence"
)

func CreateShortURL(url string, dataMap *persistence.SMap) string {
	if shortURLFound := dataMap.Get(url); shortURLFound != "" {
		return shortURLFound
	}

	randomShortURL := randomCode(5)
	isDuplicate := dataMap.CheckForId(randomShortURL)
	for isDuplicate {
		randomShortURL = randomCode(5)
		isDuplicate = dataMap.CheckForId(randomShortURL)
	}
	dataMap.Set(url, randomShortURL)

	return randomShortURL
}

func randomCode(n int) string {
	output := ""
	for range n {
		output = output + string(allowedLetters[rand.Intn(len(allowedLetters))])
	}

	return output
}

func GetMetrics(dataMap *persistence.SMap) map[string]int64 {
	hostLinksMap := make(map[string]int64)
	return hostLinksMap
}

func CheckForValidShortURL(id string) bool {
	re := regexp.MustCompile(`[a-zA-Z0-9]{5}`)
	return re.MatchString(id)
}
