# Cacher 💽

Open-source micro-service written in GoLang. This micro-service saves values into memory.

### Installer

##### Linux and macOS:

Use: 
``
bash ./start.sh
``

### Security

This micro-service uses API Key in header HTTP request 

|  Type   |   Endpoints              |            Description            |
| ------- | ------------------------ | --------------------------------- |
| GET     | /query/cache             | Returns value by key              |
| GET     | /query/cache/validation  | Verifies if value exist           |
| POST    | /mutation/cache          | Saves value into memory           |
| DELETE  | /mutation/cache          | Removes value from memory         |
| PUT     | /mutation/cache          | Replaces already exist value      |
| GET     | /service/ping            | Returns only plain text "Pong!"   |

Written by JustDrven