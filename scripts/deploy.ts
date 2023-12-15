import {ethers} from "hardhat";
import { task } from "hardhat/config";
import { TaskArguments } from "hardhat/types";

const cardpackFactory = "0x3a2d175B8228De11dF7B44959527D0114cEbbe62";
const categoryFactory = "0xf474b7B4d946e559f156c7d5d02e8f5Ab994bDB6"
const yearFactory = "0x687Dcd759C323dBB76D53a9D5893F9D391b93358";
const daymonthFactory = "0xb040c97359377478D51Bf0b3C5D459FC46250E43";
const craftingcardFactory = "0x3c4Ca73bFA4C4CCa9D67b7D79b23dBdc2703Fe5C";
const triggerFactory = "0x78BD6D288237b0CEC3fF4Ee95Aa76A70F660A174";
const IdentityFactory = "0x3b0e4F34873feD7d5680000D7ae1C4866A75d501"
const PredictionFactory = "0xe27B954f788FEA7C4BA36D12bA705D70017f8Bf3"

const factories = [
  'CardPackFactory', 
  // 'CategoryFactory', 'YearFactory',
  // 'DayMonthFactory', 'TriggerFactory', 'CraftingCardFactory',
  // 'IdentityFactory', 'PredictionFactory'
  ]
const deployedAddr = "0xeBbc8b54488e60bC141bc4cf4b22E310F297AFBB"
task("deploy:Factory", "Deploy Factories Smart Contract")  
  .setAction(async function (taskArguments: TaskArguments, hre) {        
    
    for(let i = 0; i < factories.length; i++) {
      const Factory = await hre.ethers.getContractFactory(factories[i]);
        // Deploy Contract
      const factory = await hre.upgrades.deployProxy(Factory);
      await factory.deployed();
  
      console.log(factories[i] + " deployed to:", factory.address);
    }
});

task("deploy:Collection", "Deploy CardPack Smart Contract")  
  .setAction(async function (taskArguments: TaskArguments, hre) {    
    
    const Collection = await hre.ethers.getContractFactory(
      "Collection"
    );

    // Deploy Contract
    const collection = await hre.upgrades.deployProxy(Collection, [
      cardpackFactory, 
      categoryFactory, 
      yearFactory, 
      daymonthFactory, 
      craftingcardFactory, 
      triggerFactory,
      IdentityFactory,
      PredictionFactory,
    ]);
    await collection.deployed();

    console.log("Collection deployed to:", collection.address);
});


task("deploy:PackOpener", "Deploy PackOpener Smart Contract")  
  
  .setAction(async function (taskArguments: TaskArguments, hre) {    

    // const { cardpack, daymonth, year, category, craftingcard, trigger} = taskArguments;
    

    const PackOpenerFactory = await hre.ethers.getContractFactory(
      "PackOpener"
    );

    // Deploy Contract
    const PackOpener = await hre.upgrades.deployProxy(PackOpenerFactory);
    await PackOpener.deployed();

    console.log("Marketplace deployed to:", PackOpener.address);
});

// task("deploy:CollectionCrafting", "Deploy CollectionCrafting Smart Contract")  
  
//   .setAction(async function (taskArguments: TaskArguments, hre) {    
  
//     const CollectionCraftingFactory = await hre.ethers.getContractFactory(
//       "CollectionCrafting"
//     );

//     // Deploy Contract
//     const CollectionCrafting = await hre.upgrades.deployProxy(CollectionCraftingFactory);
//     await CollectionCrafting.deployed();

//     console.log("Marketplace deployed to:", CollectionCrafting.address);
// });

// task("deploy:PredictionCrafting", "Deploy PredictionCrafting Smart Contract")    
//   .setAction(async function (taskArguments: TaskArguments, hre) {        
//     const PredictionCraftingFactory = await hre.ethers.getContractFactory(
//       "PredictionCrafting"
//     );

//     // Deploy Contract
//     const PredictionCrafting = await hre.upgrades.deployProxy(PredictionCraftingFactory, [indentityContract]);
//     await PredictionCrafting.deployed();

//     console.log("Marketplace deployed to:", PredictionCrafting.address);
// });

task("upgrade:contract", "Upgrade contract")
  // .addParam("address", "The deployed smart contract address")
  .setAction(async function (taskArguments: TaskArguments, hre) {
    const address = taskArguments.address;

    // Upgrade Contract
    const ContractV2 = await hre.ethers.getContractFactory("CollectionCrafting");
    const upgraded = await hre.upgrades.upgradeProxy(deployedAddr, ContractV2);
    await upgraded.deployed();

    console.log("Collection deployed to:", upgraded.address);
    console.log("ERC721A was upgraded successfully !!!");
  });