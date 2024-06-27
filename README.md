# go_key_database
A simple key value memory database in go with the ability to save/load in/from file

#### Warning
This project is under development os they are some bugs

#### Bugs

- Set/Get value other than string, bool some times don't work

### Usage

Base url: `http://localhost:8080`

- `/api/v1/register?username=test&email=test@gmail.com&password=test&type=User`

Method: GET

Register a new user and return the accound data as JSON 

- `/api/v1/:user/get?key=test&api_key={API_KEY}`

Method: GET

Get the value corresponding in the key parameter

- `/api/v1/:user/get_all?api_key={API_KEY}`

Method: GET

Get the user's data

- `/api/v1/:user/set?key=test&value=542.5234131&type=float64&api_key={API_KEY}`

Method: POST

Set a new value or update it if the key already exist

- `/api/v1/:user/stats?api_key={API_KEY}`

Method: GET

Return some statistics about the server (if the user have the permission)

- `/api/v1/:user/admin/save?api_key={API_KEY}`

Method: GET

Save the database into a file

- `/api/v1/:user/admin/load?api_key={API_KEY}`

Method: GET

Load the database from a file 