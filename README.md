# Cacher 💽

Open-source micro-service written in GoLang. This micro-service saves values into memory

### Installer

Linux and macOS:

Use: ``
bash ./start.sh
``

### Security

This micro-service uses API Key in header HTTP request 

|  Type   |   EndPoints   |            Description            |
| ------- | ------------- | --------------------------------- |
| GET     | /v1/get       | This endpoint returns value by key          |
| POST    | /v1/set       | This endpoint saves value into memory          |
| DELETE  | /v1/remove    | This endpoint removes value from memory          |
| GET     | /v1/valid     | This endpoint verifies if value exist              |
| PUT     | /v1/replace   | This endpoint replaces already exist value       |
| GET     | /v1/ping      | This endpoint returns only plain text "Pong!"         |

Written by JustDrven