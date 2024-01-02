/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedListener,
  TypedContractMethod,
} from "./common";

export interface ExampleSum3Interface extends Interface {
  getFunction(nameOrSignature: "calcSum3" | "getSum3"): FunctionFragment;

  encodeFunctionData(
    functionFragment: "calcSum3",
    values: [BigNumberish, BigNumberish, BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "getSum3", values?: undefined): string;

  decodeFunctionResult(functionFragment: "calcSum3", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "getSum3", data: BytesLike): Result;
}

export interface ExampleSum3 extends BaseContract {
  connect(runner?: ContractRunner | null): ExampleSum3;
  waitForDeployment(): Promise<this>;

  interface: ExampleSum3Interface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  calcSum3: TypedContractMethod<
    [a: BigNumberish, b: BigNumberish, c: BigNumberish],
    [void],
    "nonpayable"
  >;

  getSum3: TypedContractMethod<[], [string], "view">;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "calcSum3"
  ): TypedContractMethod<
    [a: BigNumberish, b: BigNumberish, c: BigNumberish],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "getSum3"
  ): TypedContractMethod<[], [string], "view">;

  filters: {};
}
