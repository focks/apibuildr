# ApiBuildr

apibuildr is a commandline tool for creating rest apis in golang language. 

`apibuildr` makes it easy to generate boilerplate code while adding rest apis.


It is not a framework, it uses gorilla mux server internally. 

### Installing apibuildr

```
go install github.com/focks/apibuildr/apibuildr@latest
```

### Starting A Project

Apibuildr creates a project directory suitable for a monorepo. The created repo/project would be suitable for housing multiple microservices in it.

```
apibuildr startProject <project-name> --package example.com/api
```

### Adding Your First Microservice 

It will create a rest api server with no endpoints. 

```
apibuildr addApp <microservice-name>
```


### Adding the first Api

Get into the desired microservice where you want the api to be added.

```shell
apibuildr addApi <api-name> -p /v1/hello -m POST 
```

-p or `--path` represents the api path  <br/>
-m or `--method` is the api method (by default it is get)


### Contributing 
please feel free to fork this project and raise PRs. Please include feature requests into `todos.txt` file.

**NOTE** this repository is under development. please feel free to fork and raise pull requests.
