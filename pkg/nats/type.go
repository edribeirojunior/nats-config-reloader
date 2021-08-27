package nats

type NatsCluster struct {
	ApiVs    string `json:"apiVersion"`
	Kind     string `json:"Kind"`
	Metadata `json:"metadata"`
	Spec     `json:"spec"`
}

type NatsGetCluster struct {
	ApiVs    string `json:"apiVersion"`
	Kind     string `json:"Kind"`
	Metadata `json:"metadata"`
}
type NatsConfig struct {
	Debug bool `json:"debug"`
	Trace bool `json:"trace"`
}

type Pod struct {
	EnableCfgReload bool `json:"enableConfigReload"`
	EnableMetrics   bool `json:"enableMetrics"`
	Resources
}

type Resources struct {
	Limits struct {
		Cpu    string `json:"cpu"`
		Memory string `json:"memory"`
	} `json:"limits"`
	Requests struct {
		Cpu    string `json:"cpu"`
		Memory string `json:"memory"`
	} `json:"requests"`
}

type Size int

type Version string

type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type Spec struct {
	Auth
	GatewayConfig
	NatsConfig
	Pod
	Size    `json:"size"`
	Version `json:"version"`
}

type Auth struct {
	EnableSvcAccounts bool `json:"enableServiceAccounts"`
}

type Gateway struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}
type GatewayConfig struct {
	Gateways     []Gateway
	HostPort     int    `json:"hostPort"`
	Name         string `json:"name"`
	RejectUnkown bool   `json:"rejectUnkown"`
}
