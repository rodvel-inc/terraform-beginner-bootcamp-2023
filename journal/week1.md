# Terraform Beginner Bootcamp 2023 - Week 1

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
