# awsometag

If you want to tag AWS resources via the CLI in a uniform manner, then `awsometag` might be just the tool you're looking for.

## Install it

Install `awsometag` by downloading one of the [binaries](https://github.com/mhausenblas/awsometag/releases) or,
if you have Go 1.12+ installed, you can clone this repo and build it from source.

## Use it

The `awsometag` CLI tool takes two arguments: 

1. the ARN of the resource to tag,
1. a list of comma-separated tags, each in the format `key=value`

Hence, the general usage pattern for `awsometag` is:

```sh
$ awsometag RESOURCE_ARN "TAG_KEY1=TAG_VAL1,TAG_KEY2=TAG_VAL2,..."
```

Currently, `awsometag` supports tagging resources in:

- AWS Identity and Access Management: users, roles
- Amazon Simple Storage Service:  buckets, objects
- AWS Lambda: functions
- Amazon DynamoDB: tables
- Amazon Elastic Compute Cloud: all resources
- Amazon Elastic Container Registry: repositories
- Amazon Elastic Container Service: capacity providers, clusters, tasks, task definitions, services, and container instances
- Amazon Elastic Kubernetes Service: clusters, managed node groups
