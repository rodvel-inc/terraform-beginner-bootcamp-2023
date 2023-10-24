terraform {

  cloud {
    organization = "MyTerraform-Beginner-Project"
    workspaces {
      name = "Terrahouse-melomaniac"
    }
  }

  required_providers {
#    random = {
#      source = "hashicorp/random"
#      version = "3.5.1"
#    }
  aws = {
  source = "hashicorp/aws"
  version = "5.19.0"
    }
  }
}

#provider "aws" {
  # Configuration options
#}

#provider "random" {
  # Configuration options
#}
#resource "random_string" "bucket_name" {
#  lower = true
#  upper = false
#  length = 16
#  special = false
#}

resource "aws_s3_bucket" "website_bucket" {
  bucket = var.bucket_name

  tags = {
    UserUUID        = var.user_uuid
    }
}

resource "aws_s3_bucket_website_configuration" "website_configuration" {
  bucket = aws_s3_bucket.website_bucket.bucket

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }
}

resource "aws_s3_object" "object_index" {
  bucket = aws_s3_bucket.website_bucket.bucket
  key    = "index.html"
  source = "${path.root}/public/index.html"
  etag = filemd5("${path.root}/public/index.html")
}

resource "aws_s3_object" "object_error" {
  bucket = aws_s3_bucket.website_bucket.bucket
  key    = "error.html"
  source = "${path.root}/public/error.html"
  etag = filemd5("${path.root}/public/error.html")
}

