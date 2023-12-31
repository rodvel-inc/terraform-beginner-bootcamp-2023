# Terraform Beginner Bootcamp 2023 - Week 0

## Table of Contents

- [Terraform Beginner Bootcamp 2023](#terraform-beginner-bootcamp-2023)
  * [Table of Contents](#table-of-contents)
  * [Semantic Versioning](#semantic-versioning)
  * [Setting env var for the PROJECT_ROOT](#setting-env-var-for-the-project-root)
  * [AWS CLI installation](#aws-cli-installation)
  * [Terraform Basics](#terraform-basics)
    + [Terraform Registry](#terraform-registry)
    + [Terraform Console](#terraform-console)
      - [Terraform Init](#terraform-init)
      - [Terraform Plan](#terraform-plan)
      - [Terraform Apply](#terraform-apply)
    + [Terraform Lock files](#terraform-lock-files)
    + [Terraform State files](#terraform-state-files)
    + [Terraform Directory](#terraform-directory)
    + [Terraform cloud](#terraform-cloud)
      - [Organizations, Projects and Workspaces](#organizations--projects-and-workspaces)
      - [Connect your code to Terraform cloud](#connect-your-code-to-terraform-cloud)
      - [Issues with Terraform cloud state files](#issues-with-terraform-cloud-state-files)
        * [Workaround](#workaround)

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

### Terraform cloud

Terraform state files can be kept locally to your development environment. However, it is very easy to lose them or have them compromised.

Using S3 or a cloud provider can help to keep those files remote and safer.

In this particular case, I am using Terraform Cloud (in its free tier) to keep the state files remotely.

#### Organizations, Projects and Workspaces

This, in this very order, is the hierarchical structure within Terraform Cloud.

First, create an organization; then, create a Project; lastly, create Workspaces inside a Project.

These structural elements will be used inside the Terraform main.tf file when linking code to the Terraform Cloud.

#### Connect your code to Terraform cloud

Use the following code block in your main.tf file:

```terraform
terraform {
  cloud {
    organization = "ORGANIZATION-NAME"
    workspaces {
      name = "learn-terraform-cloud-migrate"
    }
  }
```
#### Issues with Terraform cloud state files

This issue presented when using Gitpod as the IDE, since its interface does not have a browser which turned out to be necessary for the configuration of the authentication in order to invoke Terraform Cloud from the main.tf file.

Write this command on your terminal:

```bash
terraform login
```

Press "P" to print the URL that you should use externally on a browser.

Visit the URL, follow the instructions and copy the token.

In my case, CTRL+V was enough to enter the token. 

Please, note that the token will not be visible. Now, press "enter".

You should see this message now:

```bash
Retrieved token for user user-name
```

You will now see a welcome banner and you are good to go now.

##### Workaround
Since the main goal is to have your local system know the token generated by Terraform Cloud, you can create the file that is going to keep the token locally, using the suggested path:

```bash
#touch /home/gitpod/.terraform.d/credentials.tfrc.json
```

Then, edit the file with your favourite editor and include this structure:

```json
{
    "credentials": {
        "app.terraform.io":{
            "token": "YOUR-TOKEN-FROM-THE-WEBSITE"
        }
    }
}
```
Init, plan and apply the terraform code and it should work just fine.

At this point, you should see your state file in the Terraform Cloud dashboard.

I have automated thos workaround with the following bash script [/bin/generate_credentials](bin/generate_credentials)
