import { task } from "hardhat/config";

task("galva:deploy", "Deploy Galva Contracts").setAction(async (args, hre) => {
  // get signer
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deployer", await deployer.getAddress());
  // deploy galva contracts
  const registry = await (await hre.ethers.getContractFactory("GalvaRegistry"))
    .connect(deployer)
    .deploy();
  console.log("Galva deployed at", registry.address);
});
