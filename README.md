# eve-cli

eve-cli is a cli tool to interact with the eve api, to get data about the game. It is a work in progress and far from finished

# Examples

To look at a character's wallet:
## Linux
```
#install go first
go build
./eve-cli login
./eve-cli character set <<character id>>
./eve-cli wallet
```
## Windows
```
#install go first
go build
.\eve-cli.exe login
.\eve-cli.exe character set <<character id>>
.\eve-cli.exe wallet
```

# Setting up the game account
1. Login to https://developers.eveonline.com/applications
2. Create new application
3. Fill out as: name "eve-cli"
4. Connection type: Authentication & API Access
5. Permissions: select as many as you would like to have access to
6. Callback URL: http://localhost:8080/oauth2/redirect
7. Create

