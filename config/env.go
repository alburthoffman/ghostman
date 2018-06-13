package config

import (
	"os"
	"strconv"
)

const (
	Rcconf    = "GHOSTMAN_CONF"
	Tmpldir   = "GHOSTMAN_TMPLDIR"
	Maxconcur = "GHOSTMAN_MAXCONCUR"
)

func GetDefaultsFromEnv() *Defaults {
	defaults := Defaults{}

	defaults.Rcconf = os.Getenv(Rcconf)
	defaults.TmplDir = os.Getenv(Tmpldir)

	smaxconcur := os.Getenv(Maxconcur)
	if imaxconcur, err := strconv.Atoi(smaxconcur); err != nil {
		log.Errorf("invalid env var %s", Maxconcur)
	} else {
		defaults.MaxConcur = imaxconcur
	}

	log.Debugf("load defaults from env: %v", defaults)
	return &defaults
}
