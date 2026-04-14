// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";

/**
 * @title TradingGuard
 * @notice Emergency circuit breaker triggered by off-chain AI anomaly detection.
 */
contract TradingGuard is Pausable, AccessControl {
    // Role for the off-chain AI "watcher" node
    bytes32 public constant BOT_ROLE = keccak256("BOT_ROLE");
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    event EmergencyStopTriggered(address indexed initiator, string reason);
    event TradingResumed(address indexed initiator);

    constructor(address _admin, address _bot) {
        _grantRole(ADMIN_ROLE, _admin);
        _grantRole(BOT_ROLE, _bot);
        _setRoleAdmin(BOT_ROLE, ADMIN_ROLE);
    }

    /**
     * @notice Triggered by the AI backend when an anomaly is detected
     * @param reason The type of anomaly (e.g., "Flash Crash", "Toxic Flow")
     */
    function triggerCircuitBreaker(string calldata reason) external onlyRole(BOT_ROLE) {
        _pause();
        emit EmergencyStopTriggered(msg.sender, reason);
    }

    /**
     * @notice Resumes trading once the anomaly has passed. Only Admin can do this.
     */
    function resumeTrading() external onlyRole(ADMIN_ROLE) {
        _unpause();
        emit TradingResumed(msg.sender);
    }

    /**
     * @notice A modifier for your trading functions to check the guard
     */
    modifier isTradeAllowed() {
        require(!paused(), "TradingGuard: Circuit breaker active");
        _;
    }
}