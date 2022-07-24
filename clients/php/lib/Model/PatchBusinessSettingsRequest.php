<?php
/**
 * PatchBusinessSettingsRequest
 *
 * PHP version 7.4
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * 
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.0.1
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 6.1.0-SNAPSHOT
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * PatchBusinessSettingsRequest Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<string, mixed>
 */
class PatchBusinessSettingsRequest implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'PatchBusinessSettingsRequest';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'currency_label' => 'string',
        'currency_label_suffixed' => 'bool',
        'accounts_overdraftable' => 'bool',
        'billing_type' => '\OpenAPI\Client\Model\BillingType',
        'payment_providers' => 'string[]',
        'stripe_sandbox_publishable_key' => 'string',
        'stripe_sandbox_secret_key' => 'string',
        'stripe_publishable_key' => 'string',
        'stripe_secret_key' => 'string',
        'paypal_client_id' => 'string',
        'paypal_client_secret' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'currency_label' => null,
        'currency_label_suffixed' => null,
        'accounts_overdraftable' => null,
        'billing_type' => null,
        'payment_providers' => null,
        'stripe_sandbox_publishable_key' => null,
        'stripe_sandbox_secret_key' => null,
        'stripe_publishable_key' => null,
        'stripe_secret_key' => null,
        'paypal_client_id' => null,
        'paypal_client_secret' => null
    ];

    /**
      * Array of nullable properties. Used for (de)serialization
      *
      * @var boolean[]
      */
    protected static array $openAPINullables = [
        'currency_label' => false,
		'currency_label_suffixed' => false,
		'accounts_overdraftable' => false,
		'billing_type' => false,
		'payment_providers' => false,
		'stripe_sandbox_publishable_key' => false,
		'stripe_sandbox_secret_key' => false,
		'stripe_publishable_key' => false,
		'stripe_secret_key' => false,
		'paypal_client_id' => false,
		'paypal_client_secret' => false
    ];

    /**
      * If a nullable field gets set to null, insert it here
      *
      * @var boolean[]
      */
    protected array $openAPINullablesSetToNull = [];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of nullable properties
     *
     * @return array
     */
    protected static function openAPINullables(): array
    {
        return self::$openAPINullables;
    }

    /**
     * Array of nullable field names deliberately set to null
     *
     * @return boolean[]
     */
    private function getOpenAPINullablesSetToNull(): array
    {
        return $this->openAPINullablesSetToNull;
    }

    /**
     * Checks if a property is nullable
     *
     * @param string $property
     * @return bool
     */
    public static function isNullable(string $property): bool
    {
        return self::openAPINullables()[$property] ?? false;
    }

    /**
     * Checks if a nullable property is set to null.
     *
     * @param string $property
     * @return bool
     */
    public function isNullableSetToNull(string $property): bool
    {
        return in_array($property, $this->getOpenAPINullablesSetToNull(), true);
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'currency_label' => 'currencyLabel',
        'currency_label_suffixed' => 'currencyLabelSuffixed',
        'accounts_overdraftable' => 'accountsOverdraftable',
        'billing_type' => 'billingType',
        'payment_providers' => 'paymentProviders',
        'stripe_sandbox_publishable_key' => 'stripeSandboxPublishableKey',
        'stripe_sandbox_secret_key' => 'stripeSandboxSecretKey',
        'stripe_publishable_key' => 'stripePublishableKey',
        'stripe_secret_key' => 'stripeSecretKey',
        'paypal_client_id' => 'paypalClientId',
        'paypal_client_secret' => 'paypalClientSecret'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'currency_label' => 'setCurrencyLabel',
        'currency_label_suffixed' => 'setCurrencyLabelSuffixed',
        'accounts_overdraftable' => 'setAccountsOverdraftable',
        'billing_type' => 'setBillingType',
        'payment_providers' => 'setPaymentProviders',
        'stripe_sandbox_publishable_key' => 'setStripeSandboxPublishableKey',
        'stripe_sandbox_secret_key' => 'setStripeSandboxSecretKey',
        'stripe_publishable_key' => 'setStripePublishableKey',
        'stripe_secret_key' => 'setStripeSecretKey',
        'paypal_client_id' => 'setPaypalClientId',
        'paypal_client_secret' => 'setPaypalClientSecret'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'currency_label' => 'getCurrencyLabel',
        'currency_label_suffixed' => 'getCurrencyLabelSuffixed',
        'accounts_overdraftable' => 'getAccountsOverdraftable',
        'billing_type' => 'getBillingType',
        'payment_providers' => 'getPaymentProviders',
        'stripe_sandbox_publishable_key' => 'getStripeSandboxPublishableKey',
        'stripe_sandbox_secret_key' => 'getStripeSandboxSecretKey',
        'stripe_publishable_key' => 'getStripePublishableKey',
        'stripe_secret_key' => 'getStripeSecretKey',
        'paypal_client_id' => 'getPaypalClientId',
        'paypal_client_secret' => 'getPaypalClientSecret'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    public const PAYMENT_PROVIDERS_PAYPAL = 'paypal';
    public const PAYMENT_PROVIDERS_STRIPE = 'stripe';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getPaymentProvidersAllowableValues()
    {
        return [
            self::PAYMENT_PROVIDERS_PAYPAL,
            self::PAYMENT_PROVIDERS_STRIPE,
        ];
    }

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->setIfExists('currency_label', $data ?? [], null);
        $this->setIfExists('currency_label_suffixed', $data ?? [], null);
        $this->setIfExists('accounts_overdraftable', $data ?? [], null);
        $this->setIfExists('billing_type', $data ?? [], null);
        $this->setIfExists('payment_providers', $data ?? [], null);
        $this->setIfExists('stripe_sandbox_publishable_key', $data ?? [], null);
        $this->setIfExists('stripe_sandbox_secret_key', $data ?? [], null);
        $this->setIfExists('stripe_publishable_key', $data ?? [], null);
        $this->setIfExists('stripe_secret_key', $data ?? [], null);
        $this->setIfExists('paypal_client_id', $data ?? [], null);
        $this->setIfExists('paypal_client_secret', $data ?? [], null);
    }

    /**
    * Sets $this->container[$variableName] to the given data or to the given default Value; if $variableName
    * is nullable and its value is set to null in the $fields array, then mark it as "set to null" in the
    * $this->openAPINullablesSetToNull array
    *
    * @param string $variableName
    * @param array  $fields
    * @param mixed  $defaultValue
    */
    private function setIfExists(string $variableName, array $fields, $defaultValue): void
    {
        if (self::isNullable($variableName) && array_key_exists($variableName, $fields) && is_null($fields[$variableName])) {
            $this->openAPINullablesSetToNull[] = $variableName;
        }

        $this->container[$variableName] = $fields[$variableName] ?? $defaultValue;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets currency_label
     *
     * @return string|null
     */
    public function getCurrencyLabel()
    {
        return $this->container['currency_label'];
    }

    /**
     * Sets currency_label
     *
     * @param string|null $currency_label currency_label
     *
     * @return self
     */
    public function setCurrencyLabel($currency_label)
    {

        if (is_null($currency_label)) {
            throw new \InvalidArgumentException('non-nullable currency_label cannot be null');
        }

        $this->container['currency_label'] = $currency_label;

        return $this;
    }

    /**
     * Gets currency_label_suffixed
     *
     * @return bool|null
     */
    public function getCurrencyLabelSuffixed()
    {
        return $this->container['currency_label_suffixed'];
    }

    /**
     * Sets currency_label_suffixed
     *
     * @param bool|null $currency_label_suffixed currency_label_suffixed
     *
     * @return self
     */
    public function setCurrencyLabelSuffixed($currency_label_suffixed)
    {

        if (is_null($currency_label_suffixed)) {
            throw new \InvalidArgumentException('non-nullable currency_label_suffixed cannot be null');
        }

        $this->container['currency_label_suffixed'] = $currency_label_suffixed;

        return $this;
    }

    /**
     * Gets accounts_overdraftable
     *
     * @return bool|null
     */
    public function getAccountsOverdraftable()
    {
        return $this->container['accounts_overdraftable'];
    }

    /**
     * Sets accounts_overdraftable
     *
     * @param bool|null $accounts_overdraftable accounts_overdraftable
     *
     * @return self
     */
    public function setAccountsOverdraftable($accounts_overdraftable)
    {

        if (is_null($accounts_overdraftable)) {
            throw new \InvalidArgumentException('non-nullable accounts_overdraftable cannot be null');
        }

        $this->container['accounts_overdraftable'] = $accounts_overdraftable;

        return $this;
    }

    /**
     * Gets billing_type
     *
     * @return \OpenAPI\Client\Model\BillingType|null
     */
    public function getBillingType()
    {
        return $this->container['billing_type'];
    }

    /**
     * Sets billing_type
     *
     * @param \OpenAPI\Client\Model\BillingType|null $billing_type billing_type
     *
     * @return self
     */
    public function setBillingType($billing_type)
    {

        if (is_null($billing_type)) {
            throw new \InvalidArgumentException('non-nullable billing_type cannot be null');
        }

        $this->container['billing_type'] = $billing_type;

        return $this;
    }

    /**
     * Gets payment_providers
     *
     * @return string[]|null
     */
    public function getPaymentProviders()
    {
        return $this->container['payment_providers'];
    }

    /**
     * Sets payment_providers
     *
     * @param string[]|null $payment_providers payment_providers
     *
     * @return self
     */
    public function setPaymentProviders($payment_providers)
    {
        $allowedValues = $this->getPaymentProvidersAllowableValues();
        if (!is_null($payment_providers) && array_diff($payment_providers, $allowedValues)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value for 'payment_providers', must be one of '%s'",
                    implode("', '", $allowedValues)
                )
            );
        }

        if (is_null($payment_providers)) {
            throw new \InvalidArgumentException('non-nullable payment_providers cannot be null');
        }

        $this->container['payment_providers'] = $payment_providers;

        return $this;
    }

    /**
     * Gets stripe_sandbox_publishable_key
     *
     * @return string|null
     */
    public function getStripeSandboxPublishableKey()
    {
        return $this->container['stripe_sandbox_publishable_key'];
    }

    /**
     * Sets stripe_sandbox_publishable_key
     *
     * @param string|null $stripe_sandbox_publishable_key stripe_sandbox_publishable_key
     *
     * @return self
     */
    public function setStripeSandboxPublishableKey($stripe_sandbox_publishable_key)
    {

        if (is_null($stripe_sandbox_publishable_key)) {
            throw new \InvalidArgumentException('non-nullable stripe_sandbox_publishable_key cannot be null');
        }

        $this->container['stripe_sandbox_publishable_key'] = $stripe_sandbox_publishable_key;

        return $this;
    }

    /**
     * Gets stripe_sandbox_secret_key
     *
     * @return string|null
     */
    public function getStripeSandboxSecretKey()
    {
        return $this->container['stripe_sandbox_secret_key'];
    }

    /**
     * Sets stripe_sandbox_secret_key
     *
     * @param string|null $stripe_sandbox_secret_key stripe_sandbox_secret_key
     *
     * @return self
     */
    public function setStripeSandboxSecretKey($stripe_sandbox_secret_key)
    {

        if (is_null($stripe_sandbox_secret_key)) {
            throw new \InvalidArgumentException('non-nullable stripe_sandbox_secret_key cannot be null');
        }

        $this->container['stripe_sandbox_secret_key'] = $stripe_sandbox_secret_key;

        return $this;
    }

    /**
     * Gets stripe_publishable_key
     *
     * @return string|null
     */
    public function getStripePublishableKey()
    {
        return $this->container['stripe_publishable_key'];
    }

    /**
     * Sets stripe_publishable_key
     *
     * @param string|null $stripe_publishable_key stripe_publishable_key
     *
     * @return self
     */
    public function setStripePublishableKey($stripe_publishable_key)
    {

        if (is_null($stripe_publishable_key)) {
            throw new \InvalidArgumentException('non-nullable stripe_publishable_key cannot be null');
        }

        $this->container['stripe_publishable_key'] = $stripe_publishable_key;

        return $this;
    }

    /**
     * Gets stripe_secret_key
     *
     * @return string|null
     */
    public function getStripeSecretKey()
    {
        return $this->container['stripe_secret_key'];
    }

    /**
     * Sets stripe_secret_key
     *
     * @param string|null $stripe_secret_key stripe_secret_key
     *
     * @return self
     */
    public function setStripeSecretKey($stripe_secret_key)
    {

        if (is_null($stripe_secret_key)) {
            throw new \InvalidArgumentException('non-nullable stripe_secret_key cannot be null');
        }

        $this->container['stripe_secret_key'] = $stripe_secret_key;

        return $this;
    }

    /**
     * Gets paypal_client_id
     *
     * @return string|null
     */
    public function getPaypalClientId()
    {
        return $this->container['paypal_client_id'];
    }

    /**
     * Sets paypal_client_id
     *
     * @param string|null $paypal_client_id paypal_client_id
     *
     * @return self
     */
    public function setPaypalClientId($paypal_client_id)
    {

        if (is_null($paypal_client_id)) {
            throw new \InvalidArgumentException('non-nullable paypal_client_id cannot be null');
        }

        $this->container['paypal_client_id'] = $paypal_client_id;

        return $this;
    }

    /**
     * Gets paypal_client_secret
     *
     * @return string|null
     */
    public function getPaypalClientSecret()
    {
        return $this->container['paypal_client_secret'];
    }

    /**
     * Sets paypal_client_secret
     *
     * @param string|null $paypal_client_secret paypal_client_secret
     *
     * @return self
     */
    public function setPaypalClientSecret($paypal_client_secret)
    {

        if (is_null($paypal_client_secret)) {
            throw new \InvalidArgumentException('non-nullable paypal_client_secret cannot be null');
        }

        $this->container['paypal_client_secret'] = $paypal_client_secret;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset): bool
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    #[\ReturnTypeWillChange]
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value): void
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset): void
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    #[\ReturnTypeWillChange]
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}

