package structs

import (
	"fmt"
	"time"

	"github.com/jimzord12/learning-go-fantan/cmd/sections/helpers"
)

// / Declaring a Struct's Type
type BankAccount struct {
	AccountNumber string
	Balance       float64
}

type AuditInfo struct {
	CreatedAt    time.Time
	LastModified time.Time
}

type Customer struct {
	Name     string
	accounts []BankAccount
	AuditInfo
}

func Main() {
	/// Declaring a Struct
	var account1 BankAccount // Go does not like this declaration syntax :P
	account1 = BankAccount{
		AccountNumber: "Acc1",
		Balance:       250.0,
	}
	fmt.Println(account1)

	account2 := BankAccount{
		AccountNumber: "Jimzord",
		Balance:       6000.0,
	}
	fmt.Println(account2)

	account3 := &BankAccount{ // This is a pointer
		AccountNumber: "Jimzord123",
		Balance:       2500.0,
	}
	fmt.Printf("Account 3: %p\n", account3)

	/// Using Methods
	account1.DisplayBalance()
	account2.DisplayBalance()
	account3.DisplayBalance() // 👉 No need to dereference

	helpers.Divider()

	if err := account1.Withdraw(249.50); err != nil {
		fmt.Println(err)
	}
	if err := account2.Withdraw(6005); err != nil {
		fmt.Println("\n", err)
	}

	account1.DisplayBalance()
	account2.DisplayBalance()

	helpers.Divider()

	customer := Customer{
		Name: "AAAAss",
		AuditInfo: AuditInfo{
			CreatedAt:    time.Now(),
			LastModified: time.Now(),
		},
	}

	customer.AddAccount(BankAccount{AccountNumber: "Daily", Balance: 1500.0})
	customer.AddAccount(BankAccount{AccountNumber: "Saving", Balance: 50000.0})

	customer.DisplayAccountsBalance()

	helpers.Divider()

	customer2 := newCustomer("PAOK")
	customer2.AddAccount(*newBankAccount("2221114433"))
	customer2.accounts[0].Deposit(15000.0)
	customer2.accounts[0].DisplayBalance()
	customer2.accounts[0].Withdraw(2500.0)
	customer2.accounts[0].DisplayBalance()
	// Getting the size of a Struct
	// fmt.Println("Size of the Account Struct Var:", unsafe.Sizeof(account1))
	// fmt.Println("Size of the Account Struct Pointer:", unsafe.Sizeof(account3))
}

// / Declaring Struct Methods

// Value Receiver - Receives a COPY of a Struct (Immutability)
// GP (Good Practise): struct must be small and we want immutability
func (ba BankAccount) DisplayBalance() {
	fmt.Printf("[VIEW] Account: %s, Balance: %.2f\n", ba.AccountNumber, ba.Balance)
}

// Pointer Receiver - Receives a Reference of the ORIGINAL struct
// GP (Good Practise): good choice for large structs and when we want to mutate the original struct
func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.Balance < amount {
		return fmt.Errorf("[error]: Account: %s, msg: withdraw failed, reason: insufficient balance", ba.AccountNumber)
	}
	fmt.Printf("[WITHDRAW] [%.2f], from Account: %s\n", amount, ba.AccountNumber)
	ba.Balance -= amount
	return nil
}

func (c *Customer) AddAccount(account BankAccount) {
	c.accounts = append(c.accounts, account)
}

func (c Customer) DisplayAccountsBalance() {
	for _, acc := range c.accounts {
		acc.DisplayBalance()
	}
}

// / Constuctor Functions
func newCustomer(name string) *Customer {
	return &Customer{
		Name:      name,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}
}

func newBankAccount(num string) *BankAccount {
	return &BankAccount{
		AccountNumber: num,
	}
}
