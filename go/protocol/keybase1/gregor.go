// Auto-generated by avdl-compiler v1.3.19 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/gregor.avdl

package keybase1

import (
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type GetStateArg struct {
}

func (o GetStateArg) DeepCopy() GetStateArg {
	return GetStateArg{}
}

type InjectItemArg struct {
	Cat   string               `codec:"cat" json:"cat"`
	Body  string               `codec:"body" json:"body"`
	Dtime gregor1.TimeOrOffset `codec:"dtime" json:"dtime"`
}

func (o InjectItemArg) DeepCopy() InjectItemArg {
	return InjectItemArg{
		Cat:   o.Cat,
		Body:  o.Body,
		Dtime: o.Dtime.DeepCopy(),
	}
}

type DismissCategoryArg struct {
	Category gregor1.Category `codec:"category" json:"category"`
}

func (o DismissCategoryArg) DeepCopy() DismissCategoryArg {
	return DismissCategoryArg{
		Category: o.Category.DeepCopy(),
	}
}

type GregorInterface interface {
	GetState(context.Context) (gregor1.State, error)
	InjectItem(context.Context, InjectItemArg) (gregor1.MsgID, error)
	DismissCategory(context.Context, gregor1.Category) error
}

func GregorProtocol(i GregorInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.gregor",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getState": {
				MakeArg: func() interface{} {
					ret := make([]GetStateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetState(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"injectItem": {
				MakeArg: func() interface{} {
					ret := make([]InjectItemArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]InjectItemArg)
					if !ok {
						err = rpc.NewTypeError((*[]InjectItemArg)(nil), args)
						return
					}
					ret, err = i.InjectItem(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"dismissCategory": {
				MakeArg: func() interface{} {
					ret := make([]DismissCategoryArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DismissCategoryArg)
					if !ok {
						err = rpc.NewTypeError((*[]DismissCategoryArg)(nil), args)
						return
					}
					err = i.DismissCategory(ctx, (*typedArgs)[0].Category)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type GregorClient struct {
	Cli rpc.GenericClient
}

func (c GregorClient) GetState(ctx context.Context) (res gregor1.State, err error) {
	err = c.Cli.Call(ctx, "keybase.1.gregor.getState", []interface{}{GetStateArg{}}, &res)
	return
}

func (c GregorClient) InjectItem(ctx context.Context, __arg InjectItemArg) (res gregor1.MsgID, err error) {
	err = c.Cli.Call(ctx, "keybase.1.gregor.injectItem", []interface{}{__arg}, &res)
	return
}

func (c GregorClient) DismissCategory(ctx context.Context, category gregor1.Category) (err error) {
	__arg := DismissCategoryArg{Category: category}
	err = c.Cli.Call(ctx, "keybase.1.gregor.dismissCategory", []interface{}{__arg}, nil)
	return
}
