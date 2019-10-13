package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/dexidp/dex/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

func newDexClient() (api.DexClient, error) {
	host := viper.GetString("host")
	port := viper.GetString("port")
	caPath := viper.GetString("client-ca")
	certPath := viper.GetString("client-crt")
	keyPath := viper.GetString("client-key")
	cPool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, fmt.Errorf("invalid CA crt file: %s", caPath)
	}
	if cPool.AppendCertsFromPEM(caCert) != true {
		return nil, fmt.Errorf("failed to parse CA crt")
	}

	clientCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf("invalid client crt file: %s", caPath)
	}

	clientTLSConfig := &tls.Config{
		RootCAs:            cPool,
		Certificates:       []tls.Certificate{clientCert},
		InsecureSkipVerify: true,
	}
	creds := credentials.NewTLS(clientTLSConfig)

	conn, err := grpc.Dial(host + ":" + port, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("dial: %v", err)
	}
	return api.NewDexClient(conn), nil
}