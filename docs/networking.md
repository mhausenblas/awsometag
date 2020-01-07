## ELB

To tag the ALB `arn:aws:elasticloadbalancing:us-west-2:123456789102:loadbalancer/app/my-test-alb/1234567890` 
with `some=thing`, use the following:

```sh
# tag:
$ awsometag arn:aws:elasticloadbalancing:us-west-2:123456789102:loadbalancer/app/my-test-alb/1234567890 some=thing
2020/01/07 09:08:09 Tagging load balancer 'my-test-alb' of type 'app' in region 'us-west-2' with some:thing

# verify the tagging:
$ aws elbv2 describe-tags \
      --resource-arns arn:aws:elasticloadbalancing:us-west-2:123456789102:loadbalancer/app/my-test-alb/1234567890
{
    "TagDescriptions": [
        {
            "ResourceArn": "arn:aws:elasticloadbalancing:us-west-2:123456789102:loadbalancer/app/my-test-alb/1234567890",
            "Tags": [
                {
                    "Key": "some",
                    "Value": "thing"
                }
            ]
        }
    ]
}
```

!!! tip
    The original elastic load balancer uses the [API v1](https://docs.aws.amazon.com/cli/latest/reference/elb/), while ALBs and NLBs are 
    using [v2 of the API](https://docs.aws.amazon.com/cli/latest/reference/elbv2/).