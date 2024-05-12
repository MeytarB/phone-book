package types

type UserError string

const DuplicateNumberError UserError ="error: contact number already exists"
const DuplicateNameError UserError = "error: contact full name already exists"
const ValidNumberError  UserError = "error: not a valid number"
const ExistingNumberError UserError = "error: new phone number already exist"
const ExistingFullNameError UserError = "error: new full name already exist" 
const NumberNotFoundError UserError = "error: contact was not found"

func AllUserErrors () []UserError {
	var userErrors =[]UserError{
		DuplicateNumberError, 
		DuplicateNameError, 
		ValidNumberError, 
		ExistingFullNameError, 
		ExistingNumberError, 
		NumberNotFoundError }

	return userErrors;
}