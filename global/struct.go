package global

type DatabaseConfig struct {
	Database Dbconnect `yaml:"database"`
}

type Dbconnect struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
