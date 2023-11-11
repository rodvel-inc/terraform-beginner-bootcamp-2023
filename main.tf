terraform {
    required_providers {
      terratowns = {
        source = "local.providers/local/terratowns"
        #source = "home/gitpod/.terraform.d/plugins/local.providers/local/terratowns/"
        version = "1.0.0"
      }
    }
}

# this endpoint is the local terratowns, a local emulator
#provider "terratowns" {
#    endpoint = "http://localhost:4567/api"
#    user_uuid=var.teacherseat_user_uuid 
#    token="9b49b3fb-b8e9-483c-b703-97ba88eef8e0"
#}

#this is the real Terratowns
provider "terratowns" {
    endpoint = var.terratowns_endpoint
    user_uuid = var.teacherseat_user_uuid
    token=var.terratowns_access_token
}

module "terrahouse_aws" {
    source = "./modules/terrahouse_aws"
    user_uuid = var.teacherseat_user_uuid
    bucket_name = var.bucket_name
    index_html_filepath = var.index_html_filepath
    error_html_filepath = var.error_html_filepath
    content_version = var.content_version
}

resource "terratowns_home" "home" {
  name = "Metal Rodvel" 
  description = "Best Death Metal bands ever!"
  town = "missingo"
  content_version = 1
  domain_name = module.terrahouse_aws.cloudfront_url
  #domain_name = "3faasd13c.cloudfront.net"
  
}