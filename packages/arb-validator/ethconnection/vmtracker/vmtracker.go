// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vmtracker

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArbProtocolABI is the input ABI used to generate the binding from.
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateBeforeValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_tokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_destinations\",\"type\":\"address[]\"}],\"name\":\"generateLastMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"name\":\"_lastLogHash\",\"type\":\"bytes32\"},{\"name\":\"_totalMessageValueAmounts\",\"type\":\"uint256[]\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[5]\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNum\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmount\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestination\",\"type\":\"address[]\"}],\"name\":\"unanimousAssertHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"generateSentMessageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"beforeBalancesValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_dataHashes\",\"type\":\"bytes32[]\"},{\"name\":\"_tokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_destinations\",\"type\":\"address[]\"}],\"name\":\"generateLastMessageHashStub\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"},{\"name\":\"_newMessage\",\"type\":\"bytes32\"}],\"name\":\"appendInboxPendingMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_inboxHash\",\"type\":\"bytes32\"},{\"name\":\"_pendingMessages\",\"type\":\"bytes32\"}],\"name\":\"appendInboxMessages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ArbProtocolFuncSigs = map[string]string{
	"f11fcc26": "appendInboxMessages(bytes32,bytes32)",
	"d78d18ea": "appendInboxPendingMessage(bytes32,bytes32)",
	"af17d922": "beforeBalancesValid(bytes21[],uint256[])",
	"0f89fbff": "calculateBeforeValues(bytes21[],uint16[],uint256[])",
	"20903721": "generateAssertionHash(bytes32,uint32,bytes32,bytes32,bytes32,bytes32,uint256[])",
	"1914612a": "generateLastMessageHash(bytes21[],bytes,uint16[],uint256[],address[])",
	"d14cf098": "generateLastMessageHashStub(bytes21[],bytes32[],uint16[],uint256[],address[])",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"3e285598": "generatePreconditionHash(bytes32,uint64[2],bytes32,bytes21[],uint256[])",
	"505221a1": "generateSentMessageHash(address,bytes32,bytes21,uint256,address)",
	"4f930c50": "unanimousAssertHash(bytes32[5],uint64[2],bytes21[],bytes,uint16[],uint256[],address[])",
}

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
var ArbProtocolBin = "0x6121df610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b25760003560e01c80634f930c501161007b5780634f930c50146107d5578063505221a114610ade578063af17d92214610b2a578063d14cf09814610c61578063d78d18ea14610f0a578063f11fcc2614610f2d576100b2565b80624c28f6146100b75780630f89fbff1461010b5780631914612a1461030057806320903721146105ac5780633e28559814610677575b600080fd5b6100f9600480360360808110156100cd57600080fd5b5080359060208101356001600160581b03191690604081013590606001356001600160a01b0316610f50565b60408051918252519081900360200190f35b6102b06004803603606081101561012157600080fd5b810190602081018135600160201b81111561013b57600080fd5b82018360208201111561014d57600080fd5b803590602001918460208302840111600160201b8311171561016e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101bd57600080fd5b8201836020820111156101cf57600080fd5b803590602001918460208302840111600160201b831117156101f057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561023f57600080fd5b82018360208201111561025157600080fd5b803590602001918460208302840111600160201b8311171561027257600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061103c945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102ec5781810151838201526020016102d4565b505050509050019250505060405180910390f35b6100f9600480360360a081101561031657600080fd5b810190602081018135600160201b81111561033057600080fd5b82018360208201111561034257600080fd5b803590602001918460208302840111600160201b8311171561036357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103b257600080fd5b8201836020820111156103c457600080fd5b803590602001918460018302840111600160201b831117156103e557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561043757600080fd5b82018360208201111561044957600080fd5b803590602001918460208302840111600160201b8311171561046a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104b957600080fd5b8201836020820111156104cb57600080fd5b803590602001918460208302840111600160201b831117156104ec57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561053b57600080fd5b82018360208201111561054d57600080fd5b803590602001918460208302840111600160201b8311171561056e57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611221945050505050565b6100f9600480360360e08110156105c257600080fd5b81359163ffffffff6020820135169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561060657600080fd5b82018360208201111561061857600080fd5b803590602001918460208302840111600160201b8311171561063957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611464945050505050565b6100f9600480360360c081101561068d57600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194833594909390925060408101915060200135600160201b8111156106e257600080fd5b8201836020820111156106f457600080fd5b803590602001918460208302840111600160201b8311171561071557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561076457600080fd5b82018360208201111561077657600080fd5b803590602001918460208302840111600160201b8311171561079757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611503945050505050565b6100f960048036036101808110156107ec57600080fd5b810190808060a001906005806020026040519081016040528092919082600560200280828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561086257600080fd5b82018360208201111561087457600080fd5b803590602001918460208302840111600160201b8311171561089557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156108e457600080fd5b8201836020820111156108f657600080fd5b803590602001918460018302840111600160201b8311171561091757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561096957600080fd5b82018360208201111561097b57600080fd5b803590602001918460208302840111600160201b8311171561099c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156109eb57600080fd5b8201836020820111156109fd57600080fd5b803590602001918460208302840111600160201b83111715610a1e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a6d57600080fd5b820183602082011115610a7f57600080fd5b803590602001918460208302840111600160201b83111715610aa057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506115ee945050505050565b6100f9600480360360a0811015610af457600080fd5b506001600160a01b0381358116916020810135916001600160581b03196040830135169160608101359160809091013516611753565b610c4d60048036036040811015610b4057600080fd5b810190602081018135600160201b811115610b5a57600080fd5b820183602082011115610b6c57600080fd5b803590602001918460208302840111600160201b83111715610b8d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610bdc57600080fd5b820183602082011115610bee57600080fd5b803590602001918460208302840111600160201b83111715610c0f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611947945050505050565b604080519115158252519081900360200190f35b6100f9600480360360a0811015610c7757600080fd5b810190602081018135600160201b811115610c9157600080fd5b820183602082011115610ca357600080fd5b803590602001918460208302840111600160201b83111715610cc457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d1357600080fd5b820183602082011115610d2557600080fd5b803590602001918460208302840111600160201b83111715610d4657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d9557600080fd5b820183602082011115610da757600080fd5b803590602001918460208302840111600160201b83111715610dc857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610e1757600080fd5b820183602082011115610e2957600080fd5b803590602001918460208302840111600160201b83111715610e4a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610e9957600080fd5b820183602082011115610eab57600080fd5b803590602001918460208302840111600160201b83111715610ecc57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611b4c945050505050565b6100f960048036036040811015610f2057600080fd5b5080359060200135611cea565b6100f960048036036040811015610f4357600080fd5b5080359060200135611d2e565b60408051600480825260a0820190925260009160609190816020015b610f74612174565b815260200190600190039081610f6c579050509050610f9286611d48565b81600081518110610f9f57fe5b6020026020010181905250610fbc836001600160a01b0316611da4565b81600181518110610fc957fe5b6020026020010181905250610fdd84611da4565b81600281518110610fea57fe5b60209081029190910101526110086001600160581b03198616611da4565b8160038151811061101557fe5b602002602001018190525061103161102c82611dfe565b611e86565b519695505050505050565b606060008351905060608551604051908082528060200260200182016040528015611071578160200160208202803883390190505b50905060005b8281101561121757600086828151811061108d57fe5b60200260200101519050878161ffff16815181106110a757fe5b60200260200101516014601581106110bb57fe5b1a60f81b6001600160f81b031916611108578582815181106110d957fe5b6020026020010151838261ffff16815181106110f157fe5b60200260200101818151019150818152505061120e565b828161ffff168151811061111857fe5b6020026020010151600014611174576040805162461bcd60e51b815260206004820152601d60248201527f43616e277420696e636c756465204e465420746f6b656e207477696365000000604482015290519081900360640190fd5b85828151811061118057fe5b6020026020010151600014156111dd576040805162461bcd60e51b815260206004820152601f60248201527f4e465420746f6b656e206d7573742068617665206e6f6e2d7a65726f20696400604482015290519081900360640190fd5b8582815181106111e957fe5b6020026020010151838261ffff168151811061120157fe5b6020026020010181815250505b50600101611077565b5095945050505050565b6000815183511461126f576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b83518351146112bb576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b825160009081908190815b818110156114565773__$d969135829891f807aa9c34494da4ecd99$__6389df40da8b866040518363ffffffff1660e01b81526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561133c578181015183820152602001611324565b50505050905090810190601f1680156113695780820380516001836020036101000a031916815260200191505b509350505050604080518083038186803b15801561138657600080fd5b505af415801561139a573d6000803e3d6000fd5b505050506040513d60408110156113b057600080fd5b5080516020909101518a51919550935061141f9084908d908c90859081106113d457fe5b602002602001015161ffff16815181106113ea57fe5b60200260200101518a84815181106113fe57fe5b60200260200101518a858151811061141257fe5b6020026020010151610f50565b60408051602080820198909852808201839052815180820383018152606090910190915280519601959095209492506001016112c6565b505050505095945050505050565b600087878787878787604051602001808881526020018763ffffffff1663ffffffff1660e01b8152600401868152602001858152602001848152602001838152602001828051906020019060200280838360005b838110156114d05781810151838201526020016114b8565b50505050905001975050505050505050604051602081830303815290604052805190602001209050979650505050505050565b600085858260200201518660016020020151868686604051602001808781526020018667ffffffffffffffff1667ffffffffffffffff1660c01b81526008018567ffffffffffffffff1667ffffffffffffffff1660c01b8152600801848152602001838051906020019060200280838360005b8381101561158e578181015183820152602001611576565b50505050905001828051906020019060200280838360005b838110156115be5781810151838201526020016115a6565b50505050905001965050505050505060405160208183030381529060405280519060200120905095945050505050565b6000878787878787876040516020018088600560200280838360005b8381101561162257818101518382015260200161160a565b5050505090500187600260200280838360005b8381101561164d578181015183820152602001611635565b50505050905001868051906020019060200280838360005b8381101561167d578181015183820152602001611665565b5050505090500185805190602001908083835b602083106116af5780518252601f199092019160209182019101611690565b51815160209384036101000a60001901801990921691161790528751919093019287810192500280838360005b838110156116f45781810151838201526020016116dc565b50505050905001838051906020019060200280838360005b8381101561172457818101518382015260200161170c565b5050505090500182805190602001906020028083836000838110156114d05781810151838201526020016114b8565b60408051606087811b6bffffffffffffffffffffffff191660208084019190915260348301889052605483018690526001600160581b03198716607484015283518084036069018152608984018086528151919092012060048083526101298501909552600094909360a9015b6117c8612174565b8152602001906001900390816117c05790505090506117e687611d48565b816000815181106117f357fe5b602002602001018190525061180742611da4565b8160018151811061181457fe5b602002602001018190525061182843611da4565b8160028151811061183557fe5b602090810291909101015261184982611da4565b8160038151811061185657fe5b602090810291909101015260408051600480825260a08201909252606091816020015b611881612174565b81526020019060019003908161187957905050905061189f82611dfe565b816000815181106118ac57fe5b60200260200101819052506118c9856001600160a01b0316611da4565b816001815181106118d657fe5b60200260200101819052506118ea86611da4565b816002815181106118f757fe5b60209081029190910101526119156001600160581b03198816611da4565b8160038151811061192257fe5b602002602001018190525061193961102c82611dfe565b519998505050505050505050565b81516000908015806119595750806001145b15611968576001915050611b46565b60005b60018203811015611af957600085828151811061198457fe5b602002602001015160146015811061199857fe5b1a60f81b90506001600160f81b03198116611a02578582815181106119b957fe5b60200260200101516001600160581b0319168683600101815181106119da57fe5b60200260200101516001600160581b031916116119fd5760009350505050611b46565b611af0565b600160f81b6001600160f81b031982161415611ae457858281518110611a2457fe5b60200260200101516001600160581b031916868360010181518110611a4557fe5b60200260200101516001600160581b0319161080611ad35750858281518110611a6a57fe5b60200260200101516001600160581b031916868360010181518110611a8b57fe5b60200260200101516001600160581b031916148015611ad35750848281518110611ab157fe5b6020026020010151858360010181518110611ac857fe5b602002602001015111155b156119fd5760009350505050611b46565b60009350505050611b46565b5060010161196b565b50600160f81b846001830381518110611b0e57fe5b6020026020010151601460158110611b2257fe5b1a60f81b6001600160f81b0319161115611b40576000915050611b46565b60019150505b92915050565b60008351855114611b9a576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8251855114611be6576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b8151855114611c32576040805162461bcd60e51b8152602060048201526013602482015272092dce0eae840e6d2f4ca40dad2e6dac2e8c6d606b1b604482015290519081900360640190fd5b84516000908190815b81811015611cdc57611ca5898281518110611c5257fe5b60200260200101518b8a8481518110611c6757fe5b602002602001015161ffff1681518110611c7d57fe5b6020026020010151898481518110611c9157fe5b602002602001015189858151811061141257fe5b6040805160208082019790975280820183905281518082038301815260609091019091528051950194909420939250600101611c3b565b509198975050505050505050565b6000611d276040518060600160405280611d046000611da4565b8152602001611d1286611d48565b8152602001611d2085611d48565b9052611f75565b9392505050565b6000611d276040518060600160405280611d046001611da4565b611d50612174565b604080516060810182528381528151600080825260208281019094529192830191611d91565b611d7e612174565b815260200190600190039081611d765790505b508152600260209091015290505b919050565b611dac612174565b604080516060810182528381528151600080825260208281019094529192830191611ded565b611dda612174565b815260200190600190039081611dd25790505b508152600060209091015292915050565b611e06612174565b611e108251611ffd565b611e61576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b611e8e612198565b6040820151600c60ff90911610611ee0576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff16611f0d576040518060200160405280611f048460000151612004565b90529050611d9f565b604082015160ff1660021415611f325750604080516020810190915281518152611d9f565b600360ff16826040015160ff1610158015611f5657506040820151600c60ff909116105b15611f73576040518060200160405280611f048460200151612028565bfe5b6040805160038082526080820190925260009160609190816020015b611f99612174565b815260200190600190039081611f91575050805190915060005b81811015611feb57848160038110611fc757fe5b6020020151838281518110611fd857fe5b6020908102919091010152600101611fb3565b50611ff582612028565b949350505050565b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600882511115612078576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b606082516040519080825280602002602001820160405280156120a5578160200160208202803883390190505b50805190915060005b81811015612101576120be612198565b6120da8683815181106120cd57fe5b6020026020010151611e86565b905080600001518483815181106120ed57fe5b6020908102919091010152506001016120ae565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561214a578181015183820152602001612132565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72305820a46f4e4e547acc13a3ed3fde52775d00c8901055863db33668b59b6d25b02f8f64736f6c634300050a0032"

// DeployArbProtocol deploys a new Ethereum contract, binding an instance of ArbProtocol to it.
func DeployArbProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbProtocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	ArbProtocolBin = strings.Replace(ArbProtocolBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// ArbProtocol is an auto generated Go binding around an Ethereum contract.
type ArbProtocol struct {
	ArbProtocolCaller     // Read-only binding to the contract
	ArbProtocolTransactor // Write-only binding to the contract
	ArbProtocolFilterer   // Log filterer for contract events
}

// ArbProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbProtocolSession struct {
	Contract     *ArbProtocol      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbProtocolCallerSession struct {
	Contract *ArbProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ArbProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbProtocolTransactorSession struct {
	Contract     *ArbProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ArbProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbProtocolRaw struct {
	Contract *ArbProtocol // Generic contract binding to access the raw methods on
}

// ArbProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbProtocolCallerRaw struct {
	Contract *ArbProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ArbProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbProtocolTransactorRaw struct {
	Contract *ArbProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbProtocol creates a new instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocol(address common.Address, backend bind.ContractBackend) (*ArbProtocol, error) {
	contract, err := bindArbProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// NewArbProtocolCaller creates a new read-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolCaller(address common.Address, caller bind.ContractCaller) (*ArbProtocolCaller, error) {
	contract, err := bindArbProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolCaller{contract: contract}, nil
}

// NewArbProtocolTransactor creates a new write-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbProtocolTransactor, error) {
	contract, err := bindArbProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolTransactor{contract: contract}, nil
}

// NewArbProtocolFilterer creates a new log filterer instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbProtocolFilterer, error) {
	contract, err := bindArbProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolFilterer{contract: contract}, nil
}

// bindArbProtocol binds a generic wrapper to an already deployed contract.
func bindArbProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.ArbProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transact(opts, method, params...)
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) AppendInboxMessages(opts *bind.CallOpts, _inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "appendInboxMessages", _inboxHash, _pendingMessages)
	return *ret0, err
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) AppendInboxMessages(_inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxMessages(&_ArbProtocol.CallOpts, _inboxHash, _pendingMessages)
}

// AppendInboxMessages is a free data retrieval call binding the contract method 0xf11fcc26.
//
// Solidity: function appendInboxMessages(bytes32 _inboxHash, bytes32 _pendingMessages) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) AppendInboxMessages(_inboxHash [32]byte, _pendingMessages [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxMessages(&_ArbProtocol.CallOpts, _inboxHash, _pendingMessages)
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) AppendInboxPendingMessage(opts *bind.CallOpts, _pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "appendInboxPendingMessage", _pendingMessages, _newMessage)
	return *ret0, err
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) AppendInboxPendingMessage(_pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxPendingMessage(&_ArbProtocol.CallOpts, _pendingMessages, _newMessage)
}

// AppendInboxPendingMessage is a free data retrieval call binding the contract method 0xd78d18ea.
//
// Solidity: function appendInboxPendingMessage(bytes32 _pendingMessages, bytes32 _newMessage) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) AppendInboxPendingMessage(_pendingMessages [32]byte, _newMessage [32]byte) ([32]byte, error) {
	return _ArbProtocol.Contract.AppendInboxPendingMessage(&_ArbProtocol.CallOpts, _pendingMessages, _newMessage)
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCaller) BeforeBalancesValid(opts *bind.CallOpts, _tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "beforeBalancesValid", _tokenTypes, _beforeBalances)
	return *ret0, err
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCallerSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCaller) CalculateBeforeValues(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "calculateBeforeValues", _tokenTypes, _messageTokenNums, _messageAmounts)
	return *ret0, err
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCallerSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x1914612a.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHash(opts *bind.CallOpts, _tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHash", _tokenTypes, _data, _tokenNums, _amounts, _destinations)
	return *ret0, err
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x1914612a.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _data, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHash is a free data retrieval call binding the contract method 0x1914612a.
//
// Solidity: function generateLastMessageHash(bytes21[] _tokenTypes, bytes _data, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHash(_tokenTypes [][21]byte, _data []byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHash(&_ArbProtocol.CallOpts, _tokenTypes, _data, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xd14cf098.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateLastMessageHashStub(opts *bind.CallOpts, _tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateLastMessageHashStub", _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
	return *ret0, err
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xd14cf098.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
}

// GenerateLastMessageHashStub is a free data retrieval call binding the contract method 0xd14cf098.
//
// Solidity: function generateLastMessageHashStub(bytes21[] _tokenTypes, bytes32[] _dataHashes, uint16[] _tokenNums, uint256[] _amounts, address[] _destinations) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateLastMessageHashStub(_tokenTypes [][21]byte, _dataHashes [][32]byte, _tokenNums []uint16, _amounts []*big.Int, _destinations []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateLastMessageHashStub(&_ArbProtocol.CallOpts, _tokenTypes, _dataHashes, _tokenNums, _amounts, _destinations)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x505221a1.
//
// Solidity: function generateSentMessageHash(address _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, address _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateSentMessageHash(opts *bind.CallOpts, _dest common.Address, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateSentMessageHash", _dest, _data, _tokenType, _value, _sender)
	return *ret0, err
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x505221a1.
//
// Solidity: function generateSentMessageHash(address _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, address _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateSentMessageHash(_dest common.Address, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateSentMessageHash(&_ArbProtocol.CallOpts, _dest, _data, _tokenType, _value, _sender)
}

