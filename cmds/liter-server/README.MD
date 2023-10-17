
# liter-server

**liter-server** is a minecraft reverse proxy with a simple manage dashboard.

If you want to run liter server as a systemd service, we have a service file at [litermc.service](./litermc.service)


## Setup

All the data & configs are saved under current directory (in windows) or `/opt/litermc` (other systems). Please makesure that this application have permission to the directory.


Set up commands for run liter-server use your own user with right perimission
```sh
sudo mkdir -p /opt/litermc
sudo chown $(whoami) -R /opt/litermc
```

## Configs

### `config.yml`

```yml
debug: false # if you want to see debug logs, set this to true
online-mode: true
server-addr: :25565 # server address in <host>:<port>
dashboard:
    enable: false # enable dashboard
    addr: 127.0.0.1:25580 # dashboard address in <host>:<port>, use 127.0.0.1 as the host to keep it private
enable-whitelist: false # enable whitelist or not
enable-ip-whitelist: false # enable IP whitelist or not
servers: # subserver list
    - id: main # the subserver ID, must be unique
      target: 127.0.0.1:25665 # subserver connect address
      names: # domain lists that can connect to the server
        - localhost
        - anotherdomain.example.com
      handle-ping: false # response the motd below for client's server list ping
      motd: Welcome to the server # this motd will only be sent if `handle-ping` is true
      motd-failed: |- # this motd will only be sent if `handle-ping` is false and could not ping the subserver
        The Server is closed
        §aGreen text
```

### `whitelist.json`

```json
{
	"players": [
    {
      "name": "ckupen",
      "id": "7a0ba4fe-e6ec-4bfe-99fc-56bf677a15a5"
    },
    "example_user" // When you edit this file, you can also just give a string and it will automatically parse into the correct structure above
  ],
	"ips": [
		"127.0.0.1",      // IP can be a single IP
		"192.168.1.0/24", // or a IP network as well
	]
}
```

### `blacklist.json`

Same as [`whitelist.json`](#whitelist.json) above

### `users.json`

This is the user database for dashboard.  
You should **NEVER** edit this file.  
However, if you forget your user password, you can delete this file and the root user will be regenerated after reloading.

### `server.hmac.private_key`

This is the JWT signature key for dashboard.  
You should **NEVER** edit this file, and you should **keep this file PRIVATE**.  

