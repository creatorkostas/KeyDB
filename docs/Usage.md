# Usage

Base url: `http://localhost:8080`

- Register a new user and return the accound data as JSON 

```console
curl http://localhost:8080/api/v1/register?username={USERNAME}&email={EMAIL}&password={PASSWORD}&type=User
```
> Method: GET


- Get the value corresponding in the key parameter

```console
curl http://localhost:8080/api/v1/{USERNAME}/get?key=test&api_key={API_KEY}
```

> Method: GET


- Get the user's data

```console
curl http://localhost:8080/api/v1/{USERNAME}/get_all?api_key={API_KEY}
```

> Method: GET


- Set a new value or update it if the key already exist

**Available types:** string, int, float64, float32, bool

```console
curl -X POST http://localhost:8080/api/v1/{USERNAME}/set?key=test&value=542.5234131&type=float64&api_key={API_KEY}
```

> Method: POST



- Return some statistics about the server (if the user have the permission)

```console
curl http://localhost:8080/api/v1/{USERNAME}/stats?api_key={API_KEY}
```

> Method: GET


- Save the database into a file

```console
curl http://localhost:8080/api/v1/{USERNAME}/admin/save?api_key={API_KEY}
```

> Method: GET

- Load the database from a file 

```console
curl http://localhost:8080/api/v1/{USERNAME}/admin/load?api_key={API_KEY}
```

> Method: GET

