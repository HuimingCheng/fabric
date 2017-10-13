/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"github.com/pkg/errors"
)

type MSPVersion int

const (
	MSPv1_0 = iota
	MSPv1_1
)

// NewOpts represent
type NewOpts interface {
	// GetVersion returns the MSP's version to be instantiated
	GetVersion() MSPVersion
}

// NewBaseOpts is the default base type for all MSP instantiation Opts
type NewBaseOpts struct {
	Version MSPVersion
}

func (o *NewBaseOpts) GetVersion() MSPVersion {
	return o.Version
}

// BCCSPNewOpts contains the options to instantiate a new BCCSP-based (X509) MSP
type BCCSPNewOpts struct {
	NewBaseOpts
}

// IdemixNewOpts contains the options to instantiate a new Idemix-based MSP
type IdemixNewOpts struct {
	NewBaseOpts
}

// New create a new MSP instance depending on the passed Opts
func New(opts NewOpts) (MSP, error) {
	switch opts.(type) {
	case *BCCSPNewOpts:
		switch opts.GetVersion() {
		case MSPv1_0:
			return newBccspMsp()
		default:
			return nil, errors.Errorf("Invalid *BCCSPNewOpts. Version not recognized [%v]", opts.GetVersion())
		}
	case *IdemixNewOpts:
		switch opts.GetVersion() {
		case MSPv1_0:
			return newIdemixMsp()
		default:
			return nil, errors.Errorf("Invalid *IdemixNewOpts. Version not recognized [%v]", opts.GetVersion())
		}
	default:
		return nil, errors.Errorf("Invalid msp.NewOpts instance. It must be either *BCCSPNewOpts or *IdemixNewOpts. It was [%v]", opts)
	}
}