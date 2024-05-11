This app is a phone book server

it is a simple server for you to manage all of your contacts using a friendly API

each contact is represented by JSON with the following structure:
{
    "firstname":""
    "lastname":""
    "phonenumber":""
    "address":""
}

this phone book supports the following :

* Add a new contact: 
endpoint: POST http://localhost:3000/app/phonebook/add
body: contact details json 
result: error if there was

* Edit existing contact
endpoint: PATCH http://localhost:3000/app/phonebook/edit/:firstname/:lastname
params: first name and last name (should be the original details.)
body: updated contact details json 
result: error if there was
example: http://localhost:3000/app/phonebook/edit/taylor/swift

* Show all contacts
endpoint: GET http://localhost:3000/app/phonebook/find-all-contacts/:page
params: page number starting from 1
result: array of maximum 10 contacts or error if there was
note: for the next 10 contacts increase the page number by 1
example: http://localhost:3000/app/phonebook/find-all-contacts/1


* Find a contact by name
endpoint: GET http://localhost:3000/app/phonebook/find-by-name?firstname=<firstname>&lastname=<lastname>
query params: first name and last name
result: json of the contact was found or error
example: http://localhost:3000/app/phonebook/find-by-name?firstname=walt&lastname=disney

* Find a contact by number
endpoint: GET http://localhost:3000/app/phonebook/find-by-number?phonenumber=<phonenumber>
query params: phone number as string
result: json of the contact was found or error
example: http://localhost:3000/app/phonebook/find-by-number?phonenumber=050123456


* Delete contact
endpoint: DELETE http://localhost:3000/app/phonebook/delete-contact?firstname=<firstname>&lastname=<lastname>
query params: first name and last name
result: error if there was
example: http://localhost:3000/app/phonebook/find-by-number?phonenumber=0500012

* Delete all contacts
endpoint: DELETE http://localhost:3000/app/phonebook/delete-all
result: error if there was
note: be careful from that one - it will delete your details too


Some important extra notes :
* when starting the app for the first time , you will get your phone details under the name "my phone" 

* This phone book will help you not calling accidently a person you didnt mean to call :) since it doesn't allow:
1. duplicates phone number ( you can't have 2 contacts sharing the same number) 
2. 2 contacts with the same full name.


in order to run the program , run in bash:
docker build -t phone-book-app .
docker run -p 3000:3000 phone-book-app
