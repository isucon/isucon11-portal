#resource "aws_route53_record" "cname_portal-prd-xi-isucon-dev" {
#  zone_id = aws_route53_zone.public.zone_id
#  name    = "portal-prd.xi.isucon.dev"
#  type    = "CNAME"
#  ttl     = "300"
#  records = ["tmp"]
#}
#
#resource "aws_route53_record" "cname_portal-grpc-prd-x-isucon-dev" {
#  zone_id = aws_route53_zone.public.zone_id
#  name    = "portal-grpc-prd.xi.isucon.dev"
#  type    = "CNAME"
#  ttl     = "300"
#  records = ["tmp"]
#}
#
#resource "aws_route53_record" "cname_portal-isucon-net-x-isucon-dev" {
#  zone_id = aws_route53_zone.public.zone_id
#  name    = "portal-isucon-net.xi.isucon.dev"
#  type    = "CNAME"
#  ttl     = "300"
#  records = ["tmp"]
#}

data "aws_lb" "hako-isuxportal-dev-fargate" {
  name = "hako-isuxportal-dev-fargate"
}

resource "aws_route53_record" "cname_portal-dev-xi-isucon-dev" {
  for_each = {
    for k in [
      aws_route53_zone.public.zone_id,
      #aws_route53_zone.private.zone_id,
    ] : k => k
  }
  zone_id = each.value
  name    = "portal-dev.xi.isucon.dev"
  type    = "CNAME"
  ttl     = "300"
  records = [data.aws_lb.hako-isuxportal-dev-fargate.dns_name]
}

resource "aws_route53_record" "cname_portal-dev-isucon-net-xi-isucon-dev" {
  for_each = {
    for k in [
      aws_route53_zone.public.zone_id,
      #aws_route53_zone.private.zone_id,
    ] : k => k
  }
  zone_id = each.value
  name    = "portal-dev-isucon-net.xi.isucon.dev"
  type    = "CNAME"
  ttl     = "300"
  records = ["${aws_cloudfront_distribution.isuxportal-dev.domain_name}."]
}
