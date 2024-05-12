package service

import (
	"github.com/MeytarB/phone-book/types"
)

type PhoneBookService interface {
	AddContact(*types.Contact) error
	EditContact(string, string, *types.Contact) error
	ShowAllContacts(int64) ([]*types.Contact, error)
	FindContactByName(string, string) (*types.Contact, error)
	FindContactByNumber(string) (*types.Contact, error)
	DeleteContact(string, string) error
	DeleteAll() error
}


func IsUserError(err error) bool {
	allUserErrors := types.AllUserErrors()
	for _, ue := range allUserErrors {
		if string(ue) == err.Error() {
		return true
		}
	}
	return false
}							
