# OpenAPI\Client\UsageApi

All URIs are relative to http://localhost.

Method | HTTP request | Description
------------- | ------------- | -------------
[**accountsAccountIdUsageGet()**](UsageApi.md#accountsAccountIdUsageGet) | **GET** /accounts/:accountId/usage | Get an Account&#39;s current Usage Billing Rate Adjustments
[**accountsAccountIdUsagePost()**](UsageApi.md#accountsAccountIdUsagePost) | **POST** /accounts/:accountId/usage | Set an Account&#39;s Usage Billing Rate


## `accountsAccountIdUsageGet()`

```php
accountsAccountIdUsageGet(): \OpenAPI\Client\Model\AccountUsageBillingAdjustmentsResponse
```

Get an Account's current Usage Billing Rate Adjustments

Get an Account's current Usage Billing Rate Adjustments

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



$apiInstance = new OpenAPI\Client\Api\UsageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);

try {
    $result = $apiInstance->accountsAccountIdUsageGet();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsageApi->accountsAccountIdUsageGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\AccountUsageBillingAdjustmentsResponse**](../Model/AccountUsageBillingAdjustmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `accountsAccountIdUsagePost()`

```php
accountsAccountIdUsagePost($post_account_usage_billing_rate_adjustment_request): object
```

Set an Account's Usage Billing Rate

Set an Account's Usage Billing Rate

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



$apiInstance = new OpenAPI\Client\Api\UsageApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$post_account_usage_billing_rate_adjustment_request = new \OpenAPI\Client\Model\PostAccountUsageBillingRateAdjustmentRequest(); // \OpenAPI\Client\Model\PostAccountUsageBillingRateAdjustmentRequest

try {
    $result = $apiInstance->accountsAccountIdUsagePost($post_account_usage_billing_rate_adjustment_request);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UsageApi->accountsAccountIdUsagePost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **post_account_usage_billing_rate_adjustment_request** | [**\OpenAPI\Client\Model\PostAccountUsageBillingRateAdjustmentRequest**](../Model/PostAccountUsageBillingRateAdjustmentRequest.md)|  |

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)