# GO inventory application OTI

Hi! this is inventory application using dessign pattern MVC


## Requirements to run project
	
* [**Go programming language**](https://go.dev/doc/install)

## Run project

Execute project with next commands:
	
	$ go mod tidy
	$ go run .


## API Reference

### Token by header Authorization

| Key | Value     |
| :-------- | :------- | 
| `Authorization` | `Bearer eyJhbGciOiJIUzI1NiIsInR5...` | 


```javascript
fetch(url, {
  method: 'POST',
  headers: {
    'Authorization': 'Bearer ' + token,
    'Content-Type': 'application/json',
  },
})
```
---
#### Login application

```json
  POST /login
  {
	"email":"luispfcanales@gmail.com",
	"password":"1234"
  }
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required** |
| `password` | `string` | **Required** |

---
- #### Api users

```http
  GET /api/users/all
```
---
- #### Api networks

```http
  GET /api/network/all
```
```http
  GET /api/network/all/resume
```
---
- ####  Api Person

```http
  GET /api/person/:dni
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `dni`      | `int` | **Required**. DNI of item to fetch |

```http
  GET /api/person/all
```
```json
PUT /api/person/
{
	"dni":72453560,
	"first_name":"LUIS ANGEL",
	"last_name":"PFUÑO CANALES",
	"birthdate":"2023-12-14T22:52:13.202Z",
	"address":"JR TEST 222"
}
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `dni` | `string` | **Required**. DNI to update |
| `first_name` | `string` | **Required** |
| `last_name` | `string` | **Required** |
| `birthdate` | `string` | **Required** |
| `address` | `string` | **Required** |

```json
POST /api/person
{
	"dni":72453560,
	"first_name":"LUIS ANGEL",
	"last_name":"PFUÑO CANALES",
	"birthdate":"2023-12-14T22:52:13.202Z",
	"address":"JR TEST 222"
}
```
```http
  DELETE /api/person/:dni
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `dni`      | `int` | **Required**. DNI to delete |

---
- #### Api campus
```http
  GET /api/campus/:id
```
```http
  GET /api/campus/all
```
```json
PUT /api/campus
{
	"id":"a468a765-932d-49b5-8090-4dfddfrb37a2",
	"abbreviation":"test",
	"name":"put testing",
	"address":"updating",
	"state":false
}
```
```json
POST /api/campus
{
	"abbreviation":"test",
	"name":"put testing",
	"address":"updating",
	"state":false
}
```
```http
  DELETE /api/campus/id
```

---
- #### Api zone
```http
  GET /api/zone/:id
```
```http
  GET /api/zone/all
```
