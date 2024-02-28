package middleware

import (
	"google.golang.org/grpc/credentials"
)

// data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem")
func NewTLS(certFile string, keyFile string) credentials.TransportCredentials {
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		panic(err)
	}
	return creds
}
