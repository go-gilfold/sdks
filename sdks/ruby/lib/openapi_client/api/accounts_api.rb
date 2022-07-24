=begin
#

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 0.0.0

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.1.0-SNAPSHOT

=end

require 'cgi'

module OpenapiClient
  class AccountsApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # Delete an Account of a Business
    # Delete an Account of a Business
    # @param [Hash] opts the optional parameters
    # @return [Object]
    def accounts_account_id_delete(opts = {})
      data, _status_code, _headers = accounts_account_id_delete_with_http_info(opts)
      data
    end

    # Delete an Account of a Business
    # Delete an Account of a Business
    # @param [Hash] opts the optional parameters
    # @return [Array<(Object, Integer, Hash)>] Object data, response status code and response headers
    def accounts_account_id_delete_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: AccountsApi.accounts_account_id_delete ...'
      end
      # resource path
      local_var_path = '/accounts/:accountId'

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
      return_type = opts[:debug_return_type] || 'Object'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"AccountsApi.accounts_account_id_delete",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: AccountsApi#accounts_account_id_delete\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get an account for a Business
    # Get an account for a Business
    # @param [Hash] opts the optional parameters
    # @return [AccountResponse]
    def accounts_account_id_get(opts = {})
      data, _status_code, _headers = accounts_account_id_get_with_http_info(opts)
      data
    end

    # Get an account for a Business
    # Get an account for a Business
    # @param [Hash] opts the optional parameters
    # @return [Array<(AccountResponse, Integer, Hash)>] AccountResponse data, response status code and response headers
    def accounts_account_id_get_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: AccountsApi.accounts_account_id_get ...'
      end
      # resource path
      local_var_path = '/accounts/:accountId'

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
      return_type = opts[:debug_return_type] || 'AccountResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"AccountsApi.accounts_account_id_get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: AccountsApi#accounts_account_id_get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Update Account for Business
    # Update Account for Business
    # @param patch_account_request [PatchAccountRequest] 
    # @param [Hash] opts the optional parameters
    # @return [AccountResponse]
    def accounts_account_id_patch(patch_account_request, opts = {})
      data, _status_code, _headers = accounts_account_id_patch_with_http_info(patch_account_request, opts)
      data
    end

    # Update Account for Business
    # Update Account for Business
    # @param patch_account_request [PatchAccountRequest] 
    # @param [Hash] opts the optional parameters
    # @return [Array<(AccountResponse, Integer, Hash)>] AccountResponse data, response status code and response headers
    def accounts_account_id_patch_with_http_info(patch_account_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: AccountsApi.accounts_account_id_patch ...'
      end
      # verify the required parameter 'patch_account_request' is set
      if @api_client.config.client_side_validation && patch_account_request.nil?
        fail ArgumentError, "Missing the required parameter 'patch_account_request' when calling AccountsApi.accounts_account_id_patch"
      end
      # resource path
      local_var_path = '/accounts/:accountId'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(patch_account_request)

      # return_type
      return_type = opts[:debug_return_type] || 'AccountResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"AccountsApi.accounts_account_id_patch",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PATCH, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: AccountsApi#accounts_account_id_patch\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get Transactions for an Account
    # Get Transactions for an Account
    # @param [Hash] opts the optional parameters
    # @return [AllTransactionsResponse]
    def accounts_account_id_transactions_get(opts = {})
      data, _status_code, _headers = accounts_account_id_transactions_get_with_http_info(opts)
      data
    end

    # Get Transactions for an Account
    # Get Transactions for an Account
    # @param [Hash] opts the optional parameters
    # @return [Array<(AllTransactionsResponse, Integer, Hash)>] AllTransactionsResponse data, response status code and response headers
    def accounts_account_id_transactions_get_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: AccountsApi.accounts_account_id_transactions_get ...'
      end
      # resource path
      local_var_path = '/accounts/:accountId/transactions'

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
      return_type = opts[:debug_return_type] || 'AllTransactionsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"AccountsApi.accounts_account_id_transactions_get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: AccountsApi#accounts_account_id_transactions_get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get all Accounts for a Business
    # Get all Accounts for a Business
    # @param [Hash] opts the optional parameters
    # @return [AllAccountsResponse]
    def accounts_get(opts = {})
      data, _status_code, _headers = accounts_get_with_http_info(opts)
      data
    end

    # Get all Accounts for a Business
    # Get all Accounts for a Business
    # @param [Hash] opts the optional parameters
    # @return [Array<(AllAccountsResponse, Integer, Hash)>] AllAccountsResponse data, response status code and response headers
    def accounts_get_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: AccountsApi.accounts_get ...'
      end
      # resource path
      local_var_path = '/accounts'

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
      return_type = opts[:debug_return_type] || 'AllAccountsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"AccountsApi.accounts_get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: AccountsApi#accounts_get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Create an Account for a Business
    # Create an Account for a Business
    # @param post_account_request [PostAccountRequest] 
    # @param [Hash] opts the optional parameters
    # @return [AccountResponse]
    def accounts_post(post_account_request, opts = {})
      data, _status_code, _headers = accounts_post_with_http_info(post_account_request, opts)
      data
    end

    # Create an Account for a Business
    # Create an Account for a Business
    # @param post_account_request [PostAccountRequest] 
    # @param [Hash] opts the optional parameters
    # @return [Array<(AccountResponse, Integer, Hash)>] AccountResponse data, response status code and response headers
    def accounts_post_with_http_info(post_account_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: AccountsApi.accounts_post ...'
      end
      # verify the required parameter 'post_account_request' is set
      if @api_client.config.client_side_validation && post_account_request.nil?
        fail ArgumentError, "Missing the required parameter 'post_account_request' when calling AccountsApi.accounts_post"
      end
      # resource path
      local_var_path = '/accounts'

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
      post_body = opts[:debug_body] || @api_client.object_to_http_body(post_account_request)

      # return_type
      return_type = opts[:debug_return_type] || 'AccountResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"AccountsApi.accounts_post",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: AccountsApi#accounts_post\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get stats for all accounts
    # Get stats for all accounts
    # @param [Hash] opts the optional parameters
    # @return [AccountStatsResponse]
    def accounts_stats_get(opts = {})
      data, _status_code, _headers = accounts_stats_get_with_http_info(opts)
      data
    end

    # Get stats for all accounts
    # Get stats for all accounts
    # @param [Hash] opts the optional parameters
    # @return [Array<(AccountStatsResponse, Integer, Hash)>] AccountStatsResponse data, response status code and response headers
    def accounts_stats_get_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: AccountsApi.accounts_stats_get ...'
      end
      # resource path
      local_var_path = '/accounts/stats'

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
      return_type = opts[:debug_return_type] || 'AccountStatsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"AccountsApi.accounts_stats_get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: AccountsApi#accounts_stats_get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
