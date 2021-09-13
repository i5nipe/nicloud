package nicloud

import (
	. "github.com/logrusorgru/aurora/v3"
	"github.com/projectdiscovery/gologger"
)

const banner = ` ███▄    █  ██▓ ▄████▄  ██▓     ▒█████   █    ██▓█████▄ 
 ██ ▀█   █ ▓██▒▒██▀ ▀█ ▓██▒    ▒██▒  ██▒ ██  ▓██▒██▀ ██▌
▓██  ▀█ ██▒▒██▒▒▓█    ▄▒██░    ▒██░  ██▒▓██  ▒██░██   █▌
▓██▒  ▐▌██▒░██░▒▓▓▄ ▄██▒██░    ▒██   ██░▓▓█  ░██░▓█▄   ▌
▒██░   ▓██░░██░▒ ▓███▀ ░██████▒░ ████▓▒░▒▒█████▓░▒████▓ 
░ ▒░   ▒ ▒ ░▓  ░ ░▒ ▒  ░ ▒░▓  ░░ ▒░▒░▒░ ░▒▓▒ ▒ ▒ ▒▒▓  ▒ 
░ ░░   ░ ▒░ ▒ ░  ░  ▒  ░ ░ ▒  ░  ░ ▒ ▒░ ░░▒░ ░ ░ ░ ▒  ▒ 
   ░   ░ ░  ▒ ░░         ░ ░   ░ ░ ░ ▒   ░░░ ░ ░ ░ ░  ░ 
         ░  ░  ░ ░         ░  ░    ░ ░     ░       ░    
               ░                                 ░      
`

// Version is the current version of nicloud
const Version = `v1.0.6`

// showBanner is used to show the banner to the user
func Banner() {
	gologger.Print().Msgf("%s", Magenta(banner))
	gologger.Print().Msgf("\t\t\t\t\tMade with %s by nipe\n", Red("<3"))

}
