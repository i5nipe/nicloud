package nicloud

import (
	"fmt"

	. "github.com/logrusorgru/aurora/v3"
	log "github.com/projectdiscovery/gologger"
)

const (
	baseUrlGCS = "https://www.googleapis.com/storage/v1/b/"
)

func BruteGcs(file, company string, threads int) {
	perm, err := gerarLista(file, company, "gct")
	if err != nil {
		return
	}

	log.Info().Msgf("%d Generated URLs for %s (GCS)", len(perm), Magenta("Google Cloud Storage").Bold())

	urls := make(chan string, threads)
	resu := make(chan Results, threads)

	for c := 0; c < threads; c++ {
		go getAWS(urls, resu, client)
	}

	go func() {
		for _, aaa := range perm {
			urls <- fmt.Sprintf("%s%s", baseUrlGCS, aaa)
		}
	}()

	for i := 0; i < len(perm); i++ {
		resp := <-resu
		switch resp.StatusCode {
		case 0:
			log.Debug().Str("url", resp.Url).Msg(fmt.Sprintf("%v\n", Red(0)))
		case 404:
			log.Debug().Str("url", resp.Url).Msg(fmt.Sprintf("%v\n", Red(404)))
		case 400:
			log.Debug().Str("url", resp.Url).Msg(fmt.Sprintf("%v\n", Red(400)))
		case 401:
			log.Debug().Str("url", resp.Url).Msg(fmt.Sprintf("%v\n", Red(401)))
		case 403:
			fmt.Printf("%s\t%d\n", resp.Url, Magenta(resp.StatusCode))
		case 200:
			fmt.Printf("%s\t%d\n", resp.Url, Green(resp.StatusCode))
		default:
			fmt.Printf("%s\t%d\n", resp.Url, Yellow(resp.StatusCode))
		}
	}

	defer close(urls)
	defer close(resu)
}
