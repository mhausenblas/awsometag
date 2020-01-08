## Secrets Manager

To tag the secret `arn:aws:secretsmanager:us-west-2:123456789012:secret:mysecret-12345` 
with `some=thing`, use the following:

```sh
# tag:
$ awsometag arn:aws:secretsmanager:us-west-2:123456789012:secret:mysecret-12345 some=thing
2020/01/08 09:27:28 Tagging secret 'mysecret' in region 'us-west-2' with some:thing

# verify the tagging:
$ aws secretsmanager describe-secret \
      --secret-id mysecret
{
    "ARN": "arn:aws:secretsmanager:us-west-2:123456789012:secret:mysecret-12345",
    "Name": "mysecret",
    "Tags": [
        {
            "Key": "some",
            "Value": "other"
        }
    ],
    ...
}
```