package tracer

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gregpr07/bridge-decoder/utils"
)

type Tracer struct {
	config TracerConfig
	client *rpc.Client
	arg    traceArg
	latest uint64
}

func New(ctx context.Context, rawurl string, config TracerConfig) (*Tracer, error) {
	client, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}

	arg := makeTraceArg(config)

	return &Tracer{config, client, arg, 0}, nil
}

type blockInfo struct {
	Transactions []string       `json:"transactions"`
	Number       hexutil.Uint64 `json:"number"`
}

type wrappedTxTrace struct {
	Result utils.TxTrace `json:"result"`
}

func (t *Tracer) TraceBlock(ctx context.Context, blockNumber string) ([]utils.TxTrace, error) {
	var block blockInfo
	err := t.client.CallContext(ctx, &block, "eth_getBlockByNumber", blockNumber, false)
	if err != nil {
		return nil, err
	}

	var wrappedTraces []wrappedTxTrace
	err = t.client.CallContext(ctx, &wrappedTraces, "debug_traceBlockByNumber", blockNumber, t.arg)
	if err != nil {
		return nil, err
	}

	t.latest = uint64(block.Number)

	traces := make([]utils.TxTrace, len(wrappedTraces))

	for i := range wrappedTraces {
		traces[i] = wrappedTraces[i].Result
		traces[i].TransactionHash = block.Transactions[i]
	}

	return traces, nil
}

func (t *Tracer) TraceLatest(ctx context.Context) ([]utils.TxTrace, error) {
	return t.TraceBlock(ctx, "latest")
}

func (t *Tracer) LatestBlock(ctx context.Context) (uint64, error) {
	var result hexutil.Uint64
	err := t.client.CallContext(ctx, &result, "eth_blockNumber")
	return uint64(result), err
}

// If no tracing has been runned, will only trace the latest; otherwise trace to latest
func (t *Tracer) TraceToLatest(ctx context.Context, c chan []utils.TxTrace) error {
	if t.latest == 0 {
		result, err := t.TraceLatest(ctx)
		if err != nil {
			return err
		}
		c <- result
	} else {
		latest, err := t.LatestBlock(ctx)
		if err != nil {
			return err
		}

		for i := t.latest + 1; i <= latest; i++ {
			result, err := t.TraceBlock(ctx, hexutil.EncodeUint64(i))
			if err != nil {
				return err
			}
			c <- result
		}
	}
	return nil
}

// func (t *Tracer) TraceLive(ctx context.Context) error {

// }

type traceArg struct {
	Tracer string `json:"tracer"`
}

// func makeTraceArg(config TracerConfig) ([]byte, error) {
// 	var arg traceArg
// 	if config.ty == NativeTracer {
// 		arg = traceArg{NativeTracerContent}
// 	} else { // JSTracer
// 		arg = traceArg{JSTracerContent}
// 	}
// 	return json.Marshal(arg)
// }

func makeTraceArg(config TracerConfig) traceArg {
	if config.ty == NativeTracer {
		return traceArg{NativeTracerContent}
	} else { // JSTracer
		return traceArg{JSTracerContent}
	}
}
