package server

type Options struct {
	Url               string        `json:"url" option:"Service Endpoint Url"`
	//Port              int           `json:"port"`
	//Trace              bool          `json:"-"`
	//Debug              bool          `json:"-"`
	//NoLog              bool          `json:"-"`
	//NoSigs             bool          `json:"-"`
	//Logtime            bool          `json:"-"`
	//MaxConn            int           `json:"max_connections"`
	//Users              []*User       `json:"-"`
	//Username           string        `json:"-"`ç
	//Password           string        `json:"-"`
	//Authorization      string        `json:"-"`
	//PingInterval       time.Duration `json:"ping_interval"`
	//MaxPingsOut        int           `json:"ping_max"`
	//HTTPHost           string        `json:"http_host"`
	//HTTPPort           int           `json:"http_port"`
	//HTTPSPort          int           `json:"https_port"`
	//AuthTimeout        float64       `json:"auth_timeout"`
	//MaxControlLine     int           `json:"max_control_line"`
	//MaxPayload         int           `json:"max_payload"`
	//MaxPending         int           `json:"max_pending_size"`
	//ClusterHost        string        `json:"addr"`
	//ClusterPort        int           `json:"cluster_port"`
	//ClusterUsername    string        `json:"-"`
	//ClusterPassword    string        `json:"-"`
	//ClusterAuthTimeout float64       `json:"auth_timeout"`
	//ClusterTLSTimeout  float64       `json:"-"`
	//ClusterTLSConfig   *tls.Config   `json:"-"`
	//ClusterListenStr   string        `json:"-"`
	//ClusterNoAdvertise bool          `json:"-"`
	//ProfPort           int           `json:"-"`
	//PidFile            string        `json:"-"`
	//LogFile            string        `json:"-"`
	//Syslog             bool          `json:"-"`
	//RemoteSyslog       string        `json:"-"`
	//Routes             []*url.URL    `json:"-"`
	//RoutesStr          string        `json:"-"`
	//TLSTimeout         float64       `json:"tls_timeout"`
	//TLS                bool          `json:"-"`
	//TLSVerify          bool          `json:"-"`
	//TLSCert            string        `json:"-"`
	//TLSKey             string        `json:"-"`
	//TLSCaCert          string        `json:"-"`
	//TLSConfig          *tls.Config   `json:"-"`
}