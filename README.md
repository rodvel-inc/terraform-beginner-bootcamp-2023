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