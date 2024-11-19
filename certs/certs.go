package certs

import (
	"crypto/tls"
	"crypto/x509"
	"errors"

	"google.golang.org/grpc/credentials"
)

// NewServerTransportCredentials creates and returns a new
// credentials.TransportCredentials using the given certificate information
// with a strong TLS server configuration.
func NewServerTransportCredentials(caCert, cert, key []byte) (credentials.TransportCredentials, error) {
	const isServer = true
	return newTransportCredentials(caCert, cert, key, isServer)
}

// NewClientTransportCredentials creates and returns a new
// credentials.TransportCredentials using the given certificate information
// with a strong TLS client configuration.
func NewClientTransportCredentials(caCert, cert, key []byte) (credentials.TransportCredentials, error) {
	const isServer = false
	return newTransportCredentials(caCert, cert, key, isServer)
}

func newTransportCredentials(caCert, cert, key []byte, isServer bool) (credentials.TransportCredentials, error) {
	certificate, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	capool := x509.NewCertPool()
	if !capool.AppendCertsFromPEM(caCert) {
		return nil, errors.New("cannot append ca cert to ca pool")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		MinVersion:   tls.VersionTLS13,
		CipherSuites: []uint16{
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
		},
	}

	if isServer {
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		tlsConfig.ClientCAs = capool
	} else {
		tlsConfig.RootCAs = capool
	}

	return credentials.NewTLS(tlsConfig), nil
}
