package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "my-bucket", nil)
		if err != nil {
			return err
		}
		mostRecent := true
		amiResult, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			Filters: []ec2.GetAmiFilter{
				{
					Name:   "name",
					Values: []string{"amzn-ami-hvm-*"},
				},
			},
			Owners:     []string{"137112412989"},
			MostRecent: &mostRecent,
		})
		if err != nil {
			return err
		}

		instance, err := ec2.NewInstance(ctx, "deploy-instance", &ec2.InstanceArgs{
			InstanceType: pulumi.String("t2.micro"),
			Ami:          pulumi.String(amiResult.Id),
		})

		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		ctx.Export("instance ip", instance.PublicIp)
		return nil
	})
}
