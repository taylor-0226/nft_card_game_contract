import { ethers } from "ethers";
import * as dotenv from "dotenv";
import CollectionABI from "../artifacts/contracts/Collection/collection.sol/Collection.json";
import packABI from "../artifacts/contracts/CardPack/cardPack.sol/CardPack.json"
import CategoryABI from "../artifacts/contracts/Parts/category.sol/Category.json"
import YearABI from "../artifacts/contracts/Parts/year.sol/Year.json";
import DayMonthABI from "../artifacts/contracts/Parts/dayMonth.sol/DayMonth.json"
import CraftingCardABI from "../artifacts/contracts/Parts/craftingCard.sol/CraftingCard.json";
import TriggerABI from "../artifacts/contracts/Parts/triggers.sol/Trigger.json";
import packOpenerABI from "../artifacts/contracts/OpenPack/packOpener.sol/PackOpener.json"
import identityABI from "../artifacts/contracts/Crafting/craftingIdentity.sol/CollectionCrafting.json";
import predictionABI from "../artifacts/contracts/Crafting/craftingPredictions.sol/PredictionCrafting.json";
import packFactoryABI from "../artifacts/contracts/Factories/cardpackFactory.sol/CardPackFactory.json";

dotenv.config()

const main = async() => {
    const collectionAddr = '0x549E02778d338A170032C0a771786ab5ecf77E0C';    
    const packAddr = "0xEddaC99edca0e45B374834D85a82c40a48C8B5e9"
    const categoryAddr = "0x1983C5bFE96e7C3e68a1c99013136CcBEF5b34e2"
    const yearAddr = "0xd7685C6daf96f96828c1E54A573b1f07A34D64b4"
    const dayMAddr = "0x0B92E457ADf89F2d047f3bCAad93606443a6c9F0"
    const craftingCAddr = "0xAB39C5EF933163439fe420546Fa5d211D781a66A"
    const triggerAddr = "0x3D8462aF5b432C2583a0bF274955085b67A88eC5"
    const packOpenerAddr = "0x5D6A8216590cEEeFd149f3CB26ad096B8516bB00"
    const identityAddr = "0xA1Be241F68686093e79BD9aB7e73115D51174341"
    const predictionAddr = "0xc341bf513aB1Bd89bED828F18c2560A158ae89ac"

    const packFAddr = "0x3a2d175B8228De11dF7B44959527D0114cEbbe62";

    const provider = new ethers.providers.JsonRpcProvider(process.env.MUMBAI_URL);
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY as string, provider);    

    const packFactory = new ethers.Contract(packFAddr, packFactoryABI.abi, wallet);
    const collection = new ethers.Contract(collectionAddr, CollectionABI.abi, wallet);    

    // await packFactory.changeAdmin(collectionAddr);
    // await collection.createCollection("https://example.com/collection/0", 300, 200, 100, 600)

    // const pack = new ethers.Contract(packAddr, packABI.abi, wallet)
    // const cateogry = new ethers.Contract(categoryAddr, CategoryABI.abi, provider)
    // const year = new ethers.Contract(yearAddr, YearABI.abi, provider)
    // const daymonth = new ethers.Contract(dayMAddr, DayMonthABI.abi, provider)
    // const craftingCard = new ethers.Contract(craftingCAddr, CraftingCardABI.abi, provider)
    // const triggerCard = new ethers.Contract(triggerAddr, TriggerABI.abi, provider)
    // const packOpener = new ethers.Contract(packOpenerAddr, packOpenerABI.abi, wallet)
    // const identity = new ethers.Contract(identityAddr, identityABI.abi, wallet)
    // const prediction = new ethers.Contract(predictionAddr, predictionABI.abi, wallet)

    // console.log('minter category', await cateogry.ownerOf(1))
    // console.log('minter year', await year.name())
    // console.log('minter daymonth', await daymonth.name())
    console.log('pack factory', await packFactory.owner())
    // console.log('collection', await collection.name())
    // console.log('minter crafting card', await craftingCard.tokenCounter())
    // console.log('minter trigger', await triggerCard.ownerOf(0))
    // console.log('token owner', await pack.isOpened(1))
    // console.log('token owner', await pack.ownerOf(0))
    
    // console.log('identity owner', await identity.ownerOf(0))
}

main().then()