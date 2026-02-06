package nicloud

import (
	"crypto/tls"
	"fmt"
	"os/exec"
	"strings"

	. "github.com/logrusorgru/aurora/v3"
	log "github.com/projectdiscovery/gologger"
	"github.com/valyala/fasthttp"
)

const (
	baseUrlAWS = ".s3.amazonaws.com/"
)

var client = fasthttp.Client{TLSConfig: &tls.Config{InsecureSkipVerify: true}}

type Results struct {
	Url              string
	StatusCode       int
	IsPermissionDenied bool
}

func BruteAWS(company, filename string, threads int) {
	urls, err := gerarLista(filename, company, "aws")
	if err != nil {
		return
	}
	log.Info().Msgf("%d Generated URLs for %s (AWS)", len(urls), Magenta("Amazon Web Service").Bold())

	results := make(chan Results, threads)
	domain := make(chan string, threads)

	for w := 1; w < threads; w++ {
		go getAWS(domain, results, client, "aws")
	}

	go func() {
		for _, aaa := range urls {
			domain <- fmt.Sprintf("https://%s%s", aaa, baseUrlAWS)
		}
	}()

	for i := 0; i < len(urls); i++ {
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
			test_aws_list(resp)
		default:
			fmt.Printf("%s\t%d\n", resp.Url, Yellow(resp.StatusCode))
		}
	}

	defer close(domain)
	defer close(results)

}
func Check_s3(bucket string) {
	fmt.Printf("%s Testing %s\n", Magenta("[*]"), Cyan(bucket))
	outl, _ := exec.Command("aws", "s3", "ls", bucket, "--no-sign-reques").Output()
	fmt.Printf("\t Directory listing:\n\t\t%v\n", string(outl))
}

func test_aws_list(resp Results) {
	bucket := strings.Replace(strings.Replace(resp.Url, "https://", "", -1), baseUrlAWS, "", -1)
	out, _ := exec.Command("aws", "s3", "ls", bucket, "--no-sign-reques").Output()
	fmt.Println(string(out))
}

func getAWS(url chan string, results chan Results, client fasthttp.Client, requestType string) {
	for j := range url {
		log.Debug().Str("URL", j).Msg("Making GET")
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()

		req.SetRequestURI(j)

		if err := client.Do(req, resp); err != nil {
			log.Debug().Msg(fmt.Sprintf("Error in get of %s\n", j))
			results <- Results{j, 0, false}
			continue
		}
		permissionDenied := false
		if requestType == "fire" && resp.StatusCode() == 403 {
			if string(resp.Body()) == "{\n  \"error\" : \"Permission denied\"\n}" {
				permissionDenied = true
			}
		}
		results <- Results{j, resp.StatusCode(), permissionDenied}

		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}
}
