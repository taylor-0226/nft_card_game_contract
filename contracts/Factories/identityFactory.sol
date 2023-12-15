// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Crafting/craftingIdentity.sol";

contract IdentityFactory is Initializable, OwnableUpgradeable {    
    address public adminWallet;

    event IdentityCreated(address category);

    modifier onlyAdmin() {
        require(msg.sender == adminWallet, "You must be the admin contract");
        _;
    }

    function initialize() external initializer {                        
        __Ownable_init();                
    }

    function changeAdmin(address _newAdmin) public onlyOwner {
        adminWallet = _newAdmin;
    }

    function createCraftingIdentity(uint256 collectionId, address owner) external 
        onlyAdmin returns(address identity){        

        CollectionCrafting newIdentity = new CollectionCrafting(collectionId, owner);
        identity = address(newIdentity);

        newIdentity.transferOwnership(owner);            

        emit IdentityCreated(identity);

        return (identity);
    }   
}