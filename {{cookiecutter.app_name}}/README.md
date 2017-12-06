# {{cookiecutter.app_name}}

{{cookiecutter.short_description}}

## Building

### Mac/Linux

0) set GOROOT environment variable
1) Install Go and Make
2) make

### Docker

0) set GOROOT environment variable
1) Install Docker, Go and Make
2) make docker


## Running

### Mac/Linux

```
./{{cookiecutter.app_name}}
```

### Docker

```
docker pull {{cookiecutter.docker_location}}/{{cookiecutter.app_name}}:latest
docker run -i -t -p 8080:8080 {{cookiecutter.docker_location}}/{{cookiecutter.app_name}}
```

## Testing

run either the docker container or the raw application binary

```
curl http://localhost:8080/health
```

---
Copyright © {{cookiecutter.release_date.split('-')[0]}}, {{ cookiecutter.full_name }}
