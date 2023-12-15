// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title A contract for interacting with the CardPack contract.
interface ICardPack {
    
    function mint(
        address _to,
        uint256 _quantity,
        uint256 _tier,
        string memory _tokenURI
    ) external returns (uint256);

    /// @notice Changes the status of a card pack to "opened".
    /// @param tokenId The ID of the token to be opened.
    function changeToOpened(uint256 tokenId) external;

    /// @notice Fetches the owner of a particular card pack.
    /// @param tokenId The ID of the token whose owner is to be fetched.
    /// @return The address of the owner.
    function ownerOf(uint256 tokenId) external view returns (address);

    /// @notice Checks whether a card pack is opened or not.
    /// @param tokenId The ID of the token to be checked.
    /// @return A boolean indicating whether the card pack is opened or not.
    function isOpened(uint256 tokenId) external view returns (bool);
}