=begin
#

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 0.0.0

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.1.0-SNAPSHOT

=end

# Common files
require 'openapi_client/api_client'
require 'openapi_client/api_error'
require 'openapi_client/version'
require 'openapi_client/configuration'

# Models
require 'openapi_client/models/account_response'
require 'openapi_client/models/account_stats_response'
require 'openapi_client/models/account_stats_response_daily_inner'
require 'openapi_client/models/account_usage_billing_adjustments_response'
require 'openapi_client/models/account_usage_billing_adjustments_response_adjustments_inner'
require 'openapi_client/models/all_accounts_response'
require 'openapi_client/models/all_transactions_response'
require 'openapi_client/models/business_response'
require 'openapi_client/models/business_settings_response'
require 'openapi_client/models/error_response'
require 'openapi_client/models/patch_account_request'
require 'openapi_client/models/patch_business_request'
require 'openapi_client/models/patch_business_settings_request'
require 'openapi_client/models/patch_transaction_request'
require 'openapi_client/models/post_account_request'
require 'openapi_client/models/post_account_usage_billing_rate_adjustment_request'
require 'openapi_client/models/post_transaction_request'
require 'openapi_client/models/transaction_response'
require 'openapi_client/models/transaction_stats_response'
require 'openapi_client/models/transaction_stats_response_daily_inner'

# APIs
require 'openapi_client/api/accounts_api'
require 'openapi_client/api/businesses_api'
require 'openapi_client/api/transactions_api'
require 'openapi_client/api/usage_api'

module OpenapiClient
  class << self
    # Customize default settings for the SDK using block.
    #   OpenapiClient.configure do |config|
    #     config.username = "xxx"
    #     config.password = "xxx"
    #   end
    # If no block given, return the default Configuration object.
    def configure
      if block_given?
        yield(Configuration.default)
      else
        Configuration.default
      end
    end
  end
end
