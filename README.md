# awsometag

After reading and reflecting on [Allocate AWS Costs with Resource Tags](https://medium.com/@davidevanpaulis/allocate-aws-costs-with-resource-tags-277de240487f) I asked myself: is there a CLI tool out there, ideally with a great UX, that allows me to tag *any* AWS resource in a uniform manner? After some searching around I concluded that the answer is likely "No" and so I set out to fill this gap with `awsometag`.

## Install it
Install it by downloading one of the [binaries](https://github.com/mhausenblas/awsometag/releases) or, if you have Go 1.12+ installed, you can clone this repo and build it from source.

## Use it

The general usage pattern is:

```sh
$ awsometag RESOURCEARN TAG_KEY1=TAG_VAL1,TAG_KEY2=TAG_VAL2,...
```

For example, to tag the bucket `arn:aws:s3:::abucket` with `thats=cool` you would use the following command:

```sh
$ awsometag arn:aws:s3:::abucket thats=cool
```

Or maybe you want to tag the IAM user `arn:aws:iam::123456789012:user/abc` with `nice=person` and `they=oweme`? Then you'd want to use the following:

```sh
$ awsometag arn:aws:iam::123456789012:user/abc "nice=person, they=oweme"
```