# Usage

**Base url: `http://localhost:8080`**

---
## Everyone

### Register a new user
- Register a new user and return the accound data as JSON 

```/api/v1/register?username={USERNAME}&email={EMAIL}&password={PASSWORD}&type=User```
> Method: GET

Fields:

| Field    | Required | Accepted values        | Description                  |
| :------- | :------: | :--------------------- | ---------------------------- |
| username |    ✅     | Any string             | The username for the account |
| email    |    ✅     | Any string             | The email for the account    |
| password |    ✅     | Any string             | The password for the account |
| type     |    ✅     | Admin \| User \| Guest | Type of the account          |

Response:
```json
{
  "response": {
    "Account": {
      "Username": "test",
      "Api_key": "07c4a18c3648329d",
      "Email": "test@gmail",
      "Password": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
      "Active": true,
      "Tokens": 100000,
      "Rate_reset": 21600000000000,
      "Burst_time": 100000000,
      "Burst_tokens": 10,
      "Tier": {
        "Type": 1,
        "Rules": {
          "Admin": false,
          "Add": true,
          "Get": true,
          "Change_password": true,
          "Change_api_key": true,
          "Analytics": false
        }
      }
    }
}
```
---
## Users
### Get data
- Get data saved in the database as JSON 

```api/v1/{USERNAME}/get?api_key={API_KEY}&key={DATA_KEY}```
> Method: GET

Fields:

| Field   | Required | Accepted values | Description                                                                              |
| :------ | :------: | --------------- | ---------------------------------------------------------------------------------------- |
| api_key |    ✅     | Any string      | The api key related to the account (in the response from the register prossess)          |
| key     |          | Any string      | The key under wich the data was saved. If it is not provided it will return all the data |


Response:
```json

```
---

### Set data
- Save data in the database 

```api/v1/{USERNAME}/set?api_key={API_KEY}&key={DATA_KEY}&type={DATATYPE}&value={DATA}```
> Method: GET

Fields:

| Field   | Required | Accepted values                             | Description                                                                     |
| :------ | :------: | ------------------------------------------- | ------------------------------------------------------------------------------- |
| api_key |    ✅     | Any string                                  | The api key related to the account (in the response from the register prossess) |
| key     |    ✅     | Any string                                  | The key under wich the data will be saved                                       |
| type    |    ✅     | int \| string \| float32 \| float64 \| bool | The type of the data to be saved                                                |
| value   |    ✅     | The type specified in 'type' field          | The data itself to be saved                                                     |


Response:
```json

```
---
### Get server stats
- Get server stats (response times, codes)

> [!IMPORTANT]  
> The user MUST have the permission set to their account to do so!

```api/v1/{USERNAME}/stats?api_key={API_KEY}```
> Method: GET

Fields:

| Field   | Required | Accepted values                             | Description                                                                     |
| :------ | :------: | ------------------------------------------- | ------------------------------------------------------------------------------- |
| api_key |    ✅     | Any string                                  | The api key related to the account (in the response from the register prossess) |



Response:
```json
{
  "gin.latency": {
    "15m.rate": 0.16949789568319829,
    "1m.rate": 0.06972056737346893,
    "5m.rate": 0.12728177813430272,
    "75%": 96838,
    "95%": 70867629,
    "99%": 70867629,
    "99.9%": 70867629,
    "count": 12,
    "max": 70867629,
    "mean": 5946960.916666667,
    "mean.rate": 0.05767223893543906,
    "median": 24957,
    "min": 3953,
    "stddev": 19574384.583412822
  },
  "gin.request": {
    "15m.rate": 0.16949789568319829,
    "1m.rate": 0.06972056737346893,
    "5m.rate": 0.12728177813430272,
    "count": 12,
    "mean.rate": 0.05767223058224958
  },
  "gin.status.200": {
    "15m.rate": 0.16882843197669722,
    "1m.rate": 0.033796546698952164,
    "5m.rate": 0.12255321263915472,
    "count": 6,
    "mean.rate": 0.03289352342563399
  },
  "gin.status.401": {
    "15m.rate": 0.1841198930725079,
    "1m.rate": 0.06961079550955261,
    "5m.rate": 0.15665677386557786,
    "count": 3,
    "mean.rate": 0.03344879382509197
  },
  "gin.status.404": {
    "15m.rate": 0.16035602629375786,
    "1m.rate": 0.02010052002917449,
    "5m.rate": 0.10418355129063632,
    "count": 2,
    "mean.rate": 0.009618772913926352
  },
  "gin.status.500": {
    "15m.rate": 0.15837791326735629,
    "1m.rate": 0.0060394766844637125,
    "5m.rate": 0.09931706075828191,
    "count": 1,
    "mean.rate": 0.0048060204797703825
  }
}
```
---

## Administrators

### Save to disk
- Save all in memory data to disk

```api/v1/{USERNAME}/admin/save?api_key={API_KEY}```
> Method: GET

Fields:

| Field   | Required | Accepted values                             | Description                                                                     |
| :------ | :------: | ------------------------------------------- | ------------------------------------------------------------------------------- |
| api_key |    ✅     | Any string                                  | The api key related to the account (in the response from the register prossess) |


Response:
```json

```
---
### Load from disk
- Load the saved data from disk to memory 

> [!WARNING]  
> All the current data in memory will be overwriten.

```api/v1/{USERNAME}/admin/load?api_key={API_KEY}```
> Method: GET

Fields:

| Field   | Required | Accepted values                             | Description                                                                     |
| :------ | :------: | ------------------------------------------- | ------------------------------------------------------------------------------- |
| api_key |    ✅     | Any string                                  | The api key related to the account (in the response from the register prossess) |


Response:
```json

```
---
### Disable Administrator account creation
- Disable the account type `Admin` in the registration endpoint

```api/v1/{USERNAME}/admin/disableAdmin?api_key={API_KEY}```
> Method: GET

Fields:

| Field   | Required | Accepted values                             | Description                                                                     |
| :------ | :------: | ------------------------------------------- | ------------------------------------------------------------------------------- |
| api_key |    ✅     | Any string                                  | The api key related to the account (in the response from the register prossess) |


Response:
```json
{
  "response": "Admin register disabled"
}
```
---
### Enable Administrator account creation
- Enable the account type `Admin` in the registration endpoint

```api/v1/{USERNAME}/admin/enableAdmin?api_key={API_KEY}```
> Method: GET

Fields:

| Field   | Required | Accepted values                             | Description                                                                     |
| :------ | :------: | ------------------------------------------- | ------------------------------------------------------------------------------- |
| api_key |    ✅     | Any string                                  | The api key related to the account (in the response from the register prossess) |


Response:
```json
{
  "response": "Admin register enabled"
}
```
---
