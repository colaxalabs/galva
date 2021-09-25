import { ethers } from "hardhat";
import { Signer, Contract } from "ethers";

import { signParcel } from "../utils";

const { expect } = require("chai");

const dummyTitle = "111/v0/43x/50300";
const dummyTitle2 = "v0/xyz/50300";
const zeroRights = ethers.constants.Zero;
const addressZero = ethers.constants.AddressZero;
const parcelSize = ethers.utils.parseUnits("0.75");

let accounts: Signer[], registry: Contract, sender: Signer;

async function setupContracts() {
  // get signers
  accounts = await ethers.getSigners();
  sender = accounts[2];

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

describe("GalvaRegistry:attestProperty", () => {
  before("setup contracts", setupContracts);

  it("should correctly upload land parcel", async () => {
    // sign property
    const { attestor, signer } = await signParcel(dummyTitle, sender);
    await registry.attestProperty(dummyTitle, parcelSize, attestor, signer);

    const consumedRight = await registry.blockchainConsumedRights();
    const accountRights = await registry.addressRights(signer);
    expect(consumedRight).to.eq(parcelSize);
    expect(accountRights).to.eq(parcelSize);
  });

  it("should correctly reject duplicate property attestation", async () => {
    // sign property
    const { attestor, signer } = await signParcel(dummyTitle, sender);
    await expect(
      registry.attestProperty(dummyTitle, parcelSize, attestor, signer)
    ).to.reverted;
  });
});

describe("GalvaRegistry:recordExists", () => {
  before("setup contracts", setupContracts);

  it("should correctly assert if property already exists", async () => {
    // sign property
    const { attestor, signer } = await signParcel(dummyTitle, sender);
    // let's not attest this property
    const { attestor: anotherAttestor } = await signParcel(dummyTitle2, sender);

    await registry.attestProperty(dummyTitle, parcelSize, attestor, signer);

    expect(await registry.recordExists(attestor)).to.be.true;
    expect(await registry.recordExists(anotherAttestor)).to.be.false;
  });
});

describe("GalvaRegistry:addressRights", () => {
  before("setup contracts", setupContracts);

  it("should correctly return rights for an account", async () => {
    expect(await registry.addressRights(await sender.getAddress())).to.eq(
      zeroRights
    );
    expect(await registry.addressRights(await accounts[4].getAddress())).to.eq(
      zeroRights
    );
  });
});

describe("GalvaRegistry:blockchainConsumedRights", () => {
  before("setup contracts", setupContracts);

  it("should correctly return total blockchain consumed rights", async () => {
    expect(await registry.blockchainConsumedRights()).to.eq(zeroRights);
  });
});

describe("GalvaRegistry:claimOwnership", () => {
  before("setup contracts", setupContracts);

  it("should correctly claim ownership to property", async () => {
    // sign property
    const { attestor, signer } = await signParcel(dummyTitle, sender);
    // register property
    await registry.attestProperty(dummyTitle, parcelSize, attestor, signer);
    // split property signature
    const sig = ethers.utils.splitSignature(attestor);
    const owner = await registry.claimOwnership(
      dummyTitle,
      sig.v,
      sig.r,
      sig.s
    );
    expect(owner).to.eq(signer);
  });
});
