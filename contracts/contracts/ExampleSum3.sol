pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract ExampleSum3 {
    function sum3(uint256 a, uint256 b, uint256 c) public view returns (bytes memory h) {
        (bool ok, bytes memory out) = address(0x0b).staticcall(abi.encode(a,b,c));
        require(ok, "precompile call failed");

        console.logString("log out:");
        console.logBytes(out);

//        h = abi.decode(out, (bytes));
        h = out;
    }
}
