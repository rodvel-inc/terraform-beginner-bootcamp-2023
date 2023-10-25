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

variable "index_html_filepath" {
    description = "The filepath for index.html"
    type = string

    validation {
      condition = fileexists(var.index_html_filepath)
      error_message = "The provided path for index.html does not exist."
    }
}

variable "error_html_filepath" {
    description = "The filepath for error.html"
    type = string

    validation {
      condition = fileexists(var.error_html_filepath)
      error_message = "The provided path for error.html does not exist."
    }
}

variable "content_version" {
  description = "The version of the content"
  type        = number
  default     = 1
  validation {
    condition     = can(regex("^[1-9][0-9]*$", var.content_version))
    error_message = "Content version must be a positive integer starting at 1."
  }
}
