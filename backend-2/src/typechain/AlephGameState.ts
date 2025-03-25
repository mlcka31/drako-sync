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
  EventFragment,
  AddressLike,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedLogDescription,
  TypedListener,
  TypedContractMethod,
} from "./common";

export declare namespace AlephGameState {
  export type MessageStruct = {
    messagePrice: BigNumberish;
    sender: AddressLike;
    content: string;
    timestamp: BigNumberish;
  };

  export type MessageStructOutput = [
    messagePrice: bigint,
    sender: string,
    content: string,
    timestamp: bigint
  ] & {
    messagePrice: bigint;
    sender: string;
    content: string;
    timestamp: bigint;
  };
}

export interface AlephGameStateInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "PRECISION"
      | "abort"
      | "adminAddress"
      | "aiAgentAddress"
      | "coeffIncrease"
      | "gameState"
      | "getMessages"
      | "initGame"
      | "messagePrice"
      | "messages"
      | "payoutPrize"
      | "prizePool"
      | "prompt"
      | "promptSet"
      | "reply"
      | "sendMessage"
      | "setPrompt"
      | "stopGameAndWithdraw"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "AIAgentSetAddress"
      | "AgentReplied"
      | "MessageSent"
      | "PrizePayout"
  ): EventFragment;

  encodeFunctionData(functionFragment: "PRECISION", values?: undefined): string;
  encodeFunctionData(functionFragment: "abort", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "adminAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "aiAgentAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "coeffIncrease",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "gameState", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "getMessages",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "initGame",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "messagePrice",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "messages",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "payoutPrize",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "prizePool", values?: undefined): string;
  encodeFunctionData(functionFragment: "prompt", values?: undefined): string;
  encodeFunctionData(functionFragment: "promptSet", values?: undefined): string;
  encodeFunctionData(functionFragment: "reply", values: [string]): string;
  encodeFunctionData(functionFragment: "sendMessage", values: [string]): string;
  encodeFunctionData(functionFragment: "setPrompt", values: [string]): string;
  encodeFunctionData(
    functionFragment: "stopGameAndWithdraw",
    values?: undefined
  ): string;

  decodeFunctionResult(functionFragment: "PRECISION", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "abort", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "adminAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "aiAgentAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "coeffIncrease",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "gameState", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getMessages",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "initGame", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "messagePrice",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "messages", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "payoutPrize",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "prizePool", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "prompt", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "promptSet", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "reply", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "sendMessage",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "setPrompt", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "stopGameAndWithdraw",
    data: BytesLike
  ): Result;
}

