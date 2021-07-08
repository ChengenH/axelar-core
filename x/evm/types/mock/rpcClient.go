// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	context "context"
	"github.com/axelarnetwork/axelar-core/x/evm/types"
	evm "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	evmTypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"
)

// Ensure, that RPCClientMock does implement types.RPCClient.
// If this is not the case, regenerate this file with moq.
var _ types.RPCClient = &RPCClientMock{}

// RPCClientMock is a mock implementation of types.RPCClient.
//
// 	func TestSomethingThatUsesRPCClient(t *testing.T) {
//
// 		// make and configure a mocked types.RPCClient
// 		mockedRPCClient := &RPCClientMock{
// 			BlockNumberFunc: func(ctx context.Context) (uint64, error) {
// 				panic("mock out the BlockNumber method")
// 			},
// 			ChainIDFunc: func(ctx context.Context) (*big.Int, error) {
// 				panic("mock out the ChainID method")
// 			},
// 			EstimateGasFunc: func(ctx context.Context, msg evm.CallMsg) (uint64, error) {
// 				panic("mock out the EstimateGas method")
// 			},
// 			PendingNonceAtFunc: func(ctx context.Context, account common.Address) (uint64, error) {
// 				panic("mock out the PendingNonceAt method")
// 			},
// 			SendAndSignTransactionFunc: func(ctx context.Context, msg evm.CallMsg) (string, error) {
// 				panic("mock out the SendAndSignTransaction method")
// 			},
// 			SendTransactionFunc: func(ctx context.Context, tx *evmTypes.Transaction) error {
// 				panic("mock out the SendTransaction method")
// 			},
// 			SuggestGasPriceFunc: func(ctx context.Context) (*big.Int, error) {
// 				panic("mock out the SuggestGasPrice method")
// 			},
// 			TransactionReceiptFunc: func(ctx context.Context, txHash common.Hash) (*evmTypes.Receipt, error) {
// 				panic("mock out the TransactionReceipt method")
// 			},
// 		}
//
// 		// use mockedRPCClient in code that requires types.RPCClient
// 		// and then make assertions.
//
// 	}
type RPCClientMock struct {
	// BlockNumberFunc mocks the BlockNumber method.
	BlockNumberFunc func(ctx context.Context) (uint64, error)

	// ChainIDFunc mocks the ChainID method.
	ChainIDFunc func(ctx context.Context) (*big.Int, error)

	// EstimateGasFunc mocks the EstimateGas method.
	EstimateGasFunc func(ctx context.Context, msg evm.CallMsg) (uint64, error)

	// PendingNonceAtFunc mocks the PendingNonceAt method.
	PendingNonceAtFunc func(ctx context.Context, account common.Address) (uint64, error)

	// SendAndSignTransactionFunc mocks the SendAndSignTransaction method.
	SendAndSignTransactionFunc func(ctx context.Context, msg evm.CallMsg) (string, error)

	// SendTransactionFunc mocks the SendTransaction method.
	SendTransactionFunc func(ctx context.Context, tx *evmTypes.Transaction) error

	// SuggestGasPriceFunc mocks the SuggestGasPrice method.
	SuggestGasPriceFunc func(ctx context.Context) (*big.Int, error)

	// TransactionReceiptFunc mocks the TransactionReceipt method.
	TransactionReceiptFunc func(ctx context.Context, txHash common.Hash) (*evmTypes.Receipt, error)

	// calls tracks calls to the methods.
	calls struct {
		// BlockNumber holds details about calls to the BlockNumber method.
		BlockNumber []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// ChainID holds details about calls to the ChainID method.
		ChainID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// EstimateGas holds details about calls to the EstimateGas method.
		EstimateGas []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Msg is the msg argument value.
			Msg evm.CallMsg
		}
		// PendingNonceAt holds details about calls to the PendingNonceAt method.
		PendingNonceAt []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Account is the account argument value.
			Account common.Address
		}
		// SendAndSignTransaction holds details about calls to the SendAndSignTransaction method.
		SendAndSignTransaction []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Msg is the msg argument value.
			Msg evm.CallMsg
		}
		// SendTransaction holds details about calls to the SendTransaction method.
		SendTransaction []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx *evmTypes.Transaction
		}
		// SuggestGasPrice holds details about calls to the SuggestGasPrice method.
		SuggestGasPrice []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// TransactionReceipt holds details about calls to the TransactionReceipt method.
		TransactionReceipt []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// TxHash is the txHash argument value.
			TxHash common.Hash
		}
	}
	lockBlockNumber            sync.RWMutex
	lockChainID                sync.RWMutex
	lockEstimateGas            sync.RWMutex
	lockPendingNonceAt         sync.RWMutex
	lockSendAndSignTransaction sync.RWMutex
	lockSendTransaction        sync.RWMutex
	lockSuggestGasPrice        sync.RWMutex
	lockTransactionReceipt     sync.RWMutex
}

