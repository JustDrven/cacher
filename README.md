# Cacher 💽

Easy microservice for saving caches into memory

### Installer

Only send the command ``
bash ./start.sh
``

### Security

This microservice uses API Key in request header

|  Type   |   EndPoints   |            Description            |
| ------- | ------------- | --------------------------------- |
| GET     | /v1/get       | Return cache by your key          |
| POST    | /v1/set       | Insert cache into memory          |
| DELETE  | /v1/remove    | Delete cache from memory          |
| GET     | /v1/valid     | Check if cache exist              |
| PUT     | /v1/replice   | Replace already exist cache       |
| GET     | /v1/ping      | Return plain text "Pong!"         |

Written by JustDrven