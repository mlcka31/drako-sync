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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_messagePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_coeffIncrease\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aiAgent\",\"type\":\"address\"}],\"name\":\"AIAgentSetAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AgentReplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizePayout\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aiAgentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coeffIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameState\",\"outputs\":[{\"internalType\":\"enumAlephGameState.GameState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCoeffIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGameState\",\"outputs\":[{\"internalType\":\"enumAlephGameState.GameState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messagePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structAlephGameState.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrecision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrizePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getmessagePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aiAgentAddress\",\"type\":\"address\"}],\"name\":\"initGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messagePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messagePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payoutPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"reply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aiAgentAddress\",\"type\":\"address\"}],\"name\":\"setAIAgentAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminAddress\",\"type\":\"address\"}],\"name\":\"setAdminAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_coeffIncrease\",\"type\":\"uint256\"}],\"name\":\"setCoeffIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumAlephGameState.GameState\",\"name\":\"_gameState\",\"type\":\"uint8\"}],\"name\":\"setGameState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_prizePool\",\"type\":\"uint256\"}],\"name\":\"setPrizePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_messagePrice\",\"type\":\"uint256\"}],\"name\":\"setmessagePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060405161141e38038061141e833981016040819052602b916058565b600580546001600160a01b0319163317905560029190915560039081555f805460ff191690911790556079565b5f5f604083850312156068575f5ffd5b505080516020909101519092909150565b611398806100865f395ff3fe608060405260043610610147575f3560e01c8063719ce73e116100b3578063aaf5eb681161006d578063aaf5eb6814610353578063acce9c241461036e578063b7d0628b14610383578063d1f9c24d146103a2578063fb29d20a146103ba578063fc6f9468146103f1575f5ffd5b8063719ce73e146102c957806375139aa0146102de578063884bf67c146102f25780639670c0bc14610306578063a3e2db8d14610320578063a8a68b4f1461033f575f5ffd5b806354cc76a71161010457806354cc76a7146102095780635ebe7c721461022c5780635ff6cbf31461024b5780636a4966b81461026c5780636a5568be1461028b5780636dd047f3146102aa575f5ffd5b80630d80fefd1461014b57806313d25eda1461018357806318172d21146101995780632a60c8fb146101b85780632c1e816d146101d7578063469c8110146101f6575b5f5ffd5b348015610156575f5ffd5b5061016a610165366004610ed5565b610410565b60405161017a9493929190610f1a565b60405180910390f35b34801561018e575f5ffd5b506101976104db565b005b3480156101a4575f5ffd5b506101976101b3366004610ed5565b610787565b3480156101c3575f5ffd5b506101976101d2366004610f54565b6107b6565b3480156101e2575f5ffd5b506101976101f1366004610f54565b610835565b610197610204366004610f95565b610881565b348015610214575f5ffd5b5061021e60025481565b60405190815260200161017a565b348015610237575f5ffd5b50610197610246366004610f95565b610aae565b348015610256575f5ffd5b5061025f610c8c565b60405161017a9190611048565b348015610277575f5ffd5b50610197610286366004610ed5565b610da3565b348015610296575f5ffd5b506101976102a5366004610f54565b610dd2565b3480156102b5575f5ffd5b506101976102c43660046110df565b610e5e565b3480156102d4575f5ffd5b5061021e60015481565b3480156102e9575f5ffd5b5060035461021e565b3480156102fd575f5ffd5b5060015461021e565b348015610311575f5ffd5b50670de0b6b3a764000061021e565b34801561032b575f5ffd5b5061019761033a366004610ed5565b610ea6565b34801561034a575f5ffd5b5060025461021e565b34801561035e575f5ffd5b5061021e670de0b6b3a764000081565b348015610379575f5ffd5b5061021e60035481565b34801561038e575f5ffd5b505f5460ff165b60405161017a9190611111565b3480156103ad575f5ffd5b505f546103959060ff1681565b3480156103c5575f5ffd5b506006546103d9906001600160a01b031681565b6040516001600160a01b03909116815260200161017a565b3480156103fc575f5ffd5b506005546103d9906001600160a01b031681565b6004818154811061041f575f80fd5b5f9182526020909120600490910201805460018201546002830180549294506001600160a01b03909116929161045490611137565b80601f016020809104026020016040519081016040528092919081815260200182805461048090611137565b80156104cb5780601f106104a2576101008083540402835291602001916104cb565b820191905f5260205f20905b8154815290600101906020018083116104ae57829003601f168201915b5050505050908060030154905084565b6005546001600160a01b0316331461050e5760405162461bcd60e51b81526004016105059061116f565b60405180910390fd5b5f5f5460ff166003811115610525576105256110fd565b146105725760405162461bcd60e51b815260206004820152601a60248201527f47616d65206973206e6f742070656e64696e67207061796f75740000000000006044820152606401610505565b6004546105b95760405162461bcd60e51b81526020600482015260156024820152744e6f206d6573736167657320617661696c61626c6560581b6044820152606401610505565b600654600480545f926001600160a01b031691906105d9906001906111c4565b815481106105e9576105e96111dd565b5f9182526020909120600160049092020101546001600160a01b031614801561061457506004546001105b1561065c5760048054610629906002906111c4565b81548110610639576106396111dd565b5f9182526020909120600160049092020101546001600160a01b0316905061069b565b6004805461066c906001906111c4565b8154811061067c5761067c6111dd565b5f9182526020909120600160049092020101546001600160a01b031690505b600180545f91829055815460ff191660021782556040519091906001600160a01b0384169083908381818185875af1925050503d805f81146106f8576040519150601f19603f3d011682016040523d82523d5f602084013e6106fd565b606091505b50509050806107405760405162461bcd60e51b815260206004820152600f60248201526e151c985b9cd9995c8819985a5b1959608a1b6044820152606401610505565b604080516001600160a01b0385168152602081018490527f504d6fcd7fd6584455567b4c515b7efa2a27fb623104bd49f98b4e9a2ccd7efb910160405180910390a1505050565b6005546001600160a01b031633146107b15760405162461bcd60e51b81526004016105059061116f565b600355565b6005546001600160a01b031633146107e05760405162461bcd60e51b81526004016105059061116f565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527fcb3ed5b60a03dc79bfa565ea813bd93572e94f8666e5b22f643b39d2e5f02930906020015b60405180910390a150565b6005546001600160a01b0316331461085f5760405162461bcd60e51b81526004016105059061116f565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b5f5f5460ff166003811115610898576108986110fd565b146108f75760405162461bcd60e51b815260206004820152602960248201527f47616d65206973206e6f7420616374697665206f72206974206973206e6f74206044820152683cb7bab9103a3ab93760b91b6064820152608401610505565b600254341461093e5760405162461bcd60e51b815260206004820152601360248201527257726f6e67206d65737361676520707269636560681b6044820152606401610505565b6040805160808101825260025481523360208201908152918101838152426060830152600480546001810182555f829052835191027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b810191825593517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c850180546001600160a01b03929092166001600160a01b03199092169190911790559051919290917f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19d90910190610a13908261123d565b506060820151816003015550503460015f828254610a3191906112f8565b9091555050600354600254670de0b6b3a764000091610a4f9161130b565b610a599190611322565b60025560405133907f4816bdd3b9d39acdd4da0ae7412d5523081e916c6a56e481063e475b86ef7c5b90610a909084904290611341565b60405180910390a25f80546001919060ff191682805b021790555050565b6006546001600160a01b03163314610b085760405162461bcd60e51b815260206004820152601760248201527f4f6e6c79204149206167656e742063616e207265706c790000000000000000006044820152606401610505565b60015f5460ff166003811115610b2057610b206110fd565b14610b6d5760405162461bcd60e51b815260206004820152601a60248201527f47616d65206973206e6f7420696e20676f696e672073746174650000000000006044820152606401610505565b604080516080810182525f8082526006546001600160a01b03908116602084019081529383018581524260608501526004805460018101825593819052845193027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b810193845594517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c86018054919093166001600160a01b03199091161790915551919290917f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19d90910190610c43908261123d565b50606091909101516003909101555f805460ff191690556040517f5aa9063abb57f078ec42ba6732d0c8631e1cb79dba57ba9b74b76bc0658755b69061082a9083904290611341565b60606004805480602002602001604051908101604052809291908181526020015f905b82821015610d9a575f84815260209081902060408051608081018252600486029092018054835260018101546001600160a01b03169383019390935260028301805492939291840191610d0190611137565b80601f0160208091040260200160405190810160405280929190818152602001828054610d2d90611137565b8015610d785780601f10610d4f57610100808354040283529160200191610d78565b820191905f5260205f20905b815481529060010190602001808311610d5b57829003601f168201915b5050505050815260200160038201548152505081526020019060010190610caf565b50505050905090565b6005546001600160a01b03163314610dcd5760405162461bcd60e51b81526004016105059061116f565b600155565b6005546001600160a01b03163314610dfc5760405162461bcd60e51b81526004016105059061116f565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527fcb3ed5b60a03dc79bfa565ea813bd93572e94f8666e5b22f643b39d2e5f029309060200160405180910390a15f8054819060ff1916600182610aa6565b6005546001600160a01b03163314610e885760405162461bcd60e51b81526004016105059061116f565b5f805482919060ff19166001836003811115610aa657610aa66110fd565b6005546001600160a01b03163314610ed05760405162461bcd60e51b81526004016105059061116f565b600255565b5f60208284031215610ee5575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b8481526001600160a01b03841660208201526080604082018190525f90610f4390830185610eec565b905082606083015295945050505050565b5f60208284031215610f64575f5ffd5b81356001600160a01b0381168114610f7a575f5ffd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610fa5575f5ffd5b813567ffffffffffffffff811115610fbb575f5ffd5b8201601f81018413610fcb575f5ffd5b803567ffffffffffffffff811115610fe557610fe5610f81565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561101457611014610f81565b60405281815282820160200186101561102b575f5ffd5b816020840160208301375f91810160200191909152949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156110d357603f1987860301845281518051865260018060a01b0360208201511660208701526040810151608060408801526110b06080880182610eec565b60609283015197909201969096529450602093840193919091019060010161106e565b50929695505050505050565b5f602082840312156110ef575f5ffd5b813560048110610f7a575f5ffd5b634e487b7160e01b5f52602160045260245ffd5b602081016004831061113157634e487b7160e01b5f52602160045260245ffd5b91905290565b600181811c9082168061114b57607f821691505b60208210810361116957634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526021908201527f4f6e6c792061646d696e2063616e2063616c6c20746869732066756e6374696f6040820152603760f91b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b818103818111156111d7576111d76111b0565b92915050565b634e487b7160e01b5f52603260045260245ffd5b601f82111561123857805f5260205f20601f840160051c810160208510156112165750805b601f840160051c820191505b81811015611235575f8155600101611222565b50505b505050565b815167ffffffffffffffff81111561125757611257610f81565b61126b816112658454611137565b846111f1565b6020601f82116001811461129d575f83156112865750848201515b5f19600385901b1c1916600184901b178455611235565b5f84815260208120601f198516915b828110156112cc57878501518255602094850194600190920191016112ac565b50848210156112e957868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156111d7576111d76111b0565b80820281158282048414176111d7576111d76111b0565b5f8261133c57634e487b7160e01b5f52601260045260245ffd5b500490565b604081525f6113536040830185610eec565b9050826020830152939250505056fea2646970667358221220078607a377a0f16e596e26af905f45ee00adb9aba359440f8b8592f2b582cee064736f6c634300081d0033",
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

// GetCoeffIncrease is a free data retrieval call binding the contract method 0x75139aa0.
//
// Solidity: function getCoeffIncrease() view returns(uint256)
func (_Contract *ContractCaller) GetCoeffIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCoeffIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCoeffIncrease is a free data retrieval call binding the contract method 0x75139aa0.
//
// Solidity: function getCoeffIncrease() view returns(uint256)
func (_Contract *ContractSession) GetCoeffIncrease() (*big.Int, error) {
	return _Contract.Contract.GetCoeffIncrease(&_Contract.CallOpts)
}

// GetCoeffIncrease is a free data retrieval call binding the contract method 0x75139aa0.
//
// Solidity: function getCoeffIncrease() view returns(uint256)
func (_Contract *ContractCallerSession) GetCoeffIncrease() (*big.Int, error) {
	return _Contract.Contract.GetCoeffIncrease(&_Contract.CallOpts)
}

// GetGameState is a free data retrieval call binding the contract method 0xb7d0628b.
//
// Solidity: function getGameState() view returns(uint8)
func (_Contract *ContractCaller) GetGameState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGameState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetGameState is a free data retrieval call binding the contract method 0xb7d0628b.
//
// Solidity: function getGameState() view returns(uint8)
func (_Contract *ContractSession) GetGameState() (uint8, error) {
	return _Contract.Contract.GetGameState(&_Contract.CallOpts)
}

