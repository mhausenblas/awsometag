## RDS

To tag the database `arn:aws:rds:eu-west-1:123456789012:db:mydb` 
with `some=thing`, use the following:

```sh
# tag:
$ awsometag arn:aws:rds:eu-west-1:123456789012:db:mydb some=thing
2020/01/09 16:36:24 Tagging RDS database 'mydb' in region 'eu-west-1' with some:thing

# verify the tagging:
$ aws rds list-tags-for-resource \
      --resource-name arn:aws:rds:eu-west-1:123456789012:db:mydb
{
    "TagList": [
        {
            "Key": "some",
            "Value": "thing"
        }
    ]
}
```