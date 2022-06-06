resource "aws_s3_bucket" "s3_input" {
  bucket        = "poc-diver-input"
  force_destroy = true
}

resource "aws_s3_bucket_acl" "s3_input_acl" {
  bucket = aws_s3_bucket.s3_input.id
  acl    = "private"
}

resource "aws_s3_bucket_notification" "s3_file_created" {
  bucket = aws_s3_bucket.s3_input.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.diver.arn
    events              = ["s3:ObjectCreated:*"]
    filter_suffix       = ".csv"
  }

  depends_on = [aws_lambda_permission.allow_bucket]
}

resource "aws_s3_bucket" "s3_output" {
  bucket        = "poc-diver-output"
  force_destroy = true
}

resource "aws_s3_bucket_acl" "s3_output_acl" {
  bucket = aws_s3_bucket.s3_output.id
  acl    = "private"
}