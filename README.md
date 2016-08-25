# dpli
DeferPanic.Net Unikernel IaaS Client

Website: https://deferpanic.com

[![wercker status](https://app.wercker.com/status/44ace461901cac92c53e919de5d7e5e2/s/master "wercker status")](https://app.wercker.com/project/bykey/44ace461901cac92c53e919de5d7e5e2)

## Authentication:

You'll need an account so go sign up at https://deferpanic.net/signup .
Then you can find your API token @ https://deferpanic.net/home/settings.

You can either pass your token in with the -token flag or you maybe
stick it in ~/.dprc.

Example if using the token flag:
```
./dpcli --token="TOKEN"
```

- [Installing](#user-content-build)
- [Projects](#user-content-projects)
- [Instances](#user-content-instances)
- [Volumes](#user-content-volumes)
- [IPs](#user-content-ips)
- [Backups](#user-content-backups)
- [Resources](#user-content-resources)
- [Addons](#user-content-addons)
- [Languages](#user-content-languages)
 
## Dependencies:
```
go get gopkg.in/alecthomas/kingpin.v2
go get  github.com/olekukonko/tablewriter
```

## Build
```
go build
```

## Projects


### Create from Repository
```
./dpcli projects new myproject php https://github.com/deferpanic/php_example
```

### Create from Image Binary
```
```

### Remove
```
./dpcli --token="TOKEN" projects delete myproject
```

### List

```
./dpcli --token="TOKEN" projects list
```

### View Log

```
./dpcli --token="TOKEN" projects log myproject
```

### Download image

```
./dpcli projects download myproject
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

### List all Instances

```
./dpcli --token="TOKEN" instances list
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

### List By Domain

```
./dpcli --token="TOKEN" volumes list --domain=stephen-anna.deferpanic.net
```

### List by Image Name

```
./dpcli volumes list --name=text
```

### Download Volume
```
./dpcli volumes download id
```

Note: To download a volume you currently need to pause the instance but we are fixing this soon.

## IPs

### Request

```
./dpcli --token="TOKEN" ips request
```

### Release

```
./dpcli --token="TOKEN" ips release 1.2.3.4
```

### Attach

```
./dpcli --token="TOKEN" ips attach 1.2.3.4 stephen-anna.deferpanic.net
```

### Detach

```
./dpcli --token="TOKEN" ips detach 1.2.3.4
```

### List

```
./dpcli --token="TOKEN" ips list
```

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

To see all your provisioned resources:
```
./dpcli --token="TOKEN" resources list
```

To see the resources attached to a given project:
```
./dpcli --token="TOKEN" resources list myproject
```

### Available

To see the available resources:
```
./dpcli --token="TOKEN" resources available
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

[Static Website](https://github.com/deferpanic/html_example)

[Php](https://github.com/deferpanic/php_example)

[Php with Redis](https://github.com/vsukhin/phprumpredis)

[Php with Mysql](https://github.com/vsukhin/phprumpmysql)

[Ruby and Sinatra](https://github.com/deferpanic/ruby_example)

[Python](https://github.com/deferpanic/python_example)

[Node Js](https://github.com/vsukhin/nodejsrump)

[Go](https://github.com/deferpanic/go_example)

[Go with DeferPanic client](https://github.com/vsukhin/dpexample)

## Developing
Go 1.6 is Required.
