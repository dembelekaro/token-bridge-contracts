/* Generated by ts-generator ver. 0.0.8 */
/* tslint:disable */

import { Contract, ContractTransaction, EventFilter, Signer } from 'ethers';
import { Listener, Provider } from 'ethers/providers';
import { Arrayish, BigNumber, BigNumberish, Interface } from 'ethers/utils';
import { TransactionOverrides, TypedEventDescription, TypedFunctionDescription } from '.';

interface GlobalPendingInboxInterface extends Interface {
    functions: {
        onERC721Received: TypedFunctionDescription<{
            encode([, _from, _tokenId]: [string, string, BigNumberish, Arrayish]): string;
        }>;

        depositERC20: TypedFunctionDescription<{
            encode([_tokenContract, _value]: [string, BigNumberish]): string;
        }>;

        withdrawERC20: TypedFunctionDescription<{
            encode([_tokenContract, _value]: [string, BigNumberish]): string;
        }>;

        depositEth: TypedFunctionDescription<{
            encode([_destination]: [string]): string;
        }>;

        withdrawEth: TypedFunctionDescription<{
            encode([_value]: [BigNumberish]): string;
        }>;

        depositERC721: TypedFunctionDescription<{
            encode([_tokenContract, _tokenId]: [string, BigNumberish]): string;
        }>;

        withdrawERC721: TypedFunctionDescription<{
            encode([_tokenContract, _tokenId]: [string, BigNumberish]): string;
        }>;

        pullPendingMessages: TypedFunctionDescription<{ encode([]: []): string }>;

        sendMessages: TypedFunctionDescription<{
            encode([_messages]: [Arrayish]): string;
        }>;

        registerForInbox: TypedFunctionDescription<{ encode([]: []): string }>;

        sendMessage: TypedFunctionDescription<{
            encode([_destination, _tokenType, _amount, _data]: [string, Arrayish, BigNumberish, Arrayish]): string;
        }>;

        forwardMessage: TypedFunctionDescription<{
            encode([_destination, _tokenType, _amount, _data, _signature]: [
                string,
                Arrayish,
                BigNumberish,
                Arrayish,
                Arrayish,
            ]): string;
        }>;

        sendEthMessage: TypedFunctionDescription<{
            encode([_destination, _data]: [string, Arrayish]): string;
        }>;
    };

    events: {
        MessageDelivered: TypedEventDescription<{
            encodeTopics([vmId, sender, tokenType, value, data]: [string | null, null, null, null, null]): string[];
        }>;
    };
}

export class GlobalPendingInbox extends Contract {
    connect(signerOrProvider: Signer | Provider | string): GlobalPendingInbox;
    attach(addressOrName: string): GlobalPendingInbox;
    deployed(): Promise<GlobalPendingInbox>;

    on(event: EventFilter | string, listener: Listener): GlobalPendingInbox;
    once(event: EventFilter | string, listener: Listener): GlobalPendingInbox;
    addListener(eventName: EventFilter | string, listener: Listener): GlobalPendingInbox;
    removeAllListeners(eventName: EventFilter | string): GlobalPendingInbox;
    removeListener(eventName: any, listener: Listener): GlobalPendingInbox;

    interface: GlobalPendingInboxInterface;

    functions: {
        getNFTTokens(
            _owner: string,
        ): Promise<{
            0: (string)[];
            1: (BigNumber)[];
        }>;

        getTokenBalances(
            _owner: string,
        ): Promise<{
            0: (string)[];
            1: (BigNumber)[];
        }>;

        getTokenBalance(_tokenContract: string, _owner: string): Promise<BigNumber>;

        hasNFT(_tokenContract: string, _owner: string, _tokenId: BigNumberish): Promise<boolean>;

        generateSentMessageHash(
            _dest: string,
            _data: Arrayish,
            _tokenType: Arrayish,
            _value: BigNumberish,
            _sender: string,
        ): Promise<string>;

        onERC721Received(
            arg0: string,
            _from: string,
            _tokenId: BigNumberish,
            arg3: Arrayish,
            overrides?: TransactionOverrides,
        ): Promise<ContractTransaction>;

        depositERC20(
            _tokenContract: string,
            _value: BigNumberish,
            overrides?: TransactionOverrides,
        ): Promise<ContractTransaction>;

        withdrawERC20(
            _tokenContract: string,
            _value: BigNumberish,
            overrides?: TransactionOverrides,
        ): Promise<ContractTransaction>;

        depositEth(_destination: string, overrides?: TransactionOverrides): Promise<ContractTransaction>;

        withdrawEth(_value: BigNumberish, overrides?: TransactionOverrides): Promise<ContractTransaction>;

        depositERC721(
            _tokenContract: string,
            _tokenId: BigNumberish,
            overrides?: TransactionOverrides,
        ): Promise<ContractTransaction>;

        withdrawERC721(
            _tokenContract: string,
            _tokenId: BigNumberish,
            overrides?: TransactionOverrides,
        ): Promise<ContractTransaction>;

        pullPendingMessages(overrides?: TransactionOverrides): Promise<ContractTransaction>;

        sendMessages(_messages: Arrayish, overrides?: TransactionOverrides): Promise<ContractTransaction>;

        registerForInbox(overrides?: TransactionOverrides): Promise<ContractTransaction>;

        sendMessage(
            _destination: string,
            _tokenType: Arrayish,
            _amount: BigNumberish,
            _data: Arrayish,
            overrides?: TransactionOverrides,
        ): Promise<ContractTransaction>;

        forwardMessage(
            _destination: string,
            _tokenType: Arrayish,
            _amount: BigNumberish,
            _data: Arrayish,
            _signature: Arrayish,
            overrides?: TransactionOverrides,
        ): Promise<ContractTransaction>;

        sendEthMessage(
            _destination: string,
            _data: Arrayish,
            overrides?: TransactionOverrides,
        ): Promise<ContractTransaction>;
    };

    filters: {
        MessageDelivered(vmId: string | null, sender: null, tokenType: null, value: null, data: null): EventFilter;
    };

    estimate: {
        onERC721Received(arg0: string, _from: string, _tokenId: BigNumberish, arg3: Arrayish): Promise<BigNumber>;

        depositERC20(_tokenContract: string, _value: BigNumberish): Promise<BigNumber>;

        withdrawERC20(_tokenContract: string, _value: BigNumberish): Promise<BigNumber>;

        depositEth(_destination: string): Promise<BigNumber>;

        withdrawEth(_value: BigNumberish): Promise<BigNumber>;

        depositERC721(_tokenContract: string, _tokenId: BigNumberish): Promise<BigNumber>;

        withdrawERC721(_tokenContract: string, _tokenId: BigNumberish): Promise<BigNumber>;

        pullPendingMessages(): Promise<BigNumber>;

        sendMessages(_messages: Arrayish): Promise<BigNumber>;

        registerForInbox(): Promise<BigNumber>;

        sendMessage(
            _destination: string,
            _tokenType: Arrayish,
            _amount: BigNumberish,
            _data: Arrayish,
        ): Promise<BigNumber>;

        forwardMessage(
            _destination: string,
            _tokenType: Arrayish,
            _amount: BigNumberish,
            _data: Arrayish,
            _signature: Arrayish,
        ): Promise<BigNumber>;

        sendEthMessage(_destination: string, _data: Arrayish): Promise<BigNumber>;
    };
}
