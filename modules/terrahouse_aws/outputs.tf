output "bucket_name" {
    value = aws_s3_bucket.website_bucket.bucket
}

output "website_endpoint" {
    description = "Website Endpoint"
    value = aws_s3_bucket_website_configuration.website_configuration.website_endpoint
  
}

output "cloudfront_url" {
    description = "CDN domain name"
    value = aws_cloudfront_distribution.s3_distribution.domain_name
}