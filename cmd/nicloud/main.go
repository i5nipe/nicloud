package main

import (
	"flag"
	"fmt"
	"os/user"

	nicloud "github.com/i5nipe/nicloud/pkg"
	log "github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

var (
	company = flag.String("d", "", "Company name for brute force")
	// debugmode  = flag.Bool("b", false, "Debug mode")
	silentmode = flag.Bool("s", false, "Enable silent mode")
	threads    = flag.Int("t", 100, "Number of concurrent threads")
	wordlis    = flag.String("w", "~/.nipe/cloud.txt", "Path to the wordlist")

	test_aws = flag.Bool("aws", false, "Brute Force on Amazon Web Services")
	test_dos = flag.Bool("dos", false, "Brute Force on DigitalOcean Space")
	test_gcs = flag.Bool("gcs", false, "Brute Force on Google Cloud Storage")
)

func main() {
	flag.Parse()
	/*
		if *debugmode {
			log.DefaultLogger.SetMaxLevel(levels.LevelDebug)
		}
	*/
	if *silentmode {
		log.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	} else {
		nicloud.Banner()
	}
	if *company == "" {
		fmt.Println("Usage of nicloud:")
		flag.PrintDefaults()
		return
	}
	if *wordlis == "~/.nipe/cloud.txt" {
		user, err := user.Current()
		if err != nil {
			log.Fatal().Msg(fmt.Sprintf("%s", err))
		}
		*wordlis = fmt.Sprintf("%s/.nipe/cloud.txt", user.HomeDir)
	}

	// No arguments
	if !*test_aws && !*test_gcs && !*test_dos {
		nicloud.BruteAWS(*company, *wordlis, *threads)
		nicloud.BruteGcs(*wordlis, *company, *threads)
		nicloud.BruteDOS(*company, *wordlis, *threads)
	}

	if *test_aws {
		nicloud.BruteAWS(*company, *wordlis, *threads)
	}
	if *test_gcs {
		nicloud.BruteGcs(*wordlis, *company, *threads)
	}
	if *test_dos {
		nicloud.BruteDOS(*company, *wordlis, *threads)
	}

}
