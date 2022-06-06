variable "lambda_function_name" {
  default = "poc_diver"
}

variable "lambda_filename" {
  default = "../build/main.zip"
}

variable "lambda_handler" {
  default = "main"
}