// BlockNumber calls BlockNumberFunc.
func (mock *RPCClientMock) BlockNumber(ctx context.Context) (uint64, error) {
	if mock.BlockNumberFunc == nil {
		panic("RPCClientMock.BlockNumberFunc: method is nil but RPCClient.BlockNumber was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockBlockNumber.Lock()
	mock.calls.BlockNumber = append(mock.calls.BlockNumber, callInfo)
	mock.lockBlockNumber.Unlock()
	return mock.BlockNumberFunc(ctx)
}

// BlockNumberCalls gets all the calls that were made to BlockNumber.
// Check the length with:
//     len(mockedRPCClient.BlockNumberCalls())
func (mock *RPCClientMock) BlockNumberCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockBlockNumber.RLock()
	calls = mock.calls.BlockNumber
	mock.lockBlockNumber.RUnlock()
	return calls
}

// ChainID calls ChainIDFunc.
func (mock *RPCClientMock) ChainID(ctx context.Context) (*big.Int, error) {
	if mock.ChainIDFunc == nil {
		panic("RPCClientMock.ChainIDFunc: method is nil but RPCClient.ChainID was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockChainID.Lock()
	mock.calls.ChainID = append(mock.calls.ChainID, callInfo)
	mock.lockChainID.Unlock()
	return mock.ChainIDFunc(ctx)
}

// ChainIDCalls gets all the calls that were made to ChainID.
// Check the length with:
//     len(mockedRPCClient.ChainIDCalls())
func (mock *RPCClientMock) ChainIDCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockChainID.RLock()
	calls = mock.calls.ChainID
	mock.lockChainID.RUnlock()
	return calls
}

// EstimateGas calls EstimateGasFunc.
func (mock *RPCClientMock) EstimateGas(ctx context.Context, msg evm.CallMsg) (uint64, error) {
	if mock.EstimateGasFunc == nil {
		panic("RPCClientMock.EstimateGasFunc: method is nil but RPCClient.EstimateGas was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Msg evm.CallMsg
	}{
		Ctx: ctx,
		Msg: msg,
	}
	mock.lockEstimateGas.Lock()
	mock.calls.EstimateGas = append(mock.calls.EstimateGas, callInfo)
	mock.lockEstimateGas.Unlock()
	return mock.EstimateGasFunc(ctx, msg)
}

// EstimateGasCalls gets all the calls that were made to EstimateGas.
// Check the length with:
//     len(mockedRPCClient.EstimateGasCalls())
func (mock *RPCClientMock) EstimateGasCalls() []struct {
	Ctx context.Context
	Msg evm.CallMsg
} {
	var calls []struct {
		Ctx context.Context
		Msg evm.CallMsg
	}
	mock.lockEstimateGas.RLock()
	calls = mock.calls.EstimateGas
	mock.lockEstimateGas.RUnlock()
	return calls
}

// PendingNonceAt calls PendingNonceAtFunc.
func (mock *RPCClientMock) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	if mock.PendingNonceAtFunc == nil {
		panic("RPCClientMock.PendingNonceAtFunc: method is nil but RPCClient.PendingNonceAt was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Account common.Address
	}{
		Ctx:     ctx,
		Account: account,
	}
	mock.lockPendingNonceAt.Lock()
	mock.calls.PendingNonceAt = append(mock.calls.PendingNonceAt, callInfo)
	mock.lockPendingNonceAt.Unlock()
	return mock.PendingNonceAtFunc(ctx, account)
}

// PendingNonceAtCalls gets all the calls that were made to PendingNonceAt.
// Check the length with:
//     len(mockedRPCClient.PendingNonceAtCalls())
func (mock *RPCClientMock) PendingNonceAtCalls() []struct {
	Ctx     context.Context
	Account common.Address
} {
	var calls []struct {
		Ctx     context.Context
		Account common.Address
	}
	mock.lockPendingNonceAt.RLock()
	calls = mock.calls.PendingNonceAt
	mock.lockPendingNonceAt.RUnlock()
	return calls
}

// SendAndSignTransaction calls SendAndSignTransactionFunc.
func (mock *RPCClientMock) SendAndSignTransaction(ctx context.Context, msg evm.CallMsg) (string, error) {
	if mock.SendAndSignTransactionFunc == nil {
		panic("RPCClientMock.SendAndSignTransactionFunc: method is nil but RPCClient.SendAndSignTransaction was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Msg evm.CallMsg
	}{
		Ctx: ctx,
		Msg: msg,
	}
	mock.lockSendAndSignTransaction.Lock()
	mock.calls.SendAndSignTransaction = append(mock.calls.SendAndSignTransaction, callInfo)
	mock.lockSendAndSignTransaction.Unlock()
	return mock.SendAndSignTransactionFunc(ctx, msg)
}

// SendAndSignTransactionCalls gets all the calls that were made to SendAndSignTransaction.
// Check the length with:
//     len(mockedRPCClient.SendAndSignTransactionCalls())
func (mock *RPCClientMock) SendAndSignTransactionCalls() []struct {
	Ctx context.Context
	Msg evm.CallMsg
} {
	var calls []struct {
		Ctx context.Context
		Msg evm.CallMsg
	}
	mock.lockSendAndSignTransaction.RLock()
	calls = mock.calls.SendAndSignTransaction
	mock.lockSendAndSignTransaction.RUnlock()
	return calls
}

// SendTransaction calls SendTransactionFunc.
func (mock *RPCClientMock) SendTransaction(ctx context.Context, tx *evmTypes.Transaction) error {
	if mock.SendTransactionFunc == nil {
		panic("RPCClientMock.SendTransactionFunc: method is nil but RPCClient.SendTransaction was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Tx  *evmTypes.Transaction
	}{
		Ctx: ctx,
		Tx:  tx,
	}
	mock.lockSendTransaction.Lock()
	mock.calls.SendTransaction = append(mock.calls.SendTransaction, callInfo)
	mock.lockSendTransaction.Unlock()
	return mock.SendTransactionFunc(ctx, tx)
}

// SendTransactionCalls gets all the calls that were made to SendTransaction.
// Check the length with:
//     len(mockedRPCClient.SendTransactionCalls())
func (mock *RPCClientMock) SendTransactionCalls() []struct {
	Ctx context.Context
	Tx  *evmTypes.Transaction
} {
	var calls []struct {
		Ctx context.Context
		Tx  *evmTypes.Transaction
	}
	mock.lockSendTransaction.RLock()
	calls = mock.calls.SendTransaction
	mock.lockSendTransaction.RUnlock()
	return calls
}

// SuggestGasPrice calls SuggestGasPriceFunc.
func (mock *RPCClientMock) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	if mock.SuggestGasPriceFunc == nil {
		panic("RPCClientMock.SuggestGasPriceFunc: method is nil but RPCClient.SuggestGasPrice was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockSuggestGasPrice.Lock()
	mock.calls.SuggestGasPrice = append(mock.calls.SuggestGasPrice, callInfo)
	mock.lockSuggestGasPrice.Unlock()
	return mock.SuggestGasPriceFunc(ctx)
}

// SuggestGasPriceCalls gets all the calls that were made to SuggestGasPrice.
// Check the length with:
//     len(mockedRPCClient.SuggestGasPriceCalls())
func (mock *RPCClientMock) SuggestGasPriceCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockSuggestGasPrice.RLock()
	calls = mock.calls.SuggestGasPrice
	mock.lockSuggestGasPrice.RUnlock()
	return calls
}

// TransactionReceipt calls TransactionReceiptFunc.
func (mock *RPCClientMock) TransactionReceipt(ctx context.Context, txHash common.Hash) (*evmTypes.Receipt, error) {
	if mock.TransactionReceiptFunc == nil {
		panic("RPCClientMock.TransactionReceiptFunc: method is nil but RPCClient.TransactionReceipt was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		TxHash common.Hash
	}{
		Ctx:    ctx,
		TxHash: txHash,
	}
	mock.lockTransactionReceipt.Lock()
	mock.calls.TransactionReceipt = append(mock.calls.TransactionReceipt, callInfo)
	mock.lockTransactionReceipt.Unlock()
	return mock.TransactionReceiptFunc(ctx, txHash)
}

// TransactionReceiptCalls gets all the calls that were made to TransactionReceipt.
// Check the length with:
//     len(mockedRPCClient.TransactionReceiptCalls())
func (mock *RPCClientMock) TransactionReceiptCalls() []struct {
	Ctx    context.Context
	TxHash common.Hash
} {
	var calls []struct {
		Ctx    context.Context
		TxHash common.Hash
	}
	mock.lockTransactionReceipt.RLock()
	calls = mock.calls.TransactionReceipt
	mock.lockTransactionReceipt.RUnlock()
	return calls
}
