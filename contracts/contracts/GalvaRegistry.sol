// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "../interfaces/IGalva.sol";

contract GalvaRegistry is IGalva {
    // Land model
    struct Record {
        string title;
        uint area;
        bytes owner;
    }

    // State
    mapping(bytes => Record) private records;
    mapping(bytes => bool) private nonces;
    mapping(address => uint) private rights;
    uint private tokenizedRights;

    // Modifiers(called before function call)
    modifier taken(bytes memory node) {
        require(!nonces[node], "exists: record exists");
        _;
    }

    /**
        * @notice Register property
        * @dev Attest to new property ownership
        * @param title property title
        * @param area approximation area
        * @param node property signature
        * @param sender who is uploading property
    */
    function attestProperty(
        string memory title,
        uint area,
        bytes memory node,
        address sender
    ) external override taken(node) {
        records[node] = Record({
            title: title,
            area: area,
            owner: node
        });
        // mark title signature as already used
        nonces[node] = true;
        // rights accumulated by the blockchain
        tokenizedRights += area;
        // rights accumulated by property owner
        rights[sender] += area;
    }

    /**
        * @notice We assume a cryptographical truth
        * @dev Checks if signature is consumed by the blockchain
        * @param node property signature
        * @return bool
    */
    function recordExists(bytes memory node) external view override returns (bool) {
        return nonces[node] != false;
    }

    /**
        * @notice Address consumed rights
        * @dev Return accumulated rights for an address
        * @param who for who
        * @return uint
    */
    function addressRights(address who) external view override returns (uint) {
        return rights[who];
    }

    /**
        * @notice Blockhain consumed rights
        * @dev Returns rights accumulated by the blockchain
        * @return uint
    */
    function blockchainConsumedRights() external view override returns (uint) {
        return tokenizedRights;
    }

    /**
        * @notice Claim property ownership
        * @param title_number property title number
        * @param v parity of the co-ordinate of r
        * @param r the x co-ordinate of r
        * @param s the s value of the signature
        * @return address
    */
    function claimOwnership(string memory title_number, uint8 v, bytes32 r, bytes32 s) external pure override returns (address) {
        bytes32 message = recreateMsg(title_number);
        address claimer = ecrecover(message, v, r, s);
        return claimer;
    }

    // reconstruct eth_sign mechanism
    function prefixed(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }

    // Recreate signed msg on-chain
    function recreateMsg(string memory title_number) internal pure returns (bytes32) {
        // encode params
        bytes32 payload = keccak256(abi.encode(title_number));

        bytes32 message = prefixed(payload);

        return message;
    }
}
