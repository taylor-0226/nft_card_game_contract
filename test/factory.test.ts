import { ethers, upgrades } from "hardhat";
import { Contract, Signer } from "ethers";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";

import packABI from "../artifacts/contracts/CardPack/cardPack.sol/CardPack.json"
import cardABI from "../artifacts/contracts/Parts/category.sol/Category.json"
import CraftingCardABI from "../artifacts/contracts/Parts/craftingCard.sol/CraftingCard.json";
import identityABI from "../artifacts/contracts/Crafting/craftingIdentity.sol/CollectionCrafting.json";
import predictionABI from "../artifacts/contracts/Crafting/craftingPredictions.sol/PredictionCrafting.json";


const collectionCreated = new ethers.utils.Interface(["event CollectionCreated(uint256 tokenId, address pack, address category, address year, address daymonth, address crafting, address trigger, address identity, address prediction)"]);

const PACK_TIER = {
  STANDARD: 0,
  PREMIUM: 1,
  ELITE: 2,
}

describe("Category", () => {
  let categoryFactory: Contract;
  let yearFactory: Contract;
  let daymonthFactory: Contract;
  let triggerFactory: Contract;
  let craftingFactory: Contract;
  let cardPackFactory: Contract;
  let identityFactory: Contract;
  let predictionFactory: Contract;
  let packOpener: Contract;
  let collection: Contract;
  let owner: Signer;
  let newOnwer: Signer;
  let ownerAddress: string;
  

  beforeEach(async () => {
    [owner, newOnwer] = await ethers.getSigners();    
    ownerAddress = await owner.getAddress()

    const CardPackFactory = await ethers.getContractFactory("CardPackFactory")
    cardPackFactory = await upgrades.deployProxy(CardPackFactory);
    await cardPackFactory.deployed()

    const CategoryFactory = await ethers.getContractFactory("CategoryFactory");    
    categoryFactory = await upgrades.deployProxy(CategoryFactory);
    await categoryFactory.deployed()

    const YearFactory = await ethers.getContractFactory("YearFactory");    
    yearFactory = await upgrades.deployProxy(YearFactory);
    await yearFactory.deployed()

    const DayMonthFactory = await ethers.getContractFactory("DayMonthFactory");    
    daymonthFactory = await upgrades.deployProxy(DayMonthFactory);
    await daymonthFactory.deployed()

    const TriggerFactory = await ethers.getContractFactory("TriggerFactory");    
    triggerFactory = await upgrades.deployProxy(TriggerFactory);
    await triggerFactory.deployed()

    const CraftingFactory = await ethers.getContractFactory("CraftingCardFactory");    
    craftingFactory = await upgrades.deployProxy(CraftingFactory);
    await categoryFactory.deployed()

    const IdentityFactory = await ethers.getContractFactory("IdentityFactory");
    identityFactory = await upgrades.deployProxy(IdentityFactory);
    await identityFactory.deployed()

    const PredictionFactory = await ethers.getContractFactory("PredictionFactory");
    predictionFactory = await upgrades.deployProxy(PredictionFactory);
    await predictionFactory.deployed()

    const Collection = await ethers.getContractFactory("Collection");    
    collection = await upgrades.deployProxy(Collection, [
      cardPackFactory.address, 
      categoryFactory.address, 
      yearFactory.address, 
      daymonthFactory.address, 
      craftingFactory.address, 
      triggerFactory.address,
      identityFactory.address,
      predictionFactory.address,
    ]);
    await collection.deployed()        

    const PackOpener = await ethers.getContractFactory("PackOpener")
    packOpener = await upgrades.deployProxy(PackOpener)
    await packOpener.deployed()    
    
    console.log('collection address', collection.address, await collection.owner())
    console.log('pack factory', cardPackFactory.address)
    console.log('category factory', categoryFactory.address)
    console.log('year factory', yearFactory.address)
    console.log('daymonth factory', daymonthFactory.address)
    console.log('trigger factory', triggerFactory.address)
    console.log('crafting factory', craftingFactory.address)
    console.log('identity factory', identityFactory.address)
    console.log('prediction factory', predictionFactory.address)
    console.log('packopener address', packOpener.address)
  });

  describe("createCategory", () => {
    it("should create a new category collection", async () => {

        
        await cardPackFactory.connect(owner).changeAdmin(collection.address);
        await categoryFactory.connect(owner).changeAdmin(collection.address);
        await yearFactory.connect(owner).changeAdmin(collection.address);
        await daymonthFactory.connect(owner).changeAdmin(collection.address);
        await craftingFactory.connect(owner).changeAdmin(collection.address);
        await triggerFactory.connect(owner).changeAdmin(collection.address);
        await identityFactory.connect(owner).changeAdmin(collection.address);
        await predictionFactory.connect(owner).changeAdmin(collection.address);
        
        await collection.connect(owner).changeMinter(ownerAddress);
        const collectionTx = await collection.connect(owner).createCollection(
          "https://example.com/collection/0",
          // configuration for cateogry          
          [300, 200, 100],
        )
        
        const collectionReceipt = await ethers.provider.getTransactionReceipt(collectionTx.hash);                
        
        console.log('logs', collectionReceipt.logs.length)
        const event = collectionCreated.parseLog(collectionReceipt.logs[collectionReceipt.logs.length-1]);        
        
        console.log('event result', event.args)
        const pack = new ethers.Contract(event.args.pack, packABI.abi, owner)  
        const category = new ethers.Contract(event.args.category, cardABI.abi, owner)
        const year = new ethers.Contract(event.args.year, cardABI.abi, owner)
        const daymonth = new ethers.Contract(event.args.daymonth, cardABI.abi, owner)
        const crafting = new ethers.Contract(event.args.crafting, CraftingCardABI.abi, owner)
        const trigger = new ethers.Contract(event.args.trigger, cardABI.abi, owner)
        const identity = new ethers.Contract(event.args.identity, identityABI.abi, owner)
        const prediction = new ethers.Contract(event.args.prediction, predictionABI.abi, owner)

        console.log('pack info', pack.address, await pack.owner());
        console.log('category info', category.address);
        console.log('year info', year.address);
        console.log('daymonth info', daymonth.address)
        console.log('crafting info', crafting.address, await crafting.name())
        console.log('trigger info', trigger.address, await trigger.name())
        console.log('identity info', identity.address)
        console.log('prediction info', prediction.address)        
        
        await pack.connect(owner).changeMinter(ownerAddress);
        await pack.connect(owner).changeOpener(packOpener.address);
        await pack.connect(owner).mint(ownerAddress, 1, PACK_TIER.STANDARD, 'https://example.com/standard_pack/0');        

        await category.changeMinter(packOpener.address);
        await daymonth.changeMinter(packOpener.address);
        await crafting.changeMinter(packOpener.address);
        await year.changeMinter(packOpener.address);
        await trigger.changeMinter(packOpener.address);

        await category.changeCrafter(identity.address);
        await year.changeCrafter(identity.address);
        await daymonth.changeCrafter(identity.address);
        await crafting.changeCrafter(identity.address)
        await crafting.changePredictor(prediction.address)
        await trigger.changeCrafter(prediction.address);
        await identity.changeCrafter(prediction.address);
        

        await packOpener.openPack(
          0,
          pack.address,
          category.address,
          year.address,
          daymonth.address,          
          crafting.address,
          trigger.address,          
          ['https://exmaple.com/collection/0/cateogry/0'],
          ['https://exmaple.com/collection/0/year/0'],
          ['https://exmaple.com/collection/0/daymonth/0'],
          ['https://exmaple.com/collection/0/crafting/0', 'https://exmaple.com/collection/0/crafting/1'],
          ['https://exmaple.com/collection/0/trigger/0'],
        )

        console.log('category card owner', await category.ownerOf(0))
        console.log('year card owner', await year.ownerOf(0))
        console.log('daymonth card owner', await daymonth.ownerOf(0))
        console.log('crafting card owner', await crafting.ownerOf(0), await crafting.ownerOf(1))
        console.log('trigger card owner', await category.ownerOf(0))

        await identity.craftCollection(
          daymonth.address,
          year.address,
          category.address,
          crafting.address,
          0,
          0,
          0,
          0,
          'https://example.com/collectin/0/identity/0'
        )

        await prediction.craftCollection(
          identity.address,
          trigger.address,
          crafting.address,
          0,
          [0],
          1,
          'https://example.com/collection/0/prediction/0'
        )

        console.log('identity token owner', await identity.ownerOf(0))
        console.log('prediction owner', await prediction.ownerOf(0))
    });
  })
})