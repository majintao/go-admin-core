package config

type Grpc struct {
	Registry Registry   // 注册etcd源用的
	Clients map[string]*GrpcClient // 注册不同的rpc服务
}

type Registry struct {
	ConnectTimeout int
	Endpoints []string
	Secure bool
	Prefix string
}


type GrpcClient struct {
	Address string
	BalancerName string
	Block bool
	DialTimeout int
}

// GrpcConfig 配置
var GrpcConfig = new(Grpc)
