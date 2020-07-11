/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* eslint-env node, mocha */

import { ethers } from '@nomiclabs/buidler'
import { Signer } from 'ethers'
import * as chai from 'chai'
import chaiAsPromised from 'chai-as-promised'
import { MessageTester } from '../build/types/MessageTester'
import { ArbValue, Message } from 'arb-provider-ethers'

chai.use(chaiAsPromised)

const { assert, expect } = chai

let accounts: Signer[]
let messageTester: MessageTester

const token = '0xc7711f36b2C13E00821fFD9EC54B04A60AEfbd1b'
const dest = '0xFeCd3992654bFC565c3aFc6C4d7b14dCe603EbF5'
const sender = '0x7073c616a8A3F277Ea4511fCe9EBB2656a1b87B8'
const value = ethers.utils.bigNumberify(563356543)

describe('Messages', async () => {
  before(async () => {
    accounts = await ethers.getSigners()

    const MessageTester = await ethers.getContractFactory('MessageTester')
    messageTester = (await MessageTester.deploy()) as MessageTester
    await messageTester.deployed()
  })

  it('can handle eth messages', async () => {
    const msg = new Message.EthMessage(dest, value)
    const outMsg = new Message.OutgoingMessage(msg, sender)
    const msgData = ArbValue.marshal(outMsg.asValue())
    const {
      0: valid,
      1: offset,
      2: kind,
      3: calculatedSender,
      4: data,
    } = await messageTester.unmarshalOutgoingMessage(msgData, 0)

    assert.isTrue(valid, 'did not deserialize outgoing message correctly')
    assert.equal(offset, ethers.utils.bigNumberify(msgData.length))
    assert.equal(kind, msg.kind, 'Incorrect message type')
    assert.equal(calculatedSender, sender, 'Incorrect sender')

    const { valid: valid2, message } = await messageTester.parseEthMessage(data)
    assert.isTrue(valid2, 'did not parse eth message correctly')
    assert.equal(message.dest, dest, 'Incorrect dest')
    expect(message.value, 'Incorrect value').to.equal(value)
  })

  it('can handle erc20 messages', async () => {
    const msg = new Message.ERC20Message(token, dest, value)
    const outMsg = new Message.OutgoingMessage(msg, sender)
    const msgData = ArbValue.marshal(outMsg.asValue())
    const {
      0: valid,
      1: offset,
      2: kind,
      3: calculatedSender,
      4: data,
    } = await messageTester.unmarshalOutgoingMessage(msgData, 0)

    assert.isTrue(valid, 'did not deserialize outgoing message correctly')
    assert.equal(offset, ethers.utils.bigNumberify(msgData.length))
    assert.equal(kind, msg.kind, 'Incorrect message type')
    assert.equal(calculatedSender, sender, 'Incorrect sender')

    const { valid: valid2, message } = await messageTester.parseERC20Message(
      data
    )
    assert.isTrue(valid2, 'did not parse eth message correctly')
    assert.equal(message.dest, dest, 'Incorrect dest')
    expect(message.value, 'Incorrect value').to.equal(value)
  })

  it('can handle erc721 messages', async () => {
    const msg = new Message.ERC721Message(token, dest, value)
    const outMsg = new Message.OutgoingMessage(msg, sender)
    const msgData = ArbValue.marshal(outMsg.asValue())
    const {
      0: valid,
      1: offset,
      2: kind,
      3: calculatedSender,
      4: data,
    } = await messageTester.unmarshalOutgoingMessage(msgData, 0)

    assert.isTrue(valid, 'did not deserialize outgoing message correctly')
    assert.equal(offset, ethers.utils.bigNumberify(msgData.length))
    assert.equal(kind, msg.kind, 'Incorrect message type')
    assert.equal(calculatedSender, sender, 'Incorrect sender')

    const { valid: valid2, message } = await messageTester.parseERC721Message(
      data
    )
    assert.isTrue(valid2, 'did not parse eth message correctly')
    assert.equal(message.dest, dest, 'Incorrect dest')
    expect(message.id, 'Incorrect id').to.equal(value)
  })
})
