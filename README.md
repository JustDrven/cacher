# Cacher 💽

Easy microservice for saving caches into memory

### Installer

Only send the command ``
bash ./start.sh
``

### Security

This microservice uses API Key in request header

|   EndPoints   |            Description            |
| ------------- | --------------------------------  |
| /v1/get       | Return cache by your value        |
| /v1/set       | Insert cache into memory          |
| /v1/remove    | Delete cache from memory          |
| /v1/valid     | Check if cache exist              |
| /v1/replice   | Replace already exist cache       |

Written by JustDrven