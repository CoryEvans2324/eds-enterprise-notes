# eds-enterprise-notes

Notes app for our Enterprise Software Developement class.


|package|docs link|
|:--|:--|
|net/http docs|https://pkg.go.dev/net/http|
|html/template docs|https://pkg.go.dev/html/template|
|gorilla/mux|https://pkg.go.dev/github.com/gorilla/mux|
|database/sql|https://pkg.go.dev/database/sql|


## Front end interface
The front end will be html, css & js. [tailwindcss](https://tailwindcss.com/docs) will be used for styling.

NodeJS & NPM needs to be installed.
Build the css with `npm run build`

## Information
### Valid text patterns for search and analysis features
- a sentence with a given prefix and/or suffix.
- a phone number with a given area code and optionally a consecutive sequence of numbers that are part
of that number.
- an email address on a domain that is only partially provided.
- text that contains at least three of the following case-insensitive words: meeting, minutes, agenda,
action, attendees, apologies.
- a word in all capitals of three characters or more.


## Server Routes / Functions

|route|method|request type|response type|description|
|:--|:-:|:-:|:-:|:--|
|`/`|`GET`|http|html|index page|

### Account 
|route|method|request type|response type|description|
|:--|:-:|:-:|:-:|:--|
|`/user/`|`GET`|http|html|View your account|
|`/user/create`|`GET`|http|html|html for creating an account|
|`/user/create`|`POST`|form|redirect|Creates an account and signs in|
|`/user/signin`|`GET`|http|html|html for signing in|
|`/user/signin`|`POST`|form|redirect|Signs into an account|
|`/user/signout`|`GET`|http|redirect|Signs out of an account|
|`/user/edit`|`GET`|http|html|recieve account settings edit page|
|`/user/edit`|`POST`|form|redirect|overwrite/change account settings|
|`/user/search`|`POST`|json|json|Search for users based on their username|

### Notes
|route|method|request type|response type|description|
|:--|:-:|:-:|:-:|:--|
|`/notes/`|`GET`|http|html|view all notes|
|`/notes/search`|`POST`|json|json|search for notes by providing a text pattern|
|`/notes/create`|`GET`|http|html|A form to create a note|
|`/notes/create`|`POST`|form|redirect|creates a note|
|`/notes/<id>`|`GET`|http|html|view a note|
|`/notes/<id>/edit`|`GET`|http|html|A form to edit note data|
|`/notes/<id>/edit`|`POST`|form|redirect|edits/overwrites a note's content or settings|
|`/notes/<id>/share`|`POST`|json|json|Shared a note with a list of users|
|`/notes/<id>/share`|`DELETE`|json|json|remove access from a list of users|



## Datatypes / Tables
### Account / User
|Attribute|Type|
|--:|:--|
|ID|`int`|
|Username|`string`|
|Password|`string`|
|ShareNewNotes|`bool`|

### Note
|Attribute|Type|
|--:|:--|
|ID|`int`|
|Title|`string`|
|Completion Time|`Date`|
|Status|`string`|
|DelegationUserID|`int`|
|Owner|`int`|

### Note Access
|Attribute|Type|
|--:|:--|
|NoteID|`int`|
|UserID|`int`|
|Permissions|`int`|

### User Pre-Shared Friends
|Attribute|Type|
|--:|:--|
|UserID|`int`|
|FriendID|`int`|
|Permissions|`int`|
