import { ethers } from "hardhat";
import { Signer, Contract } from "ethers";
const { expect } = require("chai");

const dummyTitle = "111/v0/43x/50300";
const zeroRights = ethers.constants.Zero;
const addressZero = ethers.constants.AddressZero;

let accounts: Signer[], registry: Contract;

async function setupContracts() {
  // get signers
  accounts = await ethers.getSigners();

  //deploy registry contract
  registry = await (await ethers.getContractFactory("GalvaRegistry"))
    .connect(accounts[0])
    .deploy();
}

describe("GalvaRegistry", () => {
  before("setup contracts", setupContracts);

  it("should always reject ether send to registry", async () => {
    const user = accounts[1]; // account that will send tx
    await expect(user.sendTransaction({ to: registry.address, value: 1 })).to.be
      .reverted;
  });
});

describe("GalvaRegistry:Init", () => {
  before("setup contracts", setupContracts);

  it("should correctly get blockchain consumed rights", async () => {
    const blockchainRights = await registry.blockchainConsumedRights();

    expect(blockchainRights).to.equal(zeroRights);
  });

  it("should correctly get address consumed rights", async () => {
    const forWho = await accounts[2].getAddress();

    expect(await registry.addressRights(forWho)).to.equal(zeroRights);
  });
});
