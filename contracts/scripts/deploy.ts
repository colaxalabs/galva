import { task } from "hardhat/config";
import { Contract } from "ethers";
import fs from "fs";

task("galva:deploy", "Deploy Galva Contracts").setAction(async (args, hre) => {
  // get signer
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deployer", await deployer.getAddress());
  // deploy galva contracts
  const registry = await (await hre.ethers.getContractFactory("GalvaRegistry"))
    .connect(deployer)
    .deploy();
  console.log("Galva deployed at", registry.address);
  // save contract client-side
  saveAbiClientSide(registry);
});

function saveAbiClientSide(registry: Contract) {
  const dir = "/home/lomolo/Projects/galva/client/contracts";
  // check if path exists
  if (!fs.existsSync(dir)) {
    // create path
    fs.mkdirSync(dir);
  }

  // write contract deployed addresses
  fs.writeFileSync(
    dir + "/contracts.json",
    JSON.stringify(
      {
        Registry: registry.address,
      },
      undefined,
      2
    )
  );

  // write contracts artifacts
  const registryArtifacts = require(__dirname +
    "/../artifacts/contracts/GalvaRegistry.sol/GalvaRegistry.json");
  fs.writeFileSync(
    dir + "/Registry.json",
    JSON.stringify(registryArtifacts, null, 1)
  );
  console.log("Registry artifacts saved!");
}
