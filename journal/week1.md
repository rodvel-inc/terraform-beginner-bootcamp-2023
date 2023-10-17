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
