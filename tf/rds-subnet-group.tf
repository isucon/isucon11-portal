resource "aws_db_subnet_group" "isuxportal" {
  name        = "isuxportal"
  description = "isuxportal"

  subnet_ids = aws_subnet.private.*.id
}
