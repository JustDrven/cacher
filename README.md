# Cacher 💽

Open-source micro-service written in GoLang. This micro-service saves values into memory

### Installer

##### Linux and macOS:

Use: ``
bash ./start.sh
``

### Security

This micro-service uses API Key in header HTTP request 

|  Type   |   EndPoints   |            Description            |
| ------- | ------------- | --------------------------------- |
| GET     | /v1/get       | Returns value by key              |
| POST    | /v1/set       | Saves value into memory           |
| DELETE  | /v1/remove    | Removes value from memory         |
| GET     | /v1/valid     | Verifies if value exist           |
| PUT     | /v1/replace   | Replaces already exist value      |
| GET     | /v1/ping      | Returns only plain text "Pong!"   |

Written by JustDrven