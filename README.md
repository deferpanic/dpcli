# dpli
DeferPanic.Net Unikernel IaaS Client

You'll need an account so go sign up at https://deferpanic.net/signup .
Then you can find your API token @ https://deferpanic.net/home/settings
. Replace all occurences of TOKEN with your token in the examples
provided below.

You can either pass your token in with the -token flag or you maybe
stick it in ~/.dprc.
pi

- [Installing](#user-content-build)
- [Projects](#user-content-projects)
- [Instances](#user-content-instances)
- [Volumes](#user-content-volumes)
- [Backups](#user-content-backups)
- [Resources](#user-content-resources)
- [Addons](#user-content-addons)
- [Languages](#user-content-languages)
 
## Dependencies:
```
go get gopkg.in/alecthomas/kingpin.v2
```

## Build
```
go build
```

## Projects

We currently don't support uploading custom images/volumes but we plan
on doing so in the very near future.

### Create from Repository
```
./dpcli --token="TOKEN" projects new myproject php https://github.com/vsukhin/phprump
```

### Create from Image Binary
```
```

### List

```
./dpcli --token="TOKEN" projects list
```

### View Log

```
./dpcli --token="TOKEN" projects log myproject
```

## Instances

### Create

```
./dpcli --token="TOKEN" instances new myproject
```

### Log of Individual Instance
```
./dpcli --token="TOKEN" instances log gregory-jennifer.deferpanic.net
```

### List of Instances attached to Image/Project

```
./dpcli --token="TOKEN" instances list myproject
```

### Pause

```
./dpcli --token="TOKEN" instances pause gregory-jennifer.deferpanic.net
```

### Resume

```
./dpcli --token="TOKEN" instances resume gregory-jennifer.deferpanic.net
```

### ScaleUp

```
./dpcli --token="TOKEN" instances scaleup myproject
```

### ScaleDown

```
./dpcli --token="TOKEN" instances scaledown myproject gregory-jennifer.deferpanic.net
```

## Volumes

## Backups

### Save

Pause the instance
(Live Backups coming soon.)
```
./dpcli --token="TOKEN" instances pause gregory-jennifer.deferpanic.net
```

Back it up
```
./dpcli --token="TOKEN" backups save myproject gregory-jennifer.deferpanic.net
```

### Restore

Pause the instance
```
./dpcli --token="TOKEN" instances pause gregory-jennifer.deferpanic.net
```

Back it up
```
./dpcli --token="TOKEN" backups restore myproject gregory-jennifer.deferpanic.net
```

### List

```
./dpcli --token="TOKEN" backups list
```

## Languages

### List

```
./dpcli --token="TOKEN" languages
```

## Resources

### Create

### List

To see the available resources:
```
./dpcli --token="TOKEN" resources list myproject
```

## Addons

To see the available addons:
```
./dpcli --token="TOKEN" addons
```

## Status
```
./dpcli status
```

## Version
```
```


## Examples


[Php with Redis](https://github.com/vsukhin/phprumpredis)

[Php with Mysql](https://github.com/vsukhin/phprumpmysql)

[Ruby and Sinatra](https://github.com/vsukhin/rubysinatrarump)

[Node Js](https://github.com/vsukhin/nodejsrump)

[Go](https://github.com/vsukhin/gorump)

[Go with DeferPanic client](https://github.com/vsukhin/dpexample)

[Php](https://github.com/vsukhin/phprump)
