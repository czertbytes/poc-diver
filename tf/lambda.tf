resource "aws_lambda_function" "diver" {
  function_name = var.lambda_function_name
  filename      = var.lambda_filename
  role          = aws_iam_role.lambda.arn
  handler       = var.lambda_handler
  runtime       = "go1.x"
  memory_size   = 128
  timeout       = 60

  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs,
    aws_cloudwatch_log_group.lambda_logs,
  ]
}

resource "aws_lambda_permission" "allow_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.diver.arn
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.s3_input.arn
}