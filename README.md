# cab-management-portal

### Installation steps

```
Go 1.12 Required
1. Clone the project
2. Update app/utilEntities/config.go with mysql config
3. run -> go mod vendor
4. run -> go build -mod=vendor 
5. run -> export APP_NAME=test_app;TIER=dev;PORT=8080 //Environment variable
6. run -> ./cab-management-portal
```

### DB Design
