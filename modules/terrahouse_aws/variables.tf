variable "user_uuid" {
    description = "The UUID for the user"
    type = string
validation {
    condition     = can(regex("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", var.user_uuid))
    error_message = "Invalid user_uuid format. It should be in the form of UUID."
  }
}

variable "bucket_name" {
    description = "The name of the S3 bucket"
    type = string

    validation {
        condition = (
            length(var.bucket_name) >= 3 && length(var.bucket_name) <= 63 &&
            can(regex("^[a-z0-9][a-z0-9-.]*[a-z0-9]$", var.bucket_name))
        )
        error_message = "The bucket name does not comply with the naming standard."
    }
}  