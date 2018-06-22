# ghostman [![Build Status](https://travis-ci.org/newly12/ghostman.svg?branch=master)](https://travis-ci.org/newly12/ghostman) [![Go Report Card](https://goreportcard.com/badge/github.com/newly12/ghostman)](https://goreportcard.com/report/github.com/newly12/ghostman)

concurrent and template based curl written in golang

## Usage

```bash
$ export GHOSTMAN_TMPLDIR=examples/templates
$ ./ghostman -f examples/localhost run get index
12:48:54.456 GetDefaultsFromEnv ▶ DEBU 001 load defaults from env: {/Users/<user>/go/src/github.com/newly12/ghostman/examples/config/rc.yml {examples/templates 2000}}
12:48:54.457 run ▶ DEBU 002 load host file
12:48:54.458 run ▶ DEBU 003 creating channels
12:48:54.458 run ▶ DEBU 004 generate requests
12:48:54.458 run ▶ DEBU 005 process request channel
12:48:54.458 Execute ▶ DEBU 006 working on http://localhost:80/
12:48:54.461 Println ▶ INFO 007 http://localhost:80/ - 200 <!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>

# 100 max concurrent
$ for i in `seq 1 3`;do time GHOSTMAN_MAXCONCUR=100 ./ghostman -f examples/localhost_5k run get index > /dev/null 2>&1;done
GHOSTMAN_MAXCONCUR=100 ./ghostman -f examples/localhost_5k run get index >  2>  1.42s user 1.42s system 76% cpu 3.697 total
GHOSTMAN_MAXCONCUR=100 ./ghostman -f examples/localhost_5k run get index >  2>  1.49s user 1.62s system 52% cpu 5.969 total
GHOSTMAN_MAXCONCUR=100 ./ghostman -f examples/localhost_5k run get index >  2>  1.35s user 1.99s system 81% cpu 4.090 total

# 5000 max concurrent
$ for i in `seq 1 3`;do time GHOSTMAN_MAXCONCUR=5000 ./ghostman -f examples/localhost_5k run get index > /dev/null 2>&1;done
GHOSTMAN_MAXCONCUR=5000 ./ghostman -f examples/localhost_5k run get index >  2  1.31s user 1.16s system 208% cpu 1.186 total
GHOSTMAN_MAXCONCUR=5000 ./ghostman -f examples/localhost_5k run get index >  2  1.37s user 1.72s system 107% cpu 2.877 total
GHOSTMAN_MAXCONCUR=5000 ./ghostman -f examples/localhost_5k run get index >  2  1.47s user 1.48s system 104% cpu 2.809 total
```
