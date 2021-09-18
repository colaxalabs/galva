/**
 * @type import('hardhat/config').HardhatUserConfig
 */
import "@nomiclabs/hardhat-waffle";
import { HardhatUserConfig } from "hardhat/config";

const config: HardhatUserConfig = {
  solidity: {
    version: "0.7.3",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  paths: {
    sources: "./src/contracts",
  },
  mocha: {
    timeout: 20000,
  },
};

export default config;
