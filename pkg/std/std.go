package std

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"github.com/jkcfg/jk/pkg/__std"
	"github.com/jkcfg/jk/pkg/__std/lib"
	"github.com/jkcfg/jk/pkg/deferred"

	flatbuffers "github.com/google/flatbuffers/go"
)

// Module returns the std source corresponding to the 'std' import
func Module(path string) []byte {
	data, _ := lib.ReadAll(path)
	return data
}

type sender interface {
	SendBytes([]byte) error
}

// RPCFunc is a function that can be registered for dispatch. The
// arguments each will either be []byte or JSON values; any result
// returned will be serialised to JSON.
type RPCFunc func([]interface{}) (interface{}, error)

// ExecuteOptions global input parameters to the standards library.
type ExecuteOptions struct {
	// Verbose indicates if some operations, such as write, should print out what
	// they are doing.
	Verbose bool
	// Parameters is a structured set of input parameters.
	Parameters Params
	// OutputDirectory is a directory used by any file producing functions as the
	// base directory to output files to.
	OutputDirectory string
	// Root is topmost directory under which file reads are allowed
	Root ReadBase
	// DryRun instructs standard library functions to not complete operations that
	// would mutate something (eg. std.write()).
	DryRun bool
	// Methods is where RPC methods are registered.
	Methods map[string]RPCFunc
}

func toBool(b byte) bool {
	if b == 0 {
		return false
	}
	return true
}

// stdError builds an Error flatbuffer we can return to the javascript side.
func stdError(b *flatbuffers.Builder, err error) flatbuffers.UOffsetT {
	off := b.CreateString(err.Error())
	__std.ErrorStart(b)
	__std.ErrorAddMessage(b, off)
	return __std.ErrorEnd(b)
}

// Execute parses a message from v8 and execute the corresponding function.
func Execute(msg []byte, res sender, options ExecuteOptions) []byte {
	message := __std.GetRootAsMessage(msg, 0)

	union := flatbuffers.Table{}
	if !message.Args(&union) {
		log.Fatal("could not parse Message")
	}

	switch message.ArgsType() {
	case __std.ArgsWriteArgs:
		args := __std.WriteArgs{}
		args.Init(union.Bytes, union.Pos)

		// Weave options.OutputDirectory in there.
		path := string(args.Path())
		if path != "" && !filepath.IsAbs(path) {
			path = filepath.Join(options.OutputDirectory, path)
		}

		if path != "" && options.Verbose {
			fmt.Printf("wrote %s\n", path)
		}

		if options.DryRun {
			break
		}

		write(args.Value(), path, args.Format(), int(args.Indent()), toBool(args.Overwrite()))
		return nil

	case __std.ArgsReadArgs:
		args := __std.ReadArgs{}
		args.Init(union.Bytes, union.Pos)
		path := string(args.Path())
		if path != "" && options.Verbose {
			fmt.Printf("read %s\n", path)
		}
		module := string(args.Module())
		ser := deferred.Register(func() ([]byte, error) {
			return options.Root.Read(path, args.Format(), args.Encoding(), module)
		}, sendFunc(res.SendBytes))
		return deferredResponse(ser)

	case __std.ArgsRPCArgs:
		args := __std.RPCArgs{}
		args.Init(union.Bytes, union.Pos)
		method := string(args.Method())
		rpcfn, ok := options.Methods[method]
		if !ok {
			return deferredError("RPC method not found: " + method)
		}

		numArgs := args.ArgsLength()
		arguments := make([]interface{}, numArgs)

		for i := 0; i < numArgs; i++ {
			arg := __std.RPCArg{}
			var argUnion flatbuffers.Table
			if !args.Args(&arg, i) || !arg.Arg(&argUnion) {
				return deferredError(fmt.Sprintf("could not decode arguments[%d]", i))
			}

			switch arg.ArgType() {
			case __std.RPCValueRPCSerialised:
				serialised := __std.RPCSerialised{}
				serialised.Init(argUnion.Bytes, argUnion.Pos)
				if err := json.Unmarshal(serialised.Value(), &arguments[i]); err != nil {
					return deferredError(fmt.Sprintf("could not parse serialised arguments[%d]: %s", i, err.Error()))
				}
			case __std.RPCValueData:
				bytes := __std.Data{}
				bytes.Init(argUnion.Bytes, argUnion.Pos)
				arguments[i] = bytes.BytesBytes()
			}
		}

		if args.Sync() == 1 {
			result, err := rpcfn(arguments)
			if err != nil {
				return rpcError(err.Error())
			}
			bytes, err := json.Marshal(result)
			if err != nil {
				return rpcError(err.Error())
			}
			return rpcData(bytes)
		}
		ser := deferred.Register(func() ([]byte, error) {
			result, err := rpcfn(arguments)
			if err != nil {
				return nil, err
			}
			return json.Marshal(result)
		}, sendFunc(res.SendBytes))
		return deferredResponse(ser)

	case __std.ArgsParseArgs:
		args := __std.ParseArgs{}
		args.Init(union.Bytes, union.Pos)
		out, err := Parse(args.Chars(), args.Format())
		return parseUnparseResponse(out, err)

	case __std.ArgsUnparseArgs:
		args := __std.UnparseArgs{}
		args.Init(union.Bytes, union.Pos)
		out, err := Unparse(args.Object(), args.Format())
		return parseUnparseResponse(out, err)

	case __std.ArgsFileInfoArgs:
		args := __std.FileInfoArgs{}
		args.Init(union.Bytes, union.Pos)
		return options.Root.FileInfo(string(args.Path()), string(args.Module()))
	case __std.ArgsListArgs:
		args := __std.ListArgs{}
		args.Init(union.Bytes, union.Pos)
		return options.Root.DirectoryListing(string(args.Path()), string(args.Module()))

	case __std.ArgsParamArgs:
		args := __std.ParamArgs{}
		args.Init(union.Bytes, union.Pos)

		// return buffer.
		b := flatbuffers.NewBuilder(512)
		var (
			off  flatbuffers.UOffsetT
			kind byte
		)

		json, err := param(options.Parameters, __std.ParamType(args.Type()), string(args.Path()), string(args.DefaultValue()))
		if err != nil {
			kind = __std.ParamRetvalError
			off = stdError(b, err)
		} else {
			kind = __std.ParamRetvalParamValue
			jsonOffset := b.CreateString(string(json))
			__std.ParamValueStart(b)
			__std.ParamValueAddJson(b, jsonOffset)
			off = __std.ParamValueEnd(b)
		}

		__std.ParamResponseStart(b)
		__std.ParamResponseAddRetvalType(b, kind)
		__std.ParamResponseAddRetval(b, off)
		responseOffset := __std.ParamResponseEnd(b)
		b.Finish(responseOffset)
		return b.FinishedBytes()

	default:
		log.Fatalf("unknown Message (%d)", message.ArgsType())
	}

	return nil
}

