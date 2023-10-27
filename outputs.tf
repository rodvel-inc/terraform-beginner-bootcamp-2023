output "s3_bucket_name" {
  description = "The name of the bucket for the website"
  value = module.terrahouse_aws.bucket_name
}

output "s3_website_endpoint" {
  description = "s3 static Website endpoint"
  value = module.terrahouse_aws.website_endpoint
}

output "cloudfront_url" {
  value = module.terrahouse_aws.cloudfront_url
}