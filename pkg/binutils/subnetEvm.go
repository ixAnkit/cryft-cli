// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package binutils

import (
	"path/filepath"

	"github.com/ixAnkit/cryft-cli/pkg/application"
	"github.com/ixAnkit/cryft-cli/pkg/constants"
)

func SetupSubnetEVM(app *application.Avalanche, subnetEVMVersion string) (string, error) {
	// Check if already installed
	binDir := app.GetSubnetEVMBinDir()
	subDir := filepath.Join(binDir, subnetEVMBinPrefix+subnetEVMVersion)

	installer := NewInstaller()
	downloader := NewSubnetEVMDownloader()
	vmDir, err := InstallBinary(
		app,
		subnetEVMVersion,
		binDir,
		subDir,
		subnetEVMBinPrefix,
		constants.AvaLabsOrg,
		constants.SubnetEVMRepoName,
		downloader,
		installer,
	)
	return filepath.Join(vmDir, constants.SubnetEVMBin), err
}