// GenerateSentMessageHash is a free data retrieval call binding the contract method 0x505221a1.
//
// Solidity: function generateSentMessageHash(address _dest, bytes32 _data, bytes21 _tokenType, uint256 _value, address _sender) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateSentMessageHash(_dest common.Address, _data [32]byte, _tokenType [21]byte, _value *big.Int, _sender common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateSentMessageHash(&_ArbProtocol.CallOpts, _dest, _data, _tokenType, _value, _sender)
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x4f930c50.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, address[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) UnanimousAssertHash(opts *bind.CallOpts, _fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "unanimousAssertHash", _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
	return *ret0, err
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x4f930c50.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, address[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) UnanimousAssertHash(_fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.UnanimousAssertHash(&_ArbProtocol.CallOpts, _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// UnanimousAssertHash is a free data retrieval call binding the contract method 0x4f930c50.
//
// Solidity: function unanimousAssertHash(bytes32[5] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNum, uint256[] _messageAmount, address[] _messageDestination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) UnanimousAssertHash(_fields [5][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNum []uint16, _messageAmount []*big.Int, _messageDestination []common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.UnanimousAssertHash(&_ArbProtocol.CallOpts, _fields, _timeBounds, _tokenTypes, _messageData, _messageTokenNum, _messageAmount, _messageDestination)
}

// ArbValueABI is the input ABI used to generate the binding from.
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValidValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidValueHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeValueHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbValueFuncSigs maps the 4-byte function signature to its string representation.
var ArbValueFuncSigs = map[string]string{
	"89df40da": "deserializeValidValueHash(bytes,uint256)",
	"8f346036": "deserializeValueHash(bytes)",
	"1f3d4d4e": "getNextValidValue(bytes,uint256)",
	"53409fab": "hashCodePointBasicValue(uint8,bytes32)",
	"264f384b": "hashCodePointImmediateValue(uint8,bytes32,bytes32)",
	"364df277": "hashEmptyTuple()",
	"1667b411": "hashIntValue(uint256)",
	"b2b9dc62": "isValidTupleSize(uint256)",
}

// ArbValueBin is the compiled bytecode used for deploying new contracts.
var ArbValueBin = "0x610d71610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100925760003560e01c806353409fab1161006557806353409fab1461022157806389df40da146102475780638f34603614610308578063b2b9dc62146103ae57610092565b80631667b411146100975780631f3d4d4e146100c6578063264f384b146101ed578063364df27714610219575b600080fd5b6100b4600480360360208110156100ad57600080fd5b50356103df565b60408051918252519081900360200190f35b61016e600480360360408110156100dc57600080fd5b8101906020810181356401000000008111156100f757600080fd5b82018360208201111561010957600080fd5b8035906020019184600183028401116401000000008311171561012b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610405915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101b1578181015183820152602001610199565b50505050905090810190601f1680156101de5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100b46004803603606081101561020357600080fd5b5060ff813516906020810135906040013561049a565b6100b46104ec565b6100b46004803603604081101561023757600080fd5b5060ff813516906020013561055f565b6102ef6004803603604081101561025d57600080fd5b81019060208101813564010000000081111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111640100000000831117156102ac57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506105a6915050565b6040805192835260208301919091528051918290030190f35b6100b46004803603602081101561031e57600080fd5b81019060208101813564010000000081111561033957600080fd5b82018360208201111561034b57600080fd5b8035906020019184600183028401116401000000008311171561036d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610631945050505050565b6103cb600480360360208110156103c457600080fd5b50356106b5565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b60006060600080610414610d06565b61041e87876106bc565b919450925090508215610478576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161048c888880840363ffffffff61081116565b945094505050509250929050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b83811015610538578181015183820152602001610520565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6000806000806105b4610d06565b6105be87876106bc565b919450925090508215610618576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161062282610891565b51909890975095505050505050565b6000808061063d610d06565b6106488560006106bc565b9194509250905082156106a2576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b6106ab81610891565b5195945050505050565b6008101590565b6000806106c7610d06565b8451841061071c576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b6000849050600086828151811061072f57fe5b016020015160019092019160f81c9050600081610771576107508884610980565b9093509050600083610761836109a7565b9197509550935061080a92505050565b60ff821660021415610798576107878884610980565b909350905060008361076183610a01565b600360ff8316108015906107af5750600c60ff8316105b156107eb576002198201606060006107c8838c88610a5b565b9097509250905080866107da84610b16565b98509850985050505050505061080a565b8160ff166127100160006107ff60006109a7565b919750955093505050505b9250925092565b60608183018451101561082357600080fd5b60608215801561083e57604051915060208201604052610888565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561087757805183526020928301920161085f565b5050858452601f01601f1916604052505b50949350505050565b610899610d2a565b6040820151600c60ff909116106108eb576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff1661091857604051806020016040528061090f84600001516103df565b90529050610400565b604082015160ff166002141561093d5750604080516020810190915281518152610400565b600360ff16826040015160ff161015801561096157506040820151600c60ff909116105b1561097e57604051806020016040528061090f8460200151610b9e565bfe5b6000808281610995868363ffffffff610cea16565b60209290920196919550909350505050565b6109af610d06565b6040805160608101825283815281516000808252602082810190945291928301916109f0565b6109dd610d06565b8152602001906001900390816109d55790505b508152600060209091015292915050565b610a09610d06565b604080516060810182528381528151600080825260208281019094529192830191610a4a565b610a37610d06565b815260200190600190039081610a2f5790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610aa657816020015b610a93610d06565b815260200190600190039081610a8b5790505b50905060005b8960ff168160ff161015610b0057610ac489856106bc565b8451859060ff8616908110610ad557fe5b6020908102919091010152945092508215610af857509094509092509050610b0d565b600101610aac565b5060009550919350909150505b93509350939050565b610b1e610d06565b610b2882516106b5565b610b79576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b6000600882511115610bee576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610c1b578160200160208202803883390190505b50805190915060005b81811015610c7757610c34610d2a565b610c50868381518110610c4357fe5b6020026020010151610891565b90508060000151848381518110610c6357fe5b602090810291909101015250600101610c24565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610cc0578181015183820152602001610ca8565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60008160200183511015610cfd57600080fd5b50016020015190565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72305820df9accba692330c50b5a0c6ce3f31382e1447f19a4cbff3ad62e18c40a9b21de64736f6c634300050a0032"

// DeployArbValue deploys a new Ethereum contract, binding an instance of ArbValue to it.
func DeployArbValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbValue, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// ArbValue is an auto generated Go binding around an Ethereum contract.
type ArbValue struct {
	ArbValueCaller     // Read-only binding to the contract
	ArbValueTransactor // Write-only binding to the contract
	ArbValueFilterer   // Log filterer for contract events
}

// ArbValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbValueSession struct {
	Contract     *ArbValue         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbValueCallerSession struct {
	Contract *ArbValueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbValueTransactorSession struct {
	Contract     *ArbValueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbValueRaw struct {
	Contract *ArbValue // Generic contract binding to access the raw methods on
}

// ArbValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbValueCallerRaw struct {
	Contract *ArbValueCaller // Generic read-only contract binding to access the raw methods on
}

// ArbValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbValueTransactorRaw struct {
	Contract *ArbValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbValue creates a new instance of ArbValue, bound to a specific deployed contract.
func NewArbValue(address common.Address, backend bind.ContractBackend) (*ArbValue, error) {
	contract, err := bindArbValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// NewArbValueCaller creates a new read-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueCaller(address common.Address, caller bind.ContractCaller) (*ArbValueCaller, error) {
	contract, err := bindArbValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueCaller{contract: contract}, nil
}

// NewArbValueTransactor creates a new write-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbValueTransactor, error) {
	contract, err := bindArbValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueTransactor{contract: contract}, nil
}

// NewArbValueFilterer creates a new log filterer instance of ArbValue, bound to a specific deployed contract.
func NewArbValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbValueFilterer, error) {
	contract, err := bindArbValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbValueFilterer{contract: contract}, nil
}

// bindArbValue binds a generic wrapper to an already deployed contract.
func bindArbValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.ArbValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transact(opts, method, params...)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValidValueHash(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "deserializeValidValueHash", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValueHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "deserializeValueHash", data)
	return *ret0, err
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCaller) GetNextValidValue(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "getNextValidValue", data, offset)
	return *ret0, *ret1, err
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCallerSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointBasicValue(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointBasicValue", opcode, nextCodePoint)
	return *ret0, err
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointImmediateValue(opts *bind.CallOpts, opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointImmediateValue", opcode, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashIntValue(opts *bind.CallOpts, val *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashIntValue", val)
	return *ret0, err
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCaller) IsValidTupleSize(opts *bind.CallOpts, size *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "isValidTupleSize", size)
	return *ret0, err
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCallerSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058202a4b33a7382bd41fc01a7b408d0a3491e17cc0d680148cb3a692e8a4446bc59d64736f6c634300050a0032"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// DebugPrintABI is the input ABI used to generate the binding from.
const DebugPrintABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"b32\",\"type\":\"bytes32\"}],\"name\":\"bytes32string\",\"outputs\":[{\"name\":\"out\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// DebugPrintFuncSigs maps the 4-byte function signature to its string representation.
var DebugPrintFuncSigs = map[string]string{
	"252fb38d": "bytes32string(bytes32)",
}

// DebugPrintBin is the compiled bytecode used for deploying new contracts.
var DebugPrintBin = "0x610202610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063252fb38d1461003a575b600080fd5b6100576004803603602081101561005057600080fd5b50356100cc565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610091578181015183820152602001610079565b50505050905090810190601f1680156100be5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051818152606081810183529182919060208201818038833901905050905060005b602081101561019457600084826020811061010757fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b61012d8261019b565b85856002028151811061013c57fe5b60200101906001600160f81b031916908160001a90535061015c8161019b565b85856002026001018151811061016e57fe5b60200101906001600160f81b031916908160001a90535050600190920191506100f09050565b5092915050565b6000600a60f883901c10156101bb578160f81c60300160f81b90506101c8565b8160f81c60570160f81b90505b91905056fea265627a7a723058208d6eb52933501def4197f8e418e23dd408f84f76f9352d754d71b995c624db2264736f6c634300050a0032"

// DeployDebugPrint deploys a new Ethereum contract, binding an instance of DebugPrint to it.
func DeployDebugPrint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DebugPrint, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DebugPrintBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// DebugPrint is an auto generated Go binding around an Ethereum contract.
type DebugPrint struct {
	DebugPrintCaller     // Read-only binding to the contract
	DebugPrintTransactor // Write-only binding to the contract
	DebugPrintFilterer   // Log filterer for contract events
}

// DebugPrintCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebugPrintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebugPrintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebugPrintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebugPrintSession struct {
	Contract     *DebugPrint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebugPrintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebugPrintCallerSession struct {
	Contract *DebugPrintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DebugPrintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebugPrintTransactorSession struct {
	Contract     *DebugPrintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DebugPrintRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebugPrintRaw struct {
	Contract *DebugPrint // Generic contract binding to access the raw methods on
}

// DebugPrintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebugPrintCallerRaw struct {
	Contract *DebugPrintCaller // Generic read-only contract binding to access the raw methods on
}

// DebugPrintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebugPrintTransactorRaw struct {
	Contract *DebugPrintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebugPrint creates a new instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrint(address common.Address, backend bind.ContractBackend) (*DebugPrint, error) {
	contract, err := bindDebugPrint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// NewDebugPrintCaller creates a new read-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintCaller(address common.Address, caller bind.ContractCaller) (*DebugPrintCaller, error) {
	contract, err := bindDebugPrint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintCaller{contract: contract}, nil
}

// NewDebugPrintTransactor creates a new write-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintTransactor(address common.Address, transactor bind.ContractTransactor) (*DebugPrintTransactor, error) {
	contract, err := bindDebugPrint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintTransactor{contract: contract}, nil
}

// NewDebugPrintFilterer creates a new log filterer instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintFilterer(address common.Address, filterer bind.ContractFilterer) (*DebugPrintFilterer, error) {
	contract, err := bindDebugPrint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebugPrintFilterer{contract: contract}, nil
}

// bindDebugPrint binds a generic wrapper to an already deployed contract.
func bindDebugPrint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.DebugPrintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transact(opts, method, params...)
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintCaller) Bytes32string(opts *bind.CallOpts, b32 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DebugPrint.contract.Call(opts, out, "bytes32string", b32)
	return *ret0, err
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintSession) Bytes32string(b32 [32]byte) (string, error) {
	return _DebugPrint.Contract.Bytes32string(&_DebugPrint.CallOpts, b32)
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintCallerSession) Bytes32string(b32 [32]byte) (string, error) {
	return _DebugPrint.Contract.Bytes32string(&_DebugPrint.CallOpts, b32)
}

// DisputableABI is the input ABI used to generate the binding from.
const DisputableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"}],\"name\":\"withinTimeBounds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fields\",\"type\":\"bytes32[3]\"},{\"indexed\":false,\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"indexed\":false,\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"lastMessageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"logsAccHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"}]"

// DisputableFuncSigs maps the 4-byte function signature to its string representation.
var DisputableFuncSigs = map[string]string{
	"924e7b37": "confirmDisputableAsserted(VM.Data storage,bytes32,bytes32,uint32,bytes21[],bytes,uint16[],uint256[],address[],bytes32)",
	"5af93c7a": "initiateChallenge(VM.Data storage,bytes32)",
	"c97c8eec": "pendingDisputableAssert(VM.Data storage,bytes32[4],uint32,uint64[2],bytes21[],bytes32[],uint16[],uint256[],address[])",
	"42c0787e": "withinTimeBounds(uint64[2])",
}

// DisputableBin is the compiled bytecode used for deploying new contracts.
var DisputableBin = "0x6120eb610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806342c0787e1461005b5780635af93c7a146100ba578063924e7b37146100ec578063c97c8eec146103c6575b600080fd5b6100a66004803603604081101561007157600080fd5b604080518082018252918301929181830191839060029083908390808284376000920191909152509194506106ee9350505050565b604080519115158252519081900360200190f35b8180156100c657600080fd5b506100ea600480360360408110156100dd57600080fd5b5080359060200135610720565b005b8180156100f857600080fd5b506100ea600480360361014081101561011057600080fd5b81359160208101359160408201359163ffffffff6060820135169181019060a081016080820135600160201b81111561014857600080fd5b82018360208201111561015a57600080fd5b803590602001918460208302840111600160201b8311171561017b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101ca57600080fd5b8201836020820111156101dc57600080fd5b803590602001918460018302840111600160201b831117156101fd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561024f57600080fd5b82018360208201111561026157600080fd5b803590602001918460208302840111600160201b8311171561028257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102d157600080fd5b8201836020820111156102e357600080fd5b803590602001918460208302840111600160201b8311171561030457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561035357600080fd5b82018360208201111561036557600080fd5b803590602001918460208302840111600160201b8311171561038657600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506109ab915050565b8180156103d257600080fd5b506100ea60048036036101a08110156103ea57600080fd5b604080516080818101909252833593928301929160a0830191906020840190600490839083908082843760009201919091525050604080518082018252929563ffffffff853516959094909360608201935091602090910190600290839083908082843760009201919091525091949392602081019250359050600160201b81111561047557600080fd5b82018360208201111561048757600080fd5b803590602001918460208302840111600160201b831117156104a857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104f757600080fd5b82018360208201111561050957600080fd5b803590602001918460208302840111600160201b8311171561052a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561057957600080fd5b82018360208201111561058b57600080fd5b803590602001918460208302840111600160201b831117156105ac57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156105fb57600080fd5b82018360208201111561060d57600080fd5b803590602001918460208302840111600160201b8311171561062e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561067d57600080fd5b82018360208201111561068f57600080fd5b803590602001918460208302840111600160201b831117156106b057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610a13945050505050565b805160009067ffffffffffffffff16431080159061071a5750602082015167ffffffffffffffff164311155b92915050565b60048201546001600160a01b031633141561076c5760405162461bcd60e51b8152600401808060200182810382526021815260200180611f956021913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156107bb57600080fd5b505af41580156107cf573d6000803e3d6000fd5b505050506040513d60208110156107e557600080fd5b50516108225760405162461bcd60e51b81526004018080602001828103825260268152602001806120646026913960400191505060405180910390fd5b60026006830154600160501b900460ff16600381111561083e57fe5b1461087a5760405162461bcd60e51b815260040180806020018281038252602f815260200180611ece602f913960400191505060405180910390fd5b3360009081526020839052604090205460058301546001600160801b031611156108d55760405162461bcd60e51b8152600401808060200182810382526027815260200180611e6e6027913960400191505060405180910390fd5b816002015481146109175760405162461bcd60e51b8152600401808060200182810382526039815260200180611e956039913960400191505060405180910390fd5b600582015433600090815260208490526040812080546001600160801b0390931690920390915560028301556006820180546001919060ff60501b1916600160501b83021790555060068201805460ff60581b1916600160581b1790556040805133815290517f255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b79181900360200190a15050565b610a078a6040518060a001604052808c81526020018b81526020018a63ffffffff1681526020018981526020016040518060a001604052808a815260200189815260200188815260200187815260200186815250815250610abc565b50505050505050505050565b610ab1896040518061016001604052808b600060048110610a3057fe5b602002015181526020018b600160048110610a4757fe5b602002015181526020018b600260048110610a5e57fe5b602002015181526020018a63ffffffff1681526020018981526020018881526020018781526020018681526020018581526020018481526020018b600360048110610aa557fe5b6020020151905261122b565b505050505050505050565b60026006830154600160501b900460ff166003811115610ad857fe5b14610b145760405162461bcd60e51b8152600401808060200182810382526022815260200180611f4f6022913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610b6357600080fd5b505af4158015610b77573d6000803e3d6000fd5b505050506040513d6020811015610b8d57600080fd5b505115610bcb5760405162461bcd60e51b8152600401808060200182810382526024815260200180611f2b6024913960400191505060405180910390fd5b8160020154816000015173__$9836fa7140e5a33041d4b827682e675a30$__632090372184602001518560400151600073__$9836fa7140e5a33041d4b827682e675a30$__631914612a89606001518a60800151600001518b60800151602001518c60800151604001518d60800151606001516040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015610c9d578181015183820152602001610c85565b5050505090500186810385528a818151815260200191508051906020019080838360005b83811015610cd9578181015183820152602001610cc1565b50505050905090810190601f168015610d065780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015610d3b578181015183820152602001610d23565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015610d7a578181015183820152602001610d62565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015610db9578181015183820152602001610da1565b505050509050019a505050505050505050505060206040518083038186803b158015610de457600080fd5b505af4158015610df8573d6000803e3d6000fd5b505050506040513d6020811015610e0e57600080fd5b8101908080519060200190929190505050600089608001516080015173__$9836fa7140e5a33041d4b827682e675a30$__630f89fbff8c606001518d60800151602001518e60800151604001516040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015610eb2578181015183820152602001610e9a565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015610ef1578181015183820152602001610ed9565b50505050905001848103825285818151815260200191508051906020019060200280838360005b83811015610f30578181015183820152602001610f18565b50505050905001965050505050505060006040518083038186803b158015610f5757600080fd5b505af4158015610f6b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610f9457600080fd5b810190808051600160201b811115610fab57600080fd5b82016020810184811115610fbe57600080fd5b81518560208202830111600160201b82111715610fda57600080fd5b505060405160e08c811b6001600160e01b0319168252600482018c815263ffffffff8c166024840152604483018b9052606483018a90526084830189905260a4830188905260c48301918252835160e484015283519396509450925061010401906020808601910280838360005b83811015611060578181015183820152602001611048565b505050509050019850505050505050505060206040518083038186803b15801561108957600080fd5b505af415801561109d573d6000803e3d6000fd5b505050506040513d60208110156110b357600080fd5b505160408051602081810194909452808201929092528051808303820181526060909201905280519101201461111a5760405162461bcd60e51b8152600401808060200182810382526039815260200180611e956039913960400191505060405180910390fd5b600582015460048301546001600160a01b0316600090815260208490526040902054611154916001600160801b031663ffffffff611e0c16565b6004808401546001600160a01b0316600090815260208581526040808320949094558401518351633ad2660b60e21b81529283018690526024830152915173__$8e266570c8a7fb2aaac83b3e040afaf9e1$__9263eb49982c9260448082019391829003018186803b1580156111c957600080fd5b505af41580156111dd573d6000803e3d6000fd5b5050506020808301516080808501510151604080519283529282015281517f4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb93509081900390910190a15050565b60016006830154600160501b900460ff16600381111561124757fe5b146112835760405162461bcd60e51b815260040180806020018281038252602d81526020018061208a602d913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__632a3e0a97836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156112d257600080fd5b505af41580156112e6573d6000803e3d6000fd5b505050506040513d60208110156112fc57600080fd5b5051158015611383575073__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63e2fe93ca836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561135557600080fd5b505af4158015611369573d6000803e3d6000fd5b505050506040513d602081101561137f57600080fd5b5051155b6113be5760405162461bcd60e51b815260040180806020018281038252603e815260200180611fb6603e913960400191505060405180910390fd5b6006820154600160581b900460ff16156114095760405162461bcd60e51b815260040180806020018281038252602e815260200180611efd602e913960400191505060405180910390fd5b3360009081526020839052604090205460058301546001600160801b031611156114645760405162461bcd60e51b81526004018080602001828103825260278152602001806120166027913960400191505060405180910390fd5b6006820154606082015163ffffffff600160201b9092048216911611156114d2576040805162461bcd60e51b815260206004820152601f60248201527f547269656420746f206578656375746520746f6f206d616e7920737465707300604482015290519081900360640190fd5b6114df81608001516106ee565b61151a5760405162461bcd60e51b8152600401808060200182810382526024815260200180611f716024913960400191505060405180910390fd5b600182015481511461155d5760405162461bcd60e51b815260040180806020018281038252602781526020018061203d6027913960400191505060405180910390fd5b80602001518260030154146115a35760405162461bcd60e51b8152600401808060200182810382526022815260200180611ff46022913960400191505060405180910390fd5b606073__$9836fa7140e5a33041d4b827682e675a30$__630f89fbff8360a001518460e001518561010001516040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b8381101561162657818101518382015260200161160e565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561166557818101518382015260200161164d565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156116a457818101518382015260200161168c565b50505050905001965050505050505060006040518083038186803b1580156116cb57600080fd5b505af41580156116df573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561170857600080fd5b810190808051600160201b81111561171f57600080fd5b8201602081018481111561173257600080fd5b81518560208202830111600160201b8211171561174e57600080fd5b50506040805163a3a162cb60e01b815260048101899052905191955073__$8e266570c8a7fb2aaac83b3e040afaf9e1$__945063a3a162cb93506024808201935060009291829003018186803b1580156117a757600080fd5b505af41580156117bb573d6000803e3d6000fd5b50505050600073__$9836fa7140e5a33041d4b827682e675a30$__63d14cf0988460a001518560c001518660e001518761010001518861012001516040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b8381101561185557818101518382015260200161183d565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b8381101561189457818101518382015260200161187c565b50505050905001868103845289818151815260200191508051906020019060200280838360005b838110156118d35781810151838201526020016118bb565b50505050905001868103835288818151815260200191508051906020019060200280838360005b838110156119125781810151838201526020016118fa565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015611951578181015183820152602001611939565b505050509050019a505050505050505050505060206040518083038186803b15801561197c57600080fd5b505af4158015611990573d6000803e3d6000fd5b505050506040513d60208110156119a657600080fd5b505183516080850151602086015160a0870151604080516307c50ab360e31b81526004810186815296975073__$9836fa7140e5a33041d4b827682e675a30$__96633e28559896959493928a9260240190869080838360005b83811015611a175781810151838201526020016119ff565b505050509050018481526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015611a64578181015183820152602001611a4c565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611aa3578181015183820152602001611a8b565b5050505090500197505050505050505060206040518083038186803b158015611acb57600080fd5b505af4158015611adf573d6000803e3d6000fd5b505050506040513d6020811015611af557600080fd5b505160408085015160608601516101408701519251632090372160e01b81526004810183815263ffffffff83166024830152600060448301819052606483018890526084830181905260a4830186905260e060c48401908152895160e4850152895173__$9836fa7140e5a33041d4b827682e675a30$__97632090372197969593948b94869492938e93916101040190602085810191028083838a5b83811015611ba9578181015183820152602001611b91565b505050509050019850505050505050505060206040518083038186803b158015611bd257600080fd5b505af4158015611be6573d6000803e3d6000fd5b505050506040513d6020811015611bfc57600080fd5b50516040805160208181019490945280820192909252805180830382018152606090920181528151918301919091206002808801919091556005870154336000818152948990529290932080546001600160801b039094169093039092556004860180546001600160a01b031916909117905560068501805460ff60501b1916600160501b8302179055507f5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f6040518060600160405280856000015181526020018560200151815260200185604001518152503385608001518660a00151876060015186896101400151896040518089600360200280838360005b83811015611d0f578181015183820152602001611cf7565b505050506001600160a01b038b1692019182525060200187604080838360005b83811015611d47578181015183820152602001611d2f565b50505050905001806020018663ffffffff1663ffffffff16815260200185815260200184815260200180602001838103835288818151815260200191508051906020019060200280838360005b83811015611dac578181015183820152602001611d94565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611deb578181015183820152602001611dd3565b505050509050019a505050505050505050505060405180910390a150505050565b600082820183811015611e66576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe4368616c6c656e67657220646964206e6f74206861766520656e6f75676820657363726f776564507265636f6e646974696f6e20616e6420617373657274696f6e20646f206e6f74206d617463682070656e64696e6720617373657274696f6e417373657274696f6e206d7573742062652070656e64696e6720746f20696e697469617465206368616c6c656e676543616e206f6e6c792064697370757461626c6520617373657274206966206e6f7420696e206368616c6c656e6765417373657274696f6e206973207374696c6c2070656e64696e67206368616c6c656e6765564d20646f6573206e6f74206861766520617373657274696f6e2070656e64696e67507265636f6e646974696f6e3a206e6f742077697468696e2074696d6520626f756e64734368616c6c656e676520776173206372656174656420627920617373657274657243616e206f6e6c792064697370757461626c6520617373657274206966206d616368696e65206973206e6f74206572726f726564206f722068616c746564507265636f6e646974696f6e3a20696e626f7820646f6573206e6f74206d6174636856616c696461746f7220646f6573206e6f74206861766520726571756972656420657363726f77507265636f6e646974696f6e3a207374617465206861736820646f6573206e6f74206d617463684368616c6c656e676520646964206e6f7420636f6d65206265666f726520646561646c696e6543616e206f6e6c792064697370757461626c65206173736572742066726f6d2077616974696e67207374617465a265627a7a72305820ab08dee8199f1a992b78bfa0b6cefb269e87d70e45f31c38dbe84707e1fd1d6d64736f6c634300050a0032"

