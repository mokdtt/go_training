package bank_test

import (
	"fmt"
	"testing"

	bank "go_training/ch09/ex01"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		bank.Withdraw(50)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		bank.Withdraw(500)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 250; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
