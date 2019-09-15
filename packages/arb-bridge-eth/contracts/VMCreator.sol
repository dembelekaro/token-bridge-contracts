/*
 * Copyright 2019, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "./vm/VMTracker.sol";


contract VMCreator {
    event VMCreated(
        address vmAddress,
        uint32 gracePeriod,
        uint128 escrowRequired,
        address escrowCurrency,
        uint32 maxExecutionSteps,
        bytes32 vmState,
        address owner,
        address[] validators
    );

    address globalInboxAddress;
    address challengeManagerAddress;

    constructor(address _globalInboxAddress, address _challengeManagerAddress) public {
        globalInboxAddress = _globalInboxAddress;
        challengeManagerAddress = _challengeManagerAddress;
    }

    function launchVM(
        bytes32 _vmState,
        uint32 _gracePeriod,
        uint32 _maxExecutionSteps,
        uint128 _escrowRequired,
        address _escrowCurrency,
        address payable _owner,
        address[] memory _validatorKeys
    )
        public
    {
        VMTracker vm = new VMTracker(
            _vmState,
            _gracePeriod,
            _maxExecutionSteps,
            _escrowRequired,
            _escrowCurrency,
            _owner,
            challengeManagerAddress,
            globalInboxAddress,
            _validatorKeys
        );
        emit VMCreated(
            address(vm),
            _gracePeriod,
            _escrowRequired,
            _escrowCurrency,
            _maxExecutionSteps,
            _vmState,
            _owner,
            _validatorKeys
        );
    }
}
