package x509

/*
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include "openssl/x509.h"
#cgo pkg-config: openssl
*/
import "C"

type Certificate struct {
  ptr *C.X509
}

func ParseCertificate(pembuf []byte) (*Certificate, error) {
  return &Certificate{}, nil
}

func ParseCertificates(pembuf []byte) ([]*Certificate, error) {
  return []*Certificate{}, nil
}
