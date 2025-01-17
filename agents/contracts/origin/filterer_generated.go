// autogenerated file

package origin

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// IOriginFilterer ...
type IOriginFilterer interface {
	// FilterDispatched is a free log retrieval operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
	//
	// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
	FilterDispatched(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginDispatchedIterator, error)
	// WatchDispatched is a free log subscription operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
	//
	// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
	WatchDispatched(opts *bind.WatchOpts, sink chan<- *OriginDispatched, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error)
	// ParseDispatched is a log parse operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
	//
	// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
	ParseDispatched(log types.Log) (*OriginDispatched, error)
	// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
	//
	// Solidity: event Initialized(uint8 version)
	FilterInitialized(opts *bind.FilterOpts) (*OriginInitializedIterator, error)
	// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
	//
	// Solidity: event Initialized(uint8 version)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *OriginInitialized) (event.Subscription, error)
	// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
	//
	// Solidity: event Initialized(uint8 version)
	ParseInitialized(log types.Log) (*OriginInitialized, error)
	// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
	//
	// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OriginOwnershipTransferredIterator, error)
	// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
	//
	// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OriginOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
	//
	// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
	ParseOwnershipTransferred(log types.Log) (*OriginOwnershipTransferred, error)
	// FilterSent is a free log retrieval operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
	//
	// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
	FilterSent(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginSentIterator, error)
	// WatchSent is a free log subscription operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
	//
	// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
	WatchSent(opts *bind.WatchOpts, sink chan<- *OriginSent, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error)
	// ParseSent is a log parse operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
	//
	// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
	ParseSent(log types.Log) (*OriginSent, error)
	// FilterStateSaved is a free log retrieval operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
	//
	// Solidity: event StateSaved(bytes state)
	FilterStateSaved(opts *bind.FilterOpts) (*OriginStateSavedIterator, error)
	// WatchStateSaved is a free log subscription operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
	//
	// Solidity: event StateSaved(bytes state)
	WatchStateSaved(opts *bind.WatchOpts, sink chan<- *OriginStateSaved) (event.Subscription, error)
	// ParseStateSaved is a log parse operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
	//
	// Solidity: event StateSaved(bytes state)
	ParseStateSaved(log types.Log) (*OriginStateSaved, error)
}
