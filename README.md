# Terraform Beginner Bootcamp 2023

## Semantic Versioning

Tagging will use semantic versioning.

> [semver.org](https://semver.org/)

Given a version number MAJOR.MINOR.PATCH (e.g. 1.0.1), increment the:

-    ***MAJOR*** version when you make incompatible API changes
-    ***MINOR*** version when you add functionality in a backward compatible manner
-    ***PATCH*** version when you make backward compatible bug fixes

Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.

## Setting env var for the PROJECT_ROOT

The command 'env' will list all the environment variables for the workspace I am working on.

## AWS CLI installation

This tools is installed for the project via the bash script [`./bin/install_aws_cli`](./bin/install_aws_cli)

[Getting Started Install (AWS CLI)](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
[AWS CLI env vars](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)

We can check if our AWS credentials are configured correctly by running the following CLI command:

```sh
aws sts get-caller-identity
```
If successful, you should see a JSON payload return that looks like this:
```json
{
    "UserId": "AIDAUJ77UCDUSKEXAMPLE",
    "Account": "123456789012",
    "Arn": "arn:aws:iam::123456789012:user/administrador"
}
```
Commit changes and that's it.

## Terraform Basics

### Terraform Registry

Terraform sources their providers and modules from the Terraform registry which are located at [registry.terraform.io](https://registry.terraform.io/)

- **Providers**: it is an interface to APIs that will allow to create resources on Terraform.
- **Modules**: are a way to make large amounts of terraform code modular, portable and shareable.

### Terraform Console

We can see a list of all the Terraform commands by simply typing `terraform`

#### Terraform Init
`terraform init`

At the start of a new terraform project, the execution of this command will cause the binaries for the selected provider to be downloaded.

#### Terraform Plan
`terraform plan`

This will generate a changeset about the state of our infrastructure and what it will look like once deployed.

#### Terraform Apply
`terraform apply`

This will run a plan and pass the changeset to be executed by terraform.

### Terraform Lock files

`.terraform.lock.hcl` contains the locked versioning for the providers or modules.
This file **should be committed** to your Version Control System.

### Terraform State files
`.terraform.tfstate` contains the information about the current state of your infrastructre. 
This file **should not be committed** to your Version Control System.

### Terraform Directory
`.terraform.tfstate` contains the downloaded binaries that correspond to the provider used in the project.