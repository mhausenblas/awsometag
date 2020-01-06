## IAM

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

!!!note
    In the same way you can tag IAM users, you can tag IAM roles.

## S3

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

!!! warning
    If the S3 ARN does not contain the region, then you MUST provide the desired
    target region via the `S3_ENDPOINT_REGION` environment variable. For example, in above case it would be: `S3_ENDPOINT_REGION=us-west-2 awsometag arn:aws:s3:::abucket us-west-2 thats=cool`.

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

## Lambda

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

## DynamoDB

To tag the DynamoDB table `arn:aws:dynamodb:us-west-2:123456789102:table/TheTable` 
with `some=thing`, use:

```sh
# tag:
$ awsometag arn:aws:dynamodb:us-west-2:123456789102:table/TheTable some=thing
2020/01/06 05:35:48 Tagging DynamoDB table 'TheTable' in region 'us-west-2' with some:thing

# verify the tagging:
$ aws dynamodb list-tags-of-resource \
      --resource-arn arn:aws:dynamodb:us-west-2:123456789102:table/TheTable
{
    "Tags": [
        {
            "Key": "some",
            "Value": "thing"
        }
    ]
}
```

## EC2

To tag the EC2 instance `arn:aws:ec2:us-west-2:123456789102:instance/i-123456789` 
with `some=thing`, use:

```sh
# tag:
$ awsometag arn:aws:ec2:us-west-2:123456789102:instance/i-123456789 some=thing
2020/01/06 06:15:42 Tagging EC2 resource 'i-123456789' of type 'instance' in region 'us-west-2' with some:thing

# verify the tagging:
$ aws ec2 describe-tags \
      --filters "Name=resource-id,Values=i-123456789"
{
    "Tags": [
        {
            "Key": "some",
            "ResourceId": "i-123456789",
            "ResourceType": "instance",
            "Value": "thing"
        }
    ]
}
```

!!! note "EC2 resource IDs vs. ARNs"

    The EC2 service defines a range of 
    [resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonec2.html#amazonec2-resources-for-iam-policies),
    from instances to volumes to VPCs. All of them are supported and you'll need to compose the ARNs yourself.