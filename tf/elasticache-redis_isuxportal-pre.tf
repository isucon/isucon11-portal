#resource "aws_elasticache_replication_group" "isuxportal-pre" {
#  replication_group_id          = "isuxportal-pre"
#  replication_group_description = "Redis for isuxportal-pre"
#  node_type                     = "cache.t3.micro"
#  engine_version                = "5.0.6"
#  number_cache_clusters         = 2
#  port                          = 6379
#  parameter_group_name          = "default.redis5.0"
#
#  automatic_failover_enabled = true
#  subnet_group_name          = aws_elasticache_subnet_group.isuxportal.name
#
#  security_group_ids = [
#    aws_security_group.default.id,
#    aws_security_group.redis.id,
#  ]
#
#  maintenance_window         = "mon:16:30-mon:17:30"
#  snapshot_window            = "15:00-16:00"
#  snapshot_retention_limit   = 7
#  auto_minor_version_upgrade = false
#  apply_immediately          = true
#
#}
#
