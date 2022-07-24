=begin
#

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 0.0.1

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.1.0-SNAPSHOT

=end

require 'cgi'

module OpenapiClient
  class UsageApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # Get an Account's current Usage Billing Rate Adjustments
    # Get an Account's current Usage Billing Rate Adjustments
    # @param [Hash] opts the optional parameters
    # @return [AccountUsageBillingAdjustmentsResponse]
    def accounts_account_id_usage_get(opts = {})
      data, _status_code, _headers = accounts_account_id_usage_get_with_http_info(opts)
      data
    end

    # Get an Account&#39;s current Usage Billing Rate Adjustments
    # Get an Account&#39;s current Usage Billing Rate Adjustments
    # @param [Hash] opts the optional parameters
    # @return [Array<(AccountUsageBillingAdjustmentsResponse, Integer, Hash)>] AccountUsageBillingAdjustmentsResponse data, response status code and response headers
    def accounts_account_id_usage_get_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: UsageApi.accounts_account_id_usage_get ...'
      end
      # resource path
      local_var_path = '/accounts/:accountId/usage'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'AccountUsageBillingAdjustmentsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"UsageApi.accounts_account_id_usage_get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: UsageApi#accounts_account_id_usage_get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Set an Account's Usage Billing Rate
    # Set an Account's Usage Billing Rate
    # @param post_account_usage_billing_rate_adjustment_request [PostAccountUsageBillingRateAdjustmentRequest] 
    # @param [Hash] opts the optional parameters
    # @return [Object]
    def accounts_account_id_usage_post(post_account_usage_billing_rate_adjustment_request, opts = {})
      data, _status_code, _headers = accounts_account_id_usage_post_with_http_info(post_account_usage_billing_rate_adjustment_request, opts)
      data
    end

    # Set an Account&#39;s Usage Billing Rate
    # Set an Account&#39;s Usage Billing Rate
    # @param post_account_usage_billing_rate_adjustment_request [PostAccountUsageBillingRateAdjustmentRequest] 
    # @param [Hash] opts the optional parameters
    # @return [Array<(Object, Integer, Hash)>] Object data, response status code and response headers
    def accounts_account_id_usage_post_with_http_info(post_account_usage_billing_rate_adjustment_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: UsageApi.accounts_account_id_usage_post ...'
      end
      # verify the required parameter 'post_account_usage_billing_rate_adjustment_request' is set
      if @api_client.config.client_side_validation && post_account_usage_billing_rate_adjustment_request.nil?
        fail ArgumentError, "Missing the required parameter 'post_account_usage_billing_rate_adjustment_request' when calling UsageApi.accounts_account_id_usage_post"
      end
      # resource path
      local_var_path = '/accounts/:accountId/usage'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(post_account_usage_billing_rate_adjustment_request)

      # return_type
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"UsageApi.accounts_account_id_usage_post",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: UsageApi#accounts_account_id_usage_post\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end