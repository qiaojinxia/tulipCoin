package utils

import (
	"encoding/json"
	"errors"
	"reflect"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

type NetError error
type StackError error

type ErrorInfo struct {
	Code int
	Name string
}

var ErrorsMap = make(map[reflect.Type]*ErrorInfo)
func init(){
	ErrorsMap[reflect.TypeOf(NetErroWarp("").error)] = &ErrorInfo{
		Code: 100,
		Name: "NetError",
	}
	ErrorsMap[reflect.TypeOf(StackErrorWarp("").error)] = &ErrorInfo{
		Code: 200,
		Name: "StackError",
	}
}


type TError struct {
	error

}
func StackErrorWarp(msg string) *TError{
	return &TError{error:StackError(errors.New(msg))}
}

func NetErroWarp(msg string) *TError{
	return &TError{error:NetError(errors.New(msg))}
}

type ErrroMsg struct {
	Code int `json:"code"`
	ErrorType string `json:"error_type"`
	Msg string `json:"msg"`
}
func(t *TError) ConverToJsonInfo() string{
	info := ErrorsMap[reflect.TypeOf(t.error)]
	emsg := &ErrroMsg{
		Code: info.Code,
		ErrorType: info.Name,
		Msg:  t.error.Error(),
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