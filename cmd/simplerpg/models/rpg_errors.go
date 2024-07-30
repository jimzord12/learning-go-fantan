package models

import (
	"fmt"
)

type LowStaminaError struct {
	Message string
}

func (e *LowStaminaError) Error() string {
	return e.Message
}

func NewLowStaminaError(msg string) *LowStaminaError {
	return &LowStaminaError{
		Message: msg,
	}
}

///////////////////////////////////////////////////////////

type WrongBattleActionError struct {
	Message string
}

func (e *WrongBattleActionError) Error() string {
	return e.Message
}

func NewWrongBattleActionError(wants []BattleAction, got BattleAction) *WrongBattleActionError {
	msg := fmt.Sprintf("[ERROR]: Wrong Battle Action, WANTS: (%s), GOT: (%s)\n", wants, got)
	fmt.Println(msg)
	return &WrongBattleActionError{
		Message: msg,
	}
}

///////////////////////////////////////////////////////////

type WrongItemType struct {
	Message string
}

func (e *WrongItemType) Error() string {
	return e.Message
}

func NewWrongItemType(wants []ItemType, got ItemType) *WrongItemType {
	return &WrongItemType{
		Message: fmt.Sprintf("[ERROR]: Wrong Item Type, WANTS: (%s), GOT: (%s)\n", wants, got),
	}
}

///////////////////////////////////////////////////////////

type NotCorrectlyInitError struct {
	Message string
}

func (e *NotCorrectlyInitError) Error() string {
	return e.Message
}

func NewNotCorrectlyInitError(object any) *NotCorrectlyInitError {
	return &NotCorrectlyInitError{
		Message: fmt.Sprintf("[ERROR]: The object of type (%T) is not correctly initialized", object),
	}
}

///////////////////////////////////////////////////////////
