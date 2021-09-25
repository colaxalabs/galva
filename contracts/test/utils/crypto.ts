import { ethers } from "hardhat";
import { BigNumberish, Signer } from "ethers";

const abiCoder = ethers.utils.defaultAbiCoder;

// Sign property details
export async function signParcel(title: string, account: Signer) {
  // encode property details
  const payload = abiCoder.encode(["string"], [title]);
  const payloadHash = ethers.utils.keccak256(payload); // hash payload
  // generate 32 byte Uint8Array for signer
  const payloadMessage = ethers.utils.arrayify(payloadHash);
  // sign message
  const attestor = await account.signMessage(payloadMessage);
  // recover signer
  const sign = ethers.utils.splitSignature(attestor);
  const signer = ethers.utils.verifyMessage(payloadMessage, sign);
  return { attestor, signer };
}
