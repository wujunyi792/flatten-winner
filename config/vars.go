package config

type GlobalConfig struct {
	MODE        string
	ProgramName string
	AUTHOR      string
	VERSION     string
	PORT        string
	SQL         struct {
		Use    bool
		Config struct {
			TYPE     string
			IP       string
			PORT     string
			USER     string
			PASSWORD string
			DATABASE string
		}
	}
	REDIS struct {
		Use    bool
		Config struct {
			IP       string
			PORT     string
			PASSWORD string
			DB       int
		}
	}
	SERVICE struct {
		PASSKEY string
	}
}
