/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type { Signer, ContractDeployTransaction, ContractRunner } from "ethers";
import type { NonPayableOverrides } from "../common";
import type { ExampleSum3, ExampleSum3Interface } from "../ExampleSum3";

const _abi = [
  {
    inputs: [
      {
        internalType: "uint256",
        name: "a",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "b",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "c",
        type: "uint256",
      },
    ],
    name: "calcSum3",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "getSum3",
    outputs: [
      {
        internalType: "bytes",
        name: "",
        type: "bytes",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b5061053f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638e2a5b251461003b578063a104464e14610057575b600080fd5b61005560048036038101906100509190610282565b610075565b005b61005f61015f565b60405161006c9190610365565b60405180910390f35b600073030000000000000000000000000000000000000473ffffffffffffffffffffffffffffffffffffffff168484846040516020016100b793929190610396565b6040516020818303038152906040526040516100d39190610409565b6000604051808303816000865af19150503d8060008114610110576040519150601f19603f3d011682016040523d82523d6000602084013e610115565b606091505b5050905080610159576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610150906104a3565b60405180910390fd5b50505050565b606060008073030000000000000000000000000000000000000473ffffffffffffffffffffffffffffffffffffffff1660405160200161019e906104e9565b6040516020818303038152906040526040516101ba9190610409565b600060405180830381855afa9150503d80600081146101f5576040519150601f19603f3d011682016040523d82523d6000602084013e6101fa565b606091505b50915091508161023f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610236906104a3565b60405180910390fd5b809250505090565b600080fd5b6000819050919050565b61025f8161024c565b811461026a57600080fd5b50565b60008135905061027c81610256565b92915050565b60008060006060848603121561029b5761029a610247565b5b60006102a98682870161026d565b93505060206102ba8682870161026d565b92505060406102cb8682870161026d565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b60005b8381101561030f5780820151818401526020810190506102f4565b60008484015250505050565b6000601f19601f8301169050919050565b6000610337826102d5565b61034181856102e0565b93506103518185602086016102f1565b61035a8161031b565b840191505092915050565b6000602082019050818103600083015261037f818461032c565b905092915050565b6103908161024c565b82525050565b60006060820190506103ab6000830186610387565b6103b86020830185610387565b6103c56040830184610387565b949350505050565b600081905092915050565b60006103e3826102d5565b6103ed81856103cd565b93506103fd8185602086016102f1565b80840191505092915050565b600061041582846103d8565b915081905092915050565b600082825260208201905092915050565b7f63616c6c20746f20707265636f6d70696c656420636f6e74726163742066616960008201527f6c65640000000000000000000000000000000000000000000000000000000000602082015250565b600061048d602383610420565b915061049882610431565b604082019050919050565b600060208201905081810360008301526104bc81610480565b9050919050565b50565b60006104d3600083610420565b91506104de826104c3565b600082019050919050565b60006020820190508181036000830152610502816104c6565b905091905056fea26469706673582212200c14e4922a326cc0fd4c69c34dc73955dd9a31bb466a3c0f96da885f31ab1c4764736f6c63430008130033";

type ExampleSum3ConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: ExampleSum3ConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class ExampleSum3__factory extends ContractFactory {
  constructor(...args: ExampleSum3ConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(overrides || {});
  }
  override deploy(overrides?: NonPayableOverrides & { from?: string }) {
    return super.deploy(overrides || {}) as Promise<
      ExampleSum3 & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): ExampleSum3__factory {
    return super.connect(runner) as ExampleSum3__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): ExampleSum3Interface {
    return new Interface(_abi) as ExampleSum3Interface;
  }
  static connect(address: string, runner?: ContractRunner | null): ExampleSum3 {
    return new Contract(address, _abi, runner) as unknown as ExampleSum3;
  }
}
