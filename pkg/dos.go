package nicloud

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
	log "github.com/projectdiscovery/gologger"
)

const (
	baseUrlDOS = ".digitaloceanspaces.com"
)

func BruteDOS(company, filename string, threads int) {
	urls, err := gerarLista(filename, company, "dos")
	if err != nil {
		return
	}
	log.Info().Msgf("%d Generated URLs for %s (DOS)", len(urls)*5, Magenta("DigitalOcean Spaces").Bold())

	results := make(chan Results, threads)
	domain := make(chan string, threads)
	regions := []string{"nyc3", "ams3", "sgp1", "sfo2", "fra1"}

	for w := 1; w < threads; w++ {
		go getAWS(domain, results, client)
	}

	go func() {
		for _, aaa := range urls {
			for _, region := range regions {
				domain <- fmt.Sprintf("https://%s.%s%s/", aaa, region, baseUrlDOS)
			}
		}
	}()

	for i := 0; i < len(urls)*5; i++ {
		resp := <-results
		switch resp.StatusCode {
		case 0:
			log.Debug().Str("url", resp.Url).Msg(fmt.Sprintf("%v\n", Red(404)))
		case 404:
			log.Debug().Str("url", resp.Url).Msg(fmt.Sprintf("%v\n", Red(404)))
		case 403:
			fmt.Printf("%s\t%d\n", resp.Url, Magenta(resp.StatusCode))
		case 200:
			fmt.Printf("%s\t%d\n", resp.Url, Green(resp.StatusCode))
		default:
			fmt.Printf("%s\t%d\n", resp.Url, Yellow(resp.StatusCode))
		}
	}

	defer close(domain)
	defer close(results)
}
