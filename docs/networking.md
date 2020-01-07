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

!!! note
    In the same way that you can tag an ALB, you can tag an NLB. You can tell the
    difference in the type by looking at the ARN: `loadbalancer/app` vs `loadbalancer/net`.

If you want to tag a classic LB, say `my-test-clb`, then you'd need to assemble the
ARN yourself (based on region, account ID) since CLBs don't deal with ARNs explicitly.
Then, use:

```sh
# tag
$ awsometag arn:aws:elasticloadbalancing:us-west-2:123456789102:loadbalancer/my-test-clb some=thing
2020/01/07 11:31:46 Tagging classic load balancer 'my-test-clb' in region 'us-west-2' with some:thing

# verify the tagging:
$ aws elb describe-tags \
      --load-balancer-names my-test-clb
{
    "TagDescriptions": [
        {
            "LoadBalancerName": "my-test-clb",
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
    The classic ELB (or: Classic Load Balancer, CLB for short) uses the [API v1](https://docs.aws.amazon.com/cli/latest/reference/elb/),
    while "next generation" ELBs, that is, ALBs and NLBs, are using [v2 of the API](https://docs.aws.amazon.com/cli/latest/reference/elbv2/).