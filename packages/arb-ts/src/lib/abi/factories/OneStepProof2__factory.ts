/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer } from 'ethers'
import { Provider, TransactionRequest } from '@ethersproject/providers'
import { Contract, ContractFactory, Overrides } from '@ethersproject/contracts'

import type { OneStepProof2 } from '../OneStepProof2'

export class OneStepProof2__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer)
  }

  deploy(overrides?: Overrides): Promise<OneStepProof2> {
    return super.deploy(overrides || {}) as Promise<OneStepProof2>
  }
  getDeployTransaction(overrides?: Overrides): TransactionRequest {
    return super.getDeployTransaction(overrides || {})
  }
  attach(address: string): OneStepProof2 {
    return super.attach(address) as OneStepProof2
  }
  connect(signer: Signer): OneStepProof2__factory {
    return super.connect(signer) as OneStepProof2__factory
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): OneStepProof2 {
    return new Contract(address, _abi, signerOrProvider) as OneStepProof2
  }
}

const _abi = [
  {
    inputs: [
      {
        internalType: 'address[2]',
        name: 'bridges',
        type: 'address[2]',
      },
      {
        internalType: 'uint256',
        name: 'initialMessagesRead',
        type: 'uint256',
      },
      {
        internalType: 'bytes32[2]',
        name: 'accs',
        type: 'bytes32[2]',
      },
      {
        internalType: 'bytes',
        name: 'proof',
        type: 'bytes',
      },
      {
        internalType: 'bytes',
        name: 'bproof',
        type: 'bytes',
      },
    ],
    name: 'executeStep',
    outputs: [
      {
        internalType: 'uint64',
        name: 'gas',
        type: 'uint64',
      },
      {
        internalType: 'uint256',
        name: 'afterMessagesRead',
        type: 'uint256',
      },
      {
        internalType: 'bytes32[4]',
        name: 'fields',
        type: 'bytes32[4]',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address[2]',
        name: 'bridges',
        type: 'address[2]',
      },
      {
        internalType: 'uint256',
        name: 'initialMessagesRead',
        type: 'uint256',
      },
      {
        internalType: 'bytes32[2]',
        name: 'accs',
        type: 'bytes32[2]',
      },
      {
        internalType: 'bytes',
        name: 'proof',
        type: 'bytes',
      },
      {
        internalType: 'bytes',
        name: 'bproof',
        type: 'bytes',
      },
    ],
    name: 'executeStepDebug',
    outputs: [
      {
        internalType: 'string',
        name: 'startMachine',
        type: 'string',
      },
      {
        internalType: 'string',
        name: 'afterMachine',
        type: 'string',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'bytes',
        name: 'proof',
        type: 'bytes',
      },
    ],
    name: 'parseProof',
    outputs: [
      {
        internalType: 'bytes32[]',
        name: '',
        type: 'bytes32[]',
      },
      {
        internalType: 'bytes32[]',
        name: '',
        type: 'bytes32[]',
      },
      {
        internalType: 'bytes32[]',
        name: '',
        type: 'bytes32[]',
      },
      {
        internalType: 'bytes32[]',
        name: '',
        type: 'bytes32[]',
      },
    ],
    stateMutability: 'pure',
    type: 'function',
  },
]

