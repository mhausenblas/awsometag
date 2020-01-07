If you want to tag AWS resources via the CLI in a uniform manner, then `awsometag` might be just the tool you're looking for.

## Install it

Install `awsometag` by downloading one of the [binaries](https://github.com/mhausenblas/awsometag/releases) or,
if you have Go 1.12+ installed, you can clone this repo and build it from source using `make bin`.

??? example "Install on macOS"
    For example, to install `awsometag` from binary on macOS you would do:

    ```sh
    curl -L https://github.com/mhausenblas/awsometag/releases/latest/download/awsometag_darwin_amd64.tar.gz \
        -o awsometag.tar.gz && \
        tar xvzf awsometag.tar.gz awsometag && \
        mv awsometag /usr/local/bin && \
        rm awsometag*
    ```

Supported platforms:

- Linux (both Intel and ARM)
- macOS
- Windows

## Use it

The `awsometag` CLI tool takes two arguments: 

1. the ARN of the resource to tag,
1. a list of comma-separated tags, each in the format `key=value`

Hence, the general usage pattern for `awsometag` is:

```sh
$ awsometag RESOURCE_ARN "TAG_KEY1=TAG_VAL1,TAG_KEY2=TAG_VAL2,..."
```

## Supported resources

Currently, `awsometag` supports tagging resources for the following services:

**Fundamental** services:

- AWS Identity and Access Management:
    - users
    - roles
- Amazon Simple Storage Service:
    - buckets
    - objects
- AWS Lambda: functions
- Amazon DynamoDB: tables
- Amazon Elastic Compute Cloud: all resources
  
**Networking** services:

- Elastic Load Balancing:
    - Classic LBs
    - ALBs
    - NLBs

**Container** services:

- Amazon Elastic Container Registry: repositories
- Amazon Elastic Container Service:
    - capacity providers
    - clusters
    - tasks and task definitions
    - services
    - container instances
- Amazon Elastic Kubernetes Service:
    - clusters
    - managed node groups
