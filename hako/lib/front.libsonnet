local utils = import './utils.libsonnet';

{
  container: {
    cpu: 64,
    memory: 32,
    image_tag: 'latest',
    log_configuration: utils.awsLogs('front'),
  },
  script: {
    local default = {
      type: 'nginx_front',
      s3: {
        region: 'ap-northeast-1',
        bucket: 'isucon11-misc',
        prefix: 'hako/nginx-config',
      },
    },
    default: default {
      locations: {
        '/': {
          https_type: 'always',
        },
      },
    },
  },
}
