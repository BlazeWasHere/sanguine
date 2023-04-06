// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentManagerEvents} from "../events/AgentManagerEvents.sol";
import {IAgentManager} from "../interfaces/IAgentManager.sol";
import {ISystemRegistry} from "../interfaces/ISystemRegistry.sol";
import {AgentFlag, AgentStatus, SlashStatus} from "../libs/Structures.sol";
import {SystemContract} from "../system/SystemContract.sol";

abstract contract AgentManager is SystemContract, AgentManagerEvents, IAgentManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    ISystemRegistry public origin;

    ISystemRegistry public destination;

    // agent => (bool isSlashed, address prover)
    mapping(address => SlashStatus) public slashStatus;

    /// @dev gap for upgrade safety
    uint256[47] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line func-name-mixedcase
    function __AgentManager_init(ISystemRegistry origin_, ISystemRegistry destination_) internal onlyInitializing {
        origin = origin_;
        destination = destination_;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentManager
    // solhint-disable-next-line ordering
    function agentRoot() external view virtual returns (bytes32);

    /// @inheritdoc IAgentManager
    function agentStatus(address agent) external view returns (AgentStatus memory status) {
        status = _agentStatus(agent);
        // If agent was proven to commit fraud, but their slashing wasn't completed,
        // return the Fraudulent flag instead
        if (slashStatus[agent].isSlashed && status.flag != AgentFlag.Slashed) {
            status.flag = AgentFlag.Fraudulent;
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Checks and initiates the slashing of an agent.
    /// Should be called, after one of registries confirmed fraud committed by the agent.
    function _initiateSlashing(uint32 domain, address agent, address prover) internal {
        // Check that Agent hasn't been already slashed
        require(!slashStatus[agent].isSlashed, "Already slashed");
        // Check that agent is Active/Unstaking and that the domains match
        AgentStatus memory status = _agentStatus(agent);
        require(
            (status.flag == AgentFlag.Active || status.flag == AgentFlag.Unstaking) && status.domain == domain,
            "Slashing could not be initiated"
        );
        slashStatus[agent] = SlashStatus({isSlashed: true, prover: prover});
        emit StatusUpdated(AgentFlag.Fraudulent, domain, agent);
    }

    /// @dev Notifies a given set of local registries about the slashed agent.
    /// Set is defined by a bitmask, eg: DESTINATION | ORIGIN
    function _notifySlashing(uint256 registryMask, uint32 domain, address agent, address prover) internal {
        // Notify Destination, if requested
        if (registryMask & DESTINATION != 0) destination.managerSlash(domain, agent, prover);
        // Notify Origin, if requested
        if (registryMask & ORIGIN != 0) origin.managerSlash(domain, agent, prover);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Generates leaf to be saved in the Agent Merkle Tree
    function _agentLeaf(AgentFlag flag, uint32 domain, address agent) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(flag, domain, agent));
    }

    /// @dev Returns the last known status for the agent from the Agent Merkle Tree.
    function _agentStatus(address agent) internal view virtual returns (AgentStatus memory);
}