const _bytecode =
  '0x608060405234801561001057600080fd5b506140de806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806347dda1d614610046578063793deea314610162578063eba67f6e14610329575b600080fd5b610112600480360360e081101561005c57600080fd5b604082013590606083019083018360c0810160a0820135600160201b81111561008457600080fd5b82018360208201111561009657600080fd5b803590602001918460018302840111600160201b831117156100b757600080fd5b919390929091602081019035600160201b8111156100d457600080fd5b8201836020820111156100e657600080fd5b803590602001918460018302840111600160201b8311171561010757600080fd5b5090925090506104d3565b604080516001600160401b03851681526020810184905290810182608080838360005b8381101561014d578181015183820152602001610135565b50505050905001935050505060405180910390f35b6102066004803603602081101561017857600080fd5b810190602081018135600160201b81111561019257600080fd5b8201836020820111156101a457600080fd5b803590602001918460018302840111600160201b831117156101c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610592945050505050565b6040518080602001806020018060200180602001858103855289818151815260200191508051906020019060200280838360005b8381101561025257818101518382015260200161023a565b50505050905001858103845288818151815260200191508051906020019060200280838360005b83811015610291578181015183820152602001610279565b50505050905001858103835287818151815260200191508051906020019060200280838360005b838110156102d05781810151838201526020016102b8565b50505050905001858103825286818151815260200191508051906020019060200280838360005b8381101561030f5781810151838201526020016102f7565b505050509050019850505050505050505060405180910390f35b6103f5600480360360e081101561033f57600080fd5b604082013590606083019083018360c0810160a0820135600160201b81111561036757600080fd5b82018360208201111561037957600080fd5b803590602001918460018302840111600160201b8311171561039a57600080fd5b919390929091602081019035600160201b8111156103b757600080fd5b8201836020820111156103c957600080fd5b803590602001918460018302840111600160201b831117156103ea57600080fd5b5090925090506105ce565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561043657818101518382015260200161041e565b50505050905090810190601f1680156104635780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561049657818101518382015260200161047e565b50505050905090810190601f1680156104c35780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6000806104de613ece565b6104e6613eec565b61056a8a8a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b9081908401838280828437600081840152601f19601f820116905080830192505050505050508f610693565b905061057581610b4a565b61057e81610f60565b935093509350509750975097945050505050565b6060806060806105a0613f7e565b6105a986610fc3565b80516020820151604083015160609093015191975095509093509150505b9193509193565b6060806105d9613eec565b61065d898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8d018190048102820181019092528b815292508b91508a9081908401838280828437600081840152601f19601f820116905080830192505050505050508e610693565b905061066881610b4a565b61067581604001516110bb565b925061068481606001516110bb565b91505097509795505050505050565b61069b613eec565b6000846000815181106106aa57fe5b602001015160f81c60f81b60f81c90506000856001815181106106c957fe5b602001015160f81c60f81b60f81c90506000866002815181106106e857fe5b016020015160f81c9050600360606004840160ff166001600160401b038111801561071257600080fd5b5060405190808252806020026020018201604052801561074c57816020015b610739613fa6565b8152602001906001900390816107315790505b50905060608360040160ff166001600160401b038111801561076d57600080fd5b506040519080825280602002602001820160405280156107a757816020015b610794613fa6565b81526020019060019003908161078c5790505b50905060005b8560ff168110156107e3576107c28b856113c4565b8483815181106107ce57fe5b602090810291909101015293506001016107ad565b5060005b8460ff1681101561081d576107fc8b856113c4565b83838151811061080857fe5b602090810291909101015293506001016107e7565b50610826613fe3565b6108308b85611586565b809250819550505060008b858151811061084657fe5b01602001516001959095019460f81c905061085f613eec565b6001600160a01b038b35811682526020808d0135909116908201526040810183905261088a83611626565b6060820152608081018f90528d3560a08201526020808f013560c0830152600060e0830181905260408051808201825260ff8c811682528185018a905261010086019190915281518083019092528a8116825292810187905261012084015283821660018114610140850152918b1661016084015261018083018f90526101c083018e90526101e08301526101a08201879052158061092c57508160ff166001145b6040518060400160405280600b81526020016a04241445f494d4d5f5459560ac1b815250906109d95760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561099e578181015183820152602001610986565b50505050905090810190601f1680156109cb5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506109e2613fa6565b60ff8316610a03576109fc8a83604001516000015161168f565b9050610aa3565b6000865111604051806040016040528060068152602001654e4f5f494d4d60d01b81525090610a735760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561099e578181015183820152602001610986565b50610aa08a8360400151600001518860018d0360ff1681518110610a9357fe5b60200260200101516116f3565b90505b610aac81611779565b60408301515260005b838a0360ff16811015610af457610aec878281518110610ad157fe5b602002602001015184604001516118e690919063ffffffff16565b600101610ab5565b5060005b8860ff16811015610b3557610b2d868281518110610b1257fe5b6020026020010151846040015161190090919063ffffffff16565b600101610af8565b50909f9e505050505050505050505050505050565b6000806000614041610b6385610160015160ff1661191a565b93509350935093506000841180610b7d5750846101400151155b8015610b8f5750610100850151518410155b80610bb757508461014001518015610ba5575083155b8015610bb75750610100850151516001145b6040518060400160405280600a815260200169535441434b5f4d414e5960b01b81525090610c265760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561099e578181015183820152602001610986565b50610120850151516040805180820190915260088152674155585f4d414e5960c01b602082015290841015610c9c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561099e578181015183820152602001610986565b5061010085015151841115610d5a57610cbb610cb6611a2a565b611779565b610ccc866060015160200151611779565b146040518060400160405280600d81526020016c535441434b5f4d495353494e4760981b81525090610d3f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561099e578181015183820152602001610986565b50610d4b856005611a71565b50610d5585611ae6565b610e1a565b61012085015151831115610df557610d73610cb6611a2a565b610d84866060015160400151611779565b146040518060400160405280600b81526020016a4155585f4d495353494e4760a81b81525090610d3f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561099e578181015183820152602001610986565b610dff8583611a71565b15610e0d57610d5585611ae6565b610e1a858263ffffffff16565b846101e0015115610ebf5760408051600160f81b6020808301919091526000602183018190526022808401919091528351808403909101815260429092019092528051910120606086015160c001511415610e8157610e7c8560600151611af1565b610ebf565b60006101e0860152606085015160c081015190526101408501518015610ea5575083155b610eb457610100850151600090525b610120850151600090525b60005b61010086015151811015610f0b57610f03866101000151602001518281518110610ee857fe5b602002602001015187606001516118e690919063ffffffff16565b600101610ec2565b5060005b61012086015151811015610f5857610f50866101200151602001518281518110610f3557fe5b6020026020010151876060015161190090919063ffffffff16565b600101610f0f565b505050505050565b600080610f6b613ece565b8360e0015184608001516040518060800160405280610f8d8860400151611afb565b8152602001610f9f8860600151611afb565b81526020018760a0015181526020018760c001518152509250925092509193909250565b610fcb613f7e565b606061100a8384600081518110610fde57fe5b602001015160f81c60f81b85600181518110610ff657fe5b01602001516001600160f81b031916611bbf565b90506060611037848560018151811061101f57fe5b602001015160f81c60f81b86600281518110610ff657fe5b90506060611064858660028151811061104c57fe5b602001015160f81c60f81b87600381518110610ff657fe5b90506060611091868760038151811061107957fe5b602001015160f81c60f81b88600481518110610ff657fe5b6040805160808101825295865260208601949094529284019190915250606082015290505b919050565b60606110ca8260000151611c5c565b6110df6110da8460200151611779565b611c5c565b6110ef6110da8560400151611779565b6110ff6110da8660600151611779565b61110f6110da8760800151611779565b61111c8760a00151611d2b565b6111298860c00151611c5c565b60405160200180806709ac2c6d0d2dcca560c31b81525060080188805190602001908083835b6020831061116e5780518252601f19909201916020918201910161114f565b51815160209384036101000a60001901801990921691161790526216100560e91b9190930190815289516003909101928a0191508083835b602083106111c55780518252601f1990920191602091820191016111a6565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528851600390910192890191508083835b6020831061121c5780518252601f1990920191602091820191016111fd565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528751600390910192880191508083835b602083106112735780518252601f199092019160209182019101611254565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528651600390910192870191508083835b602083106112ca5780518252601f1990920191602091820191016112ab565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528551600390910192860191508083835b602083106113215780518252601f199092019160209182019101611302565b51815160209384036101000a60001901801990921691161790526216100560e91b919093019081528451600390910192850191508083835b602083106113785780518252601f199092019160209182019101611359565b5181516020939093036101000a600019018019909116921691909117905261148560f11b92019182525060408051808303601d19018152600290920190529a9950505050505050505050565b60006113ce613fa6565b83518310611414576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806114218686611e05565b9150915061142d611e2c565b60ff168160ff1614156114615760006114468784611e31565b90935090508261145582611ea5565b9450945050505061157f565b611469611f65565b60ff168160ff16141561148b576114808683611f6a565b93509350505061157f565b61149361200c565b60ff168160ff1614156114bb5760006114ac8784611e31565b90935090508261145582612011565b6114c36120fd565b60ff168160ff1614156114da576114808683612102565b6114e2612197565b60ff168160ff161015801561150357506114fa61219c565b60ff168160ff16105b1561153f576000611512612197565b8203905060606115238289866121a1565b90945090508361153282612249565b955095505050505061157f565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b9250929050565b6000611590613fe3565b611598613fe3565b600060e08201819052806115ac8787611e31565b90965091506115bb8787612102565b602085015295506115cc8787612102565b604085015295506115dd87876113c4565b606085015295506115ee87876113c4565b608085015295506115ff8787611e31565b60a085015295506116108787611e31565b92845260c0840192909252509590945092505050565b61162e613fe3565b60405180610100016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e001518152509050919050565b611697613fa6565b6040805160608101825260ff8516815260208082018590528251600080825291810184526116ea938301916116e2565b6116cf613fa6565b8152602001906001900390816116c75790505b50905261238a565b90505b92915050565b6116fb613fa6565b604080516001808252818301909252606091816020015b61171a613fa6565b815260200190600190039081611712579050509050828160008151811061173d57fe5b602002602001018190525061176e60405180606001604052808760ff1681526020018681526020018381525061238a565b9150505b9392505050565b6000611783611e2c565b60ff16826080015160ff1614156117a657815161179f9061241a565b90506110b6565b6117ae611f65565b60ff16826080015160ff1614156117cc5761179f826020015161243e565b6117d46120fd565b60ff16826080015160ff1614156117f657815160a083015161179f919061253b565b6117fe612197565b60ff16826080015160ff16141561183757611817613fa6565b611824836040015161258c565b905061182f81611779565b9150506110b6565b61183f612704565b60ff16826080015160ff161415611858575080516110b6565b61186061200c565b60ff16826080015160ff1614156118a5575060608082015160408051607b602080830191909152818301939093528151808203830181529301905281519101206110b6565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b6118f4826020015182612709565b82602001819052505050565b61190e826040015182612709565b82604001819052505050565b6000808061404160a185141561193f57506002925060009150600a90506127876105c7565b60a285141561195d57506002925060009150600a905061283e6105c7565b60a385141561197b57506002925060009150600a90506128d26105c7565b60a485141561199957506003925060009150606490506129666105c7565b60a58514156119b75750600392506000915060649050612a576105c7565b60a68514156119d55750600392506000915060649050612b2b6105c7565b60708514156119f35750600292506000915060649050612bed6105c7565b60405162461bcd60e51b815260040180806020018281038252602c81526020018061407d602c913960400191505060405180910390fd5b611a32613fa6565b60408051600080825260208201909252611a6c91611a66565b611a53613fa6565b815260200190600190039081611a4b5790505b50612249565b905090565b6000816001600160401b0316836060015160a001511015611ab6575060e0820180516005016001600160401b03169052606082015160001960a09091015260016116ed565b5060e0820180516001600160401b039083018116909152606083015160a0018051918316909103905260006116ed565b60016101e090910152565b600160e090910152565b600060028260e001511415611b12575060006110b6565b60018260e001511415611b27575060016110b6565b81516020830151611b3790611779565b611b448460400151611779565b611b518560600151611779565b611b5e8660800151611779565b8660a001518760c00151604051602001808881526020018781526020018681526020018581526020018481526020018381526020018281526020019750505050505050506040516020818303038152906040528051906020012090506110b6565b606060f883811c9083901c81900360ff169082826001600160401b0381118015611be857600080fd5b50604051908082528060200260200182016040528015611c12578160200160208202803683370190505b50905060005b83811015611c5157611c2f88828501602002612d89565b60001b828281518110611c3e57fe5b6020908102919091010152600101611c18565b509695505050505050565b60408051818152606081810183529182919060208201818036833701905050905060005b6020811015611d24576000848260208110611c9757fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b611cbd82612dc9565b858560020281518110611ccc57fe5b60200101906001600160f81b031916908160001a905350611cec81612dc9565b858560020260010181518110611cfe57fe5b60200101906001600160f81b031916908160001a9053505060019092019150611c809050565b5092915050565b60608180611d525750506040805180820190915260018152600360fc1b60208201526110b6565b8060005b8115611d6a57600101600a82049150611d56565b6060816001600160401b0381118015611d8257600080fd5b506040519080825280601f01601f191660200182016040528015611dad576020820181803683370190505b50905060001982015b8415611dfb57600a850660300160f81b82828060019003935081518110611dd957fe5b60200101906001600160f81b031916908160001a905350600a85049450611db6565b5095945050505050565b60008082600101848481518110611e1857fe5b016020015190925060f81c90509250929050565b600090565b60008082845110158015611e49575060208385510310155b611e86576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301611e9a858563ffffffff612dfa16565b915091509250929050565b611ead613fa6565b6040805160c0810182528381528151606081018352600080825260208083018290528451828152808201865293949085019390830191611f03565b611ef0613fa6565b815260200190600190039081611ee85790505b50905281526020016000604051908082528060200260200182016040528015611f4657816020015b611f33613fa6565b815260200190600190039081611f2b5790505b5081526000602082018190526040820152600160609091015292915050565b600190565b6000611f74613fa6565b82600080611f80613fa6565b6000611f8c8986611e05565b9095509350611f9b8986611e05565b9095509250600160ff85161415611fbc57611fb689866113c4565b90955091505b611fc68986612e53565b9095509050600160ff85161415611ff15784611fe38483856116f3565b96509650505050505061157f565b84611ffc848361168f565b9650965050505050509250929050565b600c90565b612019613fa6565b6040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b815260200160006001600160401b038111801561206057600080fd5b5060405190808252806020026020018201604052801561209a57816020015b612087613fa6565b81526020019060019003908161207f5790505b509052815260200160006040519080825280602002602001820160405280156120dd57816020015b6120ca613fa6565b8152602001906001900390816120c25790505b50815260208101849052600c604082015260016060909101529050919050565b600290565b600061210c613fa6565b82845110158015612121575060408385510310155b61215e576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60008061216b8686612e53565b909450915061217a8685611e31565b90945090508361218a8383612e6a565b9350935050509250929050565b600390565b600d90565b60006060828160ff87166001600160401b03811180156121c057600080fd5b506040519080825280602002602001820160405280156121fa57816020015b6121e7613fa6565b8152602001906001900390816121df5790505b50905060005b8760ff168160ff16101561223c5761221887846113c4565b838360ff168151811061222757fe5b60209081029190910101529250600101612200565b5090969095509350505050565b612251613fa6565b61225b8251612f29565b6122ac576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156122e3578381815181106122c657fe5b602002602001015160a001518201915080806001019150506122b1565b506040518060c00160405280600081526020016040518060600160405280600060ff1681526020016000801b815260200160006001600160401b038111801561232b57600080fd5b5060405190808252806020026020018201604052801561236557816020015b612352613fa6565b81526020019060019003908161234a5790505b5090528152602081019490945260006040850152600360608501526080909301525090565b612392613fa6565b6040518060c001604052806000815260200183815260200160006001600160401b03811180156123c157600080fd5b506040519080825280602002602001820160405280156123fb57816020015b6123e8613fa6565b8152602001906001900390816123e05790505b5081526000602082015260016040820181905260609091015292915050565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061244f57fe5b6040820151516124b457612461611f65565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b9093166021850152602280850191909152825180850390910181526042909301909152815191012090506110b6565b6124bc611f65565b82600001516124e284604001516000815181106124d557fe5b6020026020010151611779565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b6000612545612197565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b612594613fa6565b6008825111156125e2576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516001600160401b03811180156125fb57600080fd5b50604051908082528060200260200182016040528015612625578160200160208202803683370190505b508051909150600160005b82811015612688576126478682815181106124d557fe5b84828151811061265357fe5b60200260200101818152505085818151811061266b57fe5b602002602001015160a00151820191508080600101915050612630565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156126cd5781810151838201526020016126b5565b50505050905001925050506040516020818303038152906040528051906020012090506126fa8183612e6a565b9695505050505050565b606490565b612711613fa6565b6040805160028082526060828101909352816020015b61272f613fa6565b815260200190600190039081612727579050509050828160008151811061275257fe5b6020026020010181905250838160018151811061276b57fe5b602002602001018190525061277f8161258c565b949350505050565b61278f613fa6565b61279d826101000151612f30565b90506127a7613fa6565b6127b5836101000151612f30565b90506127c082612f72565b15806127d257506127d081612f90565b155b156127e7576127e083612f9d565b505061283b565b8151600160401b116127fc576127e083612f9d565b600061281e82606001518460000151612819876101c00151610fc3565b612fa6565b905061283784610100015161283283611ea5565b612fc8565b5050505b50565b612846613fa6565b612854826101000151612f30565b905061285e613fa6565b61286c836101000151612f30565b905061287782612f72565b1580612889575061288781612f90565b155b15612897576127e083612f9d565b815167fffffffffffffff9116128b0576127e083612f9d565b600061281e826060015184600001516128cd876101c00151610fc3565b612ff2565b6128da613fa6565b6128e8826101000151612f30565b90506128f2613fa6565b612900836101000151612f30565b905061290b82612f72565b158061291d575061291b81612f90565b155b1561292b576127e083612f9d565b815167ffffffffffffffe111612944576127e083612f9d565b600061281e82606001518460000151612961876101c00151610fc3565b613151565b61296e613fa6565b61297c826101000151612f30565b9050612986613fa6565b612994836101000151612f30565b905061299e613fa6565b6129ac846101000151612f30565b90506129b783612f72565b15806129c957506129c782613284565b155b806129da57506129d881612f90565b155b156129f0576129e884612f9d565b50505061283b565b8251600160401b111580612a075750815161010011155b15612a15576129e884612f9d565b6000612a3c826060015185600001518560000151612a37896101c00151610fc3565b61328f565b9050612a5085610100015161283283612011565b5050505050565b612a5f613fa6565b612a6d826101000151612f30565b9050612a77613fa6565b612a85836101000151612f30565b9050612a8f613fa6565b612a9d846101000151612f30565b9050612aa883612f72565b1580612aba5750612ab882613284565b155b80612acb5750612ac981612f90565b155b15612ad9576129e884612f9d565b825167fffffffffffffff9111580612af657508151600160401b11155b15612b04576129e884612f9d565b6000612a3c826060015185600001518560000151612b26896101c00151610fc3565b6132d8565b612b33613fa6565b612b41826101000151612f30565b9050612b4b613fa6565b612b59836101000151612f30565b9050612b63613fa6565b612b71846101000151612f30565b9050612b7c83612f72565b1580612b8e5750612b8c82613284565b155b80612b9f5750612b9d81612f90565b155b15612bad576129e884612f9d565b825167ffffffffffffffe111612bc6576129e884612f9d565b6000612a3c826060015185600001518560000151612be8896101c00151610fc3565b613421565b612bf5613fa6565b612c03826101000151612f30565b9050612c0d613fa6565b612c1b836101000151612f30565b9050612c2682612f72565b1580612c385750612c3681612f90565b155b15612c46576127e083612f9d565b81516127101080612c5657508151155b15612c64576127e083612f9d565b82610180015151836101a001511415612cdf57612c9781606001518360000151612c92866101c00151610fc3565b6134f3565b15612cd6576040805162461bcd60e51b815260206004820152600a602482015269084aa8cbe988a9c8ea8960b31b604482015290519081900360640190fd5b6127e083612f9d565b6101a083015182516101808501516000612cfa828585613547565b905080612d0686611779565b14612d45576040805162461bcd60e51b815260206004820152600a60248201526915d493d391d7d4d1539160b21b604482015290519081900360640190fd5b5090910160209081019190912060a0850180516040805180860192909252818101939093528251808203840181526060909101909252815191909201209052505050565b600080805b6020811015612dc157600882901b91508481850181518110612dac57fe5b016020015160f81c9190911790600101612d8e565b509392505050565b6000600a60f883901c1015612de9578160f81c60300160f81b90506110b6565b8160f81c60570160f81b90506110b6565b60008160200183511015612e4a576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b60008060208301611e9a858563ffffffff612dfa16565b612e72613fa6565b6040805160c0810182528481528151606081018352600080825260208083018290528451828152808201865293949085019390830191612ec8565b612eb5613fa6565b815260200190600190039081612ead5790505b50905281526020016000604051908082528060200260200182016040528015612f0b57816020015b612ef8613fa6565b815260200190600190039081612ef05790505b50815260006020820152600260408201526060019290925250919050565b6008101590565b612f38613fa6565b612f40613fa6565b8260200151600184600001510381518110612f5757fe5b60209081029190910101518351600019018452915050919050565b608081015160009060ff161580156116ed57505051600160401b1190565b6080015160ff16600c1490565b61283b81611ae6565b600061277f612fbe856020865b048560000151613579565b6020855b066136e7565b808260200151836000015181518110612fdd57fe5b60209081029190910101525080516001019052565b60408051600880825281830190925260009160609190602082018180368337019050509050600061302c866020875b048660000151613579565b90506020808606600801106130ff576000613053876020885b046001018760400151613579565b905060005b6018601f8816600803018110156130a757613079838260208a5b06016136e7565b60f81b84828151811061308857fe5b60200101906001600160f81b031916908160001a905350600101613058565b506018601f8716600803015b60088110156130f8576130ca826020898401612fc2565b60f81b8482815181106130d957fe5b60200101906001600160f81b031916908160001a9053506001016130b3565b5050613148565b60005b6008811015613146576131188282602089613072565b60f81b83828151811061312757fe5b60200101906001600160f81b031916908160001a905350600101613102565b505b6126fa826136f4565b60408051602080825281830190925260009160609190602082018180368337019050509050600061318486602087613021565b905060208086066020011061323d5760006131a187602088613045565b905060005b601f87166020038110156131ef576131c1838260208a613072565b60f81b8482815181106131d057fe5b60200101906001600160f81b031916908160001a9053506001016131a6565b50601f86166008035b60208110156130f85761320f826020898401612fc2565b60f81b84828151811061321e57fe5b60200101906001600160f81b031916908160001a9053506001016131f8565b60005b6020811015613146576132568282602089613072565b60f81b83828151811061326557fe5b60200101906001600160f81b031916908160001a905350600101613240565b6080015160ff161590565b60008061329e86602087612fb3565b905060006132b082602088068761372a565b905060006132cc88602089048488600001518960200151613769565b98975050505050505050565b600060606132e584613809565b905060006132f587602088613021565b9050602080870660080111156133d75760005b6018601f881660080301811015613350576133468260208984010685846018018151811061333257fe5b01602001516001600160f81b031916613873565b9150600101613308565b5061336a876020885b048387600001518860200151613769565b9650600061337a88602089613045565b90506018601f8816600803015b60088110156133b3576133a98260208a84010686846018018151811061333257fe5b9150600101613387565b506133cf88602089046001018388604001518960600151613769565b975050613416565b60005b6008811015613406576133fc828260208a060185846018018151811061333257fe5b91506001016133da565b5061341387602088613359565b96505b509495945050505050565b6000606061342e84613809565b9050600061343e87602088613021565b9050602080870660200111156134d05760005b601f871660200381101561348057613476828260208a5b060185848151811061333257fe5b9150600101613451565b5061348d87602088613359565b9650600061349d88602089613045565b9050601f87166020035b60208110156133b3576134c68260208a84010686848151811061333257fe5b91506001016134a7565b60005b6020811015613406576134e9828260208a613468565b91506001016134d3565b60008061350285602086612fb3565b9050601f84165b60208110156135345761351c82826136e7565b1561352c57600092505050611772565b600101613509565b5061176e8560208604856000015161388f565b60008061355f858561355886613a0b565b6001613a36565b50855190915061176e90613574607b84613b0a565b613b0a565b60008151600014156135e25761358f600061241a565b84146135da576040805162461bcd60e51b815260206004820152601560248201527432bc3832b1ba32b21032b6b83a3c90313ab33332b960591b604482015290519081900360640190fd5b506000611772565b6000613601836000815181106135f457fe5b602002602001015161241a565b905060015b835181101561366b57846001166001141561363f5761363884828151811061362a57fe5b602002602001015183613b0a565b915061365f565b61365c8285838151811061364f57fe5b6020026020010151613b0a565b91505b600194851c9401613606565b508481146136b8576040805162461bcd60e51b8152602060048201526015602482015274195e1c1958dd19590818dbdc9c9958dd081c9bdbdd605a1b604482015290519081900360640190fd5b83156136c8575060009050611772565b826000815181106136d557fe5b60200260200101519150509392505050565b601f036008021c60ff1690565b600080805b8351811015611d2457600882901b915083818151811061371557fe5b016020015160f81c91909117906001016136f9565b6000606061373785613809565b90508260f81b81858151811061374957fe5b60200101906001600160f81b031916908160001a90535061176e816136f4565b600081516003146137bb576040805162461bcd60e51b81526020600482015260176024820152762120a22fa727a926a0a624ad20aa24a7a72fa82927a7a360491b604482015290519081900360640190fd5b6126fa86868686866000815181106137cf57fe5b602002602001015160001c876001815181106137e757fe5b6020026020010151886002815181106137fc57fe5b6020026020010151613b36565b6040805160208082528183019092526060918391839160208201818036833701905050905060005b6020811015612dc1578260f81b8282601f038151811061384d57fe5b60200101906001600160f81b031916908160001a90535060089290921c91600101613831565b6000606061388085613809565b90508281858151811061374957fe5b60008151600014156138f8576138a5600061241a565b84146138f0576040805162461bcd60e51b815260206004820152601560248201527432bc3832b1ba32b21032b6b83a3c90313ab33332b960591b604482015290519081900360640190fd5b506001611772565b600061390a836000815181106135f457fe5b905060016060613918613db0565b905060015b85518110156139ac5786600116600114156139565761394f86828151811061394157fe5b602002602001015185613b0a565b93506139a0565b6139668487838151811061364f57fe5b935082801561399d575081600182038151811061397f57fe5b602002602001015186828151811061399357fe5b6020026020010151145b92505b600196871c960161391d565b508683146139f9576040805162461bcd60e51b8152602060048201526015602482015274195e1c1958dd19590818dbdc9c9958dd081c9bdbdd605a1b604482015290519081900360640190fd5b8515611dfb5760019350505050611772565b600060018211613a1d575060016110b6565b613a2c60026001840104613a0b565b60020290506110b6565b60008060208411613a8a5785518510613a5e57613a53600061241a565b600191509150613b01565b6000613a72613a6d8888613e51565b61241a565b905080613a7f600061241a565b909350149050613b01565b600080613aa3886002880489016002895b046000613a36565b91509150808015613ab15750845b15613ad057613ac588886002890488613a36565b935093505050613b01565b600080613ae08a8a60028b613a9b565b91509150613aee8285613b0a565b818015613af85750835b95509550505050505b94509492505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080613b428761241a565b9050613b4f898988613579565b506060613b5a613db0565b905060018751036001901b8910613c1e5787613b7a578992505050613da5565b6000613b858a613ea8565b88519091505b60018203811015613bb357613ba98c84600184038151811061364f57fe5b9b50600101613b8b565b5060015b60018203811015613c09578a60011660011415613be757613be083600183038151811061394157fe5b9350613bfd565b613bfa8484600184038151811061364f57fe5b93505b60019a8b1c9a01613bb7565b50613c148b84613b0a565b9350505050613da5565b60015b8751811015613c9e5760008a600116600114613c3d5783613c52565b888281518110613c4957fe5b60200260200101515b905060008b600116600114613c7a57898381518110613c6d57fe5b6020026020010151613c7c565b845b9050613c888282613b0a565b60019c8d1c9c909550929092019150613c219050565b508715613cad57509050613da5565b600086613cbb575084613d30565b818781518110613cc757fe5b6020026020010151851415613d23576040805162461bcd60e51b815260206004820152601c60248201527f726967687420737562747265652063616e6e6f74206265207a65726f00000000604482015290519081900360640190fd5b613d2d8686613b0a565b90505b80875b60018a5103811015613d5957613d4f8285838151811061364f57fe5b9150600101613d33565b50838114613d9f576040805162461bcd60e51b815260206004820152600e60248201526d0caf0e0cac6e8cac840dac2e8c6d60931b604482015290519081900360640190fd5b50925050505b979650505050505050565b60408051818152610820810182526060918291906020820161080080368337019050509050613ddf600061241a565b81600081518110613dec57fe5b602090810291909101015260015b6040811015613e4b57613e2c826001830381518110613e1557fe5b602002602001015183600184038151811061364f57fe5b828281518110613e3857fe5b6020908102919091010152600101613dfa565b50905090565b600080805b6020811015612dc157600882901b91506000818501865111613e79576000613e97565b8582860181518110613e8757fe5b01602001516001600160f81b0319165b60f81c929092179150600101613e56565b600081613eb7575060016110b6565b613ec4600183901c613ea8565b60010190506110b6565b60405180608001604052806004906020820280368337509192915050565b604080516102008101825260008082526020820152908101613f0c613fe3565b8152602001613f19613fe3565b81526000602082018190526040820181905260608201819052608082015260a001613f42614043565b8152602001613f4f614043565b81526000602082018190526040820181905260608083018190526080830182905260a083015260c09091015290565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040518060c0016040528060008152602001613fc061405d565b815260606020820181905260006040830181905290820181905260809091015290565b6040805161010081019091526000815260208101613fff613fa6565b815260200161400c613fa6565b8152602001614019613fa6565b8152602001614026613fa6565b81526000602082018190526040820181905260609091015290565bfe5b604051806040016040528060008152602001606081525090565b604080516060808201835260008083526020830152918101919091529056fe75736520616e6f7468657220636f6e747261637420746f2068616e646c65206f74686572206f70636f646573a2646970667358221220c0176e50d53f5c36931ecadcf85aa224de3042c6835311150684f15bb791338e64736f6c634300060b0033'
