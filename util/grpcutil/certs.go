package grpcutil

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

// NewTransportCredentials creates and returns a new
// credentials.TransportCredentials using the given certificate information
// with a strong TLS configuration.
func NewTransportCredentials(caCert, cert, key string) (credentials.TransportCredentials, error) {
	certificate, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(caCert)
	if err != nil {
		return nil, err
	}

	capool := x509.NewCertPool()
	if !capool.AppendCertsFromPEM(data) {
		return nil, fmt.Errorf("cannot append ca cert to ca pool")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    capool,
		MinVersion:   tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{
			tls.CurveP521,
			tls.CurveP384,
			tls.CurveP256,
		},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	return credentials.NewTLS(tlsConfig), nil
}
