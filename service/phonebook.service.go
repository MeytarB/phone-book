package service

import (
	"github.com/MeytarB/phone-book/types"
)

type PhoneBookService interface {
	AddContact(*types.ContactType) error
	EditContact(string, string, *types.ContactType) error
	ShowAllContacts(int64) ([]*types.ContactType, error)
	FindContactByName(string, string) (*types.ContactType, error)
	FindContactByNumber(string) (*types.ContactType, error)
	DeleteContact(string, string) error
	DeleteAll() error
}
