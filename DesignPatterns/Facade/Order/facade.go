package facade

import "fmt"


type OrderFacade struct {
	account *Account
	wallet *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger *Ledger
}

func NewOrderFacade(accountID string, code int) *OrderFacade {
	fmt.Println("Starting create an account")
	orderFacade := &OrderFacade{
		account: NewAccount(accountID),
		wallet: NewWallet(),
		securityCode: NewSecurityCode(code),
		notification: &Notification{},
		ledger: &Ledger{},
	}
	fmt.Println("Account created successfully")
	return orderFacade
}

func (orderFacade *OrderFacade) AddMoneyToWallet(accountID string, code, amount int) error {
	fmt.Println("Start to add money to the wallet")
	if err := orderFacade.account.CheckAccount(accountID); err != nil {
		return err
	}
	if err := orderFacade.securityCode.CheckSecurityCode(code); err != nil {
		return err
	}
	orderFacade.wallet.CreditWallet(amount)
	orderFacade.notification.SendWalletCreditNotification()
	orderFacade.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (orderFacade *OrderFacade) DeductMoneyFromWallet(accountID string, code, amount int) error {
	fmt.Println("Start debit money from wallet")
	if err := orderFacade.account.CheckAccount(accountID); err != nil {
		return err
	}
	if err := orderFacade.securityCode.CheckSecurityCode(code); err != nil {
		return err
	}
	if err := orderFacade.wallet.DebitWallet(amount); err != nil {
		return err
	}
	orderFacade.notification.SendWalletDebitNotification()
	orderFacade.ledger.MakeEntry(accountID, "debit", amount)
	return nil
}
