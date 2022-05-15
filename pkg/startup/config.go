package startup

type Config struct {
	DBpath      string `json:"dbpath"`
	Servicepath string `json:"servicepath"`
	Logfilename string `json:"logfilename"`
}