// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import '../CardPack/cardPack.sol';

contract CardPackFactory is Initializable, OwnableUpgradeable {    
    address public adminWallet;

    event cardPackCreated(address cardpack);

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

    function createCardPack(        
        uint256[] memory _packLimits,
        address owner
    ) external onlyAdmin returns(address cardpack){        

        CardPack newCardPackCategory = new CardPack(_packLimits);
        cardpack = address(newCardPackCategory);

        newCardPackCategory.transferOwnership(owner);            

        emit cardPackCreated(cardpack);

        return (cardpack);
    }   
}