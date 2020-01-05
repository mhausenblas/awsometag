# awsometag

After reading and reflecting on [Allocate AWS Costs with Resource Tags](https://medium.com/@davidevanpaulis/allocate-aws-costs-with-resource-tags-277de240487f) 
I asked myself: is there a CLI tool out there, ideally with a great UX, 
that allows me to tag *any* AWS resource in a uniform manner? 
After some searching around I concluded that the answer is likely "No" 
and so I set out to fill this gap with `awsometag`.

But, I hear you say, there are [AWS Resource Groups](https://docs.aws.amazon.com/ARG/latest/userguide/welcome.html) (ARG), already!
Isn't this re-inventing the wheel?

Nope. And the reason is two-fold: on the one hand, you can consider `awsometag` as
a sort of enabler for ARGs; you don't have to think about the grouping aspect up-front,
only tag your resources. Second, ARGs are very powerful, supporting more than simple
tags, including CloudFormation stack-based queries and a set of permissions 
along with the ARG. In a sense, if one was to compare `awsometag` to ARGs, it would
be roughly offering the functionality of the [Tag Editor](https://docs.aws.amazon.com/ARG/latest/userguide/tag-editor.html).

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

1. Fundamental services
   - <a href="#iam" title="AWS Identity and Access Management">IAM</a>: users, roles
   - <a href="#s3" title="Amazon Simple Storage Service">S3</a>:  buckets, objects
   - <a href="#lambda" title="AWS Lambda">Lambda</a>: functions
1. Container services
   - <a href="#ecr" title="Amazon Elastic Container Registry">ECR</a>: repositories
   - <a href="#ecs" title="Amazon Elastic Container Service">ECS</a>: capacity providers, clusters, tasks, task definitions, services, and container instances
   - <a href="#eks" title="Amazon Elastic Kubernetes Service">EKS</a>: clusters, managed node groups

:metal: and here are all the nitty gritty details …

### IAM

If you want to tag the IAM user `arn:aws:iam::123456789012:user/abc` with
`nice=person` and `they=oweme` then you'd want to use the following:

```sh
# tag:
$ awsometag arn:aws:iam::123456789012:user/abc "nice=person, they=oweme"

# verify the tagging:
$ aws iam list-user-tags \
      --user-name abc
{
    "Tags": [
        {
            "Key": "nice",
            "Value": "person"
        },
        {
            "Key": "they",
            "Value": "oweme"
        }
    ],
    "IsTruncated": false
}
```

Note: the same works for IAM roles.

### S3

To tag the bucket `arn:aws:s3:::abucket` with `thats=cool` you would use:

```sh
# tag:
$ awsometag arn:aws:s3:us-west-2::abucket thats=cool
2020/01/04 13:54:32 Tagging S3 bucket 'abucket' in region 'us-west-2' with thats:cool

# verify the tagging:
$ aws s3api get-bucket-tagging \
      --bucket abucket
{
    "TagSet": [
        {
            "Key": "thats",
            "Value": "cool"
        }
    ]
}
```

:fire:

Note: if the S3 ARN does not contain the region, then you MUST provide the desired
target region via the `S3_ENDPOINT_REGION` environment variable. For example, in 
above case it would be: `S3_ENDPOINT_REGION=us-west-2 awsometag arn:aws:s3:::abucket us-west-2 thats=cool`.

Tagging works the same for objects in a bucket: let's tag the object with the key 
`input.csv` residing in the bucket `abucket` with `this=aswell`:

```sh
# tag:
$ S3_ENDPOINT_REGION=us-west-2 awsometag arn:aws:s3:::abucket/input.csv this=aswell
2020/01/05 07:03:50 Tagging S3 object 'input.csv' in bucket 'abucket' with this:aswell

# verify the tagging:
$ aws s3api get-object-tagging \
     --bucket abucket \
     --key input.csv
{
    "TagSet": [
        {
            "Key": "this",
            "Value": "aswell"
        }
    ]
}
```

### Lambda

To tag the Lambda function `arn:aws:lambda:us-west-2:123456789102:function:coolapp-TheFunc-1234567` 
with `server=less`, use:

```sh
# tag:
$ awsometag arn:aws:lambda:us-west-2:123456789102:function:coolapp-TheFunc-1234567 server=less
2020/01/05 14:16:47 Tagging Lambda function 'coolapp-TheFunc-1234567' in region 'us-west-2' with server:less

# verify the tagging:
$ aws lambda list-tags \
      --resource arn:aws:lambda:us-west-2:123456789102:function:coolapp-TheFunc-1234567
{
    "Tags": {
        "server": "less"   
    }
}
```


### ECR

To tag the ECR repo `arn:aws:ecr:us-east-1:123456789102:repository/somerepo` 
with `my=containers`, use the following:

```sh
# tag:
$ awsometag arn:aws:ecr:us-east-1:123456789102:repository/somerepo my=containers
2020/01/05 09:43:03 Tagging ECR repository 'somerepo' in region 'us-east-1' with my:containers

# verify the tagging:
$ aws ecr list-tags-for-resource \
      --resource-arn arn:aws:ecr:us-east-1:123456789102:repository/somerepo
{
    "tags": [
        {
            "Key": "my",
            "Value": "containers"
        }
    ]
}
```

### ECS

To tag the ECS task definition `arn:aws:ecs:us-west-2:123456789102:task-definition/nginx:3` 
with `my=containers`, use the following:

```sh
# tag:
$ awsometag arn:aws:ecs:us-west-2:123456789102:task-definition/nginx:3 my=containers
2020/01/05 13:13:44 Tagging ECS task definition 'nginx:3' in region 'us-west-2' with my:containers

# verify the tagging:
$ aws ecs list-tags-for-resource \
      --resource-arn arn:aws:ecs:us-west-2:123456789102:task-definition/nginx:3
{
    "tags": [
        {
            "key": "some",
            "value": "thing"
        }
    ]
}
```

### EKS

To tag the EKS cluster `arn:aws:eks:us-west-2:123456789102:cluster/somecluster` 
with `my=containers`, use the following:

```sh
# tag:
$ awsometag arn:aws:eks:us-west-2:123456789102:cluster/somecluster my=containers
2020/01/05 08:26:03 Tagging EKS cluster 'somecluster' in region 'us-west-1' with my:containers

# verify the tagging:
$ aws eks list-tags-for-resource \
      --resource-arn arn:aws:eks:us-west-1:123456789102:cluster/somecluster
{
    "tags": {
        "my": "containers"
    }
}
```

Note: the same works for tagging [managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).