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
Install it by downloading one of the [binaries](https://github.com/mhausenblas/awsometag/releases) or,
if you have Go 1.12+ installed, you can clone this repo and build it from source.

## Use it

`awsometag` takes exactly three arguments: 

1. the ARN of the resource to tag,
1. the region, and 
1. a list of comma-separated tags in the format `key=value`

That is, the general usage pattern is as follows:

```sh
$ awsometag RESOURCE_ARN REGION "TAG_KEY1=TAG_VAL1,TAG_KEY2=TAG_VAL2,..."
```

For example, to tag the bucket `arn:aws:s3:::abucket` with `thats=cool` you 
would use the following command:

```sh
$ awsometag arn:aws:s3:::abucket us-west-2 thats=cool
2020/01/04 13:54:32 Tagging S3 bucket 'abucket' with thats:cool

$ aws s3api get-bucket-tagging --bucket abucket
{
    "TagSet": [
        {
            "Key": "thats",
            "Value": "cool"
        }
    ]
}
```

Or maybe you want to tag the IAM user `arn:aws:iam::123456789012:user/abc` with
`nice=person` and `they=oweme`? Then you'd want to use the following:

```sh
$ awsometag arn:aws:iam::123456789012:user/abc eu-west-1 "nice=person, they=oweme"

$ aws iam list-user-tags --user-name abc
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
