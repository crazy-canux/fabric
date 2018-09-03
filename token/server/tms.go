/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package server

import "github.com/hyperledger/fabric/protos/token"

//go:generate counterfeiter -o mock/issuer.go -fake-name Issuer . Issuer

// An Issuer creates token import requests.
type Issuer interface {
	// Issue creates an import request transaction.
	RequestImport(tokensToIssue []*token.TokenToIssue) (*token.TokenTransaction, error)
}

//go:generate counterfeiter -o mock/transactor.go -fake-name Transactor . Transactor

// Transactor allows to operate on issued tokens
type Transactor interface {
	// ListUnspentTokens returns a slice of unspent tokens owned by this transactor
	ListUnspentTokens() ([]*token.TokenOutput, error)
}

//go:generate counterfeiter -o mock/tms_manager.go -fake-name TMSManager . TMSManager

type TMSManager interface {
	// GetIssuer returns an Issuer bound to the passed channel and whose credential
	// is the tuple (privateCredential, publicCredential).
	GetIssuer(channel string, privateCredential, publicCredential []byte) (Issuer, error)

	// GetTransactor returns a Transactor bound to the passed channel and whose credential
	// is the tuple (privateCredential, publicCredential).
	GetTransactor(channel string, privateCredential, publicCredential []byte) (Transactor, error)
}
