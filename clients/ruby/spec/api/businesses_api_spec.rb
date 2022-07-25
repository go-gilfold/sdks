=begin
#

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 0.0.0

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.1.0-SNAPSHOT

=end

require 'spec_helper'
require 'json'

# Unit tests for OpenapiClient::BusinessesApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'BusinessesApi' do
  before do
    # run before each test
    @api_instance = OpenapiClient::BusinessesApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of BusinessesApi' do
    it 'should create an instance of BusinessesApi' do
      expect(@api_instance).to be_instance_of(OpenapiClient::BusinessesApi)
    end
  end

  # unit tests for business_business_id_settings_get
  # Get Settings for a business
  # Get Settings for a business
  # @param [Hash] opts the optional parameters
  # @return [BusinessSettingsResponse]
  describe 'business_business_id_settings_get test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for businesses_business_id_get
  # Get Business Details
  #  Use this to get the name and other business details. Here is a link: [to google](https://google.com). Let&#39;s see if it works. 
  # @param [Hash] opts the optional parameters
  # @return [BusinessResponse]
  describe 'businesses_business_id_get test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for businesses_business_id_patch
  # Update Business Details
  # Update Business Details
  # @param patch_business_request 
  # @param [Hash] opts the optional parameters
  # @return [BusinessResponse]
  describe 'businesses_business_id_patch test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
