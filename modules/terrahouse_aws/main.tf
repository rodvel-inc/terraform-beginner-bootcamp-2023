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

