package bootstrap

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func AutoloadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	s, e := readln(reader)
	for e == nil {
		loadConfigFromEnv(s)
		s, e = readln(reader)
	}
}

func readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func loadConfigFromEnv(s string) {
	var conf = strings.Split(s, "=")
	if len(conf) > 1 {
		file, err := os.OpenFile("conf/app.conf", os.O_APPEND|os.O_RDWR, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		switch conf[0] {
		case "APP_NAME":
			writeConfFile(file, "appname", conf[1])
			return
		case "HTTP_PORT":
			writeConfFile(file, "httpport", conf[1])
			return
		case "RUN_MODE":
			writeConfFile(file, "runmode", conf[1])
			return
		case "ENABLE_XSRF":
			writeConfFile(file, "EnableXSRF", conf[1])
			return
		case "XSRF_KEY":
			writeConfFile(file, "XSRFKey", conf[1])
			return
		case "XSRF_EXPIRE":
			writeConfFile(file, "XSRFExpire", conf[1])
			return
		case "SESSION_ON":
			writeConfFile(file, "sessionon", conf[1])
			return
		case "AUTO_RENDER":
			writeConfFile(file, "autorender", conf[1])
			return
		case "RECOVER_PANNIC":
			writeConfFile(file, "recoverpanic", conf[1])
			return
		case "VIEW_PATH":
			writeConfFile(file, "viewspath", conf[1])
			return
		default:
			writeConfFile(file, strings.ToLower(conf[0]), conf[1])
			return
		}
	}
}

func writeConfFile(file *os.File, key, value string) {
	r := bufio.NewReader(file)
	s, e := readln(r)
	for e == nil {
		if appConf := strings.Split(s, "="); strings.TrimSpace(appConf[0]) == key {
			return
		}
		s, e = readln(r)
	}
	file.WriteString(key + " = " + strings.TrimSpace(value) + "\n")
}
