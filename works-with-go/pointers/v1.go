/* In Go, when you call a function or a method the arguments are copied. */
/* When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
   To fix this with pointers. Pointers let us point to some values and then let us change them.
   So rather than taking a copy of the Wallet, we take a pointer to the wallet so we can change it.
   the code below using (*w) is absolutely valid. However, the makers of Go deemed this notation cumbersome,
   so the language permits us to write w.balance, without explicit dereference.

   Pointers
   1. Go copies values when you pass them to functions/methods so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
   2. The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference.
   Examples could be very large data or perhaps things you intend only to have one instance of (like database connection pools).
*/

package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	startingBalance := Bitcoin(20)
	w := Wallet{startingBalance}
	w.Deposit(Bitcoin(100))
	w.Withdraw(Bitcoin(50))
	w.Deposit(Bitcoin(60))
	fmt.Println(w.Balance())
}
