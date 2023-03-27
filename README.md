# LineBot

* Overview
    * Receive and push messages to users by LineBot.
    * Save the user info and message in MongoDB.
    * Start the service with docker.

## Start service
You will need to set the __*LineBot channel secret and token*__ when you first make up the service.

```
# pull docker images and build service
make build

# start app and mongodb
make up

# shut down all service
make down
```

## API
**Query messages by userID**
```
[GET] 
(url):8000/messages?userID=user001
```

**Receive user info and messages from line webhook and save in mongodb**
```
[POST] 
(url):8000/callback
```

**Push messages to user by linebot**
```
[POST] 
(url):8000/pushMessage
required body {
    userID  string
    text    string
}
```
