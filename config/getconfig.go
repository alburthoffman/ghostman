package config

func GetDefaults() *Defaults {
	envDefaults := GetDefaultsFromEnv()
	configDefaults := &Defaults{}
	if envDefaults.Rcconf != "" {
		configDefaults = GetDefaultsFromConfig(envDefaults.Rcconf)
	}

	if envDefaults.TmplDir == "" {
		envDefaults.TmplDir = configDefaults.TmplDir
	}
	if envDefaults.MaxConcur == 0 {
		envDefaults.MaxConcur = configDefaults.MaxConcur
	}
	return envDefaults
}
