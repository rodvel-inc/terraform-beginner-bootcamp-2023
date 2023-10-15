variable "user_uuid" {
    description = "The UUID for the user"
    type = string
validation {
    condition     = can(regex("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", var.user_uuid))
    error_message = "Invalid user_uuid format. It should be in the form of UUID."
  }
}