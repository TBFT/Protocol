// Copyright (c) 2018 NEC Laboratories Europe GmbH.
//
// Authors: Sergey Fedorov <sergey.fedorov@neclab.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package sgx provides Go interface to SGX USIG implementation.
package sgx

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/binary"
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger-labs/minbft/usig"
)

// USIG implements USIG interface around USIGEnclave.
type USIG struct {
	*USIGEnclave
}

var _ usig.USIG = new(USIG)

// New creates a new instance of SGXUSIG. It is a wrapper around
// NewUSIGEnclave(). See NewUSIGEnclave() for more details. Note that
// the created instance has to be disposed with Destroy() method, e.g.
// using defer.
func New(enclaveFile string, sealedKey []byte, key []byte) (*USIG, error) {
	enclave, err := NewUSIGEnclave(enclaveFile, sealedKey, key)
	if err != nil {
		return nil, fmt.Errorf("failed to create USIG enclave: %v", err)
	}
	return &USIG{enclave}, nil
}

// CreateUI creates a unique identifier assigned to the message.
func (u *USIG) CreateUI(message []byte) (*usig.UI, error) {
	counter, signature, encryptedShares, encryptedSecretH, err := u.USIGEnclave.CreateUI(messageDigest(message))
	if err != nil {
		return nil, err
	}
	return &usig.UI{
		Counter: counter,
		Cert:    MakeCert(u.epoch, signature, encryptedShares, encryptedSecretH),
	}, nil
}

// ID returns the USIG instance identity.
func (u *USIG) ID() []byte {
	id, err := MakeID(u.Epoch(), u.PublicKey())
	if err != nil {
		panic(err)
	}
	return id
}

func (u *USIG) VerifyUI(message []byte, ui *usig.UI, usigID []byte) ([]byte,
	error) {
	if message == nil || ui == nil || usigID == nil || ui.Cert == nil {
		log.Println("Error in SGX! ui is nil")
	}

	_, signature, encryptedShares, encryptedSecretH, err := ParseCert(ui.Cert)
	if err != nil {
		return nil, fmt.Errorf("failed to parse UI cert: %s", err)
	}
	if signature == nil || encryptedShares == nil || encryptedSecretH == nil {
		log.Println("Error in SGX! something is nil!")
	}

	shares, secretHash, err := u.USIGEnclave.VerifyUI(messageDigest(message), signature, encryptedShares, encryptedSecretH)

	if err != nil {
		return nil, fmt.Errorf("invalid UI: %s", err)
	}
	bb, er := SharesSecretHashToBytes(shares, secretHash)
	if shares == nil {
		log.Println("Error in SGX! something is nil")
	}
	if er != nil {
		return bb, er
	}
	return bb, nil

}
func SharesSecretHashToBytes(share []byte, secretHash Digest) ([]byte, error) {
	type verifier struct {
		S  []byte
		SH []byte
	}
	v := verifier{
		S:  share,
		SH: secretHash[:],
	}
	ret, err := asn1.Marshal(v)
	if err != nil {
		panic(err)
		return nil, err
	}
	return ret, nil
}

func messageDigest(message []byte) Digest {
	return sha256.Sum256(message)
}

// MakeID composes a USIG identity which is 64-bit big-endian encoded
// epoch value followed by public key serialized in PKIX format.
func MakeID(epoch uint64, publicKey interface{}) ([]byte, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize public key: %s", err)
	}

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, epoch); err != nil {
		panic(err)
	}
	if err := binary.Write(buf, binary.BigEndian, publicKeyBytes); err != nil {
		panic(err)
	}

	return buf.Bytes(), nil
}

// ParseID breaks a USIG identity down to epoch value and public key.
func ParseID(usigID []byte) (epoch uint64, pubKey crypto.PublicKey, err error) {
	buf := bytes.NewBuffer(usigID)

	err = binary.Read(buf, binary.BigEndian, &epoch)
	if err != nil {
		return uint64(0), nil, fmt.Errorf("failed to extract epoch from USIG ID: %s", err)
	}

	pubKey, err = x509.ParsePKIXPublicKey(buf.Bytes())
	if err != nil {
		return uint64(0), nil, fmt.Errorf("failed to parse public key: %s", err)
	}

	return epoch, pubKey, err
}

// MakeCert composes a USIG certificate which is 64-bit big-endian
// encoded epoch value followed by serialized USIG signature.
func MakeCert(epoch uint64, signature []byte, encryptedShares []byte, encryptedSecretH []byte) []byte {
	if signature == nil || encryptedShares == nil || encryptedSecretH == nil {
		return nil
	}
	type certmaker struct {
		E   int
		S   []byte
		ES  []byte
		ESH []byte
	}
	intNum, _ := strconv.Atoi(strconv.FormatUint(epoch, 10))
	c := certmaker{
		E:   intNum,
		S:   signature,
		ES:  encryptedShares,
		ESH: encryptedSecretH,
	}
	ret, err := asn1.Marshal(c)
	if err != nil {
		panic(err)
	}
	return ret
}

// ParseCert breaks a USIG certificate down to epoch value and
// serialized USIG signature.
func ParseCert(cert []byte) (epoch uint64, signature []byte, encryptedShares []byte, encryptedSecretH []byte, err error) {
	type certparser struct {
		E   int
		S   []byte
		ES  []byte
		ESH []byte
	}
	retStruct := new(certparser)

	ret, err := asn1.Unmarshal(cert, retStruct)
	if err != nil || ret == nil {
		panic(err)
	}
	return uint64(retStruct.E), retStruct.S, retStruct.ES, retStruct.ESH, err
}
