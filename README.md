# NoobGo

NoobGo is a simple and modest web framework written in Golang

## Quick Start

### Prerequisites

Having go installed on your computer.\
[How to install Go](https://golang.org/doc/)

### Running

```bash
# Clone our repo
git clone https://github.com/khantzen/NoobGo.git noobGo

# Change directory to your cloned repo
cd noobGo

# Install dependencies
./install_dependencies.sh

# Build main.go
go build main.go

# Run main
./main
```

Go to [http://localhost:8080](http://localhost:8080) in your browser

## How to use

Here, we will describe how to display a user list from our mysql database using our framework

### Prerequisites
Having a MySql database with following table

```mysql
Create table users (
  id int primary key auto_increment,
  username varchar(250) not null,
  creationDate timestamp not null default CURRENT_TIMESTAMP
);

Insert Into users (username) Values
('Tortin'), ('Raydenawon'), ('Ecera'), ('Eoauymdyst');
```

**Actually database connection is hard coded (my bad) so make sure that it is correctly set for your environnement
in function _InitDatabase_ in _repository/database.go_**

```go
func InitDatabase() (*DB, error) {
	db, err := sql.Open("mysql", "root:root@/noobGoDb")

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
```

### Model
At this time, we have two kinds of model in NoobGo: _db_ and _view_

  * _db_ models goal is to map your database's table columns.
  * _view_ models goal is to regroup variables that you will need in your views.
  
 We will create two models files _model/db/userDb.go_ and _model/view/displayUserListView.go_
 
```
project
|__model
    |__db
        |  userDb.go
    |__view
        |  displayUserlistView.go
        |  mainView.go
```

```go
// usersDb.go
package db

import (
	"github.com/go-sql-driver/mysql"
)

type UserDb struct {
	Id           int
	Username     string
	CreationDate mysql.NullTime
}

```
```go
// displayUserListView.go
package view

import "time"

type DisplayUserListView struct {
	UserList []userDisplay
}

type userDisplay struct {
	UserName     string
	CreationDate time.Time
}
```

There is already a _mainView.go_ file in _model/db/view_ do not delete it.

### Repository
Repository package contains method having access to your database. 
Actually only mysql driver is supported.

In repositories create a new file named _userRepository.go_

```
project
|__repository
    | database.go
    | userRepository.go
```

In this file add function _GetUserList_ with following content

```go
import (
	"log"
	dbm "../model/db"
)

func (db *DB) GetUserList() (*[]dbm.UserDb) {
	rows, err := db.Query(
		"Select *" +
			" From users")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var userList = []dbm.UserDb{}

	for rows.Next() {

		var userDb = dbm.UserDb{}

		if err := rows.Scan(
			&userDb.Id,
			&userDb.Username,
			&userDb.CreationDate); err != nil {
			log.Fatal(err)
		}

		userList = append(userList, userDb)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &userList
}
```

Once it's done, go to _repository/database.go_ file.\
Here you'll see an empty interface name Repository.\
Add the method _GetUserList_ that we've just created to it.

```go
type Repository interface {
	GetUserList() (*[]db.UserDb)
}
```

### Libraries
Not implemented yet

### Controller
Controller package contains method that will be called when an url is called, 
we'll see later how to bind an url to a controller method, for now create a 
new file in controller named _user.go_

```
project
|__ controller
    |   ctrl.go
    |   user.go
    |   welcome.go
```

```go
package controller

import (
	"net/http"
	vm "../model/view"
	vr "../views"
)

func (ctrl *Ctrl) UserDisplayList(w http.ResponseWriter, r *http.Request) {
	userList := ctrl.Repository.GetUserList()

	var displayViewModel = vm.DisplayUserListView{}

	for _, u := range *userList {
		displayViewModel.UserList =
			append(displayViewModel.UserList,
				vm.UserDisplay{
					UserName:     u.Username,
					CreationDate: u.CreationDate.Time})
	}

	vr.Render("user/list", displayViewModel, w)
}
```

There is no fixed standard to name your controller's method, but I usually follow this pattern [ControllerName][Action]

### Router
_routing.go_ is part of the config package. 

```
project
|__config
    |  routing.go
```

In this file the method _SetRouting_ bind url path to controller method.\

Let's add our new route to this method

```go
handler.HandleFunc("^(?i)/user/list", ctrl.UserDisplayList)
```

So now everytime that url [http://localhost:8080/user/list](http://localhost:8080/user/list) is called,
NoobGo will execute our previous method _UserDisplayList_.


### Views

#### Html
#### Css
#### Javascript

## Demo
A demo will be available soon on another repositories

## Deploy on your server
Not finished yet