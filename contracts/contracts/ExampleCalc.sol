pragma solidity ^0.8.0;

import "hardhat/console.sol";

//contract ExampleSum3 {
//    function sum3(uint256 a, uint256 b, uint256 c) public view returns (bytes memory h) {
//        (bool ok, bytes memory out) = address(0x0b).staticcall(abi.encode(a,b,c));
//        require(ok, "precompile call failed");
//
//        console.logString("log out:");
//        console.logBytes(out);
//
////        h = abi.decode(out, (bytes));
//        h = out;
//    }
//}

interface ICalc {
    function calcSum (uint256 a, uint256 b) external;
    function calcDiff(uint256 a, uint256 b) external;

    function getSum () external view returns (uint256 result);
    function getDiff() external view returns (uint256 result);
}

address constant CALC_ADDRESS = 0x0300000000000000000000000000000000000003;

contract ExampleCalc {
    ICalc calc = ICalc(CALC_ADDRESS);

    function calcSum (uint256 a, uint256 b) external {
        calc.calcSum(a, b);
    }
    function calcDiff(uint256 a, uint256 b) external {
        calc.calcDiff(a, b);
    }

    function getSum() external view returns (uint256 result) {
        return calc.getSum();
    }
    function getDiff() external view returns (uint256 result) {
        return calc.getDiff();
    }
}