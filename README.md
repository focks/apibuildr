# ApiBuildr

apibuildr is a commandline tool for creating rest apis in golang language. 

It is not a framework, it uses gorilla mux server internally. 

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