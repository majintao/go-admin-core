package config

type MongoSource struct {
	MongoConfigs map[string]*MongoConfig
}

type MongoConfig struct {
	Dsn string
	PoolLimit int
	SocketTimeout int
}

var MongoSourceObj = new(MongoSource)
