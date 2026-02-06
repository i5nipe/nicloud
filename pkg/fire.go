package nicloud

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
	log "github.com/projectdiscovery/gologger"
)

const (
	baseUrlFIRE = ".firebaseio.com"
)

func BruteFire(company, filename string, threads int){
	urls, err := gerarLista(filename, company, "fire")
	if err != nil {
		return
	}
	log.Info().Msgf("%d Generated URLs for %s (FIRE)", len(urls)*5, Magenta("RealTime Database Firebase").Bold())
	results := make(chan Results, threads)
	domain := make(chan string, threads)
	regions := []string{"-default-rtb"}

	for w := 1; w < threads; w++ {
		go getAWS(domain, results, client)
	}

	go func() {
		for _, aaa := range urls {
			for _, region := range regions {
				domain <- fmt.Sprintf("https://%s.%s%s/users.json", aaa, region, baseUrlFIRE)
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