// deferredResponse constructs a response containing the serial number
// of the deferred value, to indicate to JavaScript that the request
// has been accepted and its success or failure will be communicated
// later.
func deferredResponse(s deferred.Serial) []byte {
	b := flatbuffers.NewBuilder(20)
	__std.DeferredStart(b)
	__std.DeferredAddSerial(b, uint64(s))
	off := __std.DeferredEnd(b)
	__std.DeferredResponseStart(b)
	__std.DeferredResponseAddRetvalType(b, __std.DeferredRetvalDeferred)
	__std.DeferredResponseAddRetval(b, off)
	off = __std.DeferredResponseEnd(b)
	b.Finish(off)
	return b.FinishedBytes()
}

// deferredError constructs a response containing an error, to
// indicate to JavaScript that request has not been accepted.
func deferredError(err string) []byte {
	b := flatbuffers.NewBuilder(512)
	msg := b.CreateString(err)
	__std.ErrorStart(b)
	__std.ErrorAddMessage(b, msg)
	off := __std.ErrorEnd(b)
	__std.DeferredResponseStart(b)
	__std.DeferredResponseAddRetvalType(b, __std.DeferredRetvalError)
	__std.DeferredResponseAddRetval(b, off)
	off = __std.DeferredResponseEnd(b)
	b.Finish(off)
	return b.FinishedBytes()
}

// rpcData encodes a successful synchronous RPC call
func rpcData(data []byte) []byte {
	b := flatbuffers.NewBuilder(1024)
	off := b.CreateByteVector(data)
	__std.DataStart(b)
	__std.DataAddBytes(b, off)
	off = __std.DataEnd(b)
	__std.RPCSyncResponseStart(b)
	__std.RPCSyncResponseAddRetvalType(b, __std.RPCSyncRetvalData)
	__std.RPCSyncResponseAddRetval(b, off)
	off = __std.RPCSyncResponseEnd(b)
	b.Finish(off)
	return b.FinishedBytes()
}

// rpcError encodes a failed synchronous RPC call
func rpcError(msg string) []byte {
	b := flatbuffers.NewBuilder(1024)
	off := b.CreateString(msg)
	__std.ErrorStart(b)
	__std.ErrorAddMessage(b, off)
	off = __std.ErrorEnd(b)
	__std.RPCSyncResponseStart(b)
	__std.RPCSyncResponseAddRetvalType(b, __std.RPCSyncRetvalError)
	__std.RPCSyncResponseAddRetval(b, off)
	off = __std.RPCSyncResponseEnd(b)
	b.Finish(off)
	return b.FinishedBytes()
}

type sendFunc func([]byte) error

func (fn sendFunc) Error(s deferred.Serial, err error) {
	b := flatbuffers.NewBuilder(512)
	off := stdError(b, err)
	__std.FulfilmentStart(b)
	__std.FulfilmentAddSerial(b, uint64(s))
	__std.FulfilmentAddValueType(b, __std.FulfilmentValueError)
	__std.FulfilmentAddValue(b, off)
	off = __std.FulfilmentEnd(b)
	b.Finish(off)
	if err := fn(b.FinishedBytes()); err != nil {
		panic(err)
	}
}

func (fn sendFunc) Data(s deferred.Serial, data []byte) {
	b := flatbuffers.NewBuilder(1024)
	off := b.CreateByteVector(data)
	__std.DataStart(b)
	__std.DataAddBytes(b, off)
	off = __std.DataEnd(b)
	__std.FulfilmentStart(b)
	__std.FulfilmentAddSerial(b, uint64(s))
	__std.FulfilmentAddValueType(b, __std.FulfilmentValueData)
	__std.FulfilmentAddValue(b, off)
	off = __std.FulfilmentEnd(b)
	b.Finish(off)
	if err := fn(b.FinishedBytes()); err != nil {
		panic(err)
	}
}

func (fn sendFunc) End(s deferred.Serial) {
	b := flatbuffers.NewBuilder(1024)
	__std.EndOfStreamStart(b)
	off := __std.EndOfStreamEnd(b)
	__std.FulfilmentStart(b)
	__std.FulfilmentAddSerial(b, uint64(s))
	__std.FulfilmentAddValueType(b, __std.FulfilmentValueEndOfStream)
	__std.FulfilmentAddValue(b, off)
	off = __std.FulfilmentEnd(b)
	b.Finish(off)
	if err := fn(b.FinishedBytes()); err != nil {
		panic(err)
	}
}
