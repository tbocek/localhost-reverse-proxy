# About
localhost-reverse-proxy is a web server for static files and a reverse proxy. 
Since many browsers have CORS restrictions, you can either setup a reverse proxy or
disable CORS. I'm now at a point where I think its easier to have a small reverse
proxy that can test frontend backend and backend.

It is assumed that your frontend and backend runs on localhost, and your frontend
is HTML/JavaScript that is stored in a specific folder.

# Installation

```bash
mkdir -p go/src/github.com/tbocek
cd go/src/github.com/tbocek
git clone https://github.com/tbocek/localhost-reverse-proxy.git
#this will build the binary and copy it to go/bin
cd localhost-reverse-proxy
go install
#now you can run the binary localhost-reverse-proxy if your go/bin is in your $PATH 
localhost-reverse-proxy -d /var/www
```

# Output

If you run localhost-reverse-proxy, it will use default settings and print 
its options:

```
Static file HTTP server and reverse proxy for localhost. This tool exists due to CORS.
  -d string
    	directory to server HTML files from, e.g. /var/www (default ".")
  -l string
    	listen on port, e.g., 8080 (default "8080")
  -r string
    	redirect port to redirect to, e.g., 8545. This willredirect http://localhost:8080/8545 to http://localhost:8545 (default "8545")
Serving: /var/www on http://localhost:8080, redirecting http://localhost:8080/8545 to http://localhost:8545
```