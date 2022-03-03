# ApiBuildr

apibuildr is a commandline tool for creating rest apis in golang language. 

`apibuildr` makes it easy to generate boilerplate code while adding rest apis.


It is not a framework, it uses gorilla mux server internally. 

### Installing apibuildr

```
go get github.com/focks/apibuildr
```

### Initializing a project

```
apibuildr init .
```
It will create a rest api server with no endpoints. 

**NOTE** this repository is under development. please feel free to fork and raise pull requests.


### Adding the first Api

```shell
apibuildr addApi apiName -p /v1/hello -m POST 
```

-p or `--path` represents the api path  <br/>
-m or `--method` is the api method (by default it is get)


### Contributing 
please feel free to fork this project and raise PRs. Please include feature requests into `todos.txt` file.