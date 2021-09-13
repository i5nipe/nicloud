package nicloud

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var (
	Bot     = "1718046352:AAGL0WvHYITBaAXRgauQ3bH3S-I-zIISSdA"
	Chat_id = "844132360"
)

func Deuboa(mensagem string) {
	params := url.Values{}
	params.Add("chat_id", Chat_id)
	params.Add("text", mensagem)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", Bot), body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
