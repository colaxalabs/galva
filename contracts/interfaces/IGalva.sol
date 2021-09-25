// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

interface IGalva {
    // Logged when property owner uploads property
    event Attest(bytes32 indexed node);

    /**
        * @notice Register property
        * @dev Attest to new property ownership
        * @param title property title
        * @param area approximation area
        * @param node property signature
        * @param sender who is uploading property
    */
    function attestProperty(string memory title, uint area, bytes memory node, address sender) external;

    /**
        * @notice We assume a cryptographical truth
        * @dev Checks if signature is consumed by the blockchain
        * @param node property signature
        * @return bool
    */
    function recordExists(bytes memory node) external returns (bool);

    /**
        * @notice Address consumed rights
        * @dev Return accumulated rights for an address
        * @param who for who
        * @return uint
    */
    function addressRights(address who) external returns (uint);

    /**
        * @notice Claim property ownership
        * @param title_number property title number
        * @param v parity of the co-ordinate of r
        * @param r the x co-ordinate of r
        * @param s the s value of the signature
        * @return address
    */
    function claimOwnership(string memory title_number, uint8 v, bytes32 r, bytes32 s) external returns (address);

    /**
        * @notice Blockhain consumed rights
        * @dev Returns rights accumulated by the blockchain
        * @return uint
    */
    function blockchainConsumedRights() external returns (uint);
}
