package types

import (
	"fmt"
)

const (
	// ModuleName defines the interchain accounts module name
	ModuleName = "interchainaccounts"

	// PortID is the default port id that the interchain accounts host submodule binds to
	PortID = "interchain-account"

	// PortPrefix is the default port prefix that the interchain accounts controller submodule binds to
	PortPrefix = "ics27-"

	// Version defines the current version for interchain accounts
	Version = "ics27-1"

	// StoreKey is the store key string for interchain accounts
	StoreKey = ModuleName

	// RouterKey is the message route for interchain accounts
	RouterKey = ModuleName

	// QuerierRoute is the querier route for interchain accounts
	QuerierRoute = ModuleName
)

var (
	// ActiveChannelKeyPrefix defines the key prefix used to store active channels
	ActiveChannelKeyPrefix = "activeChannel"

	// OwnerKeyPrefix defines the key prefix used to store interchain accounts
	OwnerKeyPrefix = "owner"

	// PortKeyPrefix defines the key prefix used to store ports
	PortKeyPrefix = "port"
)

// KeyActiveChannel creates and returns a new key used for active channels store operations
func KeyActiveChannel(connectionID, portID string) []byte {
	return []byte(fmt.Sprintf("%s/%s/%s", ActiveChannelKeyPrefix, connectionID, portID))
}

// KeyOwnerAccount creates and returns a new key used for interchain account store operations
func KeyOwnerAccount(connectionID, portID string) []byte {
	return []byte(fmt.Sprintf("%s/%s/%s", OwnerKeyPrefix, connectionID, portID))
}

// KeyPort creates and returns a new key used for port store operations
func KeyPort(portID string) []byte {
	return []byte(fmt.Sprintf("%s/%s", PortKeyPrefix, portID))
}
