package nicloud

import (
	"bufio"
	"fmt"
	"os"
)

func gerarLista(file, company, mode string) ([]string, error) {
	wordlist, err := os.Open(file)
	if err != nil {
		return []string{""}, err
	}
	defer wordlist.Close()

	var newword []string
	newword = append(newword, fmt.Sprintf("%s", company))
	scanner := bufio.NewScanner(wordlist)

	separ := []string{".", "-", "_", ""}
	if mode == "gct" {
		separ = []string{"-", "_", ""}

	} else {
		separ = []string{".", "-", "_", ""}

	}
	for scanner.Scan() {
		for _, separador := range separ {
			newword = append(newword, fmt.Sprintf("%s%s%s", company, separador, scanner.Text()))
			newword = append(newword, fmt.Sprintf("%s%s%s", scanner.Text(), separador, company))
		}
	}
	return newword, nil
}
