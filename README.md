# GoLang HTTP Server for Remote WOL Requesting from an CSV Computer List

A HTTP server who sends a Wake On LAN package on an HTTP request.

### Docker Image

[![](https://images.microbadger.com/badges/version/shawnsavour/wkup.svg)](https://hub.docker.com/repository/docker/shawnsavour/wkup "https://hub.docker.com/repository/docker/shawnsavour/wkup") [![](https://images.microbadger.com/badges/image/shawnsavour/wkup.svg)](https://hub.docker.com/repository/docker/shawnsavour/wkup "https://hub.docker.com/repository/docker/shawnsavour/wkup")

### Simple bootstrap UI for easy usage

![Screenshot](https://github.com/shawnsavour/wkup/raw/master/screenshot.PNG)

### Simple REST API to let a machine wake someone up

/api/wakeup/computer/**&lt;hostname&gt;** -  Returns a JSON object

```json
{
  "success":true,
  "message":"Succesfully Wakeup Computer Computer1 with Mac 64-07-2D-BB-BB-BF on Broadcast IP 192.168.10.254:9",
  "error":null
}
```

## Enviroment Variables

| Variable Name | Description |
| ------------- | ------------------------------------------------------------------------------- |
| WOLHTTPPORT   | Define the port on which the webserver will listen to **(Default: 8080)**       |
| WOLFILE       | Path to the Json file containing the list of hosts **(Default: ./computerlist.json)** |


## Commandline arguments

| Commandline argument | Example          | Description                                                                            |
| -------------------- | ---------------- | -------------------------------------------------------------------------------------- |
| --port               | --port 80        | Define the port on which the webserver will listen to **(Default: 8080)**              |
| --file               | --file comp.json  | Path to the CSV file containing the list of hosts **(Default: ./computerlist.json)**        |

## Computer list file CSV layout

### Columns
__&lt;name of the computer&gt;__,__&lt;mac address of the computer&gt;__,__&lt;broadcast ip to send the magic packet&gt;__


### Example
```json
name,mac,ip
[
 {
  "Name": "node-02",
  "Mac": "2c:f0:5d:33:b7:xx",
  "BroadcastIPAddress": "172.16.255.255:16"
 },
 {
  "Name": "node-01",
  "Mac": "2c:f0:5d:33:b7:xx",
  "BroadcastIPAddress": "172.16.255.255:16"
 },
 {
  "Name": "sondc",
  "Mac": "2c:f0:5d:33:b7:xx",
  "BroadcastIPAddress": "172.16.255.255:16"
 }
]
```

## Docker

**Docker Image:** shawnsavour/wkup

```
docker build -t shawnsavour/wkup .
docker run shawnsavour/wkup
```
If you want to run it on a different port (i.e.: 6969) and also want to provide the CSV file on your host:

```
docker run -p 6969:8080 -v $(pwd)/externall-file-on-host.json:/app/computerlist.json shawnsavour/wkup
```

If you want to run the WOL Webserver Process in the Webserver on a different Port:

```
# Used if you run in Network Host Mode
docker run -e "WOLHTTPPORT=9090" -p 9090:9090 -v $(pwd)/externall-file-on-host.json:/app/computerlist.json shawnsavour/wkup
```

This was a good exercise to learn Golang (and refresh my Docker skills).

Thx https://github.com/shawnsavour/WakeupServer for the WOL code, sorry that I stole it from you, because I got no clue how I can inject it into my program. :-(

**If you have any good ideas, I'm open for pull requests.**
