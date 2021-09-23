/**
 * @type import('hardhat/config').HardhatUserConfig
 */
import { HardhatUserConfig } from "hardhat/config";
import "@nomiclabs/hardhat-waffle";

require("./scripts/deploy.ts");

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
