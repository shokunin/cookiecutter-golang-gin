# cookiecutter-golang-gin

This sets up a Golang Gin web application with a health check, basic test, docker and JSON logging

## Prerequisites 

 Install cookiecutter
```
sudo pip install cookiecutter
```

## Using 

```
cookiecutter https://github.com/shokunin/cookiecutter-golang-gin.git
```

Answer the questions

cd into the app directory and run

```
make
./APPLICATION_NAME
curl localhost:8080/health
```
