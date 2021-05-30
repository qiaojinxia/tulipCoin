package utils

import (
	"encoding/json"

	"reflect"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */





type  BasicError struct {
	error string
}

func (n *BasicError) Error() string {
	return n.error
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
	ErrorsMap[reflect.TypeOf(NetErroWarp(""))] = 100
	ErrorsMap[reflect.TypeOf(StackErrorWarp(""))] = 200
	ErrorsMap[reflect.TypeOf(BusinessErrorWarp(""))] = 300
	ErrorsMap[reflect.TypeOf(MarshalErrorWarp(""))] = 400
	ErrorsMap[reflect.TypeOf(ConverErrorWarp(""))] = 500
}


func MarshalErrorWarp(msg string) error{
	return &MarshalError{BasicError{
		error: msg,
	}}
}
func StackErrorWarp(msg string) error{
	return &StackError{BasicError{
		error: msg,
	}}
}


func ConverErrorWarp(msg string) error{
	return &ConvertError{BasicError{
		error: msg,
	}}
}


func BusinessErrorWarp(msg string) error{
	return &BusinessError{BasicError{
		error: msg,
	}}
}

func NetErroWarp(msg string) error{
	return &NetError{BasicError{
		error: msg,
	}}
}

type ErrroMsg struct {
	Code int `json:"code"`
	ErrorType string `json:"error_type"`
	Msg string `json:"msg"`
}


func ConverToJsonInfo(err error) string{
	emsg := &ErrroMsg{
		Code: ErrorsMap[reflect.TypeOf(err)],
		ErrorType: reflect.TypeOf(err).Elem().Name(),
		Msg:  err.Error(),
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