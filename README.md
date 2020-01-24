
# NumberServer

A golang TCP server and TCP client implementation. The clients connect to the server and send as many messages as they want unless the type a non-numeric input. All the clients can be closed if some of them writres "terminate"

```
run make help

Usage:
  help                Print this help message
  build-server        Build the server
  build-client        Build sthe client
  run-server          Run the server
  run-client          Run the client
  docker-server-run   Build the image and run the container for the server
  docker-client-run   Build the image and run the container for the client
```
Note: There is an alternative way to run the server with docker. 
      Dockerize client is still pending.
