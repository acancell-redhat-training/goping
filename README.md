# goping
Connect: `oc port-forward pod/$podname 8080:8080`

Request: `curl -w '\n' localhost:8080/ping`

Reply: *{"message": "OK: PONG"}*
