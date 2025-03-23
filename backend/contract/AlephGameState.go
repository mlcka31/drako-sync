// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AlephGameStateMessage is an auto generated low-level Go binding around an user-defined struct.
type AlephGameStateMessage struct {
	MessagePrice *big.Int
	Sender       common.Address
	Content      string
	Timestamp    *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_messagePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_coeffIncrease\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aiAgent\",\"type\":\"address\"}],\"name\":\"AIAgentSetAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AgentReplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizePayout\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aiAgentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coeffIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameState\",\"outputs\":[{\"internalType\":\"enumAlephGameState.GameState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messagePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structAlephGameState.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aiAgentAddress\",\"type\":\"address\"}],\"name\":\"initGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messagePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messagePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payoutPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prompt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"promptSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"reply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_prompt\",\"type\":\"string\"}],\"name\":\"setPrompt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopGameAndWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051611832380380611832833981016040819052602b916059565b600680546001600160a01b031916331790556003919091556004556001805461ff001916610300179055607a565b5f5f604083850312156069575f5ffd5b505080516020909101519092909150565b6117ab806100875f395ff3fe608060405260043610610105575f3560e01c8063719ce73e11610092578063b4ac024f11610062578063b4ac024f1461028f578063d1f9c24d146102ae578063de4c5f1a146102d9578063fb29d20a146102ed578063fc6f946814610324575f5ffd5b8063719ce73e146102295780637c65d7111461023e578063aaf5eb681461025f578063acce9c241461027a575f5ffd5b806354cc76a7116100d857806354cc76a71461017e5780635ebe7c72146101a15780635ff6cbf3146101c05780636a5568be146101e1578063708f8b4b14610200575f5ffd5b80630d80fefd1461010957806313d25eda1461014157806335a063b414610157578063469c81101461016b575b5f5ffd5b348015610114575f5ffd5b506101286101233660046112b3565b610343565b60405161013894939291906112f8565b60405180910390f35b34801561014c575f5ffd5b5061015561040e565b005b348015610162575f5ffd5b50610155610730565b610155610179366004611346565b6109b6565b348015610189575f5ffd5b5061019360035481565b604051908152602001610138565b3480156101ac575f5ffd5b506101556101bb366004611346565b610bec565b3480156101cb575f5ffd5b506101d4610e05565b60405161013891906113f9565b3480156101ec575f5ffd5b506101556101fb366004611490565b610f1c565b34801561020b575f5ffd5b506001546102199060ff1681565b6040519015158152602001610138565b348015610234575f5ffd5b5061019360025481565b348015610249575f5ffd5b50610252611005565b60405161013891906114bd565b34801561026a575f5ffd5b50610193670de0b6b3a764000081565b348015610285575f5ffd5b5061019360045481565b34801561029a575f5ffd5b506101556102a9366004611346565b611090565b3480156102b9575f5ffd5b506001546102cc90610100900460ff1681565b60405161013891906114e3565b3480156102e4575f5ffd5b50610155611128565b3480156102f8575f5ffd5b5060075461030c906001600160a01b031681565b6040516001600160a01b039091168152602001610138565b34801561032f575f5ffd5b5060065461030c906001600160a01b031681565b60058181548110610352575f80fd5b5f9182526020909120600490910201805460018201546002830180549294506001600160a01b03909116929161038790611509565b80601f01602080910402602001604051908101604052809291908181526020018280546103b390611509565b80156103fe5780601f106103d5576101008083540402835291602001916103fe565b820191905f5260205f20905b8154815290600101906020018083116103e157829003601f168201915b5050505050908060030154905084565b6007546001600160a01b031633146104415760405162461bcd60e51b815260040161043890611541565b60405180910390fd5b6002600154610100900460ff16600381111561045f5761045f6114cf565b036104ac5760405162461bcd60e51b815260206004820152601a60248201527f47616d65206973206e6f742070656e64696e67207061796f75740000000000006044820152606401610438565b6003600154610100900460ff1660038111156104ca576104ca6114cf565b036105175760405162461bcd60e51b815260206004820152601a60248201527f47616d65206973206e6f742070656e64696e67207061796f75740000000000006044820152606401610438565b60055461055e5760405162461bcd60e51b81526020600482015260156024820152744e6f206d6573736167657320617661696c61626c6560581b6044820152606401610438565b600754600580545f926001600160a01b0316919061057e90600190611596565b8154811061058e5761058e6115af565b5f9182526020909120600160049092020101546001600160a01b03161480156105b957506005546001105b1561060157600580546105ce90600290611596565b815481106105de576105de6115af565b5f9182526020909120600160049092020101546001600160a01b03169050610640565b6005805461061190600190611596565b81548110610621576106216115af565b5f9182526020909120600160049092020101546001600160a01b031690505b600280545f918290556001805461ff0019166102001790556040519091906001600160a01b0384169083908381818185875af1925050503d805f81146106a1576040519150601f19603f3d011682016040523d82523d5f602084013e6106a6565b606091505b50509050806106e95760405162461bcd60e51b815260206004820152600f60248201526e151c985b9cd9995c8819985a5b1959608a1b6044820152606401610438565b604080516001600160a01b0385168152602081018490527f504d6fcd7fd6584455567b4c515b7efa2a27fb623104bd49f98b4e9a2ccd7efb910160405180910390a1505050565b6006546001600160a01b0316331461075a5760405162461bcd60e51b8152600401610438906115c3565b5f805b6005548110156108ef57600754600580546001600160a01b03909216918390811061078a5761078a6115af565b5f9182526020909120600160049092020101546001600160a01b0316148015906107f15750600654600580546001600160a01b0390921691839081106107d2576107d26115af565b5f9182526020909120600160049092020101546001600160a01b031614155b156108e7575f6005828154811061080a5761080a6115af565b5f91825260209091206004909102015490506108268184611604565b92505f6005838154811061083c5761083c6115af565b5f91825260208220600490910201600101546040516001600160a01b039091169184919081818185875af1925050503d805f8114610895576040519150601f19603f3d011682016040523d82523d5f602084013e61089a565b606091505b50509050806108e45760405162461bcd60e51b81526020600482015260166024820152751499599d5b99081d1c985b9cd9995c8819985a5b195960521b6044820152606401610438565b50505b60010161075d565b505f816002546108ff9190611596565b5f600281905560065460405192935090916001600160a01b039091169083908381818185875af1925050503d805f8114610954576040519150601f19603f3d011682016040523d82523d5f602084013e610959565b606091505b50509050806109a25760405162461bcd60e51b815260206004820152601560248201527410591b5a5b881d1c985b9cd9995c8819985a5b1959605a1b6044820152606401610438565b50506001805461ff00191661020017905550565b5f600154610100900460ff1660038111156109d3576109d36114cf565b14610a325760405162461bcd60e51b815260206004820152602960248201527f47616d65206973206e6f7420616374697665206f72206974206973206e6f74206044820152683cb7bab9103a3ab93760b91b6064820152608401610438565b6003543414610a795760405162461bcd60e51b815260206004820152601360248201527257726f6e67206d65737361676520707269636560681b6044820152606401610438565b6040805160808101825260035481523360208201908152918101838152426060830152600580546001810182555f9190915282517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0600490920291820190815593517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db1820180546001600160a01b0319166001600160a01b0390921691909117905590519192917f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db290910190610b4f9082611663565b506060820151816003015550503460025f828254610b6d9190611604565b9091555050600454600354670de0b6b3a764000091610b8b9161171e565b610b959190611735565b60035560405133907f4816bdd3b9d39acdd4da0ae7412d5523081e916c6a56e481063e475b86ef7c5b90610bcc9084904290611754565b60405180910390a260018054819061ff001916610100825b021790555050565b6007546001600160a01b03163314610c165760405162461bcd60e51b815260040161043890611541565b6007546001600160a01b03163314610c705760405162461bcd60e51b815260206004820152601760248201527f4f6e6c79204149206167656e742063616e207265706c790000000000000000006044820152606401610438565b60018054610100900460ff166003811115610c8d57610c8d6114cf565b14610cda5760405162461bcd60e51b815260206004820152601a60248201527f47616d65206973206e6f7420696e20676f696e672073746174650000000000006044820152606401610438565b604080516080810182525f8082526007546001600160a01b039081166020840190815293830185815242606085015260058054600181018255935283517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0600490940293840190815594517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db1840180546001600160a01b0319169190931617909155519192917f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db290910190610daf9082611663565b50606091909101516003909101556001805461ff00191690556040517f5aa9063abb57f078ec42ba6732d0c8631e1cb79dba57ba9b74b76bc0658755b690610dfa9083904290611754565b60405180910390a150565b60606005805480602002602001604051908101604052809291908181526020015f905b82821015610f13575f84815260209081902060408051608081018252600486029092018054835260018101546001600160a01b03169383019390935260028301805492939291840191610e7a90611509565b80601f0160208091040260200160405190810160405280929190818152602001828054610ea690611509565b8015610ef15780601f10610ec857610100808354040283529160200191610ef1565b820191905f5260205f20905b815481529060010190602001808311610ed457829003601f168201915b5050505050815260200160038201548152505081526020019060010190610e28565b50505050905090565b6006546001600160a01b03163314610f465760405162461bcd60e51b8152600401610438906115c3565b6007546001600160a01b031615610f9f5760405162461bcd60e51b815260206004820152601860248201527f47616d6520616c726561647920696e697469616c697a656400000000000000006044820152606401610438565b600780546001600160a01b0319166001600160a01b0383169081179091556040519081527fcb3ed5b60a03dc79bfa565ea813bd93572e94f8666e5b22f643b39d2e5f029309060200160405180910390a1600180545f919061ff00191661010083610be4565b5f805461101190611509565b80601f016020809104026020016040519081016040528092919081815260200182805461103d90611509565b80156110885780601f1061105f57610100808354040283529160200191611088565b820191905f5260205f20905b81548152906001019060200180831161106b57829003601f168201915b505050505081565b6007546001600160a01b031633146110ba5760405162461bcd60e51b815260040161043890611541565b60015460ff161561110d5760405162461bcd60e51b815260206004820152601b60248201527f50726f6d70742063616e206f6e6c7920626520736574206f6e636500000000006044820152606401610438565b5f6111188282611663565b50506001805460ff191681179055565b6006546001600160a01b031633146111525760405162461bcd60e51b8152600401610438906115c3565b6005546111985760405162461bcd60e51b81526020600482015260146024820152734e6f206d6573736167657320746f20636865636b60601b6044820152606401610438565b600580545f91906111ab90600190611596565b815481106111bb576111bb6115af565b905f5260205f2090600402016003015490508062093a806111dc9190611604565b4210156112225760405162461bcd60e51b815260206004820152601460248201527310d85b9b9bdd081cdd1bdc0819d85b59481e595d60621b6044820152606401610438565b600280545f9182905560065460405191926001600160a01b039091169183156108fc0291849190818181858888f19350505050158015611264573d5f5f3e3d5ffd5b506001805461020061ff001990911617905560408051338152602081018390527f504d6fcd7fd6584455567b4c515b7efa2a27fb623104bd49f98b4e9a2ccd7efb910160405180910390a15050565b5f602082840312156112c3575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b8481526001600160a01b03841660208201526080604082018190525f90611321908301856112ca565b905082606083015295945050505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215611356575f5ffd5b813567ffffffffffffffff81111561136c575f5ffd5b8201601f8101841361137c575f5ffd5b803567ffffffffffffffff81111561139657611396611332565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156113c5576113c5611332565b6040528181528282016020018610156113dc575f5ffd5b816020840160208301375f91810160200191909152949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561148457603f1987860301845281518051865260018060a01b03602082015116602087015260408101516080604088015261146160808801826112ca565b60609283015197909201969096529450602093840193919091019060010161141f565b50929695505050505050565b5f602082840312156114a0575f5ffd5b81356001600160a01b03811681146114b6575f5ffd5b9392505050565b602081525f6114b660208301846112ca565b634e487b7160e01b5f52602160045260245ffd5b602081016004831061150357634e487b7160e01b5f52602160045260245ffd5b91905290565b600181811c9082168061151d57607f821691505b60208210810361153b57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526021908201527f4f6e6c79206167656e742063616e2063616c6c20746869732066756e6374696f6040820152603760f91b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b818103818111156115a9576115a9611582565b92915050565b634e487b7160e01b5f52603260045260245ffd5b60208082526021908201527f4f6e6c792061646d696e2063616e2063616c6c20746869732066756e6374696f6040820152603760f91b606082015260800190565b808201808211156115a9576115a9611582565b601f82111561165e57805f5260205f20601f840160051c8101602085101561163c5750805b601f840160051c820191505b8181101561165b575f8155600101611648565b50505b505050565b815167ffffffffffffffff81111561167d5761167d611332565b6116918161168b8454611509565b84611617565b6020601f8211600181146116c3575f83156116ac5750848201515b5f19600385901b1c1916600184901b17845561165b565b5f84815260208120601f198516915b828110156116f257878501518255602094850194600190920191016116d2565b508482101561170f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b80820281158282048414176115a9576115a9611582565b5f8261174f57634e487b7160e01b5f52601260045260245ffd5b500490565b604081525f61176660408301856112ca565b9050826020830152939250505056fea26469706673582212207c506c117d2d1518e35bb1e68997fd2a65825d4839abf18a9be940df40e6f68464736f6c634300081d0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _messagePrice *big.Int, _coeffIncrease *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _messagePrice, _coeffIncrease)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Contract *ContractCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Contract *ContractSession) PRECISION() (*big.Int, error) {
	return _Contract.Contract.PRECISION(&_Contract.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Contract *ContractCallerSession) PRECISION() (*big.Int, error) {
	return _Contract.Contract.PRECISION(&_Contract.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_Contract *ContractCaller) AdminAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "adminAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_Contract *ContractSession) AdminAddress() (common.Address, error) {
	return _Contract.Contract.AdminAddress(&_Contract.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_Contract *ContractCallerSession) AdminAddress() (common.Address, error) {
	return _Contract.Contract.AdminAddress(&_Contract.CallOpts)
}

// AiAgentAddress is a free data retrieval call binding the contract method 0xfb29d20a.
//
// Solidity: function aiAgentAddress() view returns(address)
func (_Contract *ContractCaller) AiAgentAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "aiAgentAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AiAgentAddress is a free data retrieval call binding the contract method 0xfb29d20a.
//
// Solidity: function aiAgentAddress() view returns(address)
func (_Contract *ContractSession) AiAgentAddress() (common.Address, error) {
	return _Contract.Contract.AiAgentAddress(&_Contract.CallOpts)
}

// AiAgentAddress is a free data retrieval call binding the contract method 0xfb29d20a.
//
// Solidity: function aiAgentAddress() view returns(address)
func (_Contract *ContractCallerSession) AiAgentAddress() (common.Address, error) {
	return _Contract.Contract.AiAgentAddress(&_Contract.CallOpts)
}

// CoeffIncrease is a free data retrieval call binding the contract method 0xacce9c24.
//
// Solidity: function coeffIncrease() view returns(uint256)
func (_Contract *ContractCaller) CoeffIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "coeffIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoeffIncrease is a free data retrieval call binding the contract method 0xacce9c24.
//
// Solidity: function coeffIncrease() view returns(uint256)
func (_Contract *ContractSession) CoeffIncrease() (*big.Int, error) {
	return _Contract.Contract.CoeffIncrease(&_Contract.CallOpts)
}

// CoeffIncrease is a free data retrieval call binding the contract method 0xacce9c24.
//
// Solidity: function coeffIncrease() view returns(uint256)
func (_Contract *ContractCallerSession) CoeffIncrease() (*big.Int, error) {
	return _Contract.Contract.CoeffIncrease(&_Contract.CallOpts)
}

// GameState is a free data retrieval call binding the contract method 0xd1f9c24d.
//
// Solidity: function gameState() view returns(uint8)
func (_Contract *ContractCaller) GameState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "gameState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GameState is a free data retrieval call binding the contract method 0xd1f9c24d.
//
// Solidity: function gameState() view returns(uint8)
func (_Contract *ContractSession) GameState() (uint8, error) {
	return _Contract.Contract.GameState(&_Contract.CallOpts)
}

// GameState is a free data retrieval call binding the contract method 0xd1f9c24d.
//
// Solidity: function gameState() view returns(uint8)
func (_Contract *ContractCallerSession) GameState() (uint8, error) {
	return _Contract.Contract.GameState(&_Contract.CallOpts)
}

// GetMessages is a free data retrieval call binding the contract method 0x5ff6cbf3.
//
// Solidity: function getMessages() view returns((uint256,address,string,uint256)[])
func (_Contract *ContractCaller) GetMessages(opts *bind.CallOpts) ([]AlephGameStateMessage, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMessages")

	if err != nil {
		return *new([]AlephGameStateMessage), err
	}

	out0 := *abi.ConvertType(out[0], new([]AlephGameStateMessage)).(*[]AlephGameStateMessage)

	return out0, err

}

// GetMessages is a free data retrieval call binding the contract method 0x5ff6cbf3.
//
// Solidity: function getMessages() view returns((uint256,address,string,uint256)[])
func (_Contract *ContractSession) GetMessages() ([]AlephGameStateMessage, error) {
	return _Contract.Contract.GetMessages(&_Contract.CallOpts)
}

// GetMessages is a free data retrieval call binding the contract method 0x5ff6cbf3.
//
// Solidity: function getMessages() view returns((uint256,address,string,uint256)[])
func (_Contract *ContractCallerSession) GetMessages() ([]AlephGameStateMessage, error) {
	return _Contract.Contract.GetMessages(&_Contract.CallOpts)
}

// MessagePrice is a free data retrieval call binding the contract method 0x54cc76a7.
//
// Solidity: function messagePrice() view returns(uint256)
func (_Contract *ContractCaller) MessagePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "messagePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessagePrice is a free data retrieval call binding the contract method 0x54cc76a7.
//
// Solidity: function messagePrice() view returns(uint256)
func (_Contract *ContractSession) MessagePrice() (*big.Int, error) {
	return _Contract.Contract.MessagePrice(&_Contract.CallOpts)
}

// MessagePrice is a free data retrieval call binding the contract method 0x54cc76a7.
//
// Solidity: function messagePrice() view returns(uint256)
func (_Contract *ContractCallerSession) MessagePrice() (*big.Int, error) {
	return _Contract.Contract.MessagePrice(&_Contract.CallOpts)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(uint256 messagePrice, address sender, string content, uint256 timestamp)
func (_Contract *ContractCaller) Messages(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MessagePrice *big.Int
	Sender       common.Address
	Content      string
	Timestamp    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "messages", arg0)

	outstruct := new(struct {
		MessagePrice *big.Int
		Sender       common.Address
		Content      string
		Timestamp    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MessagePrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Sender = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Content = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(uint256 messagePrice, address sender, string content, uint256 timestamp)
func (_Contract *ContractSession) Messages(arg0 *big.Int) (struct {
	MessagePrice *big.Int
	Sender       common.Address
	Content      string
	Timestamp    *big.Int
}, error) {
	return _Contract.Contract.Messages(&_Contract.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(uint256 messagePrice, address sender, string content, uint256 timestamp)
func (_Contract *ContractCallerSession) Messages(arg0 *big.Int) (struct {
	MessagePrice *big.Int
	Sender       common.Address
	Content      string
	Timestamp    *big.Int
}, error) {
	return _Contract.Contract.Messages(&_Contract.CallOpts, arg0)
}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(uint256)
func (_Contract *ContractCaller) PrizePool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "prizePool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(uint256)
func (_Contract *ContractSession) PrizePool() (*big.Int, error) {
	return _Contract.Contract.PrizePool(&_Contract.CallOpts)
}

// PrizePool is a free data retrieval call binding the contract method 0x719ce73e.
//
// Solidity: function prizePool() view returns(uint256)
func (_Contract *ContractCallerSession) PrizePool() (*big.Int, error) {
	return _Contract.Contract.PrizePool(&_Contract.CallOpts)
}

// Prompt is a free data retrieval call binding the contract method 0x7c65d711.
//
// Solidity: function prompt() view returns(string)
func (_Contract *ContractCaller) Prompt(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "prompt")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Prompt is a free data retrieval call binding the contract method 0x7c65d711.
//
// Solidity: function prompt() view returns(string)
func (_Contract *ContractSession) Prompt() (string, error) {
	return _Contract.Contract.Prompt(&_Contract.CallOpts)
}

// Prompt is a free data retrieval call binding the contract method 0x7c65d711.
//
// Solidity: function prompt() view returns(string)
func (_Contract *ContractCallerSession) Prompt() (string, error) {
	return _Contract.Contract.Prompt(&_Contract.CallOpts)
}

// PromptSet is a free data retrieval call binding the contract method 0x708f8b4b.
//
// Solidity: function promptSet() view returns(bool)
func (_Contract *ContractCaller) PromptSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "promptSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PromptSet is a free data retrieval call binding the contract method 0x708f8b4b.
//
// Solidity: function promptSet() view returns(bool)
func (_Contract *ContractSession) PromptSet() (bool, error) {
	return _Contract.Contract.PromptSet(&_Contract.CallOpts)
}

// PromptSet is a free data retrieval call binding the contract method 0x708f8b4b.
//
// Solidity: function promptSet() view returns(bool)
func (_Contract *ContractCallerSession) PromptSet() (bool, error) {
	return _Contract.Contract.PromptSet(&_Contract.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Contract *ContractTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Contract *ContractSession) Abort() (*types.Transaction, error) {
	return _Contract.Contract.Abort(&_Contract.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Contract *ContractTransactorSession) Abort() (*types.Transaction, error) {
	return _Contract.Contract.Abort(&_Contract.TransactOpts)
}

// InitGame is a paid mutator transaction binding the contract method 0x6a5568be.
//
// Solidity: function initGame(address _aiAgentAddress) returns()
func (_Contract *ContractTransactor) InitGame(opts *bind.TransactOpts, _aiAgentAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initGame", _aiAgentAddress)
}

// InitGame is a paid mutator transaction binding the contract method 0x6a5568be.
//
// Solidity: function initGame(address _aiAgentAddress) returns()
func (_Contract *ContractSession) InitGame(_aiAgentAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.InitGame(&_Contract.TransactOpts, _aiAgentAddress)
}

// InitGame is a paid mutator transaction binding the contract method 0x6a5568be.
//
// Solidity: function initGame(address _aiAgentAddress) returns()
func (_Contract *ContractTransactorSession) InitGame(_aiAgentAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.InitGame(&_Contract.TransactOpts, _aiAgentAddress)
}

// PayoutPrize is a paid mutator transaction binding the contract method 0x13d25eda.
//
// Solidity: function payoutPrize() returns()
func (_Contract *ContractTransactor) PayoutPrize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "payoutPrize")
}

// PayoutPrize is a paid mutator transaction binding the contract method 0x13d25eda.
//
// Solidity: function payoutPrize() returns()
func (_Contract *ContractSession) PayoutPrize() (*types.Transaction, error) {
	return _Contract.Contract.PayoutPrize(&_Contract.TransactOpts)
}

// PayoutPrize is a paid mutator transaction binding the contract method 0x13d25eda.
//
// Solidity: function payoutPrize() returns()
func (_Contract *ContractTransactorSession) PayoutPrize() (*types.Transaction, error) {
	return _Contract.Contract.PayoutPrize(&_Contract.TransactOpts)
}

// Reply is a paid mutator transaction binding the contract method 0x5ebe7c72.
//
// Solidity: function reply(string _content) returns()
func (_Contract *ContractTransactor) Reply(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reply", _content)
}

// Reply is a paid mutator transaction binding the contract method 0x5ebe7c72.
//
// Solidity: function reply(string _content) returns()
func (_Contract *ContractSession) Reply(_content string) (*types.Transaction, error) {
	return _Contract.Contract.Reply(&_Contract.TransactOpts, _content)
}

// Reply is a paid mutator transaction binding the contract method 0x5ebe7c72.
//
// Solidity: function reply(string _content) returns()
func (_Contract *ContractTransactorSession) Reply(_content string) (*types.Transaction, error) {
	return _Contract.Contract.Reply(&_Contract.TransactOpts, _content)
}

// SendMessage is a paid mutator transaction binding the contract method 0x469c8110.
//
// Solidity: function sendMessage(string _content) payable returns()
func (_Contract *ContractTransactor) SendMessage(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sendMessage", _content)
}

// SendMessage is a paid mutator transaction binding the contract method 0x469c8110.
//
// Solidity: function sendMessage(string _content) payable returns()
func (_Contract *ContractSession) SendMessage(_content string) (*types.Transaction, error) {
	return _Contract.Contract.SendMessage(&_Contract.TransactOpts, _content)
}

// SendMessage is a paid mutator transaction binding the contract method 0x469c8110.
//
// Solidity: function sendMessage(string _content) payable returns()
func (_Contract *ContractTransactorSession) SendMessage(_content string) (*types.Transaction, error) {
	return _Contract.Contract.SendMessage(&_Contract.TransactOpts, _content)
}

// SetPrompt is a paid mutator transaction binding the contract method 0xb4ac024f.
//
// Solidity: function setPrompt(string _prompt) returns()
func (_Contract *ContractTransactor) SetPrompt(opts *bind.TransactOpts, _prompt string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPrompt", _prompt)
}

// SetPrompt is a paid mutator transaction binding the contract method 0xb4ac024f.
//
// Solidity: function setPrompt(string _prompt) returns()
func (_Contract *ContractSession) SetPrompt(_prompt string) (*types.Transaction, error) {
	return _Contract.Contract.SetPrompt(&_Contract.TransactOpts, _prompt)
}

// SetPrompt is a paid mutator transaction binding the contract method 0xb4ac024f.
//
// Solidity: function setPrompt(string _prompt) returns()
func (_Contract *ContractTransactorSession) SetPrompt(_prompt string) (*types.Transaction, error) {
	return _Contract.Contract.SetPrompt(&_Contract.TransactOpts, _prompt)
}

// StopGameAndWithdraw is a paid mutator transaction binding the contract method 0xde4c5f1a.
//
// Solidity: function stopGameAndWithdraw() returns()
func (_Contract *ContractTransactor) StopGameAndWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stopGameAndWithdraw")
}

// StopGameAndWithdraw is a paid mutator transaction binding the contract method 0xde4c5f1a.
//
// Solidity: function stopGameAndWithdraw() returns()
func (_Contract *ContractSession) StopGameAndWithdraw() (*types.Transaction, error) {
	return _Contract.Contract.StopGameAndWithdraw(&_Contract.TransactOpts)
}

// StopGameAndWithdraw is a paid mutator transaction binding the contract method 0xde4c5f1a.
//
// Solidity: function stopGameAndWithdraw() returns()
func (_Contract *ContractTransactorSession) StopGameAndWithdraw() (*types.Transaction, error) {
	return _Contract.Contract.StopGameAndWithdraw(&_Contract.TransactOpts)
}

// ContractAIAgentSetAddressIterator is returned from FilterAIAgentSetAddress and is used to iterate over the raw logs and unpacked data for AIAgentSetAddress events raised by the Contract contract.
type ContractAIAgentSetAddressIterator struct {
	Event *ContractAIAgentSetAddress // Event containing the contract specifics and raw log

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
func (it *ContractAIAgentSetAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAIAgentSetAddress)
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
		it.Event = new(ContractAIAgentSetAddress)
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
func (it *ContractAIAgentSetAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAIAgentSetAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAIAgentSetAddress represents a AIAgentSetAddress event raised by the Contract contract.
type ContractAIAgentSetAddress struct {
	AiAgent common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAIAgentSetAddress is a free log retrieval operation binding the contract event 0xcb3ed5b60a03dc79bfa565ea813bd93572e94f8666e5b22f643b39d2e5f02930.
//
// Solidity: event AIAgentSetAddress(address aiAgent)
func (_Contract *ContractFilterer) FilterAIAgentSetAddress(opts *bind.FilterOpts) (*ContractAIAgentSetAddressIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AIAgentSetAddress")
	if err != nil {
		return nil, err
	}
	return &ContractAIAgentSetAddressIterator{contract: _Contract.contract, event: "AIAgentSetAddress", logs: logs, sub: sub}, nil
}

// WatchAIAgentSetAddress is a free log subscription operation binding the contract event 0xcb3ed5b60a03dc79bfa565ea813bd93572e94f8666e5b22f643b39d2e5f02930.
//
// Solidity: event AIAgentSetAddress(address aiAgent)
func (_Contract *ContractFilterer) WatchAIAgentSetAddress(opts *bind.WatchOpts, sink chan<- *ContractAIAgentSetAddress) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AIAgentSetAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAIAgentSetAddress)
				if err := _Contract.contract.UnpackLog(event, "AIAgentSetAddress", log); err != nil {
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

// ParseAIAgentSetAddress is a log parse operation binding the contract event 0xcb3ed5b60a03dc79bfa565ea813bd93572e94f8666e5b22f643b39d2e5f02930.
//
// Solidity: event AIAgentSetAddress(address aiAgent)
func (_Contract *ContractFilterer) ParseAIAgentSetAddress(log types.Log) (*ContractAIAgentSetAddress, error) {
	event := new(ContractAIAgentSetAddress)
	if err := _Contract.contract.UnpackLog(event, "AIAgentSetAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAgentRepliedIterator is returned from FilterAgentReplied and is used to iterate over the raw logs and unpacked data for AgentReplied events raised by the Contract contract.
type ContractAgentRepliedIterator struct {
	Event *ContractAgentReplied // Event containing the contract specifics and raw log

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
func (it *ContractAgentRepliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAgentReplied)
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
		it.Event = new(ContractAgentReplied)
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
func (it *ContractAgentRepliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAgentRepliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAgentReplied represents a AgentReplied event raised by the Contract contract.
type ContractAgentReplied struct {
	Content   string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentReplied is a free log retrieval operation binding the contract event 0x5aa9063abb57f078ec42ba6732d0c8631e1cb79dba57ba9b74b76bc0658755b6.
//
// Solidity: event AgentReplied(string content, uint256 timestamp)
func (_Contract *ContractFilterer) FilterAgentReplied(opts *bind.FilterOpts) (*ContractAgentRepliedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AgentReplied")
	if err != nil {
		return nil, err
	}
	return &ContractAgentRepliedIterator{contract: _Contract.contract, event: "AgentReplied", logs: logs, sub: sub}, nil
}

// WatchAgentReplied is a free log subscription operation binding the contract event 0x5aa9063abb57f078ec42ba6732d0c8631e1cb79dba57ba9b74b76bc0658755b6.
//
// Solidity: event AgentReplied(string content, uint256 timestamp)
func (_Contract *ContractFilterer) WatchAgentReplied(opts *bind.WatchOpts, sink chan<- *ContractAgentReplied) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AgentReplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAgentReplied)
				if err := _Contract.contract.UnpackLog(event, "AgentReplied", log); err != nil {
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

// ParseAgentReplied is a log parse operation binding the contract event 0x5aa9063abb57f078ec42ba6732d0c8631e1cb79dba57ba9b74b76bc0658755b6.
//
// Solidity: event AgentReplied(string content, uint256 timestamp)
func (_Contract *ContractFilterer) ParseAgentReplied(log types.Log) (*ContractAgentReplied, error) {
	event := new(ContractAgentReplied)
	if err := _Contract.contract.UnpackLog(event, "AgentReplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the Contract contract.
type ContractMessageSentIterator struct {
	Event *ContractMessageSent // Event containing the contract specifics and raw log

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
func (it *ContractMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMessageSent)
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
		it.Event = new(ContractMessageSent)
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
func (it *ContractMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMessageSent represents a MessageSent event raised by the Contract contract.
type ContractMessageSent struct {
	Sender    common.Address
	Content   string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x4816bdd3b9d39acdd4da0ae7412d5523081e916c6a56e481063e475b86ef7c5b.
//
// Solidity: event MessageSent(address indexed sender, string content, uint256 timestamp)
func (_Contract *ContractFilterer) FilterMessageSent(opts *bind.FilterOpts, sender []common.Address) (*ContractMessageSentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MessageSent", senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractMessageSentIterator{contract: _Contract.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x4816bdd3b9d39acdd4da0ae7412d5523081e916c6a56e481063e475b86ef7c5b.
//
// Solidity: event MessageSent(address indexed sender, string content, uint256 timestamp)
func (_Contract *ContractFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ContractMessageSent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MessageSent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMessageSent)
				if err := _Contract.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x4816bdd3b9d39acdd4da0ae7412d5523081e916c6a56e481063e475b86ef7c5b.
//
// Solidity: event MessageSent(address indexed sender, string content, uint256 timestamp)
func (_Contract *ContractFilterer) ParseMessageSent(log types.Log) (*ContractMessageSent, error) {
	event := new(ContractMessageSent)
	if err := _Contract.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPrizePayoutIterator is returned from FilterPrizePayout and is used to iterate over the raw logs and unpacked data for PrizePayout events raised by the Contract contract.
type ContractPrizePayoutIterator struct {
	Event *ContractPrizePayout // Event containing the contract specifics and raw log

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
func (it *ContractPrizePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPrizePayout)
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
		it.Event = new(ContractPrizePayout)
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
func (it *ContractPrizePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPrizePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPrizePayout represents a PrizePayout event raised by the Contract contract.
type ContractPrizePayout struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPrizePayout is a free log retrieval operation binding the contract event 0x504d6fcd7fd6584455567b4c515b7efa2a27fb623104bd49f98b4e9a2ccd7efb.
//
// Solidity: event PrizePayout(address winner, uint256 amount)
func (_Contract *ContractFilterer) FilterPrizePayout(opts *bind.FilterOpts) (*ContractPrizePayoutIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PrizePayout")
	if err != nil {
		return nil, err
	}
	return &ContractPrizePayoutIterator{contract: _Contract.contract, event: "PrizePayout", logs: logs, sub: sub}, nil
}

// WatchPrizePayout is a free log subscription operation binding the contract event 0x504d6fcd7fd6584455567b4c515b7efa2a27fb623104bd49f98b4e9a2ccd7efb.
//
// Solidity: event PrizePayout(address winner, uint256 amount)
func (_Contract *ContractFilterer) WatchPrizePayout(opts *bind.WatchOpts, sink chan<- *ContractPrizePayout) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PrizePayout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPrizePayout)
				if err := _Contract.contract.UnpackLog(event, "PrizePayout", log); err != nil {
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

// ParsePrizePayout is a log parse operation binding the contract event 0x504d6fcd7fd6584455567b4c515b7efa2a27fb623104bd49f98b4e9a2ccd7efb.
//
// Solidity: event PrizePayout(address winner, uint256 amount)
func (_Contract *ContractFilterer) ParsePrizePayout(log types.Log) (*ContractPrizePayout, error) {
	event := new(ContractPrizePayout)
	if err := _Contract.contract.UnpackLog(event, "PrizePayout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
