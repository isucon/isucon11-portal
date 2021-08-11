require 'erb'

module CloudFormation
  TEST_ERB = ERB.new(File.read(File.join(File.dirname(__FILE__), './cf_templates/test.yaml.erb')))

  def self.template(erb, b)
    erb.result(b)
  end

  def self.test_template(zone_id)
    unless zone_id.nil?
      template(TEST_ERB, binding)
    else
      ''
    end
  end
end