// DeployDisputable deploys a new Ethereum contract, binding an instance of Disputable to it.
func DeployDisputable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Disputable, error) {
	parsed, err := abi.JSON(strings.NewReader(DisputableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	vMAddr, _, _, _ := DeployVM(auth, backend)
	DisputableBin = strings.Replace(DisputableBin, "__$8e266570c8a7fb2aaac83b3e040afaf9e1$__", vMAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	DisputableBin = strings.Replace(DisputableBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DisputableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Disputable{DisputableCaller: DisputableCaller{contract: contract}, DisputableTransactor: DisputableTransactor{contract: contract}, DisputableFilterer: DisputableFilterer{contract: contract}}, nil
}

// Disputable is an auto generated Go binding around an Ethereum contract.
type Disputable struct {
	DisputableCaller     // Read-only binding to the contract
	DisputableTransactor // Write-only binding to the contract
	DisputableFilterer   // Log filterer for contract events
}

// DisputableCaller is an auto generated read-only Go binding around an Ethereum contract.
type DisputableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DisputableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DisputableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DisputableSession struct {
	Contract     *Disputable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DisputableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DisputableCallerSession struct {
	Contract *DisputableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DisputableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DisputableTransactorSession struct {
	Contract     *DisputableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DisputableRaw is an auto generated low-level Go binding around an Ethereum contract.
type DisputableRaw struct {
	Contract *Disputable // Generic contract binding to access the raw methods on
}

// DisputableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DisputableCallerRaw struct {
	Contract *DisputableCaller // Generic read-only contract binding to access the raw methods on
}

// DisputableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DisputableTransactorRaw struct {
	Contract *DisputableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDisputable creates a new instance of Disputable, bound to a specific deployed contract.
func NewDisputable(address common.Address, backend bind.ContractBackend) (*Disputable, error) {
	contract, err := bindDisputable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Disputable{DisputableCaller: DisputableCaller{contract: contract}, DisputableTransactor: DisputableTransactor{contract: contract}, DisputableFilterer: DisputableFilterer{contract: contract}}, nil
}

// NewDisputableCaller creates a new read-only instance of Disputable, bound to a specific deployed contract.
func NewDisputableCaller(address common.Address, caller bind.ContractCaller) (*DisputableCaller, error) {
	contract, err := bindDisputable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DisputableCaller{contract: contract}, nil
}

// NewDisputableTransactor creates a new write-only instance of Disputable, bound to a specific deployed contract.
func NewDisputableTransactor(address common.Address, transactor bind.ContractTransactor) (*DisputableTransactor, error) {
	contract, err := bindDisputable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DisputableTransactor{contract: contract}, nil
}

// NewDisputableFilterer creates a new log filterer instance of Disputable, bound to a specific deployed contract.
func NewDisputableFilterer(address common.Address, filterer bind.ContractFilterer) (*DisputableFilterer, error) {
	contract, err := bindDisputable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DisputableFilterer{contract: contract}, nil
}

// bindDisputable binds a generic wrapper to an already deployed contract.
func bindDisputable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DisputableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Disputable *DisputableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Disputable.Contract.DisputableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Disputable *DisputableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Disputable.Contract.DisputableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Disputable *DisputableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Disputable.Contract.DisputableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Disputable *DisputableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Disputable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Disputable *DisputableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Disputable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Disputable *DisputableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Disputable.Contract.contract.Transact(opts, method, params...)
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_Disputable *DisputableCaller) WithinTimeBounds(opts *bind.CallOpts, _timeBounds [2]uint64) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Disputable.contract.Call(opts, out, "withinTimeBounds", _timeBounds)
	return *ret0, err
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_Disputable *DisputableSession) WithinTimeBounds(_timeBounds [2]uint64) (bool, error) {
	return _Disputable.Contract.WithinTimeBounds(&_Disputable.CallOpts, _timeBounds)
}

// WithinTimeBounds is a free data retrieval call binding the contract method 0x42c0787e.
//
// Solidity: function withinTimeBounds(uint64[2] _timeBounds) constant returns(bool)
func (_Disputable *DisputableCallerSession) WithinTimeBounds(_timeBounds [2]uint64) (bool, error) {
	return _Disputable.Contract.WithinTimeBounds(&_Disputable.CallOpts, _timeBounds)
}

// DisputableConfirmedDisputableAssertionIterator is returned from FilterConfirmedDisputableAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedDisputableAssertion events raised by the Disputable contract.
type DisputableConfirmedDisputableAssertionIterator struct {
	Event *DisputableConfirmedDisputableAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DisputableConfirmedDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputableConfirmedDisputableAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DisputableConfirmedDisputableAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DisputableConfirmedDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputableConfirmedDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputableConfirmedDisputableAssertion represents a ConfirmedDisputableAssertion event raised by the Disputable contract.
type DisputableConfirmedDisputableAssertion struct {
	NewState    [32]byte
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedDisputableAssertion is a free log retrieval operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_Disputable *DisputableFilterer) FilterConfirmedDisputableAssertion(opts *bind.FilterOpts) (*DisputableConfirmedDisputableAssertionIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &DisputableConfirmedDisputableAssertionIterator{contract: _Disputable.contract, event: "ConfirmedDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedDisputableAssertion is a free log subscription operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_Disputable *DisputableFilterer) WatchConfirmedDisputableAssertion(opts *bind.WatchOpts, sink chan<- *DisputableConfirmedDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _Disputable.contract.WatchLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputableConfirmedDisputableAssertion)
				if err := _Disputable.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmedDisputableAssertion is a log parse operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_Disputable *DisputableFilterer) ParseConfirmedDisputableAssertion(log types.Log) (*DisputableConfirmedDisputableAssertion, error) {
	event := new(DisputableConfirmedDisputableAssertion)
	if err := _Disputable.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DisputableInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the Disputable contract.
type DisputableInitiatedChallengeIterator struct {
	Event *DisputableInitiatedChallenge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DisputableInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputableInitiatedChallenge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DisputableInitiatedChallenge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DisputableInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputableInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputableInitiatedChallenge represents a InitiatedChallenge event raised by the Disputable contract.
type DisputableInitiatedChallenge struct {
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_Disputable *DisputableFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*DisputableInitiatedChallengeIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &DisputableInitiatedChallengeIterator{contract: _Disputable.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_Disputable *DisputableFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *DisputableInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _Disputable.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputableInitiatedChallenge)
				if err := _Disputable.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_Disputable *DisputableFilterer) ParseInitiatedChallenge(log types.Log) (*DisputableInitiatedChallenge, error) {
	event := new(DisputableInitiatedChallenge)
	if err := _Disputable.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DisputablePendingDisputableAssertionIterator is returned from FilterPendingDisputableAssertion and is used to iterate over the raw logs and unpacked data for PendingDisputableAssertion events raised by the Disputable contract.
type DisputablePendingDisputableAssertionIterator struct {
	Event *DisputablePendingDisputableAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DisputablePendingDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputablePendingDisputableAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DisputablePendingDisputableAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DisputablePendingDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputablePendingDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputablePendingDisputableAssertion represents a PendingDisputableAssertion event raised by the Disputable contract.
type DisputablePendingDisputableAssertion struct {
	Fields          [3][32]byte
	Asserter        common.Address
	TimeBounds      [2]uint64
	TokenTypes      [][21]byte
	NumSteps        uint32
	LastMessageHash [32]byte
	LogsAccHash     [32]byte
	Amounts         []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_Disputable *DisputableFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*DisputablePendingDisputableAssertionIterator, error) {

	logs, sub, err := _Disputable.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &DisputablePendingDisputableAssertionIterator{contract: _Disputable.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_Disputable *DisputableFilterer) WatchPendingDisputableAssertion(opts *bind.WatchOpts, sink chan<- *DisputablePendingDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _Disputable.contract.WatchLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputablePendingDisputableAssertion)
				if err := _Disputable.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_Disputable *DisputableFilterer) ParsePendingDisputableAssertion(log types.Log) (*DisputablePendingDisputableAssertion, error) {
	event := new(DisputablePendingDisputableAssertion)
	if err := _Disputable.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IChallengeManagerABI is the input ABI used to generate the binding from.
const IChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"players\",\"type\":\"address[2]\"},{\"name\":\"escrows\",\"type\":\"uint128[2]\"},{\"name\":\"challengePeriod\",\"type\":\"uint32\"},{\"name\":\"challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeManagerFuncSigs = map[string]string{
	"208e04d4": "initiateChallenge(address[2],uint128[2],uint32,bytes32)",
}

// IChallengeManager is an auto generated Go binding around an Ethereum contract.
type IChallengeManager struct {
	IChallengeManagerCaller     // Read-only binding to the contract
	IChallengeManagerTransactor // Write-only binding to the contract
	IChallengeManagerFilterer   // Log filterer for contract events
}

// IChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerSession struct {
	Contract     *IChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCallerSession struct {
	Contract *IChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerTransactorSession struct {
	Contract     *IChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerRaw struct {
	Contract *IChallengeManager // Generic contract binding to access the raw methods on
}

// IChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCallerRaw struct {
	Contract *IChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactorRaw struct {
	Contract *IChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManager creates a new instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManager(address common.Address, backend bind.ContractBackend) (*IChallengeManager, error) {
	contract, err := bindIChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCaller creates a new read-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCaller, error) {
	contract, err := bindIChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCaller{contract: contract}, nil
}

// NewIChallengeManagerTransactor creates a new write-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerTransactor, error) {
	contract, err := bindIChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerTransactor{contract: contract}, nil
}

// NewIChallengeManagerFilterer creates a new log filterer instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerFilterer, error) {
	contract, err := bindIChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerFilterer{contract: contract}, nil
}

// bindIChallengeManager binds a generic wrapper to an already deployed contract.
func bindIChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.IChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initiateChallenge", players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerSession) InitiateChallenge(players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) InitiateChallenge(players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, players, escrows, challengePeriod, challengeRoot)
}

// IGlobalPendingInboxABI is the input ABI used to generate the binding from.
const IGlobalPendingInboxABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_vmId\",\"type\":\"address\"}],\"name\":\"pullPendingMessages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"hasFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_tokenTypeNum\",\"type\":\"uint16[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_destinations\",\"type\":\"address[]\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerForInbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmId\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenType\",\"type\":\"bytes21\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"}]"

// IGlobalPendingInboxFuncSigs maps the 4-byte function signature to its string representation.
var IGlobalPendingInboxFuncSigs = map[string]string{
	"3bbc3c32": "forwardMessage(address,bytes21,uint256,bytes,bytes)",
	"acb633b6": "hasFunds(address,bytes21[],uint256[])",
	"31618a03": "pullPendingMessages(address)",
	"f3972383": "registerForInbox()",
	"3fc6eb80": "sendEthMessage(address,bytes)",
	"626cef85": "sendMessage(address,bytes21,uint256,bytes)",
	"ec22a767": "sendMessages(bytes21[],bytes,uint16[],uint256[],address[])",
}

// IGlobalPendingInbox is an auto generated Go binding around an Ethereum contract.
type IGlobalPendingInbox struct {
	IGlobalPendingInboxCaller     // Read-only binding to the contract
	IGlobalPendingInboxTransactor // Write-only binding to the contract
	IGlobalPendingInboxFilterer   // Log filterer for contract events
}

// IGlobalPendingInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGlobalPendingInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGlobalPendingInboxSession struct {
	Contract     *IGlobalPendingInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGlobalPendingInboxCallerSession struct {
	Contract *IGlobalPendingInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IGlobalPendingInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGlobalPendingInboxTransactorSession struct {
	Contract     *IGlobalPendingInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGlobalPendingInboxRaw struct {
	Contract *IGlobalPendingInbox // Generic contract binding to access the raw methods on
}

// IGlobalPendingInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCallerRaw struct {
	Contract *IGlobalPendingInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IGlobalPendingInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactorRaw struct {
	Contract *IGlobalPendingInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGlobalPendingInbox creates a new instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInbox(address common.Address, backend bind.ContractBackend) (*IGlobalPendingInbox, error) {
	contract, err := bindIGlobalPendingInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInbox{IGlobalPendingInboxCaller: IGlobalPendingInboxCaller{contract: contract}, IGlobalPendingInboxTransactor: IGlobalPendingInboxTransactor{contract: contract}, IGlobalPendingInboxFilterer: IGlobalPendingInboxFilterer{contract: contract}}, nil
}

// NewIGlobalPendingInboxCaller creates a new read-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxCaller(address common.Address, caller bind.ContractCaller) (*IGlobalPendingInboxCaller, error) {
	contract, err := bindIGlobalPendingInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxCaller{contract: contract}, nil
}

// NewIGlobalPendingInboxTransactor creates a new write-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IGlobalPendingInboxTransactor, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxTransactor{contract: contract}, nil
}

// NewIGlobalPendingInboxFilterer creates a new log filterer instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IGlobalPendingInboxFilterer, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxFilterer{contract: contract}, nil
}

// bindIGlobalPendingInbox binds a generic wrapper to an already deployed contract.
func bindIGlobalPendingInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGlobalPendingInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transact(opts, method, params...)
}

// HasFunds is a free data retrieval call binding the contract method 0xacb633b6.
//
// Solidity: function hasFunds(address _owner, bytes21[] _tokenTypes, uint256[] _amounts) constant returns(bool)
func (_IGlobalPendingInbox *IGlobalPendingInboxCaller) HasFunds(opts *bind.CallOpts, _owner common.Address, _tokenTypes [][21]byte, _amounts []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IGlobalPendingInbox.contract.Call(opts, out, "hasFunds", _owner, _tokenTypes, _amounts)
	return *ret0, err
}

// HasFunds is a free data retrieval call binding the contract method 0xacb633b6.
//
// Solidity: function hasFunds(address _owner, bytes21[] _tokenTypes, uint256[] _amounts) constant returns(bool)
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) HasFunds(_owner common.Address, _tokenTypes [][21]byte, _amounts []*big.Int) (bool, error) {
	return _IGlobalPendingInbox.Contract.HasFunds(&_IGlobalPendingInbox.CallOpts, _owner, _tokenTypes, _amounts)
}

// HasFunds is a free data retrieval call binding the contract method 0xacb633b6.
//
// Solidity: function hasFunds(address _owner, bytes21[] _tokenTypes, uint256[] _amounts) constant returns(bool)
func (_IGlobalPendingInbox *IGlobalPendingInboxCallerSession) HasFunds(_owner common.Address, _tokenTypes [][21]byte, _amounts []*big.Int) (bool, error) {
	return _IGlobalPendingInbox.Contract.HasFunds(&_IGlobalPendingInbox.CallOpts, _owner, _tokenTypes, _amounts)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0x3bbc3c32.
//
// Solidity: function forwardMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) ForwardMessage(opts *bind.TransactOpts, _destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "forwardMessage", _destination, _tokenType, _amount, _data, _signature)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0x3bbc3c32.
//
// Solidity: function forwardMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) ForwardMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data, _signature)
}

// ForwardMessage is a paid mutator transaction binding the contract method 0x3bbc3c32.
//
// Solidity: function forwardMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) ForwardMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data, _signature)
}

// PullPendingMessages is a paid mutator transaction binding the contract method 0x31618a03.
//
// Solidity: function pullPendingMessages(address _vmId) returns(bytes32)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) PullPendingMessages(opts *bind.TransactOpts, _vmId common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "pullPendingMessages", _vmId)
}

// PullPendingMessages is a paid mutator transaction binding the contract method 0x31618a03.
//
// Solidity: function pullPendingMessages(address _vmId) returns(bytes32)
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) PullPendingMessages(_vmId common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.PullPendingMessages(&_IGlobalPendingInbox.TransactOpts, _vmId)
}

// PullPendingMessages is a paid mutator transaction binding the contract method 0x31618a03.
//
// Solidity: function pullPendingMessages(address _vmId) returns(bytes32)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) PullPendingMessages(_vmId common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.PullPendingMessages(&_IGlobalPendingInbox.TransactOpts, _vmId)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) RegisterForInbox(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "registerForInbox")
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0x3fc6eb80.
//
// Solidity: function sendEthMessage(address _destination, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendEthMessage(opts *bind.TransactOpts, _destination common.Address, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendEthMessage", _destination, _data)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0x3fc6eb80.
//
// Solidity: function sendEthMessage(address _destination, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendEthMessage(_destination common.Address, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendEthMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _data)
}

// SendEthMessage is a paid mutator transaction binding the contract method 0x3fc6eb80.
//
// Solidity: function sendEthMessage(address _destination, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendEthMessage(_destination common.Address, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendEthMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x626cef85.
//
// Solidity: function sendMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendMessage(opts *bind.TransactOpts, _destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendMessage", _destination, _tokenType, _amount, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x626cef85.
//
// Solidity: function sendMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data)
}

// SendMessage is a paid mutator transaction binding the contract method 0x626cef85.
//
// Solidity: function sendMessage(address _destination, bytes21 _tokenType, uint256 _amount, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendMessage(_destination common.Address, _tokenType [21]byte, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessage(&_IGlobalPendingInbox.TransactOpts, _destination, _tokenType, _amount, _data)
}

// SendMessages is a paid mutator transaction binding the contract method 0xec22a767.
//
// Solidity: function sendMessages(bytes21[] _tokenTypes, bytes _messageData, uint16[] _tokenTypeNum, uint256[] _amounts, address[] _destinations) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendMessages(opts *bind.TransactOpts, _tokenTypes [][21]byte, _messageData []byte, _tokenTypeNum []uint16, _amounts []*big.Int, _destinations []common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendMessages", _tokenTypes, _messageData, _tokenTypeNum, _amounts, _destinations)
}

// SendMessages is a paid mutator transaction binding the contract method 0xec22a767.
//
// Solidity: function sendMessages(bytes21[] _tokenTypes, bytes _messageData, uint16[] _tokenTypeNum, uint256[] _amounts, address[] _destinations) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendMessages(_tokenTypes [][21]byte, _messageData []byte, _tokenTypeNum []uint16, _amounts []*big.Int, _destinations []common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _tokenTypes, _messageData, _tokenTypeNum, _amounts, _destinations)
}

// SendMessages is a paid mutator transaction binding the contract method 0xec22a767.
//
// Solidity: function sendMessages(bytes21[] _tokenTypes, bytes _messageData, uint16[] _tokenTypeNum, uint256[] _amounts, address[] _destinations) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendMessages(_tokenTypes [][21]byte, _messageData []byte, _tokenTypeNum []uint16, _amounts []*big.Int, _destinations []common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _tokenTypes, _messageData, _tokenTypeNum, _amounts, _destinations)
}

// IGlobalPendingInboxMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGlobalPendingInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGlobalPendingInboxMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGlobalPendingInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxMessageDelivered represents a MessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxMessageDelivered struct {
	VmId      common.Address
	Sender    common.Address
	TokenType [21]byte
	Value     *big.Int
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x4d0d890cdec30a2409c07864cb0bdbd32b2f7f57aaf8966b83df1bd2a5da3384.
//
// Solidity: event MessageDelivered(address indexed vmId, address sender, bytes21 tokenType, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterMessageDelivered(opts *bind.FilterOpts, vmId []common.Address) (*IGlobalPendingInboxMessageDeliveredIterator, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "MessageDelivered", vmIdRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x4d0d890cdec30a2409c07864cb0bdbd32b2f7f57aaf8966b83df1bd2a5da3384.
//
// Solidity: event MessageDelivered(address indexed vmId, address sender, bytes21 tokenType, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxMessageDelivered, vmId []common.Address) (event.Subscription, error) {

	var vmIdRule []interface{}
	for _, vmIdItem := range vmId {
		vmIdRule = append(vmIdRule, vmIdItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "MessageDelivered", vmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDelivered is a log parse operation binding the contract event 0x4d0d890cdec30a2409c07864cb0bdbd32b2f7f57aaf8966b83df1bd2a5da3384.
//
// Solidity: event MessageDelivered(address indexed vmId, address sender, bytes21 tokenType, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseMessageDelivered(log types.Log) (*IGlobalPendingInboxMessageDelivered, error) {
	event := new(IGlobalPendingInboxMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058203ff4034c6ea4209ece9da4562199bebb6f4a5e4818ff91583cbf2d23590e03bd64736f6c634300050a0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SigUtilsABI is the input ABI used to generate the binding from.
const SigUtilsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"countSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_signatures\",\"type\":\"bytes\"},{\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"parseSignature\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"recoverAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SigUtilsFuncSigs maps the 4-byte function signature to its string representation.
var SigUtilsFuncSigs = map[string]string{
	"33ae3ad0": "countSignatures(bytes)",
	"b31d63cc": "parseSignature(bytes,uint256)",
	"c655d7aa": "recoverAddress(bytes32,bytes)",
	"f0c8e969": "recoverAddresses(bytes32,bytes)",
}

// SigUtilsBin is the compiled bytecode used for deploying new contracts.
var SigUtilsBin = "0x610764610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806333ae3ad01461005b578063b31d63cc14610111578063c655d7aa146101d9578063f0c8e969146102a0575b600080fd5b6100ff6004803603602081101561007157600080fd5b810190602081018135600160201b81111561008b57600080fd5b82018360208201111561009d57600080fd5b803590602001918460018302840111600160201b831117156100be57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061039b945050505050565b60408051918252519081900360200190f35b6101b76004803603604081101561012757600080fd5b810190602081018135600160201b81111561014157600080fd5b82018360208201111561015357600080fd5b803590602001918460018302840111600160201b8311171561017457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506103c8915050565b6040805160ff9094168452602084019290925282820152519081900360600190f35b610284600480360360408110156101ef57600080fd5b81359190810190604081016020820135600160201b81111561021057600080fd5b82018360208201111561022257600080fd5b803590602001918460018302840111600160201b8311171561024357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610456945050505050565b604080516001600160a01b039092168252519081900360200190f35b61034b600480360360408110156102b657600080fd5b81359190810190604081016020820135600160201b8111156102d757600080fd5b8201836020820111156102e957600080fd5b803590602001918460018302840111600160201b8311171561030a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610589945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561038757818101518382015260200161036f565b505050509050019250505060405180910390f35b600060418251816103a857fe5b06156103b55760006103c2565b60418251816103c057fe5b045b92915050565b604180820283810160208101516040820151919093015160ff169291601b8410156103f457601b840193505b8360ff16601b148061040957508360ff16601c145b61044e576040805162461bcd60e51b8152602060048201526011602482015270496e636f727265637420762076616c756560781b604482015290519081900360640190fd5b509250925092565b60008060008060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081886040516020018083805190602001908083835b602083106104cc5780518252601f1990920191602091820191016104ad565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152938201905282519201919091209250610513915088905060006103c8565b6040805160008152602080820180845287905260ff8616828401526060820185905260808201849052915194995092975090955060019260a080840193601f198301929081900390910190855afa158015610572573d6000803e3d6000fd5b5050604051601f1901519998505050505050505050565b606060008060008061059a8661039b565b90506060816040519080825280602002602001820160405280156105c8578160200160208202803883390190505b50905060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090506000818a6040516020018083805190602001908083835b6020831061063b5780518252601f19909201916020918201910161061c565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190528251920191909120925060009150505b848110156107205761068b8a826103c8565b6040805160008152602080820180845288905260ff86168284015260608201859052608082018490529151949c50929a5090985060019260a080840193601f198301929081900390910190855afa1580156106ea573d6000803e3d6000fd5b5050506020604051035184828151811061070057fe5b6001600160a01b0390921660209283029190910190910152600101610679565b5091999850505050505050505056fea265627a7a72305820b6df419ca41f5fa0e34f46087ebbe225a45bf35ef03c0c5e3c6746bef674435364736f6c634300050a0032"

// DeploySigUtils deploys a new Ethereum contract, binding an instance of SigUtils to it.
func DeploySigUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// SigUtils is an auto generated Go binding around an Ethereum contract.
type SigUtils struct {
	SigUtilsCaller     // Read-only binding to the contract
	SigUtilsTransactor // Write-only binding to the contract
	SigUtilsFilterer   // Log filterer for contract events
}

// SigUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigUtilsSession struct {
	Contract     *SigUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigUtilsCallerSession struct {
	Contract *SigUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SigUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigUtilsTransactorSession struct {
	Contract     *SigUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SigUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigUtilsRaw struct {
	Contract *SigUtils // Generic contract binding to access the raw methods on
}

// SigUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigUtilsCallerRaw struct {
	Contract *SigUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SigUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigUtilsTransactorRaw struct {
	Contract *SigUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigUtils creates a new instance of SigUtils, bound to a specific deployed contract.
func NewSigUtils(address common.Address, backend bind.ContractBackend) (*SigUtils, error) {
	contract, err := bindSigUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// NewSigUtilsCaller creates a new read-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsCaller(address common.Address, caller bind.ContractCaller) (*SigUtilsCaller, error) {
	contract, err := bindSigUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsCaller{contract: contract}, nil
}

// NewSigUtilsTransactor creates a new write-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SigUtilsTransactor, error) {
	contract, err := bindSigUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTransactor{contract: contract}, nil
}

// NewSigUtilsFilterer creates a new log filterer instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SigUtilsFilterer, error) {
	contract, err := bindSigUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigUtilsFilterer{contract: contract}, nil
}

// bindSigUtils binds a generic wrapper to an already deployed contract.
func bindSigUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.SigUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transact(opts, method, params...)
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_SigUtils *SigUtilsCaller) CountSignatures(opts *bind.CallOpts, _signatures []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SigUtils.contract.Call(opts, out, "countSignatures", _signatures)
	return *ret0, err
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_SigUtils *SigUtilsSession) CountSignatures(_signatures []byte) (*big.Int, error) {
	return _SigUtils.Contract.CountSignatures(&_SigUtils.CallOpts, _signatures)
}

// CountSignatures is a free data retrieval call binding the contract method 0x33ae3ad0.
//
// Solidity: function countSignatures(bytes _signatures) constant returns(uint256)
func (_SigUtils *SigUtilsCallerSession) CountSignatures(_signatures []byte) (*big.Int, error) {
	return _SigUtils.Contract.CountSignatures(&_SigUtils.CallOpts, _signatures)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_SigUtils *SigUtilsCaller) ParseSignature(opts *bind.CallOpts, _signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	ret := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	out := ret
	err := _SigUtils.contract.Call(opts, out, "parseSignature", _signatures, _pos)
	return *ret, err
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_SigUtils *SigUtilsSession) ParseSignature(_signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigUtils.Contract.ParseSignature(&_SigUtils.CallOpts, _signatures, _pos)
}

// ParseSignature is a free data retrieval call binding the contract method 0xb31d63cc.
//
// Solidity: function parseSignature(bytes _signatures, uint256 _pos) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_SigUtils *SigUtilsCallerSession) ParseSignature(_signatures []byte, _pos *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigUtils.Contract.ParseSignature(&_SigUtils.CallOpts, _signatures, _pos)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 _messageHash, bytes _signature) constant returns(address)
func (_SigUtils *SigUtilsCaller) RecoverAddress(opts *bind.CallOpts, _messageHash [32]byte, _signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SigUtils.contract.Call(opts, out, "recoverAddress", _messageHash, _signature)
	return *ret0, err
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 _messageHash, bytes _signature) constant returns(address)
func (_SigUtils *SigUtilsSession) RecoverAddress(_messageHash [32]byte, _signature []byte) (common.Address, error) {
	return _SigUtils.Contract.RecoverAddress(&_SigUtils.CallOpts, _messageHash, _signature)
}

// RecoverAddress is a free data retrieval call binding the contract method 0xc655d7aa.
//
// Solidity: function recoverAddress(bytes32 _messageHash, bytes _signature) constant returns(address)
func (_SigUtils *SigUtilsCallerSession) RecoverAddress(_messageHash [32]byte, _signature []byte) (common.Address, error) {
	return _SigUtils.Contract.RecoverAddress(&_SigUtils.CallOpts, _messageHash, _signature)
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_SigUtils *SigUtilsCaller) RecoverAddresses(opts *bind.CallOpts, _messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _SigUtils.contract.Call(opts, out, "recoverAddresses", _messageHash, _signatures)
	return *ret0, err
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_SigUtils *SigUtilsSession) RecoverAddresses(_messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	return _SigUtils.Contract.RecoverAddresses(&_SigUtils.CallOpts, _messageHash, _signatures)
}

// RecoverAddresses is a free data retrieval call binding the contract method 0xf0c8e969.
//
// Solidity: function recoverAddresses(bytes32 _messageHash, bytes _signatures) constant returns(address[])
func (_SigUtils *SigUtilsCallerSession) RecoverAddresses(_messageHash [32]byte, _signatures []byte) ([]common.Address, error) {
	return _SigUtils.Contract.RecoverAddresses(&_SigUtils.CallOpts, _messageHash, _signatures)
}

// UnanimousABI is the input ABI used to generate the binding from.
const UnanimousABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"unanHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"PendingUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"ConfirmedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"unanHash\",\"type\":\"bytes32\"}],\"name\":\"FinalizedUnanimousAssertion\",\"type\":\"event\"}]"

// UnanimousFuncSigs maps the 4-byte function signature to its string representation.
var UnanimousFuncSigs = map[string]string{
	"319b5f67": "confirmUnanimousAsserted(VM.Data storage,bytes32,bytes32,bytes21[],bytes,uint16[],uint256[],address[])",
	"3c0ec6de": "finalizedUnanimousAssert(VM.Data storage,bytes32[3],bytes21[],bytes,uint16[],uint256[],address[],bytes)",
	"a2ee90f4": "pendingUnanimousAssert(VM.Data storage,bytes32,bytes21[],uint16[],uint256[],uint64,bytes32,bytes)",
}

// UnanimousBin is the compiled bytecode used for deploying new contracts.
var UnanimousBin = "0x611a57610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c8063319b5f67146100505780633c0ec6de1461031f578063a2ee90f414610695575b600080fd5b81801561005c57600080fd5b5061031d600480360361010081101561007457600080fd5b81359160208101359160408201359190810190608081016060820135600160201b8111156100a157600080fd5b8201836020820111156100b357600080fd5b803590602001918460208302840111600160201b831117156100d457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561012357600080fd5b82018360208201111561013557600080fd5b803590602001918460018302840111600160201b8311171561015657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156101a857600080fd5b8201836020820111156101ba57600080fd5b803590602001918460208302840111600160201b831117156101db57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561022a57600080fd5b82018360208201111561023c57600080fd5b803590602001918460208302840111600160201b8311171561025d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102ac57600080fd5b8201836020820111156102be57600080fd5b803590602001918460208302840111600160201b831117156102df57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506108ef945050505050565b005b81801561032b57600080fd5b5061031d600480360361014081101561034357600080fd5b60408051606081810190925283359392830192916080830191906020840190600390839083908082843760009201919091525091949392602081019250359050600160201b81111561039457600080fd5b8201836020820111156103a657600080fd5b803590602001918460208302840111600160201b831117156103c757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561041657600080fd5b82018360208201111561042857600080fd5b803590602001918460018302840111600160201b8311171561044957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561049b57600080fd5b8201836020820111156104ad57600080fd5b803590602001918460208302840111600160201b831117156104ce57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561051d57600080fd5b82018360208201111561052f57600080fd5b803590602001918460208302840111600160201b8311171561055057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561059f57600080fd5b8201836020820111156105b157600080fd5b803590602001918460208302840111600160201b831117156105d257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561062157600080fd5b82018360208201111561063357600080fd5b803590602001918460018302840111600160201b8311171561065457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c70945050505050565b8180156106a157600080fd5b5061031d60048036036101008110156106b957600080fd5b813591602081013591810190606081016040820135600160201b8111156106df57600080fd5b8201836020820111156106f157600080fd5b803590602001918460208302840111600160201b8311171561071257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561076157600080fd5b82018360208201111561077357600080fd5b803590602001918460208302840111600160201b8311171561079457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156107e357600080fd5b8201836020820111156107f557600080fd5b803590602001918460208302840111600160201b8311171561081657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929567ffffffffffffffff85351695602086013595919450925060608101915060400135600160201b81111561087b57600080fd5b82018360208201111561088d57600080fd5b803590602001918460018302840111600160201b831117156108ae57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610d03945050505050565b60036006890154600160501b900460ff16600381111561090b57fe5b146109475760405162461bcd60e51b81526004018080602001828103825260308152602001806118f76030913960400191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__638ab48be5896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561099657600080fd5b505af41580156109aa573d6000803e3d6000fd5b505050506040513d60208110156109c057600080fd5b5051156109fe5760405162461bcd60e51b815260040180806020018281038252603e8152602001806118b9603e913960400191505060405180910390fd5b8760020154858484898b89876040516020018085815260200184815260200183805190602001908083835b60208310610a485780518252601f199092019160209182019101610a29565b51815160209384036101000a60001901801990921691161790528551919093019285810192500280838360005b83811015610a8d578181015183820152602001610a75565b505050509050019450505050506040516020818303038152906040528051906020012060405160200180858051906020019060200280838360005b83811015610ae0578181015183820152602001610ac8565b50505050905001848051906020019060200280838360005b83811015610b10578181015183820152602001610af8565b50505050905001838051906020019060200280838360005b83811015610b40578181015183820152602001610b28565b505050509050018281526020019450505050506040516020818303038152906040528051906020012014610ba55760405162461bcd60e51b81526004018080602001828103825260348152602001806119c26034913960400191505060405180910390fd5b6003880186905560408051633ad2660b60e21b8152600481018a905260248101899052905173__$8e266570c8a7fb2aaac83b3e040afaf9e1$__9163eb49982c916044808301926000929190829003018186803b158015610c0557600080fd5b505af4158015610c19573d6000803e3d6000fd5b50505050600588015460408051600160c01b90920467ffffffffffffffff168252517fbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2916020908290030190a15050505050505050565b610cf9886040518060a001604052808a600060038110610c8c57fe5b602002015181526020018a600160038110610ca357fe5b602002015181526020018981526020016040518060a001604052808a81526020018981526020018881526020018781526020018c600260038110610ce357fe5b6020020151815250815260200184815250610d49565b5050505050505050565b610cf9886040518060e001604052808a81526020018981526020018881526020018781526020018667ffffffffffffffff168152602001858152602001848152506110ca565b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63e2fe93ca836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610d9857600080fd5b505af4158015610dac573d6000803e3d6000fd5b505050506040513d6020811015610dc257600080fd5b505115610e16576040805162461bcd60e51b815260206004820152601b60248201527f43616e2774206173736572742068616c746564206d616368696e650000000000604482015290519081900360640190fd5b60016006830154600160501b900460ff166003811115610e3257fe5b1480610e57575060026006830154600160501b900460ff166003811115610e5557fe5b145b80610e7b575060036006830154600160501b900460ff166003811115610e7957fe5b145b610eb65760405162461bcd60e51b815260040180806020018281038252602e815260200180611994602e913960400191505060405180910390fd5b6000610fab83836020015184600001518560600151600001518660600151606001516040516020018085815260200184815260200183805190602001908083835b60208310610f165780518252601f199092019160209182019101610ef7565b51815160209384036101000a60001901801990921691161790528551919093019285810192500280838360005b83811015610f5b578181015183820152602001610f43565b505050509050019450505050506040516020818303038152906040528051906020012084604001518560600151602001518660600151604001516000198860600151608001518960800151611516565b905073__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63317bf1ad846040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610ffc57600080fd5b505af4158015611010573d6000803e3d6000fd5b5050506020830151600385015550815160408051633ad2660b60e21b81526004810186905260248101929092525173__$8e266570c8a7fb2aaac83b3e040afaf9e1$__9163eb49982c916044808301926000929190829003018186803b15801561107957600080fd5b505af415801561108d573d6000803e3d6000fd5b50506040805184815290517f709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa9350908190036020019150a1505050565b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63e2fe93ca836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561111957600080fd5b505af415801561112d573d6000803e3d6000fd5b505050506040513d602081101561114357600080fd5b505115611197576040805162461bcd60e51b815260206004820152601b60248201527f43616e2774206173736572742068616c746564206d616368696e650000000000604482015290519081900360640190fd5b60016006830154600160501b900460ff1660038111156111b357fe5b14806111d8575060026006830154600160501b900460ff1660038111156111d657fe5b145b806111fc575060036006830154600160501b900460ff1660038111156111fa57fe5b145b6112375760405162461bcd60e51b815260040180806020018281038252602d8152602001806119f6602d913960400191505060405180910390fd5b600061126583836000015184602001518560400151866060015187608001518860a001518960c00151611516565b905060036006840154600160501b900460ff16600381111561128357fe5b14156112e4576005830154608083015167ffffffffffffffff600160c01b90920482169116116112e45760405162461bcd60e51b81526004018080602001828103825260428152602001806119526042913960600191505060405180910390fd5b73__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63317bf1ad846040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561133357600080fd5b505af4158015611347573d6000803e3d6000fd5b5050505073__$8e266570c8a7fb2aaac83b3e040afaf9e1$__63a3a162cb846040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561139a57600080fd5b505af41580156113ae573d6000803e3d6000fd5b5050506006840180546003925060ff60501b1916600160501b83021790555081608001518360050160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816020015182604001518360600151846000015160405160200180858051906020019060200280838360005b83811015611440578181015183820152602001611428565b50505050905001848051906020019060200280838360005b83811015611470578181015183820152602001611458565b50505050905001838051906020019060200280838360005b838110156114a0578181015183820152602001611488565b5050505091909101928352505060408051808303815260208084018084528251919092012060028a015560808801519087905267ffffffffffffffff1681830152517fc87ab2402746691bbdb443504eab6563fce71e66d5c223f63b0af3442b20851d94509081900360600192509050a1505050565b60008030898b600101548c600301548b8b8b8b60405160200180888152602001878152602001868152602001858051906020019060200280838360005b8381101561156b578181015183820152602001611553565b50505050905001848051906020019060200280838360005b8381101561159b578181015183820152602001611583565b50505050905001838051906020019060200280838360005b838110156115cb5781810151838201526020016115b3565b505050509050018267ffffffffffffffff1667ffffffffffffffff1660c01b8152600801975050505050505050604051602081830303815290604052805190602001208560405160200180846001600160a01b03166001600160a01b031660601b815260140183815260200182815260200193505050506040516020818303038152906040528051906020012090506116658a8285611672565b9998505050505050505050565b6040805163f0c8e96960e01b8152600481018481526024820192835283516044830152835160609373__$59c09a8a68cf3791d03cdf6ed66ba4d742$__9363f0c8e9699388938893919260640190602085019080838360005b838110156116e35781810151838201526020016116cb565b50505050905090810190601f1680156117105780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b15801561172e57600080fd5b505af4158015611742573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561176b57600080fd5b810190808051600160201b81111561178257600080fd5b8201602081018481111561179557600080fd5b81518560208202830111600160201b821117156117b157600080fd5b505080516006890154919550935068010000000000000000900461ffff168314915061181090505760405162461bcd60e51b815260040180806020018281038252602b815260200180611927602b913960400191505060405180910390fd5b60005b818110156118b05785600001600084838151811061182d57fe5b6020908102919091018101516001600160a01b031682528101919091526040016000206001015460ff166118a8576040805162461bcd60e51b815260206004820181905260248201527f5369676e61747572652066726f6d206e6f6e2d76616c696461746f72206b6579604482015290519081900360640190fd5b600101611813565b50505050505056fe43616e206f6e6c7920636f6e6669726d20617373657274696f6e2077686f7365206368616c6c656e676520646561646c696e65206861732070617373656443616e206f6e6c7920636f6e6669726d20696620746865726520697320612070656e64696e6720617373657274696f6e4d7573742068617665206f6e65207369676e61747572652066726f6d20656163682076616c696461746f7243616e206f6e6c79207375706572736564652070726576696f757320617373657274696f6e207769746820677265617465722073657175656e6365206e756d626572547269656420746f2066696e616c697a6520756e616e696d6f75732066726f6d20696e76616c696420737461746543616e206f6e6c7920636f6e6669726d20617373657274696f6e20746861742069732063757272656e746c792070656e64696e67547269656420746f2070656e64696e6720756e616e696d6f75732066726f6d20696e76616c6964207374617465a265627a7a723058200c3127d9fc1c37b9e98d7592b7e2eb289880b68e4645e69e1c06d274b7161cfc64736f6c634300050a0032"

// DeployUnanimous deploys a new Ethereum contract, binding an instance of Unanimous to it.
func DeployUnanimous(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Unanimous, error) {
	parsed, err := abi.JSON(strings.NewReader(UnanimousABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	sigUtilsAddr, _, _, _ := DeploySigUtils(auth, backend)
	UnanimousBin = strings.Replace(UnanimousBin, "__$59c09a8a68cf3791d03cdf6ed66ba4d742$__", sigUtilsAddr.String()[2:], -1)

	vMAddr, _, _, _ := DeployVM(auth, backend)
	UnanimousBin = strings.Replace(UnanimousBin, "__$8e266570c8a7fb2aaac83b3e040afaf9e1$__", vMAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnanimousBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Unanimous{UnanimousCaller: UnanimousCaller{contract: contract}, UnanimousTransactor: UnanimousTransactor{contract: contract}, UnanimousFilterer: UnanimousFilterer{contract: contract}}, nil
}

// Unanimous is an auto generated Go binding around an Ethereum contract.
type Unanimous struct {
	UnanimousCaller     // Read-only binding to the contract
	UnanimousTransactor // Write-only binding to the contract
	UnanimousFilterer   // Log filterer for contract events
}

// UnanimousCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnanimousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnanimousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnanimousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnanimousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnanimousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnanimousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnanimousSession struct {
	Contract     *Unanimous        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnanimousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnanimousCallerSession struct {
	Contract *UnanimousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UnanimousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnanimousTransactorSession struct {
	Contract     *UnanimousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UnanimousRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnanimousRaw struct {
	Contract *Unanimous // Generic contract binding to access the raw methods on
}

// UnanimousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnanimousCallerRaw struct {
	Contract *UnanimousCaller // Generic read-only contract binding to access the raw methods on
}

// UnanimousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnanimousTransactorRaw struct {
	Contract *UnanimousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnanimous creates a new instance of Unanimous, bound to a specific deployed contract.
func NewUnanimous(address common.Address, backend bind.ContractBackend) (*Unanimous, error) {
	contract, err := bindUnanimous(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unanimous{UnanimousCaller: UnanimousCaller{contract: contract}, UnanimousTransactor: UnanimousTransactor{contract: contract}, UnanimousFilterer: UnanimousFilterer{contract: contract}}, nil
}

// NewUnanimousCaller creates a new read-only instance of Unanimous, bound to a specific deployed contract.
func NewUnanimousCaller(address common.Address, caller bind.ContractCaller) (*UnanimousCaller, error) {
	contract, err := bindUnanimous(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnanimousCaller{contract: contract}, nil
}

// NewUnanimousTransactor creates a new write-only instance of Unanimous, bound to a specific deployed contract.
func NewUnanimousTransactor(address common.Address, transactor bind.ContractTransactor) (*UnanimousTransactor, error) {
	contract, err := bindUnanimous(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnanimousTransactor{contract: contract}, nil
}

// NewUnanimousFilterer creates a new log filterer instance of Unanimous, bound to a specific deployed contract.
func NewUnanimousFilterer(address common.Address, filterer bind.ContractFilterer) (*UnanimousFilterer, error) {
	contract, err := bindUnanimous(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnanimousFilterer{contract: contract}, nil
}

// bindUnanimous binds a generic wrapper to an already deployed contract.
func bindUnanimous(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnanimousABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unanimous *UnanimousRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Unanimous.Contract.UnanimousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unanimous *UnanimousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unanimous.Contract.UnanimousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unanimous *UnanimousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unanimous.Contract.UnanimousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unanimous *UnanimousCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Unanimous.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unanimous *UnanimousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unanimous.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unanimous *UnanimousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unanimous.Contract.contract.Transact(opts, method, params...)
}

// UnanimousConfirmedUnanimousAssertionIterator is returned from FilterConfirmedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedUnanimousAssertion events raised by the Unanimous contract.
type UnanimousConfirmedUnanimousAssertionIterator struct {
	Event *UnanimousConfirmedUnanimousAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnanimousConfirmedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnanimousConfirmedUnanimousAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnanimousConfirmedUnanimousAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnanimousConfirmedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnanimousConfirmedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnanimousConfirmedUnanimousAssertion represents a ConfirmedUnanimousAssertion event raised by the Unanimous contract.
type UnanimousConfirmedUnanimousAssertion struct {
	SequenceNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedUnanimousAssertion is a free log retrieval operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_Unanimous *UnanimousFilterer) FilterConfirmedUnanimousAssertion(opts *bind.FilterOpts) (*UnanimousConfirmedUnanimousAssertionIterator, error) {

	logs, sub, err := _Unanimous.contract.FilterLogs(opts, "ConfirmedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &UnanimousConfirmedUnanimousAssertionIterator{contract: _Unanimous.contract, event: "ConfirmedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedUnanimousAssertion is a free log subscription operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_Unanimous *UnanimousFilterer) WatchConfirmedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *UnanimousConfirmedUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _Unanimous.contract.WatchLogs(opts, "ConfirmedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnanimousConfirmedUnanimousAssertion)
				if err := _Unanimous.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmedUnanimousAssertion is a log parse operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_Unanimous *UnanimousFilterer) ParseConfirmedUnanimousAssertion(log types.Log) (*UnanimousConfirmedUnanimousAssertion, error) {
	event := new(UnanimousConfirmedUnanimousAssertion)
	if err := _Unanimous.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UnanimousFinalizedUnanimousAssertionIterator is returned from FilterFinalizedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for FinalizedUnanimousAssertion events raised by the Unanimous contract.
type UnanimousFinalizedUnanimousAssertionIterator struct {
	Event *UnanimousFinalizedUnanimousAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnanimousFinalizedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnanimousFinalizedUnanimousAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnanimousFinalizedUnanimousAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnanimousFinalizedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnanimousFinalizedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnanimousFinalizedUnanimousAssertion represents a FinalizedUnanimousAssertion event raised by the Unanimous contract.
type UnanimousFinalizedUnanimousAssertion struct {
	UnanHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizedUnanimousAssertion is a free log retrieval operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_Unanimous *UnanimousFilterer) FilterFinalizedUnanimousAssertion(opts *bind.FilterOpts) (*UnanimousFinalizedUnanimousAssertionIterator, error) {

	logs, sub, err := _Unanimous.contract.FilterLogs(opts, "FinalizedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &UnanimousFinalizedUnanimousAssertionIterator{contract: _Unanimous.contract, event: "FinalizedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchFinalizedUnanimousAssertion is a free log subscription operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_Unanimous *UnanimousFilterer) WatchFinalizedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *UnanimousFinalizedUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _Unanimous.contract.WatchLogs(opts, "FinalizedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnanimousFinalizedUnanimousAssertion)
				if err := _Unanimous.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinalizedUnanimousAssertion is a log parse operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_Unanimous *UnanimousFilterer) ParseFinalizedUnanimousAssertion(log types.Log) (*UnanimousFinalizedUnanimousAssertion, error) {
	event := new(UnanimousFinalizedUnanimousAssertion)
	if err := _Unanimous.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UnanimousPendingUnanimousAssertionIterator is returned from FilterPendingUnanimousAssertion and is used to iterate over the raw logs and unpacked data for PendingUnanimousAssertion events raised by the Unanimous contract.
type UnanimousPendingUnanimousAssertionIterator struct {
	Event *UnanimousPendingUnanimousAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnanimousPendingUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnanimousPendingUnanimousAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnanimousPendingUnanimousAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnanimousPendingUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnanimousPendingUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnanimousPendingUnanimousAssertion represents a PendingUnanimousAssertion event raised by the Unanimous contract.
type UnanimousPendingUnanimousAssertion struct {
	UnanHash    [32]byte
	SequenceNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingUnanimousAssertion is a free log retrieval operation binding the contract event 0xc87ab2402746691bbdb443504eab6563fce71e66d5c223f63b0af3442b20851d.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum)
func (_Unanimous *UnanimousFilterer) FilterPendingUnanimousAssertion(opts *bind.FilterOpts) (*UnanimousPendingUnanimousAssertionIterator, error) {

	logs, sub, err := _Unanimous.contract.FilterLogs(opts, "PendingUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &UnanimousPendingUnanimousAssertionIterator{contract: _Unanimous.contract, event: "PendingUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingUnanimousAssertion is a free log subscription operation binding the contract event 0xc87ab2402746691bbdb443504eab6563fce71e66d5c223f63b0af3442b20851d.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum)
func (_Unanimous *UnanimousFilterer) WatchPendingUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *UnanimousPendingUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _Unanimous.contract.WatchLogs(opts, "PendingUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnanimousPendingUnanimousAssertion)
				if err := _Unanimous.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingUnanimousAssertion is a log parse operation binding the contract event 0xc87ab2402746691bbdb443504eab6563fce71e66d5c223f63b0af3442b20851d.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum)
func (_Unanimous *UnanimousFilterer) ParsePendingUnanimousAssertion(log types.Log) (*UnanimousPendingUnanimousAssertion, error) {
	event := new(UnanimousPendingUnanimousAssertion)
	if err := _Unanimous.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMABI is the input ABI used to generate the binding from.
const VMABI = "[]"

// VMFuncSigs maps the 4-byte function signature to its string representation.
var VMFuncSigs = map[string]string{
	"eb49982c": "acceptAssertion(VM.Data storage,bytes32)",
	"317bf1ad": "cancelCurrentState(VM.Data storage)",
	"2a3e0a97": "isErrored(VM.Data storage)",
	"e2fe93ca": "isHalted(VM.Data storage)",
	"a3a162cb": "resetDeadline(VM.Data storage)",
	"8ab48be5": "withinDeadline(VM.Data storage)",
}

// VMBin is the compiled bytecode used for deploying new contracts.
var VMBin = "0x610390610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c80632a3e0a9714610071578063317bf1ad146100a25780638ab48be5146100ce578063a3a162cb146100eb578063e2fe93ca14610115578063eb49982c14610132575b600080fd5b61008e6004803603602081101561008757600080fd5b5035610162565b604080519115158252519081900360200190f35b8180156100ae57600080fd5b506100cc600480360360208110156100c557600080fd5b503561016c565b005b61008e600480360360208110156100e457600080fd5b503561027d565b8180156100f757600080fd5b506100cc6004803603602081101561010e57600080fd5b5035610298565b61008e6004803603602081101561012b57600080fd5b50356102d7565b81801561013e57600080fd5b506100cc6004803603604081101561015557600080fd5b50803590602001356102df565b6001908101541490565b60016006820154600160501b900460ff16600381111561018857fe5b146101f7576005810154600160801b900467ffffffffffffffff164311156101f7576040805162461bcd60e51b815260206004820152601c60248201527f43616e27742063616e63656c2066696e616c697a656420737461746500000000604482015290519081900360640190fd5b60026006820154600160501b900460ff16600381111561021357fe5b141561027a57600581015460048201546001600160a01b031660009081526020839052604090205461025c916fffffffffffffffffffffffffffffffff1663ffffffff6102fa16565b60048201546001600160a01b03166000908152602083905260409020555b50565b60050154600160801b900467ffffffffffffffff1643111590565b60068101546005909101805467ffffffffffffffff60801b1916600160801b63ffffffff909316430167ffffffffffffffff1692909202919091179055565b600101541590565b6001820155600601805460ff60501b1916600160501b179055565b600082820183811015610354576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fea265627a7a72305820e28b39e4c6fbe9445d33c8378791ef436843683d976fddb1ece89d211fb78c4a64736f6c634300050a0032"

// DeployVM deploys a new Ethereum contract, binding an instance of VM to it.
func DeployVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VM, error) {
	parsed, err := abi.JSON(strings.NewReader(VMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VM{VMCaller: VMCaller{contract: contract}, VMTransactor: VMTransactor{contract: contract}, VMFilterer: VMFilterer{contract: contract}}, nil
}

// VM is an auto generated Go binding around an Ethereum contract.
type VM struct {
	VMCaller     // Read-only binding to the contract
	VMTransactor // Write-only binding to the contract
	VMFilterer   // Log filterer for contract events
}

// VMCaller is an auto generated read-only Go binding around an Ethereum contract.
type VMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VMSession struct {
	Contract     *VM               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VMCallerSession struct {
	Contract *VMCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VMTransactorSession struct {
	Contract     *VMTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMRaw is an auto generated low-level Go binding around an Ethereum contract.
type VMRaw struct {
	Contract *VM // Generic contract binding to access the raw methods on
}

// VMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VMCallerRaw struct {
	Contract *VMCaller // Generic read-only contract binding to access the raw methods on
}

// VMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VMTransactorRaw struct {
	Contract *VMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVM creates a new instance of VM, bound to a specific deployed contract.
func NewVM(address common.Address, backend bind.ContractBackend) (*VM, error) {
	contract, err := bindVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VM{VMCaller: VMCaller{contract: contract}, VMTransactor: VMTransactor{contract: contract}, VMFilterer: VMFilterer{contract: contract}}, nil
}

// NewVMCaller creates a new read-only instance of VM, bound to a specific deployed contract.
func NewVMCaller(address common.Address, caller bind.ContractCaller) (*VMCaller, error) {
	contract, err := bindVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VMCaller{contract: contract}, nil
}

// NewVMTransactor creates a new write-only instance of VM, bound to a specific deployed contract.
func NewVMTransactor(address common.Address, transactor bind.ContractTransactor) (*VMTransactor, error) {
	contract, err := bindVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VMTransactor{contract: contract}, nil
}

// NewVMFilterer creates a new log filterer instance of VM, bound to a specific deployed contract.
func NewVMFilterer(address common.Address, filterer bind.ContractFilterer) (*VMFilterer, error) {
	contract, err := bindVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VMFilterer{contract: contract}, nil
}

// bindVM binds a generic wrapper to an already deployed contract.
func bindVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VM *VMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VM.Contract.VMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VM *VMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VM.Contract.VMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VM *VMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VM.Contract.VMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VM *VMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VM *VMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VM *VMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VM.Contract.contract.Transact(opts, method, params...)
}

// VMCreatorABI is the input ABI used to generate the binding from.
const VMCreatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"name\":\"_escrowCurrency\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_validatorKeys\",\"type\":\"address[]\"}],\"name\":\"launchVM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"name\":\"_challengeManagerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"escrowCurrency\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"vmState\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"VMCreated\",\"type\":\"event\"}]"

// VMCreatorFuncSigs maps the 4-byte function signature to its string representation.
var VMCreatorFuncSigs = map[string]string{
	"34d1efe9": "launchVM(bytes32,uint32,uint32,uint128,address,address,address[])",
}

// VMCreatorBin is the compiled bytecode used for deploying new contracts.
var VMCreatorBin = "0x608060405234801561001057600080fd5b50604051613c09380380613c098339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b03199182161790915560018054939092169216919091179055613b8f8061007a6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806334d1efe914610030575b600080fd5b610117600480360360e081101561004657600080fd5b81359163ffffffff60208201358116926040830135909116916001600160801b03606082013516916001600160a01b03608083013581169260a08101359091169181019060e0810160c08201356401000000008111156100a557600080fd5b8201836020820111156100b757600080fd5b803590602001918460208302840111640100000000831117156100d957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610119945050505050565b005b6001546000805460405191928a928a928a928a928a928a926001600160a01b03908116929116908a9061014b90610327565b89815263ffffffff808a1660208084019190915290891660408301526001600160801b03881660608301526001600160a01b03808816608084015286811660a084015285811660c0840152841660e0830152610120610100830181815284519184019190915283519091610140840191858201910280838360005b838110156101de5781810151838201526020016101c6565b505050509050019a5050505050505050505050604051809103906000f08015801561020d573d6000803e3d6000fd5b5090507fb2f02f344f7e007eb90f5ca565eac18aaaaa406a00105e0647cfa902a6978bd7818887878a8d898960405180896001600160a01b03166001600160a01b031681526020018863ffffffff1663ffffffff168152602001876001600160801b03166001600160801b03168152602001866001600160a01b03166001600160a01b031681526020018563ffffffff1663ffffffff168152602001848152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156103035781810151838201526020016102eb565b50505050905001995050505050505050505060405180910390a15050505050505050565b613826806103358339019056fe60806040523480156200001157600080fd5b50604051620038263803806200382683398181016040526101208110156200003857600080fd5b8151602083015160408401516060850151608086015160a087015160c088015160e08901516101008a018051989a97999698959794969395929491938201926401000000008111156200008a57600080fd5b820160208101848111156200009e57600080fd5b8151856020820283011164010000000082111715620000bc57600080fd5b509093505050506001600160a01b0385161562000125576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180620038056021913960400191505060405180910390fd5b600180546001600160a01b038085166001600160a01b03199283161792839055600080548783169316929092178255604080517ff39723830000000000000000000000000000000000000000000000000000000081529051939091169263f39723839260048084019391929182900301818387803b158015620001a757600080fd5b505af1158015620001bc573d6000803e3d6000fd5b5050600c80546001600160a01b0319166001600160a01b0388811691909117909155600980546001600160b01b031916918916919091179055505080516008805461ffff60401b19166801000000000000000061ffff84160217905560005b8161ffff168161ffff161015620002a3576040518060400160405280600081526020016001151581525060026000016000858461ffff16815181106200025d57fe5b6020908102919091018101516001600160a01b0316825281810192909252604001600020825181559101516001918201805460ff1916911515919091179055016200021b565b5060038a90556008805460ff60501b19169055604080517f364df277000000000000000000000000000000000000000000000000000000008152905173__$d969135829891f807aa9c34494da4ecd99$__9163364df277916004808301926020929190829003018186803b1580156200031b57600080fd5b505af415801562000330573d6000803e3d6000fd5b505050506040513d60208110156200034757600080fd5b50516005555050600780546001600160801b0319166001600160801b03969096169590951790945550506008805463ffffffff191663ffffffff9586161763ffffffff60201b19166401000000009490951693909302939093179091555061344c9150819050620003b96000396000f3fe6080604052600436106101355760003560e01c80636be00229116100ab578063aca0f3721161006f578063aca0f37214610f3f578063b99738e014610f54578063bab7237714610f87578063cfa80707146112da578063d489113a146112ef578063fdf4a2611461130457610135565b80636be0022914610900578063899b7c74146109155780638da5cb5b146109415780638fe08c2b1461095657806397e2e25614610c1b57610135565b80632782e87e116100fd5780632782e87e146102205780633a7684631461024a5780634526c5d914610300578063513164fe146105d557806360675a8714610697578063675125b9146106ac57610135565b8063023a96fe1461013a57806305b050de1461016b57806308dc89d7146101755780631865c57d146101ba57806322c091bc146101f3575b600080fd5b34801561014657600080fd5b5061014f611319565b604080516001600160a01b039092168252519081900360200190f35b610173611328565b005b34801561018157600080fd5b506101a86004803603602081101561019857600080fd5b50356001600160a01b0316611453565b60408051918252519081900360200190f35b3480156101c657600080fd5b506101cf611472565b604051808260038111156101df57fe5b60ff16815260200191505060405180910390f35b3480156101ff57600080fd5b506101736004803603608081101561021657600080fd5b5060408101611482565b34801561022c57600080fd5b506101736004803603602081101561024357600080fd5b50356115d2565b34801561025657600080fd5b5061025f6117bd565b604080518d8152602081018d90529081018b90526001600160a01b038a1660608201526001600160801b038916608082015267ffffffffffffffff80891660a0830152871660c082015263ffffffff80871660e0830152851661010082015261ffff841661012082015261014081018360038111156102da57fe5b60ff1681529115156020830152506040805191829003019b509950505050505050505050f35b34801561030c57600080fd5b50610173600480360361012081101561032457600080fd5b81359160208101359163ffffffff6040830135169190810190608081016060820135600160201b81111561035757600080fd5b82018360208201111561036957600080fd5b803590602001918460208302840111600160201b8311171561038a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103d957600080fd5b8201836020820111156103eb57600080fd5b803590602001918460018302840111600160201b8311171561040c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561045e57600080fd5b82018360208201111561047057600080fd5b803590602001918460208302840111600160201b8311171561049157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104e057600080fd5b8201836020820111156104f257600080fd5b803590602001918460208302840111600160201b8311171561051357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561056257600080fd5b82018360208201111561057457600080fd5b803590602001918460208302840111600160201b8311171561059557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550509135925061183a915050565b3480156105e157600080fd5b50610683600480360360208110156105f857600080fd5b810190602081018135600160201b81111561061257600080fd5b82018360208201111561062457600080fd5b803590602001918460208302840111600160201b8311171561064557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611c48945050505050565b604080519115158252519081900360200190f35b3480156106a357600080fd5b5061014f611cd5565b3480156106b857600080fd5b50610173600480360360e08110156106cf57600080fd5b81359190810190604081016020820135600160201b8111156106f057600080fd5b82018360208201111561070257600080fd5b803590602001918460208302840111600160201b8311171561072357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561077257600080fd5b82018360208201111561078457600080fd5b803590602001918460208302840111600160201b831117156107a557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156107f457600080fd5b82018360208201111561080657600080fd5b803590602001918460208302840111600160201b8311171561082757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929567ffffffffffffffff85351695602086013595919450925060608101915060400135600160201b81111561088c57600080fd5b82018360208201111561089e57600080fd5b803590602001918460018302840111600160201b831117156108bf57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611ce4945050505050565b34801561090c57600080fd5b5061014f6122f5565b34801561092157600080fd5b5061092a612304565b6040805161ffff9092168252519081900360200190f35b34801561094d57600080fd5b5061014f612315565b34801561096257600080fd5b50610173600480360360e081101561097957600080fd5b813591602081013591810190606081016040820135600160201b81111561099f57600080fd5b8201836020820111156109b157600080fd5b803590602001918460208302840111600160201b831117156109d257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a2157600080fd5b820183602082011115610a3357600080fd5b803590602001918460018302840111600160201b83111715610a5457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610aa657600080fd5b820183602082011115610ab857600080fd5b803590602001918460208302840111600160201b83111715610ad957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b2857600080fd5b820183602082011115610b3a57600080fd5b803590602001918460208302840111600160201b83111715610b5b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610baa57600080fd5b820183602082011115610bbc57600080fd5b803590602001918460208302840111600160201b83111715610bdd57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550612324945050505050565b348015610c2757600080fd5b506101736004803603610180811015610c3f57600080fd5b810190808060800190600480602002604051908101604052809291908260046020028082843760009201919091525050604080518082018252929563ffffffff853516959094909360608201935091602090910190600290839083908082843760009201919091525091949392602081019250359050600160201b811115610cc657600080fd5b820183602082011115610cd857600080fd5b803590602001918460208302840111600160201b83111715610cf957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d4857600080fd5b820183602082011115610d5a57600080fd5b803590602001918460208302840111600160201b83111715610d7b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610dca57600080fd5b820183602082011115610ddc57600080fd5b803590602001918460208302840111600160201b83111715610dfd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610e4c57600080fd5b820183602082011115610e5e57600080fd5b803590602001918460208302840111600160201b83111715610e7f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610ece57600080fd5b820183602082011115610ee057600080fd5b803590602001918460208302840111600160201b83111715610f0157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550612714945050505050565b348015610f4b57600080fd5b506101a8612dd5565b348015610f6057600080fd5b5061068360048036036020811015610f7757600080fd5b50356001600160a01b0316612de4565b348015610f9357600080fd5b506101736004803603610120811015610fab57600080fd5b813591602081013591810190606081016040820135600160201b811115610fd157600080fd5b820183602082011115610fe357600080fd5b803590602001918460208302840111600160201b8311171561100457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561105357600080fd5b82018360208201111561106557600080fd5b803590602001918460018302840111600160201b8311171561108657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156110d857600080fd5b8201836020820111156110ea57600080fd5b803590602001918460208302840111600160201b8311171561110b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561115a57600080fd5b82018360208201111561116c57600080fd5b803590602001918460208302840111600160201b8311171561118d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156111dc57600080fd5b8201836020820111156111ee57600080fd5b803590602001918460208302840111600160201b8311171561120f57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092958435959094909350604081019250602001359050600160201b81111561126657600080fd5b82018360208201111561127857600080fd5b803590602001918460018302840111600160201b8311171561129957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612e05945050505050565b3480156112e657600080fd5b50610173613122565b3480156112fb57600080fd5b5061014f61318b565b34801561131057600080fd5b5061014f61319a565b6000546001600160a01b031681565b3360009081526002602052604090206001015460ff1661138a576040805162461bcd60e51b815260206004820152601860248201527721b0b63632b91036bab9ba103132903b30b634b230ba37b960411b604482015290519081900360640190fd5b336000908152600260205260409020600754815434810183556001600160801b03909116118080156113ca575060075482546001600160801b0390911611155b156113f65760098054600161ffff600160a01b808404821692909201160261ffff60a01b199091161790555b600854600954600160a01b900461ffff908116600160401b9092041614801561143657506000600854600160501b900460ff16600381111561143457fe5b145b1561144f576008805460ff60501b1916600160501b1790555b5050565b6001600160a01b0381166000908152600260205260409020545b919050565b600854600160501b900460ff1690565b6000546001600160a01b031633146114cb5760405162461bcd60e51b815260040180806020018281038252602d8152602001806133c7602d913960400191505060405180910390fd5b600854600160581b900460ff166115135760405162461bcd60e51b81526004018080602001828103825260268152602001806133a16026913960400191505060405180910390fd5b6008805460ff60581b191690556115756001600160801b038235166002600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b03168152602001908152602001600020600001546131a990919063ffffffff16565b82356001600160a01b031660009081526002602081815260408320939093556115ad928401356001600160801b031691856001611535565b6001600160a01b03602093840135166000908152600290935260409092209190915550565b3360009081526002602052604090206001015460ff16611634576040805162461bcd60e51b815260206004820152601860248201527721b0b63632b91036bab9ba103132903b30b634b230ba37b960411b604482015290519081900360640190fd5b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__635af93c7a6002836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561168c57600080fd5b505af41580156116a0573d6000803e3d6000fd5b5050600080546040805180820182526006546001600160a01b03908116825233602080840191909152835180850185526007546001600160801b0316808252918101919091526008548451630823813560e21b815292909516975063208e04d496509194919363ffffffff16928892600490920191829187918190849084905b83811015611738578181015183820152602001611720565b5050505090500184600260200280838360005b8381101561176357818101518382015260200161174b565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050600060405180830381600087803b1580156117a257600080fd5b505af11580156117b6573d6000803e3d6000fd5b5050505050565b6003546004546005546006546007546008546001600160a01b03909216916001600160801b0382169167ffffffffffffffff600160801b8204811692600160c01b909204169063ffffffff80821691600160201b81049091169061ffff600160401b8204169060ff600160501b8204811691600160581b9004168c565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63924e7b3760028b8b8b8b8b8b8b8b8b6040518b63ffffffff1660e01b8152600401808b81526020018a81526020018981526020018863ffffffff1663ffffffff168152602001806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b838110156118e85781810151838201526020016118d0565b5050505090500186810385528b818151815260200191508051906020019080838360005b8381101561192457818101518382015260200161190c565b50505050905090810190601f1680156119515780820380516001836020036101000a031916815260200191505b5086810384528a5181528a51602091820191808d01910280838360005b8381101561198657818101518382015260200161196e565b50505050905001868103835289818151815260200191508051906020019060200280838360005b838110156119c55781810151838201526020016119ad565b50505050905001868103825288818151815260200191508051906020019060200280838360005b83811015611a045781810151838201526020016119ec565b505050509050019f5050505050505050505050505050505060006040518083038186803b158015611a3457600080fd5b505af4158015611a48573d6000803e3d6000fd5b50505050611a5461320a565b600160009054906101000a90046001600160a01b03166001600160a01b031663ec22a76787878787876040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015611adc578181015183820152602001611ac4565b5050505090500186810385528a818151815260200191508051906020019080838360005b83811015611b18578181015183820152602001611b00565b50505050905090810190601f168015611b455780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015611b7a578181015183820152602001611b62565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015611bb9578181015183820152602001611ba1565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015611bf8578181015183820152602001611be0565b505050509050019a5050505050505050505050600060405180830381600087803b158015611c2557600080fd5b505af1158015611c39573d6000803e3d6000fd5b50505050505050505050505050565b805160085460009190600160401b900461ffff168114611c6c57600091505061146d565b60005b81811015611ccb5760026000016000858381518110611c8a57fe5b6020908102919091018101516001600160a01b031682528101919091526040016000206001015460ff16611cc35760009250505061146d565b600101611c6f565b5060019392505050565b600b546001600160a01b031681565b606073__$9836fa7140e5a33041d4b827682e675a30$__630f89fbff8888886040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015611d5a578181015183820152602001611d42565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015611d99578181015183820152602001611d81565b50505050905001848103825285818151815260200191508051906020019060200280838360005b83811015611dd8578181015183820152602001611dc0565b50505050905001965050505050505060006040518083038186803b158015611dff57600080fd5b505af4158015611e13573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015611e3c57600080fd5b810190808051600160201b811115611e5357600080fd5b82016020810184811115611e6657600080fd5b81518560208202830111600160201b82111715611e8257600080fd5b50506040805163578bec9160e11b8152600481019182528c5160448201528c5192965073__$9836fa7140e5a33041d4b827682e675a30$__955063af17d92294508c93508692829160248101916064909101906020808801910280838360005b83811015611efa578181015183820152602001611ee2565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611f39578181015183820152602001611f21565b5050505090500194505050505060206040518083038186803b158015611f5e57600080fd5b505af4158015611f72573d6000803e3d6000fd5b505050506040513d6020811015611f8857600080fd5b5051611fc55760405162461bcd60e51b81526004018080602001828103825260248152602001806133f46024913960400191505060405180910390fd5b60015460405163565b19db60e11b815230600482018181526060602484019081528b5160648501528b516001600160a01b039095169463acb633b6948d9388939092909160448101916084909101906020808801910280838360005b83811015612039578181015183820152602001612021565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015612078578181015183820152602001612060565b505050509050019550505050505060206040518083038186803b15801561209e57600080fd5b505afa1580156120b2573d6000803e3d6000fd5b505050506040513d60208110156120c857600080fd5b505161211b576040805162461bcd60e51b815260206004820152601b60248201527f564d2068617320696e73756666696369656e742062616c616e63650000000000604482015290519081900360640190fd5b73__$caf066876633ea418098495f1e5bb4c2f5$__63a2ee90f460028a8a8a8a8a8a8a6040518963ffffffff1660e01b8152600401808981526020018881526020018060200180602001806020018767ffffffffffffffff1667ffffffffffffffff1681526020018681526020018060200185810385528b818151815260200191508051906020019060200280838360005b838110156121c55781810151838201526020016121ad565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b838110156122045781810151838201526020016121ec565b50505050905001858103835289818151815260200191508051906020019060200280838360005b8381101561224357818101518382015260200161222b565b50505050905001858103825286818151815260200191508051906020019080838360005b8381101561227f578181015183820152602001612267565b50505050905090810190601f1680156122ac5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060006040518083038186803b1580156122d357600080fd5b505af41580156122e7573d6000803e3d6000fd5b505050505050505050505050565b600a546001600160a01b031681565b600954600160a01b900461ffff1681565b600c546001600160a01b031681565b73__$caf066876633ea418098495f1e5bb4c2f5$__63319b5f676002898989898989896040518963ffffffff1660e01b815260040180898152602001888152602001878152602001806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b838110156123b85781810151838201526020016123a0565b5050505090500186810385528a818151815260200191508051906020019080838360005b838110156123f45781810151838201526020016123dc565b50505050905090810190601f1680156124215780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b8381101561245657818101518382015260200161243e565b50505050905001868103835288818151815260200191508051906020019060200280838360005b8381101561249557818101518382015260200161247d565b50505050905001868103825287818151815260200191508051906020019060200280838360005b838110156124d45781810151838201526020016124bc565b505050509050019d505050505050505050505050505060006040518083038186803b15801561250257600080fd5b505af4158015612516573d6000803e3d6000fd5b5050505061252261320a565b600160009054906101000a90046001600160a01b03166001600160a01b031663ec22a76786868686866040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b838110156125aa578181015183820152602001612592565b5050505090500186810385528a818151815260200191508051906020019080838360005b838110156125e65781810151838201526020016125ce565b50505050905090810190601f1680156126135780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015612648578181015183820152602001612630565b50505050905001868103835288818151815260200191508051906020019060200280838360005b8381101561268757818101518382015260200161266f565b50505050905001868103825287818151815260200191508051906020019060200280838360005b838110156126c65781810151838201526020016126ae565b505050509050019a5050505050505050505050600060405180830381600087803b1580156126f357600080fd5b505af1158015612707573d6000803e3d6000fd5b5050505050505050505050565b3360009081526002602052604090206001015460ff16612776576040805162461bcd60e51b815260206004820152601860248201527721b0b63632b91036bab9ba103132903b30b634b230ba37b960411b604482015290519081900360640190fd5b606073__$9836fa7140e5a33041d4b827682e675a30$__630f89fbff8786866040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b838110156127ec5781810151838201526020016127d4565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561282b578181015183820152602001612813565b50505050905001848103825285818151815260200191508051906020019060200280838360005b8381101561286a578181015183820152602001612852565b50505050905001965050505050505060006040518083038186803b15801561289157600080fd5b505af41580156128a5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156128ce57600080fd5b810190808051600160201b8111156128e557600080fd5b820160208101848111156128f857600080fd5b81518560208202830111600160201b8211171561291457600080fd5b50506040805163578bec9160e11b8152600481019182528b5160448201528b5192965073__$9836fa7140e5a33041d4b827682e675a30$__955063af17d92294508b93508692829160248101916064909101906020808801910280838360005b8381101561298c578181015183820152602001612974565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156129cb5781810151838201526020016129b3565b5050505090500194505050505060206040518083038186803b1580156129f057600080fd5b505af4158015612a04573d6000803e3d6000fd5b505050506040513d6020811015612a1a57600080fd5b5051612a575760405162461bcd60e51b81526004018080602001828103825260248152602001806133f46024913960400191505060405180910390fd5b60015460405163565b19db60e11b815230600482018181526060602484019081528a5160648501528a516001600160a01b039095169463acb633b6948c9388939092909160448101916084909101906020808801910280838360005b83811015612acb578181015183820152602001612ab3565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015612b0a578181015183820152602001612af2565b505050509050019550505050505060206040518083038186803b158015612b3057600080fd5b505afa158015612b44573d6000803e3d6000fd5b505050506040513d6020811015612b5a57600080fd5b5051612bad576040805162461bcd60e51b815260206004820152601b60248201527f564d2068617320696e73756666696369656e742062616c616e63650000000000604482015290519081900360640190fd5b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63c97c8eec60028b8b8b8b8b8b8b8b6040518a63ffffffff1660e01b8152600401808a815260200189600460200280838360005b83811015612c0e578181015183820152602001612bf6565b5050505063ffffffff8b1692019182525060200187604080838360005b83811015612c43578181015183820152602001612c2b565b50505050905001806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015612c96578181015183820152602001612c7e565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b83811015612cd5578181015183820152602001612cbd565b50505050905001868103845289818151815260200191508051906020019060200280838360005b83811015612d14578181015183820152602001612cfc565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015612d53578181015183820152602001612d3b565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015612d92578181015183820152602001612d7a565b505050509050019e50505050505050505050505050505060006040518083038186803b158015612dc157600080fd5b505af4158015611c39573d6000803e3d6000fd5b6007546001600160801b031690565b6001600160a01b031660009081526002602052604090206001015460ff1690565b73__$caf066876633ea418098495f1e5bb4c2f5$__633c0ec6de600260405180606001604052808d81526020018c8152602001868152508a8a8a8a8a896040518963ffffffff1660e01b81526004018089815260200188600360200280838360005b83811015612e7f578181015183820152602001612e67565b5050505090500180602001806020018060200180602001806020018060200187810387528d818151815260200191508051906020019060200280838360005b83811015612ed6578181015183820152602001612ebe565b5050505090500187810386528c818151815260200191508051906020019080838360005b83811015612f12578181015183820152602001612efa565b50505050905090810190601f168015612f3f5780820380516001836020036101000a031916815260200191505b5087810385528b5181528b51602091820191808e01910280838360005b83811015612f74578181015183820152602001612f5c565b5050505090500187810384528a818151815260200191508051906020019060200280838360005b83811015612fb3578181015183820152602001612f9b565b50505050905001878103835289818151815260200191508051906020019060200280838360005b83811015612ff2578181015183820152602001612fda565b50505050905001878103825288818151815260200191508051906020019080838360005b8381101561302e578181015183820152602001613016565b50505050905090810190601f16801561305b5780820380516001836020036101000a031916815260200191505b509e50505050505050505050505050505060006040518083038186803b15801561308457600080fd5b505af4158015613098573d6000803e3d6000fd5b505050506130a461320a565b60015460405163ec22a76760e01b815260a060048201908152895160a483015289516001600160a01b039093169263ec22a767928b928b928b928b928b9290918291602482019160448101916064820191608481019160c4909101906020808e01910280838360008315611adc578181015183820152602001611ac4565b600c546001600160a01b03163314613181576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6c79206f776e65722063616e2073687574646f776e2074686520564d0000604482015290519081900360640190fd5b613189613392565b565b6001546001600160a01b031681565b6009546001600160a01b031681565b600082820183811015613203576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600154604080516331618a0360e01b815230600482015290516000926001600160a01b0316916331618a0391602480830192602092919082900301818787803b15801561325657600080fd5b505af115801561326a573d6000803e3d6000fd5b505050506040513d602081101561328057600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b1580156132d157600080fd5b505af41580156132e5573d6000803e3d6000fd5b505050506040513d60208110156132fb57600080fd5b5051811461338f5773__$9836fa7140e5a33041d4b827682e675a30$__63f11fcc26600260030154836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b15801561335f57600080fd5b505af4158015613373573d6000803e3d6000fd5b505050506040513d602081101561338957600080fd5b50516005555b50565b600c546001600160a01b0316fffe564d206d75737420626520696e206368616c6c656e676520746f20636f6d706c6574652069744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e6765546f6b656e207479706573206d7573742062652076616c696420616e6420736f72746564a265627a7a7230582072f27c35176282cc55c31b69184c9daa06284f5cbe091e0506d3515f45933f0764736f6c634300050a003256616c696461746f72206465706f73697473206d75737420626520696e20455448a265627a7a72305820e475cc82424623db90bff099fc66aeb198d98292e9ae96b58065e489440ac21864736f6c634300050a0032"

// DeployVMCreator deploys a new Ethereum contract, binding an instance of VMCreator to it.
func DeployVMCreator(auth *bind.TransactOpts, backend bind.ContractBackend, _globalInboxAddress common.Address, _challengeManagerAddress common.Address) (common.Address, *types.Transaction, *VMCreator, error) {
	parsed, err := abi.JSON(strings.NewReader(VMCreatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	VMCreatorBin = strings.Replace(VMCreatorBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	VMCreatorBin = strings.Replace(VMCreatorBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	unanimousAddr, _, _, _ := DeployUnanimous(auth, backend)
	VMCreatorBin = strings.Replace(VMCreatorBin, "__$caf066876633ea418098495f1e5bb4c2f5$__", unanimousAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	VMCreatorBin = strings.Replace(VMCreatorBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VMCreatorBin), backend, _globalInboxAddress, _challengeManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VMCreator{VMCreatorCaller: VMCreatorCaller{contract: contract}, VMCreatorTransactor: VMCreatorTransactor{contract: contract}, VMCreatorFilterer: VMCreatorFilterer{contract: contract}}, nil
}

// VMCreator is an auto generated Go binding around an Ethereum contract.
type VMCreator struct {
	VMCreatorCaller     // Read-only binding to the contract
	VMCreatorTransactor // Write-only binding to the contract
	VMCreatorFilterer   // Log filterer for contract events
}

// VMCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type VMCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VMCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VMCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VMCreatorSession struct {
	Contract     *VMCreator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VMCreatorCallerSession struct {
	Contract *VMCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VMCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VMCreatorTransactorSession struct {
	Contract     *VMCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VMCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type VMCreatorRaw struct {
	Contract *VMCreator // Generic contract binding to access the raw methods on
}

// VMCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VMCreatorCallerRaw struct {
	Contract *VMCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// VMCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VMCreatorTransactorRaw struct {
	Contract *VMCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVMCreator creates a new instance of VMCreator, bound to a specific deployed contract.
func NewVMCreator(address common.Address, backend bind.ContractBackend) (*VMCreator, error) {
	contract, err := bindVMCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VMCreator{VMCreatorCaller: VMCreatorCaller{contract: contract}, VMCreatorTransactor: VMCreatorTransactor{contract: contract}, VMCreatorFilterer: VMCreatorFilterer{contract: contract}}, nil
}

// NewVMCreatorCaller creates a new read-only instance of VMCreator, bound to a specific deployed contract.
func NewVMCreatorCaller(address common.Address, caller bind.ContractCaller) (*VMCreatorCaller, error) {
	contract, err := bindVMCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VMCreatorCaller{contract: contract}, nil
}

// NewVMCreatorTransactor creates a new write-only instance of VMCreator, bound to a specific deployed contract.
func NewVMCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*VMCreatorTransactor, error) {
	contract, err := bindVMCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VMCreatorTransactor{contract: contract}, nil
}

// NewVMCreatorFilterer creates a new log filterer instance of VMCreator, bound to a specific deployed contract.
func NewVMCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*VMCreatorFilterer, error) {
	contract, err := bindVMCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VMCreatorFilterer{contract: contract}, nil
}

// bindVMCreator binds a generic wrapper to an already deployed contract.
func bindVMCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VMCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VMCreator *VMCreatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VMCreator.Contract.VMCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VMCreator *VMCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMCreator.Contract.VMCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VMCreator *VMCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VMCreator.Contract.VMCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VMCreator *VMCreatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VMCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VMCreator *VMCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VMCreator *VMCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VMCreator.Contract.contract.Transact(opts, method, params...)
}

// LaunchVM is a paid mutator transaction binding the contract method 0x34d1efe9.
//
// Solidity: function launchVM(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _escrowCurrency, address _owner, address[] _validatorKeys) returns()
func (_VMCreator *VMCreatorTransactor) LaunchVM(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _escrowCurrency common.Address, _owner common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _VMCreator.contract.Transact(opts, "launchVM", _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _escrowCurrency, _owner, _validatorKeys)
}

// LaunchVM is a paid mutator transaction binding the contract method 0x34d1efe9.
//
// Solidity: function launchVM(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _escrowCurrency, address _owner, address[] _validatorKeys) returns()
func (_VMCreator *VMCreatorSession) LaunchVM(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _escrowCurrency common.Address, _owner common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _VMCreator.Contract.LaunchVM(&_VMCreator.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _escrowCurrency, _owner, _validatorKeys)
}

// LaunchVM is a paid mutator transaction binding the contract method 0x34d1efe9.
//
// Solidity: function launchVM(bytes32 _vmState, uint32 _gracePeriod, uint32 _maxExecutionSteps, uint128 _escrowRequired, address _escrowCurrency, address _owner, address[] _validatorKeys) returns()
func (_VMCreator *VMCreatorTransactorSession) LaunchVM(_vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _escrowCurrency common.Address, _owner common.Address, _validatorKeys []common.Address) (*types.Transaction, error) {
	return _VMCreator.Contract.LaunchVM(&_VMCreator.TransactOpts, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _escrowCurrency, _owner, _validatorKeys)
}

// VMCreatorVMCreatedIterator is returned from FilterVMCreated and is used to iterate over the raw logs and unpacked data for VMCreated events raised by the VMCreator contract.
type VMCreatorVMCreatedIterator struct {
	Event *VMCreatorVMCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VMCreatorVMCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMCreatorVMCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VMCreatorVMCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VMCreatorVMCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMCreatorVMCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMCreatorVMCreated represents a VMCreated event raised by the VMCreator contract.
type VMCreatorVMCreated struct {
	VmAddress         common.Address
	GracePeriod       uint32
	EscrowRequired    *big.Int
	EscrowCurrency    common.Address
	MaxExecutionSteps uint32
	VmState           [32]byte
	Owner             common.Address
	Validators        []common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVMCreated is a free log retrieval operation binding the contract event 0xb2f02f344f7e007eb90f5ca565eac18aaaaa406a00105e0647cfa902a6978bd7.
//
// Solidity: event VMCreated(address vmAddress, uint32 gracePeriod, uint128 escrowRequired, address escrowCurrency, uint32 maxExecutionSteps, bytes32 vmState, address owner, address[] validators)
func (_VMCreator *VMCreatorFilterer) FilterVMCreated(opts *bind.FilterOpts) (*VMCreatorVMCreatedIterator, error) {

	logs, sub, err := _VMCreator.contract.FilterLogs(opts, "VMCreated")
	if err != nil {
		return nil, err
	}
	return &VMCreatorVMCreatedIterator{contract: _VMCreator.contract, event: "VMCreated", logs: logs, sub: sub}, nil
}

// WatchVMCreated is a free log subscription operation binding the contract event 0xb2f02f344f7e007eb90f5ca565eac18aaaaa406a00105e0647cfa902a6978bd7.
//
// Solidity: event VMCreated(address vmAddress, uint32 gracePeriod, uint128 escrowRequired, address escrowCurrency, uint32 maxExecutionSteps, bytes32 vmState, address owner, address[] validators)
func (_VMCreator *VMCreatorFilterer) WatchVMCreated(opts *bind.WatchOpts, sink chan<- *VMCreatorVMCreated) (event.Subscription, error) {

	logs, sub, err := _VMCreator.contract.WatchLogs(opts, "VMCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMCreatorVMCreated)
				if err := _VMCreator.contract.UnpackLog(event, "VMCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVMCreated is a log parse operation binding the contract event 0xb2f02f344f7e007eb90f5ca565eac18aaaaa406a00105e0647cfa902a6978bd7.
//
// Solidity: event VMCreated(address vmAddress, uint32 gracePeriod, uint128 escrowRequired, address escrowCurrency, uint32 maxExecutionSteps, bytes32 vmState, address owner, address[] validators)
func (_VMCreator *VMCreatorFilterer) ParseVMCreated(log types.Log) (*VMCreatorVMCreated, error) {
	event := new(VMCreatorVMCreated)
	if err := _VMCreator.contract.UnpackLog(event, "VMCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMTrackerABI is the input ABI used to generate the binding from.
const VMTrackerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"challengeManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"currentDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_players\",\"type\":\"address[2]\"},{\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_assertPreHash\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vm\",\"outputs\":[{\"name\":\"machineHash\",\"type\":\"bytes32\"},{\"name\":\"pendingHash\",\"type\":\"bytes32\"},{\"name\":\"inbox\",\"type\":\"bytes32\"},{\"name\":\"asserter\",\"type\":\"address\"},{\"name\":\"escrowRequired\",\"type\":\"uint128\"},{\"name\":\"deadline\",\"type\":\"uint64\"},{\"name\":\"sequenceNum\",\"type\":\"uint64\"},{\"name\":\"gracePeriod\",\"type\":\"uint32\"},{\"name\":\"maxExecutionSteps\",\"type\":\"uint32\"},{\"name\":\"validatorCount\",\"type\":\"uint16\"},{\"name\":\"state\",\"type\":\"uint8\"},{\"name\":\"inChallenge\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_preconditionHash\",\"type\":\"bytes32\"},{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestinations\",\"type\":\"address[]\"},{\"name\":\"_logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"confirmDisputableAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"isValidatorList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"terminateAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_unanRest\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_sequenceNum\",\"type\":\"uint64\"},{\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"pendingUnanimousAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activatedValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_newInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestinations\",\"type\":\"address[]\"}],\"name\":\"confirmUnanimousAsserted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fields\",\"type\":\"bytes32[4]\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageDataHash\",\"type\":\"bytes32[]\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestinations\",\"type\":\"address[]\"}],\"name\":\"pendingDisputableAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowRequired\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isListedValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_newInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageData\",\"type\":\"bytes\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"},{\"name\":\"_messageDestinations\",\"type\":\"address[]\"},{\"name\":\"_logsAccHash\",\"type\":\"bytes32\"},{\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"finalizedUnanimousAssert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"escrowCurrency\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"name\":\"_gracePeriod\",\"type\":\"uint32\"},{\"name\":\"_maxExecutionSteps\",\"type\":\"uint32\"},{\"name\":\"_escrowRequired\",\"type\":\"uint128\"},{\"name\":\"_escrowCurrency\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_challengeManagerAddress\",\"type\":\"address\"},{\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"name\":\"_validatorKeys\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"unanHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"PendingUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sequenceNum\",\"type\":\"uint64\"}],\"name\":\"ConfirmedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"unanHash\",\"type\":\"bytes32\"}],\"name\":\"FinalizedUnanimousAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fields\",\"type\":\"bytes32[3]\"},{\"indexed\":false,\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeBounds\",\"type\":\"uint64[2]\"},{\"indexed\":false,\"name\":\"tokenTypes\",\"type\":\"bytes21[]\"},{\"indexed\":false,\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"lastMessageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"logsAccHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"PendingDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newState\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"logsAccHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedDisputableAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"}]"

// VMTrackerFuncSigs maps the 4-byte function signature to its string representation.
var VMTrackerFuncSigs = map[string]string{
	"899b7c74": "activatedValidators()",
	"023a96fe": "challengeManager()",
	"22c091bc": "completeChallenge(address[2],uint128[2])",
	"4526c5d9": "confirmDisputableAsserted(bytes32,bytes32,uint32,bytes21[],bytes,uint16[],uint256[],address[],bytes32)",
	"8fe08c2b": "confirmUnanimousAsserted(bytes32,bytes32,bytes21[],bytes,uint16[],uint256[],address[])",
	"08dc89d7": "currentDeposit(address)",
	"fdf4a261": "escrowCurrency()",
	"aca0f372": "escrowRequired()",
	"6be00229": "exitAddress()",
	"bab72377": "finalizedUnanimousAssert(bytes32,bytes32,bytes21[],bytes,uint16[],uint256[],address[],bytes32,bytes)",
	"1865c57d": "getState()",
	"d489113a": "globalInbox()",
	"05b050de": "increaseDeposit()",
	"2782e87e": "initiateChallenge(bytes32)",
	"b99738e0": "isListedValidator(address)",
	"513164fe": "isValidatorList(address[])",
	"8da5cb5b": "owner()",
	"cfa80707": "ownerShutdown()",
	"97e2e256": "pendingDisputableAssert(bytes32[4],uint32,uint64[2],bytes21[],bytes32[],uint16[],uint256[],address[])",
	"675125b9": "pendingUnanimousAssert(bytes32,bytes21[],uint16[],uint256[],uint64,bytes32,bytes)",
	"60675a87": "terminateAddress()",
	"3a768463": "vm()",
}

// VMTrackerBin is the compiled bytecode used for deploying new contracts.
var VMTrackerBin = "0x60806040523480156200001157600080fd5b50604051620038263803806200382683398181016040526101208110156200003857600080fd5b8151602083015160408401516060850151608086015160a087015160c088015160e08901516101008a018051989a97999698959794969395929491938201926401000000008111156200008a57600080fd5b820160208101848111156200009e57600080fd5b8151856020820283011164010000000082111715620000bc57600080fd5b509093505050506001600160a01b0385161562000125576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180620038056021913960400191505060405180910390fd5b600180546001600160a01b038085166001600160a01b03199283161792839055600080548783169316929092178255604080517ff39723830000000000000000000000000000000000000000000000000000000081529051939091169263f39723839260048084019391929182900301818387803b158015620001a757600080fd5b505af1158015620001bc573d6000803e3d6000fd5b5050600c80546001600160a01b0319166001600160a01b0388811691909117909155600980546001600160b01b031916918916919091179055505080516008805461ffff60401b19166801000000000000000061ffff84160217905560005b8161ffff168161ffff161015620002a3576040518060400160405280600081526020016001151581525060026000016000858461ffff16815181106200025d57fe5b6020908102919091018101516001600160a01b0316825281810192909252604001600020825181559101516001918201805460ff1916911515919091179055016200021b565b5060038a90556008805460ff60501b19169055604080517f364df277000000000000000000000000000000000000000000000000000000008152905173__$d969135829891f807aa9c34494da4ecd99$__9163364df277916004808301926020929190829003018186803b1580156200031b57600080fd5b505af415801562000330573d6000803e3d6000fd5b505050506040513d60208110156200034757600080fd5b50516005555050600780546001600160801b0319166001600160801b03969096169590951790945550506008805463ffffffff191663ffffffff9586161763ffffffff60201b19166401000000009490951693909302939093179091555061344c9150819050620003b96000396000f3fe6080604052600436106101355760003560e01c80636be00229116100ab578063aca0f3721161006f578063aca0f37214610f3f578063b99738e014610f54578063bab7237714610f87578063cfa80707146112da578063d489113a146112ef578063fdf4a2611461130457610135565b80636be0022914610900578063899b7c74146109155780638da5cb5b146109415780638fe08c2b1461095657806397e2e25614610c1b57610135565b80632782e87e116100fd5780632782e87e146102205780633a7684631461024a5780634526c5d914610300578063513164fe146105d557806360675a8714610697578063675125b9146106ac57610135565b8063023a96fe1461013a57806305b050de1461016b57806308dc89d7146101755780631865c57d146101ba57806322c091bc146101f3575b600080fd5b34801561014657600080fd5b5061014f611319565b604080516001600160a01b039092168252519081900360200190f35b610173611328565b005b34801561018157600080fd5b506101a86004803603602081101561019857600080fd5b50356001600160a01b0316611453565b60408051918252519081900360200190f35b3480156101c657600080fd5b506101cf611472565b604051808260038111156101df57fe5b60ff16815260200191505060405180910390f35b3480156101ff57600080fd5b506101736004803603608081101561021657600080fd5b5060408101611482565b34801561022c57600080fd5b506101736004803603602081101561024357600080fd5b50356115d2565b34801561025657600080fd5b5061025f6117bd565b604080518d8152602081018d90529081018b90526001600160a01b038a1660608201526001600160801b038916608082015267ffffffffffffffff80891660a0830152871660c082015263ffffffff80871660e0830152851661010082015261ffff841661012082015261014081018360038111156102da57fe5b60ff1681529115156020830152506040805191829003019b509950505050505050505050f35b34801561030c57600080fd5b50610173600480360361012081101561032457600080fd5b81359160208101359163ffffffff6040830135169190810190608081016060820135600160201b81111561035757600080fd5b82018360208201111561036957600080fd5b803590602001918460208302840111600160201b8311171561038a57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103d957600080fd5b8201836020820111156103eb57600080fd5b803590602001918460018302840111600160201b8311171561040c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561045e57600080fd5b82018360208201111561047057600080fd5b803590602001918460208302840111600160201b8311171561049157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104e057600080fd5b8201836020820111156104f257600080fd5b803590602001918460208302840111600160201b8311171561051357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561056257600080fd5b82018360208201111561057457600080fd5b803590602001918460208302840111600160201b8311171561059557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550509135925061183a915050565b3480156105e157600080fd5b50610683600480360360208110156105f857600080fd5b810190602081018135600160201b81111561061257600080fd5b82018360208201111561062457600080fd5b803590602001918460208302840111600160201b8311171561064557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611c48945050505050565b604080519115158252519081900360200190f35b3480156106a357600080fd5b5061014f611cd5565b3480156106b857600080fd5b50610173600480360360e08110156106cf57600080fd5b81359190810190604081016020820135600160201b8111156106f057600080fd5b82018360208201111561070257600080fd5b803590602001918460208302840111600160201b8311171561072357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561077257600080fd5b82018360208201111561078457600080fd5b803590602001918460208302840111600160201b831117156107a557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156107f457600080fd5b82018360208201111561080657600080fd5b803590602001918460208302840111600160201b8311171561082757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929567ffffffffffffffff85351695602086013595919450925060608101915060400135600160201b81111561088c57600080fd5b82018360208201111561089e57600080fd5b803590602001918460018302840111600160201b831117156108bf57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611ce4945050505050565b34801561090c57600080fd5b5061014f6122f5565b34801561092157600080fd5b5061092a612304565b6040805161ffff9092168252519081900360200190f35b34801561094d57600080fd5b5061014f612315565b34801561096257600080fd5b50610173600480360360e081101561097957600080fd5b813591602081013591810190606081016040820135600160201b81111561099f57600080fd5b8201836020820111156109b157600080fd5b803590602001918460208302840111600160201b831117156109d257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a2157600080fd5b820183602082011115610a3357600080fd5b803590602001918460018302840111600160201b83111715610a5457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610aa657600080fd5b820183602082011115610ab857600080fd5b803590602001918460208302840111600160201b83111715610ad957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b2857600080fd5b820183602082011115610b3a57600080fd5b803590602001918460208302840111600160201b83111715610b5b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610baa57600080fd5b820183602082011115610bbc57600080fd5b803590602001918460208302840111600160201b83111715610bdd57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550612324945050505050565b348015610c2757600080fd5b506101736004803603610180811015610c3f57600080fd5b810190808060800190600480602002604051908101604052809291908260046020028082843760009201919091525050604080518082018252929563ffffffff853516959094909360608201935091602090910190600290839083908082843760009201919091525091949392602081019250359050600160201b811115610cc657600080fd5b820183602082011115610cd857600080fd5b803590602001918460208302840111600160201b83111715610cf957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d4857600080fd5b820183602082011115610d5a57600080fd5b803590602001918460208302840111600160201b83111715610d7b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610dca57600080fd5b820183602082011115610ddc57600080fd5b803590602001918460208302840111600160201b83111715610dfd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610e4c57600080fd5b820183602082011115610e5e57600080fd5b803590602001918460208302840111600160201b83111715610e7f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610ece57600080fd5b820183602082011115610ee057600080fd5b803590602001918460208302840111600160201b83111715610f0157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550612714945050505050565b348015610f4b57600080fd5b506101a8612dd5565b348015610f6057600080fd5b5061068360048036036020811015610f7757600080fd5b50356001600160a01b0316612de4565b348015610f9357600080fd5b506101736004803603610120811015610fab57600080fd5b813591602081013591810190606081016040820135600160201b811115610fd157600080fd5b820183602082011115610fe357600080fd5b803590602001918460208302840111600160201b8311171561100457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561105357600080fd5b82018360208201111561106557600080fd5b803590602001918460018302840111600160201b8311171561108657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156110d857600080fd5b8201836020820111156110ea57600080fd5b803590602001918460208302840111600160201b8311171561110b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561115a57600080fd5b82018360208201111561116c57600080fd5b803590602001918460208302840111600160201b8311171561118d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156111dc57600080fd5b8201836020820111156111ee57600080fd5b803590602001918460208302840111600160201b8311171561120f57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092958435959094909350604081019250602001359050600160201b81111561126657600080fd5b82018360208201111561127857600080fd5b803590602001918460018302840111600160201b8311171561129957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612e05945050505050565b3480156112e657600080fd5b50610173613122565b3480156112fb57600080fd5b5061014f61318b565b34801561131057600080fd5b5061014f61319a565b6000546001600160a01b031681565b3360009081526002602052604090206001015460ff1661138a576040805162461bcd60e51b815260206004820152601860248201527721b0b63632b91036bab9ba103132903b30b634b230ba37b960411b604482015290519081900360640190fd5b336000908152600260205260409020600754815434810183556001600160801b03909116118080156113ca575060075482546001600160801b0390911611155b156113f65760098054600161ffff600160a01b808404821692909201160261ffff60a01b199091161790555b600854600954600160a01b900461ffff908116600160401b9092041614801561143657506000600854600160501b900460ff16600381111561143457fe5b145b1561144f576008805460ff60501b1916600160501b1790555b5050565b6001600160a01b0381166000908152600260205260409020545b919050565b600854600160501b900460ff1690565b6000546001600160a01b031633146114cb5760405162461bcd60e51b815260040180806020018281038252602d8152602001806133c7602d913960400191505060405180910390fd5b600854600160581b900460ff166115135760405162461bcd60e51b81526004018080602001828103825260268152602001806133a16026913960400191505060405180910390fd5b6008805460ff60581b191690556115756001600160801b038235166002600085815b60200201356001600160a01b03166001600160a01b03166001600160a01b03168152602001908152602001600020600001546131a990919063ffffffff16565b82356001600160a01b031660009081526002602081815260408320939093556115ad928401356001600160801b031691856001611535565b6001600160a01b03602093840135166000908152600290935260409092209190915550565b3360009081526002602052604090206001015460ff16611634576040805162461bcd60e51b815260206004820152601860248201527721b0b63632b91036bab9ba103132903b30b634b230ba37b960411b604482015290519081900360640190fd5b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__635af93c7a6002836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561168c57600080fd5b505af41580156116a0573d6000803e3d6000fd5b5050600080546040805180820182526006546001600160a01b03908116825233602080840191909152835180850185526007546001600160801b0316808252918101919091526008548451630823813560e21b815292909516975063208e04d496509194919363ffffffff16928892600490920191829187918190849084905b83811015611738578181015183820152602001611720565b5050505090500184600260200280838360005b8381101561176357818101518382015260200161174b565b505050509050018363ffffffff1663ffffffff168152602001828152602001945050505050600060405180830381600087803b1580156117a257600080fd5b505af11580156117b6573d6000803e3d6000fd5b5050505050565b6003546004546005546006546007546008546001600160a01b03909216916001600160801b0382169167ffffffffffffffff600160801b8204811692600160c01b909204169063ffffffff80821691600160201b81049091169061ffff600160401b8204169060ff600160501b8204811691600160581b9004168c565b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63924e7b3760028b8b8b8b8b8b8b8b8b6040518b63ffffffff1660e01b8152600401808b81526020018a81526020018981526020018863ffffffff1663ffffffff168152602001806020018060200180602001806020018060200187815260200186810386528c818151815260200191508051906020019060200280838360005b838110156118e85781810151838201526020016118d0565b5050505090500186810385528b818151815260200191508051906020019080838360005b8381101561192457818101518382015260200161190c565b50505050905090810190601f1680156119515780820380516001836020036101000a031916815260200191505b5086810384528a5181528a51602091820191808d01910280838360005b8381101561198657818101518382015260200161196e565b50505050905001868103835289818151815260200191508051906020019060200280838360005b838110156119c55781810151838201526020016119ad565b50505050905001868103825288818151815260200191508051906020019060200280838360005b83811015611a045781810151838201526020016119ec565b505050509050019f5050505050505050505050505050505060006040518083038186803b158015611a3457600080fd5b505af4158015611a48573d6000803e3d6000fd5b50505050611a5461320a565b600160009054906101000a90046001600160a01b03166001600160a01b031663ec22a76787878787876040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015611adc578181015183820152602001611ac4565b5050505090500186810385528a818151815260200191508051906020019080838360005b83811015611b18578181015183820152602001611b00565b50505050905090810190601f168015611b455780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015611b7a578181015183820152602001611b62565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015611bb9578181015183820152602001611ba1565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015611bf8578181015183820152602001611be0565b505050509050019a5050505050505050505050600060405180830381600087803b158015611c2557600080fd5b505af1158015611c39573d6000803e3d6000fd5b50505050505050505050505050565b805160085460009190600160401b900461ffff168114611c6c57600091505061146d565b60005b81811015611ccb5760026000016000858381518110611c8a57fe5b6020908102919091018101516001600160a01b031682528101919091526040016000206001015460ff16611cc35760009250505061146d565b600101611c6f565b5060019392505050565b600b546001600160a01b031681565b606073__$9836fa7140e5a33041d4b827682e675a30$__630f89fbff8888886040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015611d5a578181015183820152602001611d42565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015611d99578181015183820152602001611d81565b50505050905001848103825285818151815260200191508051906020019060200280838360005b83811015611dd8578181015183820152602001611dc0565b50505050905001965050505050505060006040518083038186803b158015611dff57600080fd5b505af4158015611e13573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015611e3c57600080fd5b810190808051600160201b811115611e5357600080fd5b82016020810184811115611e6657600080fd5b81518560208202830111600160201b82111715611e8257600080fd5b50506040805163578bec9160e11b8152600481019182528c5160448201528c5192965073__$9836fa7140e5a33041d4b827682e675a30$__955063af17d92294508c93508692829160248101916064909101906020808801910280838360005b83811015611efa578181015183820152602001611ee2565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611f39578181015183820152602001611f21565b5050505090500194505050505060206040518083038186803b158015611f5e57600080fd5b505af4158015611f72573d6000803e3d6000fd5b505050506040513d6020811015611f8857600080fd5b5051611fc55760405162461bcd60e51b81526004018080602001828103825260248152602001806133f46024913960400191505060405180910390fd5b60015460405163565b19db60e11b815230600482018181526060602484019081528b5160648501528b516001600160a01b039095169463acb633b6948d9388939092909160448101916084909101906020808801910280838360005b83811015612039578181015183820152602001612021565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015612078578181015183820152602001612060565b505050509050019550505050505060206040518083038186803b15801561209e57600080fd5b505afa1580156120b2573d6000803e3d6000fd5b505050506040513d60208110156120c857600080fd5b505161211b576040805162461bcd60e51b815260206004820152601b60248201527f564d2068617320696e73756666696369656e742062616c616e63650000000000604482015290519081900360640190fd5b73__$caf066876633ea418098495f1e5bb4c2f5$__63a2ee90f460028a8a8a8a8a8a8a6040518963ffffffff1660e01b8152600401808981526020018881526020018060200180602001806020018767ffffffffffffffff1667ffffffffffffffff1681526020018681526020018060200185810385528b818151815260200191508051906020019060200280838360005b838110156121c55781810151838201526020016121ad565b5050505090500185810384528a818151815260200191508051906020019060200280838360005b838110156122045781810151838201526020016121ec565b50505050905001858103835289818151815260200191508051906020019060200280838360005b8381101561224357818101518382015260200161222b565b50505050905001858103825286818151815260200191508051906020019080838360005b8381101561227f578181015183820152602001612267565b50505050905090810190601f1680156122ac5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060006040518083038186803b1580156122d357600080fd5b505af41580156122e7573d6000803e3d6000fd5b505050505050505050505050565b600a546001600160a01b031681565b600954600160a01b900461ffff1681565b600c546001600160a01b031681565b73__$caf066876633ea418098495f1e5bb4c2f5$__63319b5f676002898989898989896040518963ffffffff1660e01b815260040180898152602001888152602001878152602001806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b838110156123b85781810151838201526020016123a0565b5050505090500186810385528a818151815260200191508051906020019080838360005b838110156123f45781810151838201526020016123dc565b50505050905090810190601f1680156124215780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b8381101561245657818101518382015260200161243e565b50505050905001868103835288818151815260200191508051906020019060200280838360005b8381101561249557818101518382015260200161247d565b50505050905001868103825287818151815260200191508051906020019060200280838360005b838110156124d45781810151838201526020016124bc565b505050509050019d505050505050505050505050505060006040518083038186803b15801561250257600080fd5b505af4158015612516573d6000803e3d6000fd5b5050505061252261320a565b600160009054906101000a90046001600160a01b03166001600160a01b031663ec22a76786868686866040518663ffffffff1660e01b815260040180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b838110156125aa578181015183820152602001612592565b5050505090500186810385528a818151815260200191508051906020019080838360005b838110156125e65781810151838201526020016125ce565b50505050905090810190601f1680156126135780820380516001836020036101000a031916815260200191505b508681038452895181528951602091820191808c01910280838360005b83811015612648578181015183820152602001612630565b50505050905001868103835288818151815260200191508051906020019060200280838360005b8381101561268757818101518382015260200161266f565b50505050905001868103825287818151815260200191508051906020019060200280838360005b838110156126c65781810151838201526020016126ae565b505050509050019a5050505050505050505050600060405180830381600087803b1580156126f357600080fd5b505af1158015612707573d6000803e3d6000fd5b5050505050505050505050565b3360009081526002602052604090206001015460ff16612776576040805162461bcd60e51b815260206004820152601860248201527721b0b63632b91036bab9ba103132903b30b634b230ba37b960411b604482015290519081900360640190fd5b606073__$9836fa7140e5a33041d4b827682e675a30$__630f89fbff8786866040518463ffffffff1660e01b815260040180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b838110156127ec5781810151838201526020016127d4565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561282b578181015183820152602001612813565b50505050905001848103825285818151815260200191508051906020019060200280838360005b8381101561286a578181015183820152602001612852565b50505050905001965050505050505060006040518083038186803b15801561289157600080fd5b505af41580156128a5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156128ce57600080fd5b810190808051600160201b8111156128e557600080fd5b820160208101848111156128f857600080fd5b81518560208202830111600160201b8211171561291457600080fd5b50506040805163578bec9160e11b8152600481019182528b5160448201528b5192965073__$9836fa7140e5a33041d4b827682e675a30$__955063af17d92294508b93508692829160248101916064909101906020808801910280838360005b8381101561298c578181015183820152602001612974565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156129cb5781810151838201526020016129b3565b5050505090500194505050505060206040518083038186803b1580156129f057600080fd5b505af4158015612a04573d6000803e3d6000fd5b505050506040513d6020811015612a1a57600080fd5b5051612a575760405162461bcd60e51b81526004018080602001828103825260248152602001806133f46024913960400191505060405180910390fd5b60015460405163565b19db60e11b815230600482018181526060602484019081528a5160648501528a516001600160a01b039095169463acb633b6948c9388939092909160448101916084909101906020808801910280838360005b83811015612acb578181015183820152602001612ab3565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015612b0a578181015183820152602001612af2565b505050509050019550505050505060206040518083038186803b158015612b3057600080fd5b505afa158015612b44573d6000803e3d6000fd5b505050506040513d6020811015612b5a57600080fd5b5051612bad576040805162461bcd60e51b815260206004820152601b60248201527f564d2068617320696e73756666696369656e742062616c616e63650000000000604482015290519081900360640190fd5b73__$2104f4b4ea1fa2fd2334e6605946f6eea1$__63c97c8eec60028b8b8b8b8b8b8b8b6040518a63ffffffff1660e01b8152600401808a815260200189600460200280838360005b83811015612c0e578181015183820152602001612bf6565b5050505063ffffffff8b1692019182525060200187604080838360005b83811015612c43578181015183820152602001612c2b565b50505050905001806020018060200180602001806020018060200186810386528b818151815260200191508051906020019060200280838360005b83811015612c96578181015183820152602001612c7e565b5050505090500186810385528a818151815260200191508051906020019060200280838360005b83811015612cd5578181015183820152602001612cbd565b50505050905001868103845289818151815260200191508051906020019060200280838360005b83811015612d14578181015183820152602001612cfc565b50505050905001868103835288818151815260200191508051906020019060200280838360005b83811015612d53578181015183820152602001612d3b565b50505050905001868103825287818151815260200191508051906020019060200280838360005b83811015612d92578181015183820152602001612d7a565b505050509050019e50505050505050505050505050505060006040518083038186803b158015612dc157600080fd5b505af4158015611c39573d6000803e3d6000fd5b6007546001600160801b031690565b6001600160a01b031660009081526002602052604090206001015460ff1690565b73__$caf066876633ea418098495f1e5bb4c2f5$__633c0ec6de600260405180606001604052808d81526020018c8152602001868152508a8a8a8a8a896040518963ffffffff1660e01b81526004018089815260200188600360200280838360005b83811015612e7f578181015183820152602001612e67565b5050505090500180602001806020018060200180602001806020018060200187810387528d818151815260200191508051906020019060200280838360005b83811015612ed6578181015183820152602001612ebe565b5050505090500187810386528c818151815260200191508051906020019080838360005b83811015612f12578181015183820152602001612efa565b50505050905090810190601f168015612f3f5780820380516001836020036101000a031916815260200191505b5087810385528b5181528b51602091820191808e01910280838360005b83811015612f74578181015183820152602001612f5c565b5050505090500187810384528a818151815260200191508051906020019060200280838360005b83811015612fb3578181015183820152602001612f9b565b50505050905001878103835289818151815260200191508051906020019060200280838360005b83811015612ff2578181015183820152602001612fda565b50505050905001878103825288818151815260200191508051906020019080838360005b8381101561302e578181015183820152602001613016565b50505050905090810190601f16801561305b5780820380516001836020036101000a031916815260200191505b509e50505050505050505050505050505060006040518083038186803b15801561308457600080fd5b505af4158015613098573d6000803e3d6000fd5b505050506130a461320a565b60015460405163ec22a76760e01b815260a060048201908152895160a483015289516001600160a01b039093169263ec22a767928b928b928b928b928b9290918291602482019160448101916064820191608481019160c4909101906020808e01910280838360008315611adc578181015183820152602001611ac4565b600c546001600160a01b03163314613181576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6c79206f776e65722063616e2073687574646f776e2074686520564d0000604482015290519081900360640190fd5b613189613392565b565b6001546001600160a01b031681565b6009546001600160a01b031681565b600082820183811015613203576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600154604080516331618a0360e01b815230600482015290516000926001600160a01b0316916331618a0391602480830192602092919082900301818787803b15801561325657600080fd5b505af115801561326a573d6000803e3d6000fd5b505050506040513d602081101561328057600080fd5b50516040805163364df27760e01b8152905191925073__$d969135829891f807aa9c34494da4ecd99$__9163364df27791600480820192602092909190829003018186803b1580156132d157600080fd5b505af41580156132e5573d6000803e3d6000fd5b505050506040513d60208110156132fb57600080fd5b5051811461338f5773__$9836fa7140e5a33041d4b827682e675a30$__63f11fcc26600260030154836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b15801561335f57600080fd5b505af4158015613373573d6000803e3d6000fd5b505050506040513d602081101561338957600080fd5b50516005555b50565b600c546001600160a01b0316fffe564d206d75737420626520696e206368616c6c656e676520746f20636f6d706c6574652069744f6e6c79206368616c6c656e6765206d616e616765722063616e20636f6d706c657465206368616c6c656e6765546f6b656e207479706573206d7573742062652076616c696420616e6420736f72746564a265627a7a7230582072f27c35176282cc55c31b69184c9daa06284f5cbe091e0506d3515f45933f0764736f6c634300050a003256616c696461746f72206465706f73697473206d75737420626520696e20455448"

// DeployVMTracker deploys a new Ethereum contract, binding an instance of VMTracker to it.
func DeployVMTracker(auth *bind.TransactOpts, backend bind.ContractBackend, _vmState [32]byte, _gracePeriod uint32, _maxExecutionSteps uint32, _escrowRequired *big.Int, _escrowCurrency common.Address, _owner common.Address, _challengeManagerAddress common.Address, _globalInboxAddress common.Address, _validatorKeys []common.Address) (common.Address, *types.Transaction, *VMTracker, error) {
	parsed, err := abi.JSON(strings.NewReader(VMTrackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	disputableAddr, _, _, _ := DeployDisputable(auth, backend)
	VMTrackerBin = strings.Replace(VMTrackerBin, "__$2104f4b4ea1fa2fd2334e6605946f6eea1$__", disputableAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	VMTrackerBin = strings.Replace(VMTrackerBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	unanimousAddr, _, _, _ := DeployUnanimous(auth, backend)
	VMTrackerBin = strings.Replace(VMTrackerBin, "__$caf066876633ea418098495f1e5bb4c2f5$__", unanimousAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	VMTrackerBin = strings.Replace(VMTrackerBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VMTrackerBin), backend, _vmState, _gracePeriod, _maxExecutionSteps, _escrowRequired, _escrowCurrency, _owner, _challengeManagerAddress, _globalInboxAddress, _validatorKeys)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VMTracker{VMTrackerCaller: VMTrackerCaller{contract: contract}, VMTrackerTransactor: VMTrackerTransactor{contract: contract}, VMTrackerFilterer: VMTrackerFilterer{contract: contract}}, nil
}

// VMTracker is an auto generated Go binding around an Ethereum contract.
type VMTracker struct {
	VMTrackerCaller     // Read-only binding to the contract
	VMTrackerTransactor // Write-only binding to the contract
	VMTrackerFilterer   // Log filterer for contract events
}

// VMTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VMTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VMTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VMTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VMTrackerSession struct {
	Contract     *VMTracker        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VMTrackerCallerSession struct {
	Contract *VMTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VMTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VMTrackerTransactorSession struct {
	Contract     *VMTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VMTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VMTrackerRaw struct {
	Contract *VMTracker // Generic contract binding to access the raw methods on
}

// VMTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VMTrackerCallerRaw struct {
	Contract *VMTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// VMTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VMTrackerTransactorRaw struct {
	Contract *VMTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVMTracker creates a new instance of VMTracker, bound to a specific deployed contract.
func NewVMTracker(address common.Address, backend bind.ContractBackend) (*VMTracker, error) {
	contract, err := bindVMTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VMTracker{VMTrackerCaller: VMTrackerCaller{contract: contract}, VMTrackerTransactor: VMTrackerTransactor{contract: contract}, VMTrackerFilterer: VMTrackerFilterer{contract: contract}}, nil
}

// NewVMTrackerCaller creates a new read-only instance of VMTracker, bound to a specific deployed contract.
func NewVMTrackerCaller(address common.Address, caller bind.ContractCaller) (*VMTrackerCaller, error) {
	contract, err := bindVMTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VMTrackerCaller{contract: contract}, nil
}

// NewVMTrackerTransactor creates a new write-only instance of VMTracker, bound to a specific deployed contract.
func NewVMTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*VMTrackerTransactor, error) {
	contract, err := bindVMTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VMTrackerTransactor{contract: contract}, nil
}

// NewVMTrackerFilterer creates a new log filterer instance of VMTracker, bound to a specific deployed contract.
func NewVMTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*VMTrackerFilterer, error) {
	contract, err := bindVMTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VMTrackerFilterer{contract: contract}, nil
}

// bindVMTracker binds a generic wrapper to an already deployed contract.
func bindVMTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VMTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VMTracker *VMTrackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VMTracker.Contract.VMTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VMTracker *VMTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMTracker.Contract.VMTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VMTracker *VMTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VMTracker.Contract.VMTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VMTracker *VMTrackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VMTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VMTracker *VMTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VMTracker *VMTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VMTracker.Contract.contract.Transact(opts, method, params...)
}

// ActivatedValidators is a free data retrieval call binding the contract method 0x899b7c74.
//
// Solidity: function activatedValidators() constant returns(uint16)
func (_VMTracker *VMTrackerCaller) ActivatedValidators(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "activatedValidators")
	return *ret0, err
}

// ActivatedValidators is a free data retrieval call binding the contract method 0x899b7c74.
//
// Solidity: function activatedValidators() constant returns(uint16)
func (_VMTracker *VMTrackerSession) ActivatedValidators() (uint16, error) {
	return _VMTracker.Contract.ActivatedValidators(&_VMTracker.CallOpts)
}

// ActivatedValidators is a free data retrieval call binding the contract method 0x899b7c74.
//
// Solidity: function activatedValidators() constant returns(uint16)
func (_VMTracker *VMTrackerCallerSession) ActivatedValidators() (uint16, error) {
	return _VMTracker.Contract.ActivatedValidators(&_VMTracker.CallOpts)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_VMTracker *VMTrackerCaller) ChallengeManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "challengeManager")
	return *ret0, err
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_VMTracker *VMTrackerSession) ChallengeManager() (common.Address, error) {
	return _VMTracker.Contract.ChallengeManager(&_VMTracker.CallOpts)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() constant returns(address)
func (_VMTracker *VMTrackerCallerSession) ChallengeManager() (common.Address, error) {
	return _VMTracker.Contract.ChallengeManager(&_VMTracker.CallOpts)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_VMTracker *VMTrackerCaller) CurrentDeposit(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "currentDeposit", validator)
	return *ret0, err
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_VMTracker *VMTrackerSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _VMTracker.Contract.CurrentDeposit(&_VMTracker.CallOpts, validator)
}

// CurrentDeposit is a free data retrieval call binding the contract method 0x08dc89d7.
//
// Solidity: function currentDeposit(address validator) constant returns(uint256)
func (_VMTracker *VMTrackerCallerSession) CurrentDeposit(validator common.Address) (*big.Int, error) {
	return _VMTracker.Contract.CurrentDeposit(&_VMTracker.CallOpts, validator)
}

// EscrowCurrency is a free data retrieval call binding the contract method 0xfdf4a261.
//
// Solidity: function escrowCurrency() constant returns(address)
func (_VMTracker *VMTrackerCaller) EscrowCurrency(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "escrowCurrency")
	return *ret0, err
}

// EscrowCurrency is a free data retrieval call binding the contract method 0xfdf4a261.
//
// Solidity: function escrowCurrency() constant returns(address)
func (_VMTracker *VMTrackerSession) EscrowCurrency() (common.Address, error) {
	return _VMTracker.Contract.EscrowCurrency(&_VMTracker.CallOpts)
}

// EscrowCurrency is a free data retrieval call binding the contract method 0xfdf4a261.
//
// Solidity: function escrowCurrency() constant returns(address)
func (_VMTracker *VMTrackerCallerSession) EscrowCurrency() (common.Address, error) {
	return _VMTracker.Contract.EscrowCurrency(&_VMTracker.CallOpts)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_VMTracker *VMTrackerCaller) EscrowRequired(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "escrowRequired")
	return *ret0, err
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_VMTracker *VMTrackerSession) EscrowRequired() (*big.Int, error) {
	return _VMTracker.Contract.EscrowRequired(&_VMTracker.CallOpts)
}

// EscrowRequired is a free data retrieval call binding the contract method 0xaca0f372.
//
// Solidity: function escrowRequired() constant returns(uint256)
func (_VMTracker *VMTrackerCallerSession) EscrowRequired() (*big.Int, error) {
	return _VMTracker.Contract.EscrowRequired(&_VMTracker.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_VMTracker *VMTrackerCaller) ExitAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "exitAddress")
	return *ret0, err
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_VMTracker *VMTrackerSession) ExitAddress() (common.Address, error) {
	return _VMTracker.Contract.ExitAddress(&_VMTracker.CallOpts)
}

// ExitAddress is a free data retrieval call binding the contract method 0x6be00229.
//
// Solidity: function exitAddress() constant returns(address)
func (_VMTracker *VMTrackerCallerSession) ExitAddress() (common.Address, error) {
	return _VMTracker.Contract.ExitAddress(&_VMTracker.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_VMTracker *VMTrackerCaller) GetState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "getState")
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_VMTracker *VMTrackerSession) GetState() (uint8, error) {
	return _VMTracker.Contract.GetState(&_VMTracker.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() constant returns(uint8)
func (_VMTracker *VMTrackerCallerSession) GetState() (uint8, error) {
	return _VMTracker.Contract.GetState(&_VMTracker.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_VMTracker *VMTrackerCaller) GlobalInbox(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "globalInbox")
	return *ret0, err
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_VMTracker *VMTrackerSession) GlobalInbox() (common.Address, error) {
	return _VMTracker.Contract.GlobalInbox(&_VMTracker.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() constant returns(address)
func (_VMTracker *VMTrackerCallerSession) GlobalInbox() (common.Address, error) {
	return _VMTracker.Contract.GlobalInbox(&_VMTracker.CallOpts)
}

// IsListedValidator is a free data retrieval call binding the contract method 0xb99738e0.
//
// Solidity: function isListedValidator(address validator) constant returns(bool)
func (_VMTracker *VMTrackerCaller) IsListedValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "isListedValidator", validator)
	return *ret0, err
}

// IsListedValidator is a free data retrieval call binding the contract method 0xb99738e0.
//
// Solidity: function isListedValidator(address validator) constant returns(bool)
func (_VMTracker *VMTrackerSession) IsListedValidator(validator common.Address) (bool, error) {
	return _VMTracker.Contract.IsListedValidator(&_VMTracker.CallOpts, validator)
}

// IsListedValidator is a free data retrieval call binding the contract method 0xb99738e0.
//
// Solidity: function isListedValidator(address validator) constant returns(bool)
func (_VMTracker *VMTrackerCallerSession) IsListedValidator(validator common.Address) (bool, error) {
	return _VMTracker.Contract.IsListedValidator(&_VMTracker.CallOpts, validator)
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_VMTracker *VMTrackerCaller) IsValidatorList(opts *bind.CallOpts, _validators []common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "isValidatorList", _validators)
	return *ret0, err
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_VMTracker *VMTrackerSession) IsValidatorList(_validators []common.Address) (bool, error) {
	return _VMTracker.Contract.IsValidatorList(&_VMTracker.CallOpts, _validators)
}

// IsValidatorList is a free data retrieval call binding the contract method 0x513164fe.
//
// Solidity: function isValidatorList(address[] _validators) constant returns(bool)
func (_VMTracker *VMTrackerCallerSession) IsValidatorList(_validators []common.Address) (bool, error) {
	return _VMTracker.Contract.IsValidatorList(&_VMTracker.CallOpts, _validators)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VMTracker *VMTrackerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VMTracker *VMTrackerSession) Owner() (common.Address, error) {
	return _VMTracker.Contract.Owner(&_VMTracker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VMTracker *VMTrackerCallerSession) Owner() (common.Address, error) {
	return _VMTracker.Contract.Owner(&_VMTracker.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_VMTracker *VMTrackerCaller) TerminateAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VMTracker.contract.Call(opts, out, "terminateAddress")
	return *ret0, err
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_VMTracker *VMTrackerSession) TerminateAddress() (common.Address, error) {
	return _VMTracker.Contract.TerminateAddress(&_VMTracker.CallOpts)
}

// TerminateAddress is a free data retrieval call binding the contract method 0x60675a87.
//
// Solidity: function terminateAddress() constant returns(address)
func (_VMTracker *VMTrackerCallerSession) TerminateAddress() (common.Address, error) {
	return _VMTracker.Contract.TerminateAddress(&_VMTracker.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint16 validatorCount, uint8 state, bool inChallenge)
func (_VMTracker *VMTrackerCaller) Vm(opts *bind.CallOpts) (struct {
	MachineHash       [32]byte
	PendingHash       [32]byte
	Inbox             [32]byte
	Asserter          common.Address
	EscrowRequired    *big.Int
	Deadline          uint64
	SequenceNum       uint64
	GracePeriod       uint32
	MaxExecutionSteps uint32
	ValidatorCount    uint16
	State             uint8
	InChallenge       bool
}, error) {
	ret := new(struct {
		MachineHash       [32]byte
		PendingHash       [32]byte
		Inbox             [32]byte
		Asserter          common.Address
		EscrowRequired    *big.Int
		Deadline          uint64
		SequenceNum       uint64
		GracePeriod       uint32
		MaxExecutionSteps uint32
		ValidatorCount    uint16
		State             uint8
		InChallenge       bool
	})
	out := ret
	err := _VMTracker.contract.Call(opts, out, "vm")
	return *ret, err
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint16 validatorCount, uint8 state, bool inChallenge)
func (_VMTracker *VMTrackerSession) Vm() (struct {
	MachineHash       [32]byte
	PendingHash       [32]byte
	Inbox             [32]byte
	Asserter          common.Address
	EscrowRequired    *big.Int
	Deadline          uint64
	SequenceNum       uint64
	GracePeriod       uint32
	MaxExecutionSteps uint32
	ValidatorCount    uint16
	State             uint8
	InChallenge       bool
}, error) {
	return _VMTracker.Contract.Vm(&_VMTracker.CallOpts)
}

// Vm is a free data retrieval call binding the contract method 0x3a768463.
//
// Solidity: function vm() constant returns(bytes32 machineHash, bytes32 pendingHash, bytes32 inbox, address asserter, uint128 escrowRequired, uint64 deadline, uint64 sequenceNum, uint32 gracePeriod, uint32 maxExecutionSteps, uint16 validatorCount, uint8 state, bool inChallenge)
func (_VMTracker *VMTrackerCallerSession) Vm() (struct {
	MachineHash       [32]byte
	PendingHash       [32]byte
	Inbox             [32]byte
	Asserter          common.Address
	EscrowRequired    *big.Int
	Deadline          uint64
	SequenceNum       uint64
	GracePeriod       uint32
	MaxExecutionSteps uint32
	ValidatorCount    uint16
	State             uint8
	InChallenge       bool
}, error) {
	return _VMTracker.Contract.Vm(&_VMTracker.CallOpts)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_VMTracker *VMTrackerTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_VMTracker *VMTrackerSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _VMTracker.Contract.CompleteChallenge(&_VMTracker.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_VMTracker *VMTrackerTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _VMTracker.Contract.CompleteChallenge(&_VMTracker.TransactOpts, _players, _rewards)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4526c5d9.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash) returns()
func (_VMTracker *VMTrackerTransactor) ConfirmDisputableAsserted(opts *bind.TransactOpts, _preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "confirmDisputableAsserted", _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4526c5d9.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash) returns()
func (_VMTracker *VMTrackerSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmDisputableAsserted(&_VMTracker.TransactOpts, _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash)
}

// ConfirmDisputableAsserted is a paid mutator transaction binding the contract method 0x4526c5d9.
//
// Solidity: function confirmDisputableAsserted(bytes32 _preconditionHash, bytes32 _afterHash, uint32 _numSteps, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash) returns()
func (_VMTracker *VMTrackerTransactorSession) ConfirmDisputableAsserted(_preconditionHash [32]byte, _afterHash [32]byte, _numSteps uint32, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmDisputableAsserted(&_VMTracker.TransactOpts, _preconditionHash, _afterHash, _numSteps, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0x8fe08c2b.
//
// Solidity: function confirmUnanimousAsserted(bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_VMTracker *VMTrackerTransactor) ConfirmUnanimousAsserted(opts *bind.TransactOpts, _afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "confirmUnanimousAsserted", _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0x8fe08c2b.
//
// Solidity: function confirmUnanimousAsserted(bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_VMTracker *VMTrackerSession) ConfirmUnanimousAsserted(_afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmUnanimousAsserted(&_VMTracker.TransactOpts, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// ConfirmUnanimousAsserted is a paid mutator transaction binding the contract method 0x8fe08c2b.
//
// Solidity: function confirmUnanimousAsserted(bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_VMTracker *VMTrackerTransactorSession) ConfirmUnanimousAsserted(_afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _VMTracker.Contract.ConfirmUnanimousAsserted(&_VMTracker.TransactOpts, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// FinalizedUnanimousAssert is a paid mutator transaction binding the contract method 0xbab72377.
//
// Solidity: function finalizedUnanimousAssert(bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactor) FinalizedUnanimousAssert(opts *bind.TransactOpts, _afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "finalizedUnanimousAssert", _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash, _signatures)
}

// FinalizedUnanimousAssert is a paid mutator transaction binding the contract method 0xbab72377.
//
// Solidity: function finalizedUnanimousAssert(bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerSession) FinalizedUnanimousAssert(_afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.FinalizedUnanimousAssert(&_VMTracker.TransactOpts, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash, _signatures)
}

// FinalizedUnanimousAssert is a paid mutator transaction binding the contract method 0xbab72377.
//
// Solidity: function finalizedUnanimousAssert(bytes32 _afterHash, bytes32 _newInbox, bytes21[] _tokenTypes, bytes _messageData, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactorSession) FinalizedUnanimousAssert(_afterHash [32]byte, _newInbox [32]byte, _tokenTypes [][21]byte, _messageData []byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.FinalizedUnanimousAssert(&_VMTracker.TransactOpts, _afterHash, _newInbox, _tokenTypes, _messageData, _messageTokenNums, _messageAmounts, _messageDestinations, _logsAccHash, _signatures)
}

// IncreaseDeposit is a paid mutator transaction binding the contract method 0x05b050de.
//
// Solidity: function increaseDeposit() returns()
func (_VMTracker *VMTrackerTransactor) IncreaseDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "increaseDeposit")
}

// IncreaseDeposit is a paid mutator transaction binding the contract method 0x05b050de.
//
// Solidity: function increaseDeposit() returns()
func (_VMTracker *VMTrackerSession) IncreaseDeposit() (*types.Transaction, error) {
	return _VMTracker.Contract.IncreaseDeposit(&_VMTracker.TransactOpts)
}

// IncreaseDeposit is a paid mutator transaction binding the contract method 0x05b050de.
//
// Solidity: function increaseDeposit() returns()
func (_VMTracker *VMTrackerTransactorSession) IncreaseDeposit() (*types.Transaction, error) {
	return _VMTracker.Contract.IncreaseDeposit(&_VMTracker.TransactOpts)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_VMTracker *VMTrackerTransactor) InitiateChallenge(opts *bind.TransactOpts, _assertPreHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "initiateChallenge", _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_VMTracker *VMTrackerSession) InitiateChallenge(_assertPreHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.InitiateChallenge(&_VMTracker.TransactOpts, _assertPreHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x2782e87e.
//
// Solidity: function initiateChallenge(bytes32 _assertPreHash) returns()
func (_VMTracker *VMTrackerTransactorSession) InitiateChallenge(_assertPreHash [32]byte) (*types.Transaction, error) {
	return _VMTracker.Contract.InitiateChallenge(&_VMTracker.TransactOpts, _assertPreHash)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_VMTracker *VMTrackerTransactor) OwnerShutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "ownerShutdown")
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_VMTracker *VMTrackerSession) OwnerShutdown() (*types.Transaction, error) {
	return _VMTracker.Contract.OwnerShutdown(&_VMTracker.TransactOpts)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_VMTracker *VMTrackerTransactorSession) OwnerShutdown() (*types.Transaction, error) {
	return _VMTracker.Contract.OwnerShutdown(&_VMTracker.TransactOpts)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x97e2e256.
//
// Solidity: function pendingDisputableAssert(bytes32[4] _fields, uint32 _numSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_VMTracker *VMTrackerTransactor) PendingDisputableAssert(opts *bind.TransactOpts, _fields [4][32]byte, _numSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "pendingDisputableAssert", _fields, _numSteps, _timeBounds, _tokenTypes, _messageDataHash, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x97e2e256.
//
// Solidity: function pendingDisputableAssert(bytes32[4] _fields, uint32 _numSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_VMTracker *VMTrackerSession) PendingDisputableAssert(_fields [4][32]byte, _numSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _VMTracker.Contract.PendingDisputableAssert(&_VMTracker.TransactOpts, _fields, _numSteps, _timeBounds, _tokenTypes, _messageDataHash, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// PendingDisputableAssert is a paid mutator transaction binding the contract method 0x97e2e256.
//
// Solidity: function pendingDisputableAssert(bytes32[4] _fields, uint32 _numSteps, uint64[2] _timeBounds, bytes21[] _tokenTypes, bytes32[] _messageDataHash, uint16[] _messageTokenNums, uint256[] _messageAmounts, address[] _messageDestinations) returns()
func (_VMTracker *VMTrackerTransactorSession) PendingDisputableAssert(_fields [4][32]byte, _numSteps uint32, _timeBounds [2]uint64, _tokenTypes [][21]byte, _messageDataHash [][32]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _messageDestinations []common.Address) (*types.Transaction, error) {
	return _VMTracker.Contract.PendingDisputableAssert(&_VMTracker.TransactOpts, _fields, _numSteps, _timeBounds, _tokenTypes, _messageDataHash, _messageTokenNums, _messageAmounts, _messageDestinations)
}

// PendingUnanimousAssert is a paid mutator transaction binding the contract method 0x675125b9.
//
// Solidity: function pendingUnanimousAssert(bytes32 _unanRest, bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts, uint64 _sequenceNum, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactor) PendingUnanimousAssert(opts *bind.TransactOpts, _unanRest [32]byte, _tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _sequenceNum uint64, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.contract.Transact(opts, "pendingUnanimousAssert", _unanRest, _tokenTypes, _messageTokenNums, _messageAmounts, _sequenceNum, _logsAccHash, _signatures)
}

// PendingUnanimousAssert is a paid mutator transaction binding the contract method 0x675125b9.
//
// Solidity: function pendingUnanimousAssert(bytes32 _unanRest, bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts, uint64 _sequenceNum, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerSession) PendingUnanimousAssert(_unanRest [32]byte, _tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _sequenceNum uint64, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.PendingUnanimousAssert(&_VMTracker.TransactOpts, _unanRest, _tokenTypes, _messageTokenNums, _messageAmounts, _sequenceNum, _logsAccHash, _signatures)
}

// PendingUnanimousAssert is a paid mutator transaction binding the contract method 0x675125b9.
//
// Solidity: function pendingUnanimousAssert(bytes32 _unanRest, bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts, uint64 _sequenceNum, bytes32 _logsAccHash, bytes _signatures) returns()
func (_VMTracker *VMTrackerTransactorSession) PendingUnanimousAssert(_unanRest [32]byte, _tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int, _sequenceNum uint64, _logsAccHash [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _VMTracker.Contract.PendingUnanimousAssert(&_VMTracker.TransactOpts, _unanRest, _tokenTypes, _messageTokenNums, _messageAmounts, _sequenceNum, _logsAccHash, _signatures)
}

// VMTrackerConfirmedDisputableAssertionIterator is returned from FilterConfirmedDisputableAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedDisputableAssertion events raised by the VMTracker contract.
type VMTrackerConfirmedDisputableAssertionIterator struct {
	Event *VMTrackerConfirmedDisputableAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VMTrackerConfirmedDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerConfirmedDisputableAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VMTrackerConfirmedDisputableAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VMTrackerConfirmedDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerConfirmedDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerConfirmedDisputableAssertion represents a ConfirmedDisputableAssertion event raised by the VMTracker contract.
type VMTrackerConfirmedDisputableAssertion struct {
	NewState    [32]byte
	LogsAccHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedDisputableAssertion is a free log retrieval operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_VMTracker *VMTrackerFilterer) FilterConfirmedDisputableAssertion(opts *bind.FilterOpts) (*VMTrackerConfirmedDisputableAssertionIterator, error) {

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &VMTrackerConfirmedDisputableAssertionIterator{contract: _VMTracker.contract, event: "ConfirmedDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedDisputableAssertion is a free log subscription operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_VMTracker *VMTrackerFilterer) WatchConfirmedDisputableAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerConfirmedDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "ConfirmedDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerConfirmedDisputableAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmedDisputableAssertion is a log parse operation binding the contract event 0x4a4dac0badcc6a19561138f43003082ff9638757afa521c1ed29832cd410a8bb.
//
// Solidity: event ConfirmedDisputableAssertion(bytes32 newState, bytes32 logsAccHash)
func (_VMTracker *VMTrackerFilterer) ParseConfirmedDisputableAssertion(log types.Log) (*VMTrackerConfirmedDisputableAssertion, error) {
	event := new(VMTrackerConfirmedDisputableAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "ConfirmedDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMTrackerConfirmedUnanimousAssertionIterator is returned from FilterConfirmedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedUnanimousAssertion events raised by the VMTracker contract.
type VMTrackerConfirmedUnanimousAssertionIterator struct {
	Event *VMTrackerConfirmedUnanimousAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VMTrackerConfirmedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerConfirmedUnanimousAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VMTrackerConfirmedUnanimousAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VMTrackerConfirmedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerConfirmedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerConfirmedUnanimousAssertion represents a ConfirmedUnanimousAssertion event raised by the VMTracker contract.
type VMTrackerConfirmedUnanimousAssertion struct {
	SequenceNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedUnanimousAssertion is a free log retrieval operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) FilterConfirmedUnanimousAssertion(opts *bind.FilterOpts) (*VMTrackerConfirmedUnanimousAssertionIterator, error) {

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "ConfirmedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &VMTrackerConfirmedUnanimousAssertionIterator{contract: _VMTracker.contract, event: "ConfirmedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedUnanimousAssertion is a free log subscription operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) WatchConfirmedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerConfirmedUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "ConfirmedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerConfirmedUnanimousAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmedUnanimousAssertion is a log parse operation binding the contract event 0xbecbda44e774b1f76ae07216c13391a8fd37cfef503e223f8ffa63c9a48630c2.
//
// Solidity: event ConfirmedUnanimousAssertion(uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) ParseConfirmedUnanimousAssertion(log types.Log) (*VMTrackerConfirmedUnanimousAssertion, error) {
	event := new(VMTrackerConfirmedUnanimousAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "ConfirmedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMTrackerFinalizedUnanimousAssertionIterator is returned from FilterFinalizedUnanimousAssertion and is used to iterate over the raw logs and unpacked data for FinalizedUnanimousAssertion events raised by the VMTracker contract.
type VMTrackerFinalizedUnanimousAssertionIterator struct {
	Event *VMTrackerFinalizedUnanimousAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VMTrackerFinalizedUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerFinalizedUnanimousAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VMTrackerFinalizedUnanimousAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VMTrackerFinalizedUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerFinalizedUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerFinalizedUnanimousAssertion represents a FinalizedUnanimousAssertion event raised by the VMTracker contract.
type VMTrackerFinalizedUnanimousAssertion struct {
	UnanHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizedUnanimousAssertion is a free log retrieval operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_VMTracker *VMTrackerFilterer) FilterFinalizedUnanimousAssertion(opts *bind.FilterOpts) (*VMTrackerFinalizedUnanimousAssertionIterator, error) {

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "FinalizedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &VMTrackerFinalizedUnanimousAssertionIterator{contract: _VMTracker.contract, event: "FinalizedUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchFinalizedUnanimousAssertion is a free log subscription operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_VMTracker *VMTrackerFilterer) WatchFinalizedUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerFinalizedUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "FinalizedUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerFinalizedUnanimousAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinalizedUnanimousAssertion is a log parse operation binding the contract event 0x709bbc35a8e7f8d3cf7fb672ff1e7b28dc84ff6ac29d940aeacc26f1aa463aaa.
//
// Solidity: event FinalizedUnanimousAssertion(bytes32 unanHash)
func (_VMTracker *VMTrackerFilterer) ParseFinalizedUnanimousAssertion(log types.Log) (*VMTrackerFinalizedUnanimousAssertion, error) {
	event := new(VMTrackerFinalizedUnanimousAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "FinalizedUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMTrackerInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the VMTracker contract.
type VMTrackerInitiatedChallengeIterator struct {
	Event *VMTrackerInitiatedChallenge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VMTrackerInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerInitiatedChallenge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VMTrackerInitiatedChallenge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VMTrackerInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerInitiatedChallenge represents a InitiatedChallenge event raised by the VMTracker contract.
type VMTrackerInitiatedChallenge struct {
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_VMTracker *VMTrackerFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts) (*VMTrackerInitiatedChallengeIterator, error) {

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return &VMTrackerInitiatedChallengeIterator{contract: _VMTracker.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_VMTracker *VMTrackerFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *VMTrackerInitiatedChallenge) (event.Subscription, error) {

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "InitiatedChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerInitiatedChallenge)
				if err := _VMTracker.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x255b033ec1fbcab46152fd2de20e846af1c65a63e0df0ee9c9cfe751fce2d2b7.
//
// Solidity: event InitiatedChallenge(address challenger)
func (_VMTracker *VMTrackerFilterer) ParseInitiatedChallenge(log types.Log) (*VMTrackerInitiatedChallenge, error) {
	event := new(VMTrackerInitiatedChallenge)
	if err := _VMTracker.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMTrackerPendingDisputableAssertionIterator is returned from FilterPendingDisputableAssertion and is used to iterate over the raw logs and unpacked data for PendingDisputableAssertion events raised by the VMTracker contract.
type VMTrackerPendingDisputableAssertionIterator struct {
	Event *VMTrackerPendingDisputableAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VMTrackerPendingDisputableAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerPendingDisputableAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VMTrackerPendingDisputableAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VMTrackerPendingDisputableAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerPendingDisputableAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerPendingDisputableAssertion represents a PendingDisputableAssertion event raised by the VMTracker contract.
type VMTrackerPendingDisputableAssertion struct {
	Fields          [3][32]byte
	Asserter        common.Address
	TimeBounds      [2]uint64
	TokenTypes      [][21]byte
	NumSteps        uint32
	LastMessageHash [32]byte
	LogsAccHash     [32]byte
	Amounts         []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPendingDisputableAssertion is a free log retrieval operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_VMTracker *VMTrackerFilterer) FilterPendingDisputableAssertion(opts *bind.FilterOpts) (*VMTrackerPendingDisputableAssertionIterator, error) {

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return &VMTrackerPendingDisputableAssertionIterator{contract: _VMTracker.contract, event: "PendingDisputableAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingDisputableAssertion is a free log subscription operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_VMTracker *VMTrackerFilterer) WatchPendingDisputableAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerPendingDisputableAssertion) (event.Subscription, error) {

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "PendingDisputableAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerPendingDisputableAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingDisputableAssertion is a log parse operation binding the contract event 0x5df9430f8c0d650b9ceabd2fbdfcaa42e31fd36a71c0bebdf0b47d966372d94f.
//
// Solidity: event PendingDisputableAssertion(bytes32[3] fields, address asserter, uint64[2] timeBounds, bytes21[] tokenTypes, uint32 numSteps, bytes32 lastMessageHash, bytes32 logsAccHash, uint256[] amounts)
func (_VMTracker *VMTrackerFilterer) ParsePendingDisputableAssertion(log types.Log) (*VMTrackerPendingDisputableAssertion, error) {
	event := new(VMTrackerPendingDisputableAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "PendingDisputableAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VMTrackerPendingUnanimousAssertionIterator is returned from FilterPendingUnanimousAssertion and is used to iterate over the raw logs and unpacked data for PendingUnanimousAssertion events raised by the VMTracker contract.
type VMTrackerPendingUnanimousAssertionIterator struct {
	Event *VMTrackerPendingUnanimousAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VMTrackerPendingUnanimousAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMTrackerPendingUnanimousAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VMTrackerPendingUnanimousAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VMTrackerPendingUnanimousAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMTrackerPendingUnanimousAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMTrackerPendingUnanimousAssertion represents a PendingUnanimousAssertion event raised by the VMTracker contract.
type VMTrackerPendingUnanimousAssertion struct {
	UnanHash    [32]byte
	SequenceNum uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingUnanimousAssertion is a free log retrieval operation binding the contract event 0xc87ab2402746691bbdb443504eab6563fce71e66d5c223f63b0af3442b20851d.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) FilterPendingUnanimousAssertion(opts *bind.FilterOpts) (*VMTrackerPendingUnanimousAssertionIterator, error) {

	logs, sub, err := _VMTracker.contract.FilterLogs(opts, "PendingUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return &VMTrackerPendingUnanimousAssertionIterator{contract: _VMTracker.contract, event: "PendingUnanimousAssertion", logs: logs, sub: sub}, nil
}

// WatchPendingUnanimousAssertion is a free log subscription operation binding the contract event 0xc87ab2402746691bbdb443504eab6563fce71e66d5c223f63b0af3442b20851d.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) WatchPendingUnanimousAssertion(opts *bind.WatchOpts, sink chan<- *VMTrackerPendingUnanimousAssertion) (event.Subscription, error) {

	logs, sub, err := _VMTracker.contract.WatchLogs(opts, "PendingUnanimousAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMTrackerPendingUnanimousAssertion)
				if err := _VMTracker.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingUnanimousAssertion is a log parse operation binding the contract event 0xc87ab2402746691bbdb443504eab6563fce71e66d5c223f63b0af3442b20851d.
//
// Solidity: event PendingUnanimousAssertion(bytes32 unanHash, uint64 sequenceNum)
func (_VMTracker *VMTrackerFilterer) ParsePendingUnanimousAssertion(log types.Log) (*VMTrackerPendingUnanimousAssertion, error) {
	event := new(VMTrackerPendingUnanimousAssertion)
	if err := _VMTracker.contract.UnpackLog(event, "PendingUnanimousAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}
