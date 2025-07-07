package grpc

import (
	"fmt"
	"time"
)

type GRPCServerConfig struct {
	Host           string        `json:"host" yaml:"host"`
	Port           int           `json:"port" yaml:"port"`
	MaxRecvMsgSize int           `json:"max_recv_msg_size" yaml:"max_recv_msg_size"`
	MaxSendMsgSize int           `json:"max_send_msg_size" yaml:"max_send_msg_size"`
	Timeout        time.Duration `json:"timeout" yaml:"timeout"`
	EnableTLS      bool          `json:"enable_tls" yaml:"enable_tls"`
	TLSCertFile    string        `json:"tls_cert_file" yaml:"tls_cert_file"`
	TLSKeyFile     string        `json:"tls_key_file" yaml:"tls_key_file"`
}

func (c GRPCServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
