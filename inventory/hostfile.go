package inventory

import (
	"bufio"
	"github.com/newly12/ghostman/logging"
	"os"
)

var log = logging.GetLogger()

type Hosts []string

type Hostfile struct {
	Hostfile string
	Hosts
}

func (hf *Hostfile) GetAll() []string {
	return hf.Hosts
}

func LoadHostFile(f string) *Hostfile {
	hf := Hostfile{}
	hf.Hostfile = f

	file, err := os.Open(f)
	if err != nil {
		log.Criticalf("failed to load hosts from %s: %v", f, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		h := scanner.Text()
		hf.Hosts = append(hf.Hosts, h)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return &hf
}
