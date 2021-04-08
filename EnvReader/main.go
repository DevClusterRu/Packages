package EnvReader

import (
	"github.com/kardianos/osext"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var Env map[string]string

func ConfigInit() map[string]string {

	Env = make(map[string]string)
	folderPath, err := osext.ExecutableFolder()

	if err != nil {
		log.Println("Cant find folder path", err)
	}

	file, err := os.Open(folderPath + `/config.conf`)

	if err != nil {
		file, err = os.Open("config.conf")
	}

	if err != nil {
		log.Println("Cant open config file", err)
		os.Exit(0)
	}

	b, _ := ioutil.ReadAll(file)
	rows := strings.Split(string(b), "\n")
	for _, v := range rows {
		if v == "" || strings.Index(v, "#") == 0 {
			continue
		}

		pair := strings.Split(v, "=")
		if len(pair) == 2 {
			pair[0] = strings.TrimSpace(pair[0])
			pair[1] = strings.TrimSpace(pair[1])
			Env[pair[0]] = pair[1]
		}
	}
}
