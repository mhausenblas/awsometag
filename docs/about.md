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
along with the ARG. In a sense, if one was to compare `awsometag` to ARGs, it would be offering functionality equivalent to a subset of the [Tag Editor](https://docs.aws.amazon.com/ARG/latest/userguide/tag-editor.html).