// autogenerated file

package origin

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IOriginTransactor ...
type IOriginTransactor interface {
	// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
	//
	// Solidity: function dispatch(uint32 _destination, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
	Dispatch(opts *bind.TransactOpts, _destination uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error)
	// ImproperAttestation is a paid mutator transaction binding the contract method 0x0afe7f90.
	//
	// Solidity: function improperAttestation(bytes _attestation) returns(bool)
	ImproperAttestation(opts *bind.TransactOpts, _attestation []byte) (*types.Transaction, error)
	// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
	//
	// Solidity: function initialize(address _notaryManager) returns()
	Initialize(opts *bind.TransactOpts, _notaryManager common.Address) (*types.Transaction, error)
	// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
	//
	// Solidity: function renounceOwnership() returns()
	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)
	// SetNotary is a paid mutator transaction binding the contract method 0xa394a0e6.
	//
	// Solidity: function setNotary(address _notary) returns()
	SetNotary(opts *bind.TransactOpts, _notary common.Address) (*types.Transaction, error)
	// SetNotaryManager is a paid mutator transaction binding the contract method 0xa340abc1.
	//
	// Solidity: function setNotaryManager(address _notaryManager) returns()
	SetNotaryManager(opts *bind.TransactOpts, _notaryManager common.Address) (*types.Transaction, error)
	// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
	//
	// Solidity: function setSystemRouter(address _systemRouter) returns()
	SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error)
	// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
	//
	// Solidity: function transferOwnership(address newOwner) returns()
	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)
}