# dpcli
DeferPanic.Net Unikernel IaaS Client

You'll need an account so go sign up at https://deferpanic.net/signup .
Then you can find your API token @ https://deferpanic.net/home/settings
. Replace all occurences of TOKEN with your token in the examples
provided below.

You can either pass your token in with the -token flag or you maybe
stick it in ~/.dprc.

## Build
```
go build
```

## Projects

### Create from Repository
```
./dpcli -new -token TOKEN -name bob -language php \
  -source https://github.com/vsukhin/phprump
```

### Create from Image Binary
```
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

```
./dpcli -token TOKEN -name bob -scaleup
```

### Log of Individual Instance
```
./dpcli -token TOKEN -runlog -domain robert-oklahoma
```

### List of Instances attached to Image/Project

```
./dpcli -token TOKEN -show -name bob
```

### Pause

```
./dpcli -token TOKEN -pause -domain robert-oklahoma
```

### Resume

```
./dpcli -token TOKEN -resume -domain robert-oklahoma
```

### ScaleUp

```
./dpcli -token TOKEN -name bob -scaleup
```

### ScaleDown

```
./dpcli -token TOKEN -name bob -scaledown
```

## Volumes

## Backups

### Save

### Restore

### List

```
./dpcli -token TOKEN -listbackups
```

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
./dpcli -status
```

## Version
```
./dpcli -version
```


## Examples


[Php with Redis](https://github.com/vsukhin/phprumpredis)

[Php with Mysql](https://github.com/vsukhin/phprumpmysql)

[Ruby and Sinatra](https://github.com/vsukhin/rubysinatrarump)

[Node Js](https://github.com/vsukhin/nodejsrump)

[Go](https://github.com/vsukhin/gorump)

[Go with DeferPanic client](https://github.com/vsukhin/dpexample)

[Php](https://github.com/vsukhin/phprump)
