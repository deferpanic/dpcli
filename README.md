# dpli
DeferPanic.Net Unikernel IaaS Client

Website: https://deferpanic.com

[![wercker status](https://app.wercker.com/status/44ace461901cac92c53e919de5d7e5e2/s/master "wercker status")](https://app.wercker.com/project/bykey/44ace461901cac92c53e919de5d7e5e2)

## Quick Start:

This will create an account for you and auto-log you in through the API:

```
./dpcli users create joe@bob.com password
```

## Authentication:

You'll need an account so go sign up at https://deferpanic.net/signup if you haven't created it through the API.
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
cd dpcli && go build && go install
```

## Projects


### Create from Repository
```
dpcli projects new myproject php https://github.com/deferpanic/php_example
```

### Create from Image Binary
```
```

### Fork an Existing Project

### Clone an Addon

### Remove
```
dpcli projects delete myproject
```

### List

```
dpcli projects list
```

### View Log

```
dpcli projects log myproject
```

### Download image

```
./dpcli projects download myproject
```

### Download manifest

```
./dpcli projects manifest myproject
```

## Instances

### Create

```
./dpcli instances new myproject
```

### Log of Individual Instance
```
./dpcli instances log gregory-jennifer.deferpanic.net
```

### List all Instances

```
./dpcli instances list
```

### List of Instances attached to Image/Project

```
./dpcli instances list myproject
```

### Pause

```
./dpcli instances pause gregory-jennifer.deferpanic.net
```

### Resume

```
./dpcli instances resume gregory-jennifer.deferpanic.net
```

### ScaleUp

```
./dpcli instances scaleup myproject
```

### ScaleDown

```
./dpcli instances scaledown myproject gregory-jennifer.deferpanic.net
```

## Volumes

### List By Domain

```
./dpcli volumes list --domain=stephen-anna.deferpanic.net
```

### List by Image Name

```
./dpcli volumes list --name=text
```

### Download Volume

Note: To download a volume you currently need to pause the instance.
(Live volume streaming coming soon.)

Pause the instance
```
./dpcli instances pause my_cname.deferpanic.net
```

```
./dpcli volumes download id
```

## IPs

### Request

```
./dpcli ips request
```

### Release

```
./dpcli ips release 1.2.3.4
```

### Attach

```
./dpcli ips attach 1.2.3.4 stephen-anna.deferpanic.net
```

### Detach

```
./dpcli ips detach 1.2.3.4
```

### List

```
./dpcli ips list
```

## Backups

### Save

Pause the instance
(Live Backups coming soon.)
```
./dpcli instances pause gregory-jennifer.deferpanic.net
```

Back it up
```
./dpcli backups save myproject gregory-jennifer.deferpanic.net
```

### Restore

Pause the instance
```
./dpcli instances pause gregory-jennifer.deferpanic.net
```

Back it up
```
./dpcli backups restore myproject gregory-jennifer.deferpanic.net
```

### List

```
./dpcli backups list
```

## Languages

### List

```
./dpcli languages
```

## Resources

### Create

### List

To see all your provisioned resources:
```
./dpcli resources list
```

To see the resources attached to a given project:
```
./dpcli resources list myproject
```

### Available

To see the available resources:
```
./dpcli --token="TOKEN" resources available
```

## Addons

To see the available addons:
```
dpcli --token="TOKEN" addons
```

## Status
```
dpcli status
```

## Version
```
dpcli version
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
