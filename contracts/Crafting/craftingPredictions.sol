// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

interface IPart {
    function ownerOf(uint256 tokenId) external view returns (address);

    /**
     * @dev Checks if a token has been marked as crafted.
     * @param tokenId The ID of the token to check.
     * @return A boolean indicating if the token is crafted.
     */
    function isCrafted(uint256 tokenId) external view returns (bool);

    /**
     * @dev Changes the status of a card pack to crafted.
     * @param tokenId The ID of the token to be marked as crafted.
     */
    function changeToCrafted(uint256 tokenId) external;
    function changeToCrafted2(uint256 tokenId) external;
}

contract PredictionCrafting is ERC721, Ownable {
    IPart public identitiesContract;
    IPart public triggersContract;
    IPart public craftingCardContract;
    address adminWallet;
    uint256 private tokenCounter;
    uint256 public collectionId;
    mapping(uint256 => uint256) private identityTokenIdByCollection;
    mapping(uint256 => uint256[]) private triggerTokenIdByCollection;
    mapping(uint256 => uint256) private craftingCardTokenIdByCollection;
    mapping(uint256 => bool) public crafted;
    mapping(address => bool) public won;
    /**
     * @dev Stores the token URIs for each token ID.
     */
    mapping(uint256 => string) private _tokenURIs;
    event CollectionCrafted(
        uint256 indexed collectionId,
        uint256 identityTokenId,
        uint256[] triggerTokenId,
        string tokenURI
    );

    constructor(uint256 _collectionId, address admin) 
        ERC721("Prediction", "PRD") {            
        
        collectionId = _collectionId;            
        tokenCounter = 0;
        adminWallet = admin;
    }

    modifier onlyAdmin() {
        require(msg.sender == adminWallet, "You must be the admin wallet");
        _;
    }

    /**
     * @dev Returns the URI for a given token ID.
     * @param tokenId The ID of the token to query the URI for.
     * @return A string representing the URI for the given token ID.
     */
    function tokenURI(
        uint256 tokenId
    ) public view override returns (string memory) {
        require(_exists(tokenId), "URI query for nonexistent token");
        string memory baseURI = _baseURI();
        if (bytes(baseURI).length == 0) {
            return _tokenURIs[tokenId];
        } else {
            return string(abi.encodePacked(baseURI, _tokenURIs[tokenId]));
        }
    }

    function craftCollection(
        address _identityContract,
        address _triggersContract,
        address _craftingCardContract,

        uint256 identityTokenId,
        uint256[] memory triggerTokenIds,
        uint256 craftingCardTokenId,
        string memory _tokenURI
    ) public onlyAdmin {
        triggersContract = IPart(_triggersContract);
        craftingCardContract = IPart(_craftingCardContract);
        identitiesContract = IPart(_identityContract);
        // Check ownership of identity NFT
        address identityOwner = identitiesContract.ownerOf(identityTokenId);
        // Check ownership of crafting card NFT
        address craftingCardOwner = craftingCardContract.ownerOf(
            craftingCardTokenId
        );

        // Check ownership of trigger NFTs and crafting status
        for (uint256 i = 0; i < triggerTokenIds.length; i++) {
            uint256 triggerTokenId = triggerTokenIds[i];
            address triggerOwner = triggersContract.ownerOf(triggerTokenId);
            bool triggerBool = triggersContract.isCrafted(triggerTokenId);
            require(
                triggerOwner == identityOwner,
                "Only the owner of all tokens can craft a new collection"
            );
            require(triggerBool == false, "Trigger already used");
        }

        require(
            identityOwner == craftingCardOwner,
            "Only the owner of all tokens can craft a new collection"
        );
        require(
            !identitiesContract.isCrafted(identityTokenId),
            "Identity already used"
        );
        require(
            !craftingCardContract.isCrafted(craftingCardTokenId),
            "Crafting Card already used"
        );
        
        _safeMint(msg.sender, tokenCounter);

        identitiesContract.changeToCrafted(identityTokenId);
        craftingCardContract.changeToCrafted2(craftingCardTokenId);

        for (uint256 i = 0; i < triggerTokenIds.length; i++) {
            uint256 triggerTokenId = triggerTokenIds[i];
            triggersContract.changeToCrafted(triggerTokenId);
        }

        identityTokenIdByCollection[tokenCounter] = identityTokenId;
        craftingCardTokenIdByCollection[tokenCounter] = craftingCardTokenId;

        for (uint256 i = 0; i < triggerTokenIds.length; i++) {
            uint256 triggerTokenId = triggerTokenIds[i];
            triggerTokenIdByCollection[tokenCounter].push(triggerTokenId);
        }

        tokenCounter++;
        _tokenURIs[collectionId] = _tokenURI;

        crafted[collectionId] = false;
        emit CollectionCrafted(
            collectionId,
            identityTokenId,
            triggerTokenIds,
            _tokenURI
        );
    }

    function getCollectionTokens(
        uint256 collectionId
    ) public view returns (uint256, uint256[] memory) {
        require(_exists(collectionId), "Collection does not exist");
        return (
            identityTokenIdByCollection[collectionId],
            triggerTokenIdByCollection[collectionId]
        );
    }

    function changeAdmin(address _newAdmin) public onlyOwner {
        adminWallet = _newAdmin;
    }

    function win(address _address) public onlyAdmin {
        won[_address] = true;
    }
}