export namespace AIAgentSetAddressEvent {
  export type InputTuple = [aiAgent: AddressLike];
  export type OutputTuple = [aiAgent: string];
  export interface OutputObject {
    aiAgent: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace AgentRepliedEvent {
  export type InputTuple = [content: string, timestamp: BigNumberish];
  export type OutputTuple = [content: string, timestamp: bigint];
  export interface OutputObject {
    content: string;
    timestamp: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace MessageSentEvent {
  export type InputTuple = [
    sender: AddressLike,
    content: string,
    timestamp: BigNumberish
  ];
  export type OutputTuple = [
    sender: string,
    content: string,
    timestamp: bigint
  ];
  export interface OutputObject {
    sender: string;
    content: string;
    timestamp: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace PrizePayoutEvent {
  export type InputTuple = [winner: AddressLike, amount: BigNumberish];
  export type OutputTuple = [winner: string, amount: bigint];
  export interface OutputObject {
    winner: string;
    amount: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface AlephGameState extends BaseContract {
  connect(runner?: ContractRunner | null): AlephGameState;
  waitForDeployment(): Promise<this>;

  interface: AlephGameStateInterface;

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

  PRECISION: TypedContractMethod<[], [bigint], "view">;

  abort: TypedContractMethod<[], [void], "nonpayable">;

  adminAddress: TypedContractMethod<[], [string], "view">;

  aiAgentAddress: TypedContractMethod<[], [string], "view">;

  coeffIncrease: TypedContractMethod<[], [bigint], "view">;

  gameState: TypedContractMethod<[], [bigint], "view">;

  getMessages: TypedContractMethod<
    [],
    [AlephGameState.MessageStructOutput[]],
    "view"
  >;

  initGame: TypedContractMethod<
    [_aiAgentAddress: AddressLike],
    [void],
    "nonpayable"
  >;

  messagePrice: TypedContractMethod<[], [bigint], "view">;

  messages: TypedContractMethod<
    [arg0: BigNumberish],
    [
      [bigint, string, string, bigint] & {
        messagePrice: bigint;
        sender: string;
        content: string;
        timestamp: bigint;
      }
    ],
    "view"
  >;

  payoutPrize: TypedContractMethod<[], [void], "nonpayable">;

  prizePool: TypedContractMethod<[], [bigint], "view">;

  prompt: TypedContractMethod<[], [string], "view">;

  promptSet: TypedContractMethod<[], [boolean], "view">;

  reply: TypedContractMethod<[_content: string], [void], "nonpayable">;

  sendMessage: TypedContractMethod<[_content: string], [void], "payable">;

  setPrompt: TypedContractMethod<[_prompt: string], [void], "nonpayable">;

  stopGameAndWithdraw: TypedContractMethod<[], [void], "nonpayable">;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "PRECISION"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "abort"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "adminAddress"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "aiAgentAddress"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "coeffIncrease"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "gameState"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "getMessages"
  ): TypedContractMethod<[], [AlephGameState.MessageStructOutput[]], "view">;
  getFunction(
    nameOrSignature: "initGame"
  ): TypedContractMethod<[_aiAgentAddress: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "messagePrice"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "messages"
  ): TypedContractMethod<
    [arg0: BigNumberish],
    [
      [bigint, string, string, bigint] & {
        messagePrice: bigint;
        sender: string;
        content: string;
        timestamp: bigint;
      }
    ],
    "view"
  >;
  getFunction(
    nameOrSignature: "payoutPrize"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "prizePool"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "prompt"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "promptSet"
  ): TypedContractMethod<[], [boolean], "view">;
  getFunction(
    nameOrSignature: "reply"
  ): TypedContractMethod<[_content: string], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "sendMessage"
  ): TypedContractMethod<[_content: string], [void], "payable">;
  getFunction(
    nameOrSignature: "setPrompt"
  ): TypedContractMethod<[_prompt: string], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "stopGameAndWithdraw"
  ): TypedContractMethod<[], [void], "nonpayable">;

  getEvent(
    key: "AIAgentSetAddress"
  ): TypedContractEvent<
    AIAgentSetAddressEvent.InputTuple,
    AIAgentSetAddressEvent.OutputTuple,
    AIAgentSetAddressEvent.OutputObject
  >;
  getEvent(
    key: "AgentReplied"
  ): TypedContractEvent<
    AgentRepliedEvent.InputTuple,
    AgentRepliedEvent.OutputTuple,
    AgentRepliedEvent.OutputObject
  >;
  getEvent(
    key: "MessageSent"
  ): TypedContractEvent<
    MessageSentEvent.InputTuple,
    MessageSentEvent.OutputTuple,
    MessageSentEvent.OutputObject
  >;
  getEvent(
    key: "PrizePayout"
  ): TypedContractEvent<
    PrizePayoutEvent.InputTuple,
    PrizePayoutEvent.OutputTuple,
    PrizePayoutEvent.OutputObject
  >;

  filters: {
    "AIAgentSetAddress(address)": TypedContractEvent<
      AIAgentSetAddressEvent.InputTuple,
      AIAgentSetAddressEvent.OutputTuple,
      AIAgentSetAddressEvent.OutputObject
    >;
    AIAgentSetAddress: TypedContractEvent<
      AIAgentSetAddressEvent.InputTuple,
      AIAgentSetAddressEvent.OutputTuple,
      AIAgentSetAddressEvent.OutputObject
    >;

    "AgentReplied(string,uint256)": TypedContractEvent<
      AgentRepliedEvent.InputTuple,
      AgentRepliedEvent.OutputTuple,
      AgentRepliedEvent.OutputObject
    >;
    AgentReplied: TypedContractEvent<
      AgentRepliedEvent.InputTuple,
      AgentRepliedEvent.OutputTuple,
      AgentRepliedEvent.OutputObject
    >;

    "MessageSent(address,string,uint256)": TypedContractEvent<
      MessageSentEvent.InputTuple,
      MessageSentEvent.OutputTuple,
      MessageSentEvent.OutputObject
    >;
    MessageSent: TypedContractEvent<
      MessageSentEvent.InputTuple,
      MessageSentEvent.OutputTuple,
      MessageSentEvent.OutputObject
    >;

    "PrizePayout(address,uint256)": TypedContractEvent<
      PrizePayoutEvent.InputTuple,
      PrizePayoutEvent.OutputTuple,
      PrizePayoutEvent.OutputObject
    >;
    PrizePayout: TypedContractEvent<
      PrizePayoutEvent.InputTuple,
      PrizePayoutEvent.OutputTuple,
      PrizePayoutEvent.OutputObject
    >;
  };
}
