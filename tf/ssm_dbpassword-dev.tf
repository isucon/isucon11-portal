resource "aws_ssm_parameter" "dbpassword-dev" {
  name   = "/hako/isuxportal-dev/DATABASE_PASSWORD"
  type   = "SecureString"
  key_id = aws_kms_key.isuxportal.arn
  value  = random_password.isuxportal-rds-dev.result
}
