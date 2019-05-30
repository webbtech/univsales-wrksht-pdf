package config

// defaults struct
type defaults struct {
	AWSRegion  string `yaml:"AWSRegion"`
	DBHost     string `yaml:"DBHost"`
	DBName     string `yaml:"DBName"`
	DBPassword string `yaml:"DBPassword"`
	DBUser     string `yaml:"DBUser"`
	DocAuthor  string `yaml:"DocAuthor"`
	LogoURI    string `yaml:"LogoURI"`
	S3Bucket   string `yaml:"S3Bucket"`
	SsmPath    string `yaml:"SsmPath"`
	Stage      string `yaml:"Stage"`
}

type config struct {
	AWSRegion    string
	DBConnectURL string
	DBName       string
	DocAuthor    string
	LogoURI      string
	S3Bucket     string
	Stage        StageEnvironment
}
