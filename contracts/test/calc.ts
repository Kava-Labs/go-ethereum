import {
    time,
    loadFixture,
} from "@nomicfoundation/hardhat-toolbox/network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";

const calcAbi = [
    "function getSum() external view returns (uint256)"
];

describe("calc", function() {
    // it("Should properly calculate sum of 2 numbers", async function () {
    //     const ExampleCalc = await ethers.getContractFactory("ExampleCalc")
    //     const exampleCalc = await ExampleCalc.deploy();
    //
    //     console.log(await exampleCalc.calcSum(3, 4));
    //
    //     expect(await exampleCalc.calcSum(3, 4)).to.equal("0x0000000000000000000000000000000000000000000000000000000000000007");
    // })

    it("Should properly calculate sum of 2 numbers", async function () {
        const provider = ethers.provider;

        // The Contract object
        const calcContract = new ethers.Contract("0x0300000000000000000000000000000000000003", calcAbi, provider);

        console.log(await provider.getBlockNumber());

        console.log(await calcContract.getSum());
    })
})
