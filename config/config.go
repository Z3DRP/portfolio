package config

type Configurations struct {
	ZServer       ZServerConfig
	DatabaseStore DbStoreConfig
}

type ZServerConfig struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
	Static       string
}

type DbStoreConfig struct {
	DbName string
	DbUsr  string
	DbPwd  string
}
