# Simple Magic Webserver

This is a simple http server used as a example in the presentation. It is using the `generic-makefile` also discussed in this talk. 

Once the program is running, open your http://localhost:8081/ on a browser.

Requirements:
* [Go](https://golang.org/doc/install#install) (project created with go1.9.5 linux/amd64)
* [GNU make](https://www.gnu.org/gnu/gnu.html)
* [Docker](https://www.docker.com/get-started)


`vendor/` folder was created using [glide](https://glide.sh/).

## Make targets

### 1.Run
```
make run
```

### 2.Compile
```
make magicserver
```

### 3.Clean
```
make clean
```


### 3. Build Docker image
```
make clean
```

### 3. Run Docker image
```
make docker-run
```



