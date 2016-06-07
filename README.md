# dpcli
DeferPanic.Net Unikernel IaaS Client

You'll need an account so go sign up at https://deferpanic.net/signup .
Then you can find your API token @ https://deferpanic.net/home/settings
. Replace all occurences of TOKEN with your token in the examples
provided below.

## Build
```
go build
```

## Projects

### Create
```
./dpcli -new -token TOKEN -name bob -language php \
  -source https://github.com/vsukhin/phprump
```

### List

```
./dpcli -token TOKEN -display
```

### View Log

```
./dpcli -token TOKEN -name bob -makelog
```

## Instances

### Create

### Show

### List

### Pause

### Resume

## Volumes

## Backups

### Save

### Restore

### List

## Languages

### List

```
./dpcli -token TOKEN -languages
```

## Resources

### Create

### List

To see the available resources:
```
./dpcli -token TOKEN -builtins
```

## Addons

To see the available addons:
```
./dpcli -token TOKEN -builtins
```

## Status
```
./dpcli status
```

## Examples


[Php with Redis](https://github.com/vsukhin/phprumpredis)
[Php with Mysql](https://github.com/vsukhin/phprumpmysql)
[Ruby and Sinatra](https://github.com/vsukhin/rubysinatrarump)
[Node Js](https://github.com/vsukhin/nodejsrump)
[Go](https://github.com/vsukhin/gorump)
[Go with DeferPanic client](https://github.com/vsukhin/dpexample)
[Php](https://github.com/vsukhin/phprump)
