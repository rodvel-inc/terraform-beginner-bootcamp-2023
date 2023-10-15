terraform {

  cloud {
    organization = "MyTerraform-Beginner-Project"
    workspaces {
      name = "Terrahouse-melomaniac"
    }
  }

  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.5.1"
    }
  aws = {
  source = "hashicorp/aws"
  version = "5.19.0"
    }
  }
}

provider "aws" {
  # Configuration options
}

provider "random" {
  # Configuration options
}