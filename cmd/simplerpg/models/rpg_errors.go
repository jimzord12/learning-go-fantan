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

func NewWrongBattleActionError(got BattleAction, wants ...BattleAction) *WrongBattleActionError {
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

func NewWrongItemType(got ItemType, wants ...ItemType) *WrongItemType {
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

type NotEnoughGoldError struct {
	Message string
}

func (e *NotEnoughGoldError) Error() string {
	return e.Message
}

func NewNotEnoughGoldError(itemCost int, playerBalance int) *NotEnoughGoldError {
	return &NotEnoughGoldError{
		Message: fmt.Sprintf("[ERROR]: Purchase failed. Insufficient Gold. Cost: (%d), Balance: (%d)", itemCost, playerBalance),
	}
}

///////////////////////////////////////////////////////////

type DoesNotExistInInventory struct {
	Message string
}

func (e *DoesNotExistInInventory) Error() string {
	return e.Message
}

func NewDoesNotExistInInventory(item *Item, player *Character) *DoesNotExistInInventory {
	return &DoesNotExistInInventory{
		Message: fmt.Sprintf("[ERROR]: Item Retrieval Failed. Item: \n%+v\nDoes not exist in Player's Inventory:\n%+v\n", *item, *player),
	}
}

///////////////////////////////////////////////////////////

type NilPointerError struct {
	Message string
}

func (e *NilPointerError) Error() string {
	return e.Message
}

func NewNilPointerError(paramName string) *NilPointerError {
	return &NilPointerError{
		Message: fmt.Sprintf("[ERROR]: Nil Pointer. This Param was nil: (%s)", paramName),
	}
}

///////////////////////////////////////////////////////////
