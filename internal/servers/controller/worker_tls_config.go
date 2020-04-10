package controller

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	wrapping "github.com/hashicorp/go-kms-wrapping"
	"github.com/hashicorp/watchtower/internal/cmd/base"
	"google.golang.org/protobuf/proto"
)

func (c Controller) validateWorkerTLS(hello *tls.ClientHelloInfo) (*tls.Config, error) {
	for _, p := range hello.SupportedProtos {
		switch {
		case strings.HasPrefix(p, "v1workerauth-"):
			return c.v1WorkerAuthConfig(hello.SupportedProtos)
		}
	}
	return nil, nil
}

func (c Controller) v1WorkerAuthConfig(protos []string) (*tls.Config, error) {
	var firstMatchProto string
	var encString string
	for _, p := range protos {
		if strings.HasPrefix(p, "v1workerauth-") {
			// Strip that and the number
			encString += strings.TrimPrefix(p, "v1workerauth-")[3:]
			if firstMatchProto == "" {
				firstMatchProto = p
			}
		}
	}
	if firstMatchProto == "" {
		return nil, errors.New("no matching proto found")
	}
	marshaledEncInfo, err := base64.RawStdEncoding.DecodeString(encString)
	if err != nil {
		return nil, err
	}
	encInfo := new(wrapping.EncryptedBlobInfo)
	if err := proto.Unmarshal(marshaledEncInfo, encInfo); err != nil {
		return nil, err
	}
	marshaledInfo, err := c.conf.WorkerAuthKMS.Decrypt(context.Background(), encInfo, nil)
	if err != nil {
		return nil, err
	}
	info := new(base.WorkerAuthCertInfo)
	if err := json.Unmarshal(marshaledInfo, info); err != nil {
		return nil, err
	}

	rootCAs := x509.NewCertPool()
	if ok := rootCAs.AppendCertsFromPEM(info.CACertPEM); !ok {
		return nil, errors.New("unable to add ca cert to cert pool")
	}
	tlsCert, err := tls.X509KeyPair(info.CertPEM, info.KeyPEM)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		ClientCAs:    rootCAs,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		NextProtos:   []string{firstMatchProto},
		MinVersion:   tls.VersionTLS13,
	}
	tlsConfig.BuildNameToCertificate()

	return tlsConfig, nil
}