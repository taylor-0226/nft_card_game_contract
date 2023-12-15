// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Crafting/craftingPredictions.sol";

contract PredictionFactory is Initializable, OwnableUpgradeable {    
    address public adminWallet;

    event PredictionCreated(address prediction);

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

    function createCraftingPredection(uint256 collectionId, address owner) external 
        onlyAdmin returns(address prediction){        

        PredictionCrafting newPrediction = new PredictionCrafting(collectionId, owner);
        prediction = address(newPrediction);

        newPrediction.transferOwnership(owner);            

        emit PredictionCreated(prediction);

        return (prediction);
    }   
}