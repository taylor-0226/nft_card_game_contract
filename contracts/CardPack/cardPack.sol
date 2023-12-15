// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
@title CardPack
@dev This contract represents a Card Pack ERC721 token contract, allowing users to create and manage card packs.
*/
contract CardPack is ERC721, Ownable {
    uint256 totalCounter;
    uint256[] tierCounters;    
    mapping(uint256 => bool) public opened; // Mapping to track if a card pack has been opened
    mapping(uint256 => string) private _tokenURIs; // Mapping to store the token URIs for each card pack
    
    uint256 public totalLimit;   
    uint256[] public tierLimits;     
    address public minter;
    address public opener;

    // moonpay address
    address public proxy;

    event CardPackCreated(
        uint256 indexed tokenId,
        uint256 tier,
        string tokenURI
    );
    event CardPackOpened(uint256 indexed tokenId);
    
    event LimitChanged(uint256 newLimit, string typeOf);

    /**
     * @dev Emitted when the minter contract address is changed.
     * @param previousMinter The address of the previous minter contract.
     * @param newMinter The address of the new minter contract.
     */
    event MinterChanged(
        address indexed previousMinter,
        address indexed newMinter
    );

    event OpenerChanged(
        address indexed previousOpener,
        address indexed newOpener
    );

    event PayProxyChanged(
        address indexed previousProxy,
        address indexed newProxy
    );
    /**
    @dev Initializes the CardPack contract by setting the token name and symbol.
    */   
    constructor(        
        uint256[] memory _packLimits
    ) ERC721("Card Pack", "CP") {  
        //loop through pack limits to add to total limit
        uint256 _totalAmount = 0;
        for(uint256 i = 0; i < _packLimits.length; i++) {
            _totalAmount += _packLimits[i];
        }

        tierLimits = _packLimits;

        //loop through tierLimits and set each tierCounters to 0
        for(uint256 i = 0; i < tierLimits.length; i++) {
            tierCounters.push(0);
        }
        


        totalLimit = _totalAmount;
        totalCounter = 0;
    }

    /**
     * @dev Modifier to check if a card pack with the given token ID exists.
     * @param tokenId The ID of the token to check.
     */
    modifier packExists(uint256 tokenId) {
        require(_exists(tokenId), "CardPack: Pack does not exist");
        _;
    }

    modifier tokenURICheck(string memory _tokenURI) {
        require(bytes(_tokenURI).length > 0, "TokenURI should not be empty");
        _;
    }

    modifier checkLimit() {
        require(totalLimit > totalCounter, "All the tokens have been minted");
        _;
    }

    /**
    @dev Modifier to check if the sender is the minter contract.
    */
    modifier onlyMinter() {
        require(
            msg.sender == minter || msg.sender == proxy,
            "Collection: You are not the Minting Contract :)"
        );
        _;
    }

    modifier onlyOpener() {
        require(
            msg.sender == opener,
            "Collection: You are not the Minting Contract :)"
        );
        _;
    }

    /**
    @dev Internal function to handle burning of a token.
    @param tokenId The ID of the token to be burned.
    */
    function _burn(uint256 tokenId) internal override(ERC721) {
        super._burn(tokenId);
    }

    /**
    @dev Returns the URI for a given token ID.
    @param tokenId The ID of the token to query the URI for.
    @return A string representing the URI for the given token ID.
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


    /**
    @dev Creates a new card pack token.
    @param _to The buyer address.
    @param _tokenURI The URI for the token.
    @param _tier The tier of card pack, mapped to array in collection metadata
    @return The ID of the newly created token.
    */
    function mint(
        address _to,
        uint256 _quantity,
        uint256 _tier,
        string memory _tokenURI
    ) public payable onlyMinter tokenURICheck(_tokenURI) checkLimit returns (uint256) {
        //check if tierCounters[_tier] is less than tierLimits[_tier]
        //if not, revert        
        require(_quantity > 0, "Quantity should be greater than 0");
        require((_quantity + tierCounters[_tier]) <= tierLimits[_tier], "Quantity shouldn't surpass tier limit");

        _safeMint(_to, totalCounter);
        _tokenURIs[totalCounter] = _tokenURI;
        opened[totalCounter] = false;
        
        totalCounter++;
        tierCounters[_tier]++;

        emit CardPackCreated(                
                totalCounter - 1,
                _tier,
                _tokenURI 
        );
        return totalCounter-1;
    }

    /**
    @dev Changes the status of a card pack to opened.
    @param tokenId The ID of the token to be marked as opened.
    */
    function changeToOpened(
        uint256 tokenId
    ) public onlyOpener packExists(tokenId) {
        opened[tokenId] = true;
        _burn(tokenId);
        emit CardPackOpened(tokenId);
    }

    /**
    @dev Checks if a card pack has been opened.
    @param tokenId The ID of the token to check.
    @return A boolean indicating whether the card pack has been opened.
    */
    function isOpened(
        uint256 tokenId
    ) public view packExists(tokenId) returns (bool) {
        return opened[tokenId];
    }

    function changeToTotalLimit(uint256 newLimit) public onlyOwner {
        require(newLimit > totalLimit, "New limit should be higher");
        totalLimit = newLimit;
        emit LimitChanged(newLimit, "Standard");
    }

    /**
     * @dev Changes the minter contract address.
     * @param _newMinter The new address of the minter contract.
     */
    function changeMinter(address _newMinter) public onlyOwner {
        require(_newMinter != address(0), "Minter: No zero address");
        address oldMinter = minter;
        minter = _newMinter;
        emit MinterChanged(oldMinter, _newMinter);
    }

    function changeOpener(address _newOpener) public onlyOwner {
        require(_newOpener != address(0), "Opener: No zero address");
        address oldOpener = opener;
        opener = _newOpener;
        emit MinterChanged(oldOpener, _newOpener);
    }

    function setPayProxy(address _proxy) public onlyOwner {
        require(_proxy != address(0), "Opener: No zero address");
        address oldProxy = proxy;
        proxy = _proxy;
        emit MinterChanged(oldProxy, _proxy);
    }
}
