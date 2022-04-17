package conf

type MysqlConf struct {
	DriverName string `yaml:"driverName" json:"driver_name"`
	Username   string `yaml:"username" json:"username"`
	Password   string `yaml:"password" json:"password"`
	Host       string `yaml:"host" json:"host"`
	Port       string `yaml:"port" json:"port"`
	Database   string `yaml:"database" json:"database"`
	Charset    string `yaml:"charset" json:"charset"`
	ParseTime  string `yaml:"parseTime" json:"parse_time"`
	Loc        string `yaml:"loc" json:"loc"`
}

type MinioConf struct {
	Endpoint  string `yaml:"endpoint" json:"endpoint"`
	Access    string `yaml:"access" json:"access"`
	AccessKey string `yaml:"accessKey" json:"accessKey"`
	SecretKey string `yaml:"secretKey" json:"secretKey"`
}

type ServerConf struct {
	MysqlInfo MysqlConf `yaml:"mysql" json:"mysql"`
	MinioInfo MinioConf `yaml:"minio" json:"minio"`
}
