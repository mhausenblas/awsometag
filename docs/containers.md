## ECR

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

## ECS

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

## EKS

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

!!!note
    In the same way you can tag clusters, you can tag [managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html)
    part of a cluster.