// GetGameState is a free data retrieval call binding the contract method 0xb7d0628b.
//
// Solidity: function getGameState() view returns(uint8)
func (_Contract *ContractCallerSession) GetGameState() (uint8, error) {
	return _Contract.Contract.GetGameState(&_Contract.CallOpts)
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

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() pure returns(uint256)
func (_Contract *ContractCaller) GetPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() pure returns(uint256)
func (_Contract *ContractSession) GetPrecision() (*big.Int, error) {
	return _Contract.Contract.GetPrecision(&_Contract.CallOpts)
}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() pure returns(uint256)
func (_Contract *ContractCallerSession) GetPrecision() (*big.Int, error) {
	return _Contract.Contract.GetPrecision(&_Contract.CallOpts)
}

// GetPrizePool is a free data retrieval call binding the contract method 0x884bf67c.
//
// Solidity: function getPrizePool() view returns(uint256)
func (_Contract *ContractCaller) GetPrizePool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPrizePool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrizePool is a free data retrieval call binding the contract method 0x884bf67c.
//
// Solidity: function getPrizePool() view returns(uint256)
func (_Contract *ContractSession) GetPrizePool() (*big.Int, error) {
	return _Contract.Contract.GetPrizePool(&_Contract.CallOpts)
}

// GetPrizePool is a free data retrieval call binding the contract method 0x884bf67c.
//
// Solidity: function getPrizePool() view returns(uint256)
func (_Contract *ContractCallerSession) GetPrizePool() (*big.Int, error) {
	return _Contract.Contract.GetPrizePool(&_Contract.CallOpts)
}

// GetmessagePrice is a free data retrieval call binding the contract method 0xa8a68b4f.
//
// Solidity: function getmessagePrice() view returns(uint256)
func (_Contract *ContractCaller) GetmessagePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getmessagePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetmessagePrice is a free data retrieval call binding the contract method 0xa8a68b4f.
//
// Solidity: function getmessagePrice() view returns(uint256)
func (_Contract *ContractSession) GetmessagePrice() (*big.Int, error) {
	return _Contract.Contract.GetmessagePrice(&_Contract.CallOpts)
}

// GetmessagePrice is a free data retrieval call binding the contract method 0xa8a68b4f.
//
// Solidity: function getmessagePrice() view returns(uint256)
func (_Contract *ContractCallerSession) GetmessagePrice() (*big.Int, error) {
	return _Contract.Contract.GetmessagePrice(&_Contract.CallOpts)
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

// SetAIAgentAddress is a paid mutator transaction binding the contract method 0x2a60c8fb.
//
// Solidity: function setAIAgentAddress(address _aiAgentAddress) returns()
func (_Contract *ContractTransactor) SetAIAgentAddress(opts *bind.TransactOpts, _aiAgentAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAIAgentAddress", _aiAgentAddress)
}

// SetAIAgentAddress is a paid mutator transaction binding the contract method 0x2a60c8fb.
//
// Solidity: function setAIAgentAddress(address _aiAgentAddress) returns()
func (_Contract *ContractSession) SetAIAgentAddress(_aiAgentAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAIAgentAddress(&_Contract.TransactOpts, _aiAgentAddress)
}

// SetAIAgentAddress is a paid mutator transaction binding the contract method 0x2a60c8fb.
//
// Solidity: function setAIAgentAddress(address _aiAgentAddress) returns()
func (_Contract *ContractTransactorSession) SetAIAgentAddress(_aiAgentAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAIAgentAddress(&_Contract.TransactOpts, _aiAgentAddress)
}

// SetAdminAddress is a paid mutator transaction binding the contract method 0x2c1e816d.
//
// Solidity: function setAdminAddress(address _adminAddress) returns()
func (_Contract *ContractTransactor) SetAdminAddress(opts *bind.TransactOpts, _adminAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAdminAddress", _adminAddress)
}

// SetAdminAddress is a paid mutator transaction binding the contract method 0x2c1e816d.
//
// Solidity: function setAdminAddress(address _adminAddress) returns()
func (_Contract *ContractSession) SetAdminAddress(_adminAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAdminAddress(&_Contract.TransactOpts, _adminAddress)
}

// SetAdminAddress is a paid mutator transaction binding the contract method 0x2c1e816d.
//
// Solidity: function setAdminAddress(address _adminAddress) returns()
func (_Contract *ContractTransactorSession) SetAdminAddress(_adminAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAdminAddress(&_Contract.TransactOpts, _adminAddress)
}

// SetCoeffIncrease is a paid mutator transaction binding the contract method 0x18172d21.
//
// Solidity: function setCoeffIncrease(uint256 _coeffIncrease) returns()
func (_Contract *ContractTransactor) SetCoeffIncrease(opts *bind.TransactOpts, _coeffIncrease *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCoeffIncrease", _coeffIncrease)
}

// SetCoeffIncrease is a paid mutator transaction binding the contract method 0x18172d21.
//
// Solidity: function setCoeffIncrease(uint256 _coeffIncrease) returns()
func (_Contract *ContractSession) SetCoeffIncrease(_coeffIncrease *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetCoeffIncrease(&_Contract.TransactOpts, _coeffIncrease)
}

// SetCoeffIncrease is a paid mutator transaction binding the contract method 0x18172d21.
//
// Solidity: function setCoeffIncrease(uint256 _coeffIncrease) returns()
func (_Contract *ContractTransactorSession) SetCoeffIncrease(_coeffIncrease *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetCoeffIncrease(&_Contract.TransactOpts, _coeffIncrease)
}

// SetGameState is a paid mutator transaction binding the contract method 0x6dd047f3.
//
// Solidity: function setGameState(uint8 _gameState) returns()
func (_Contract *ContractTransactor) SetGameState(opts *bind.TransactOpts, _gameState uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGameState", _gameState)
}

// SetGameState is a paid mutator transaction binding the contract method 0x6dd047f3.
//
// Solidity: function setGameState(uint8 _gameState) returns()
func (_Contract *ContractSession) SetGameState(_gameState uint8) (*types.Transaction, error) {
	return _Contract.Contract.SetGameState(&_Contract.TransactOpts, _gameState)
}

// SetGameState is a paid mutator transaction binding the contract method 0x6dd047f3.
//
// Solidity: function setGameState(uint8 _gameState) returns()
func (_Contract *ContractTransactorSession) SetGameState(_gameState uint8) (*types.Transaction, error) {
	return _Contract.Contract.SetGameState(&_Contract.TransactOpts, _gameState)
}

// SetPrizePool is a paid mutator transaction binding the contract method 0x6a4966b8.
//
// Solidity: function setPrizePool(uint256 _prizePool) returns()
func (_Contract *ContractTransactor) SetPrizePool(opts *bind.TransactOpts, _prizePool *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPrizePool", _prizePool)
}

// SetPrizePool is a paid mutator transaction binding the contract method 0x6a4966b8.
//
// Solidity: function setPrizePool(uint256 _prizePool) returns()
func (_Contract *ContractSession) SetPrizePool(_prizePool *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPrizePool(&_Contract.TransactOpts, _prizePool)
}

// SetPrizePool is a paid mutator transaction binding the contract method 0x6a4966b8.
//
// Solidity: function setPrizePool(uint256 _prizePool) returns()
func (_Contract *ContractTransactorSession) SetPrizePool(_prizePool *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPrizePool(&_Contract.TransactOpts, _prizePool)
}

// SetmessagePrice is a paid mutator transaction binding the contract method 0xa3e2db8d.
//
// Solidity: function setmessagePrice(uint256 _messagePrice) returns()
func (_Contract *ContractTransactor) SetmessagePrice(opts *bind.TransactOpts, _messagePrice *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setmessagePrice", _messagePrice)
}

// SetmessagePrice is a paid mutator transaction binding the contract method 0xa3e2db8d.
//
// Solidity: function setmessagePrice(uint256 _messagePrice) returns()
func (_Contract *ContractSession) SetmessagePrice(_messagePrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetmessagePrice(&_Contract.TransactOpts, _messagePrice)
}

// SetmessagePrice is a paid mutator transaction binding the contract method 0xa3e2db8d.
//
// Solidity: function setmessagePrice(uint256 _messagePrice) returns()
func (_Contract *ContractTransactorSession) SetmessagePrice(_messagePrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetmessagePrice(&_Contract.TransactOpts, _messagePrice)
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
