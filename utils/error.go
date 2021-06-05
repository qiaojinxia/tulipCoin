package utils

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */





type  BasicError struct {
	msg string
	error error
}

func (n *BasicError) Error() string {
	return n.error.Error()
}

func (n *BasicError) RuntimeError() {
	panic(n.error)
}

type NetError struct {
	BasicError
}

type BusinessError struct {
	BasicError
}

type MarshalError struct {
	BasicError
}

type StackError struct {
	BasicError
}

type ConvertError struct {
	BasicError
}

var ErrorsMap = make(map[reflect.Type]int)

func init(){
	ErrorsMap[reflect.TypeOf(NetErroWarp(nil,""))] = 100
	ErrorsMap[reflect.TypeOf(StackErrorWarp(nil,""))] = 200
	ErrorsMap[reflect.TypeOf(BusinessErrorWarp(nil,""))] = 300
	ErrorsMap[reflect.TypeOf(MarshalErrorWarp(nil,""))] = 400
	ErrorsMap[reflect.TypeOf(ConverErrorWarp(nil,""))] = 500
}

func MarshalErrorWarp(err error,msg string) error{
	return &MarshalError{BasicError{
		error: errors.WithStack(err) ,
		msg: msg,

	}}
}
func StackErrorWarp(err error,msg string) error{
	return &StackError{BasicError{
		error: err ,
		msg: msg,
	}}
}


func ConverErrorWarp(err error,msg string) error{
	return &ConvertError{BasicError{
		error: err ,
		msg: msg,
	}}
}


func BusinessErrorWarp(err error,msg string) error{
	return &BusinessError{BasicError{
		error:err,
		msg: msg,
	}}
}

func NetErroWarp(err error,msg string) error{
	return &NetError{BasicError{
		error:err,
		msg: msg,
	}}
}

type ErrroMsg struct {
	Code int `json:"code"`
	ErrorType string `json:"error_type"`
	Msg string `json:"msg"`
	ErrInfo string `json:"err_info"`
}


func ConverToJsonInfo(err error) string{
	object := reflect.ValueOf(err)
	msg := object.Elem().Field(0).Interface()
	emsg := &ErrroMsg{
		Code: ErrorsMap[reflect.TypeOf(err)],
		ErrorType: reflect.TypeOf(err).Elem().Name(),
		Msg: msg.(BasicError).msg,
		ErrInfo: err.Error(),

	}
	data,err := json.Marshal(emsg)
	if err != nil{
		panic(err)
	}
	return string(data)
}


type CatchHandler interface {
	Catch(e error, handler func(err error)) CatchHandler
	CatchAll(handler func(err error)) FinalHandler
	FinalHandler
}

type FinalHandler interface {
	Finally(handlers ...func())
}

type catchHandler struct {
	err error
	hasCatch bool
}
func(t *catchHandler) RequireCatch() bool {
	if t.hasCatch {
		return false
	}
	if t.err == nil {
		return false
	}
	return true
}

func (t *catchHandler) Catch(e error, handler func(err error)) CatchHandler {
	if !t.RequireCatch() {
		return t
	}

	if reflect.TypeOf(e) == reflect.TypeOf(t.err) {
		handler(t.err)
		t.hasCatch = true
	}
	return t
}

func (t *catchHandler) CatchAll(handler func(err error)) FinalHandler {
	if !t.RequireCatch() {
		return t
	}
	handler(t.err)
	t.hasCatch = true
	return t
}

func (t *catchHandler) Finally(handlers ...func()) {
	for _, handler := range handlers {
		defer handler()
	}
	err := t.err
	if err != nil && !t.hasCatch {
		panic(err)
	}
}

func Try(f func()) CatchHandler {
	t := &catchHandler{}
	defer func() {
		defer func() { //<2>
			r := recover()
			if r != nil {
				t.err = r.(error)
			}
		}()
		f() //<1>
	}()
	return t
}