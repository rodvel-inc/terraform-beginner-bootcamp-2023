# Terraform Beginner Bootcamp 2023 - Week 1# Terraform Beginner Bootcamp 2023 - Week 1

## Root Module Structure
Our module structure requires:
```
PROJECT_ROOT
├── main.tf          # mandatory primary entry point
├── variables.tf     # declaration of variables
├── outputs.tf       # declaration of outputs
├── providers.tf     # declaration of providers used within the project
├── terraform.tfvars # to assign values to the variables declared in the variables.tf file
├── README.md        # There should be a description of the module and what it should be used for.
```
  
- [Standard Module Structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure)

## Terraform Modules

### Fixing the "output missing" warning

Since the outputs.tf file was declared inside the terrahouse_aws module but not in the root module, the output was not written in the state file.

I linked the two outputs.tf files and then I used a terraform command to reflect this change of state without having to destroy and then apply the entire project again:

```sh
terraform apply -refresh-only -auto-approve
```

You can apply this plan to save these new output values to the Terraform state, without changing any
real infrastructure.

### Module Sources

Modules can have different sources; for example:
- local
- github
- terraform registry
- others

Supposse the source is Terraform Registry:

```terraform
module "consul" {
  source = "hashicorp/consul/aws"
  version = "0.1.0"
}
```
[Module Sources](https://developer.hashicorp.com/terraform/language/modules/sources)

## Uploading objects to S3 Buckets

Uploading objects to an S3 is pretty straightforward.

Suppose you need to upload the index.html to an S3 bucket:

```terraform
resource "aws_s3_object" "object" {
  bucket = "bucket_name"
  key    = "index.html"
  source = "${path.root}/folder/index.html"
  etag = filemd5("${path.root}/folder/index.html")
}
```
In the preceding example, the `index.html` file is located in a separate folder. 

The location of the folder is referenced via the `${path.root}`.

[Reference to Values in Terraform](https://developer.hashicorp.com/terraform/language/expressions/references)

See there is an `etag` attribute included in the code above. This tag helps Terraform to notice that there has been a change in the content of an object after it was originally uploaded to the bucket.

This is an MD5 hash calculated if included. If the hash changes from one `terraform apply` to the next one, this change is detected and, therefore, the file is uploaded again with its new content.

### Terraform Data Sources

[Terraform Data Sources](https://developer.hashicorp.com/terraform/language/data-sources)


### Changind the Lifecycle of resources

By default, Terraform will destroy an object / resource and then create a new one if a new configuration is detected when instructed to do so.

This behaviour can be changed with these lifecycle meta-arguments.

[The Lifecycle meta-argument](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle)