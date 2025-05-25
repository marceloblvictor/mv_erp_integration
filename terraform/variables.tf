variable "resource_group_location" {
  type        = string
  default     = "spaincentral"
  description = "Location of the resource group."
}

variable "resource_group_name" {
  type        = string
  default     = "mv-erp-integration-rg"
  description = "Name of the resource group."  
}

variable "acr_name" {
  type        = string
  default     = "mverpintegration"
  description = "Name of the azure container registry."  
}

