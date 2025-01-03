// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package subnet

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanche-cli/pkg/constants"
	"github.com/ava-labs/avalanche-cli/pkg/key"
	"github.com/ava-labs/avalanche-cli/pkg/models"
	"github.com/cryft-labs/cryftgo/ids"
	"github.com/cryft-labs/cryftgo/utils/formatting/address"
	"github.com/cryft-labs/cryftgo/vms/platformvm"
	"github.com/cryft-labs/cryftgo/vms/platformvm/txs"
	"github.com/cryft-labs/cryftgo/vms/secp256k1fx"
)

func GetOwners(network models.Network, subnetID ids.ID) ([]string, uint32, error) {
	var api string
	switch network {
	case models.Fuji:
		api = constants.FujiAPIEndpoint
	case models.Mainnet:
		api = constants.MainnetAPIEndpoint
	case models.Local:
		api = constants.LocalAPIEndpoint
	default:
		return nil, 0, fmt.Errorf("network not supported")
	}
	pClient := platformvm.NewClient(api)
	ctx := context.Background()
	txBytes, err := pClient.GetTx(ctx, subnetID)
	if err != nil {
		return nil, 0, fmt.Errorf("subnet tx %s query error: %w", subnetID, err)
	}
	var tx txs.Tx
	if _, err := txs.Codec.Unmarshal(txBytes, &tx); err != nil {
		return nil, 0, fmt.Errorf("couldn't unmarshal tx %s: %w", subnetID, err)
	}
	createSubnetTx, ok := tx.Unsigned.(*txs.CreateSubnetTx)
	if !ok {
		return nil, 0, fmt.Errorf("got unexpected type %T for subnet tx %s", tx.Unsigned, subnetID)
	}
	owner, ok := createSubnetTx.Owner.(*secp256k1fx.OutputOwners)
	if !ok {
		return nil, 0, fmt.Errorf("got unexpected type %T for subnet owners tx %s", createSubnetTx.Owner, subnetID)
	}
	controlKeys := owner.Addrs
	threshold := owner.Threshold
	networkID, err := network.NetworkID()
	if err != nil {
		return nil, 0, err
	}
	hrp := key.GetHRP(networkID)
	controlKeysStrs := []string{}
	for _, addr := range controlKeys {
		addrStr, err := address.Format("P", hrp, addr[:])
		if err != nil {
			return nil, 0, err
		}
		controlKeysStrs = append(controlKeysStrs, addrStr)
	}
	return controlKeysStrs, threshold, nil
}
