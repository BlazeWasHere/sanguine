// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { NotaryRegistry } from "../../contracts/NotaryRegistry.sol";

contract NotaryRegistryHarness is NotaryRegistry {
    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }

    function addNotary(uint32 _domain, address _notary) public returns (bool) {
        return _addNotary(_domain, _notary);
    }

    function removeNotary(uint32 _domain, address _notary) public returns (bool) {
        return _removeNotary(_domain, _notary);
    }
}