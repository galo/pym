# Pym

[Hank Pym](https://en.wikipedia.org/wiki/Hank_Pym) quantum pusher sends people to the quantum void. The Pym service sends documents to [SDM](https://github.azc.ext.hp.com/sdm/sdm), in the hope somebody can recover them.

## Building

```bash
go build
```

## Running

Launch the pym app, it will by defaul listen to port 3000.

```
./pym serve
```

You can validate the server is listening

```
curl --noproxy "*" http://127.0.0.1:3000/ping
```