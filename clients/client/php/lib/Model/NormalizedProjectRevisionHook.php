<?php
/**
 * NormalizedProjectRevisionHook
 *
 * PHP version 7.3
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.
 *
 * The version of the OpenAPI document: v0.2.0-alpha.48
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.4.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace Ory\Client\Model;

use \ArrayAccess;
use \Ory\Client\ObjectSerializer;

/**
 * NormalizedProjectRevisionHook Class Doc Comment
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class NormalizedProjectRevisionHook implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'normalizedProjectRevisionHook';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'configKey' => 'string',
        'createdAt' => '\DateTime',
        'hook' => 'string',
        'id' => 'string',
        'projectRevisionId' => 'string',
        'updatedAt' => '\DateTime',
        'webHookConfigAuthApiKeyIn' => 'string',
        'webHookConfigAuthApiKeyName' => 'string',
        'webHookConfigAuthApiKeyValue' => 'string',
        'webHookConfigAuthBasicAuthPassword' => 'string',
        'webHookConfigAuthBasicAuthUser' => 'string',
        'webHookConfigAuthType' => 'string',
        'webHookConfigBody' => 'string',
        'webHookConfigCanInterrupt' => 'bool',
        'webHookConfigMethod' => 'string',
        'webHookConfigResponseIgnore' => 'bool',
        'webHookConfigUrl' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'configKey' => null,
        'createdAt' => 'date-time',
        'hook' => null,
        'id' => 'uuid',
        'projectRevisionId' => 'uuid',
        'updatedAt' => 'date-time',
        'webHookConfigAuthApiKeyIn' => null,
        'webHookConfigAuthApiKeyName' => null,
        'webHookConfigAuthApiKeyValue' => null,
        'webHookConfigAuthBasicAuthPassword' => null,
        'webHookConfigAuthBasicAuthUser' => null,
        'webHookConfigAuthType' => null,
        'webHookConfigBody' => null,
        'webHookConfigCanInterrupt' => null,
        'webHookConfigMethod' => null,
        'webHookConfigResponseIgnore' => null,
        'webHookConfigUrl' => null
    ];

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
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'configKey' => 'config_key',
        'createdAt' => 'created_at',
        'hook' => 'hook',
        'id' => 'id',
        'projectRevisionId' => 'project_revision_id',
        'updatedAt' => 'updated_at',
        'webHookConfigAuthApiKeyIn' => 'web_hook_config_auth_api_key_in',
        'webHookConfigAuthApiKeyName' => 'web_hook_config_auth_api_key_name',
        'webHookConfigAuthApiKeyValue' => 'web_hook_config_auth_api_key_value',
        'webHookConfigAuthBasicAuthPassword' => 'web_hook_config_auth_basic_auth_password',
        'webHookConfigAuthBasicAuthUser' => 'web_hook_config_auth_basic_auth_user',
        'webHookConfigAuthType' => 'web_hook_config_auth_type',
        'webHookConfigBody' => 'web_hook_config_body',
        'webHookConfigCanInterrupt' => 'web_hook_config_can_interrupt',
        'webHookConfigMethod' => 'web_hook_config_method',
        'webHookConfigResponseIgnore' => 'web_hook_config_response_ignore',
        'webHookConfigUrl' => 'web_hook_config_url'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'configKey' => 'setConfigKey',
        'createdAt' => 'setCreatedAt',
        'hook' => 'setHook',
        'id' => 'setId',
        'projectRevisionId' => 'setProjectRevisionId',
        'updatedAt' => 'setUpdatedAt',
        'webHookConfigAuthApiKeyIn' => 'setWebHookConfigAuthApiKeyIn',
        'webHookConfigAuthApiKeyName' => 'setWebHookConfigAuthApiKeyName',
        'webHookConfigAuthApiKeyValue' => 'setWebHookConfigAuthApiKeyValue',
        'webHookConfigAuthBasicAuthPassword' => 'setWebHookConfigAuthBasicAuthPassword',
        'webHookConfigAuthBasicAuthUser' => 'setWebHookConfigAuthBasicAuthUser',
        'webHookConfigAuthType' => 'setWebHookConfigAuthType',
        'webHookConfigBody' => 'setWebHookConfigBody',
        'webHookConfigCanInterrupt' => 'setWebHookConfigCanInterrupt',
        'webHookConfigMethod' => 'setWebHookConfigMethod',
        'webHookConfigResponseIgnore' => 'setWebHookConfigResponseIgnore',
        'webHookConfigUrl' => 'setWebHookConfigUrl'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'configKey' => 'getConfigKey',
        'createdAt' => 'getCreatedAt',
        'hook' => 'getHook',
        'id' => 'getId',
        'projectRevisionId' => 'getProjectRevisionId',
        'updatedAt' => 'getUpdatedAt',
        'webHookConfigAuthApiKeyIn' => 'getWebHookConfigAuthApiKeyIn',
        'webHookConfigAuthApiKeyName' => 'getWebHookConfigAuthApiKeyName',
        'webHookConfigAuthApiKeyValue' => 'getWebHookConfigAuthApiKeyValue',
        'webHookConfigAuthBasicAuthPassword' => 'getWebHookConfigAuthBasicAuthPassword',
        'webHookConfigAuthBasicAuthUser' => 'getWebHookConfigAuthBasicAuthUser',
        'webHookConfigAuthType' => 'getWebHookConfigAuthType',
        'webHookConfigBody' => 'getWebHookConfigBody',
        'webHookConfigCanInterrupt' => 'getWebHookConfigCanInterrupt',
        'webHookConfigMethod' => 'getWebHookConfigMethod',
        'webHookConfigResponseIgnore' => 'getWebHookConfigResponseIgnore',
        'webHookConfigUrl' => 'getWebHookConfigUrl'
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
        $this->container['configKey'] = $data['configKey'] ?? null;
        $this->container['createdAt'] = $data['createdAt'] ?? null;
        $this->container['hook'] = $data['hook'] ?? null;
        $this->container['id'] = $data['id'] ?? null;
        $this->container['projectRevisionId'] = $data['projectRevisionId'] ?? null;
        $this->container['updatedAt'] = $data['updatedAt'] ?? null;
        $this->container['webHookConfigAuthApiKeyIn'] = $data['webHookConfigAuthApiKeyIn'] ?? null;
        $this->container['webHookConfigAuthApiKeyName'] = $data['webHookConfigAuthApiKeyName'] ?? null;
        $this->container['webHookConfigAuthApiKeyValue'] = $data['webHookConfigAuthApiKeyValue'] ?? null;
        $this->container['webHookConfigAuthBasicAuthPassword'] = $data['webHookConfigAuthBasicAuthPassword'] ?? null;
        $this->container['webHookConfigAuthBasicAuthUser'] = $data['webHookConfigAuthBasicAuthUser'] ?? null;
        $this->container['webHookConfigAuthType'] = $data['webHookConfigAuthType'] ?? null;
        $this->container['webHookConfigBody'] = $data['webHookConfigBody'] ?? null;
        $this->container['webHookConfigCanInterrupt'] = $data['webHookConfigCanInterrupt'] ?? null;
        $this->container['webHookConfigMethod'] = $data['webHookConfigMethod'] ?? null;
        $this->container['webHookConfigResponseIgnore'] = $data['webHookConfigResponseIgnore'] ?? null;
        $this->container['webHookConfigUrl'] = $data['webHookConfigUrl'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['configKey'] === null) {
            $invalidProperties[] = "'configKey' can't be null";
        }
        if ($this->container['hook'] === null) {
            $invalidProperties[] = "'hook' can't be null";
        }
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
     * Gets configKey
     *
     * @return string
     */
    public function getConfigKey()
    {
        return $this->container['configKey'];
    }

    /**
     * Sets configKey
     *
     * @param string $configKey The Hooks Config Key
     *
     * @return self
     */
    public function setConfigKey($configKey)
    {
        $this->container['configKey'] = $configKey;

        return $this;
    }

    /**
     * Gets createdAt
     *
     * @return \DateTime|null
     */
    public function getCreatedAt()
    {
        return $this->container['createdAt'];
    }

    /**
     * Sets createdAt
     *
     * @param \DateTime|null $createdAt The Project's Revision Creation Date
     *
     * @return self
     */
    public function setCreatedAt($createdAt)
    {
        $this->container['createdAt'] = $createdAt;

        return $this;
    }

    /**
     * Gets hook
     *
     * @return string
     */
    public function getHook()
    {
        return $this->container['hook'];
    }

    /**
     * Sets hook
     *
     * @param string $hook The Hook Type
     *
     * @return self
     */
    public function setHook($hook)
    {
        $this->container['hook'] = $hook;

        return $this;
    }

    /**
     * Gets id
     *
     * @return string|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param string|null $id ID of the entry
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets projectRevisionId
     *
     * @return string|null
     */
    public function getProjectRevisionId()
    {
        return $this->container['projectRevisionId'];
    }

    /**
     * Sets projectRevisionId
     *
     * @param string|null $projectRevisionId The Revision's ID this schema belongs to
     *
     * @return self
     */
    public function setProjectRevisionId($projectRevisionId)
    {
        $this->container['projectRevisionId'] = $projectRevisionId;

        return $this;
    }

    /**
     * Gets updatedAt
     *
     * @return \DateTime|null
     */
    public function getUpdatedAt()
    {
        return $this->container['updatedAt'];
    }

    /**
     * Sets updatedAt
     *
     * @param \DateTime|null $updatedAt Last Time Project's Revision was Updated
     *
     * @return self
     */
    public function setUpdatedAt($updatedAt)
    {
        $this->container['updatedAt'] = $updatedAt;

        return $this;
    }

    /**
     * Gets webHookConfigAuthApiKeyIn
     *
     * @return string|null
     */
    public function getWebHookConfigAuthApiKeyIn()
    {
        return $this->container['webHookConfigAuthApiKeyIn'];
    }

    /**
     * Sets webHookConfigAuthApiKeyIn
     *
     * @param string|null $webHookConfigAuthApiKeyIn Whether to send the API Key in the HTTP Header or as a HTTP Cookie
     *
     * @return self
     */
    public function setWebHookConfigAuthApiKeyIn($webHookConfigAuthApiKeyIn)
    {
        $this->container['webHookConfigAuthApiKeyIn'] = $webHookConfigAuthApiKeyIn;

        return $this;
    }

    /**
     * Gets webHookConfigAuthApiKeyName
     *
     * @return string|null
     */
    public function getWebHookConfigAuthApiKeyName()
    {
        return $this->container['webHookConfigAuthApiKeyName'];
    }

    /**
     * Sets webHookConfigAuthApiKeyName
     *
     * @param string|null $webHookConfigAuthApiKeyName The name of the api key
     *
     * @return self
     */
    public function setWebHookConfigAuthApiKeyName($webHookConfigAuthApiKeyName)
    {
        $this->container['webHookConfigAuthApiKeyName'] = $webHookConfigAuthApiKeyName;

        return $this;
    }

    /**
     * Gets webHookConfigAuthApiKeyValue
     *
     * @return string|null
     */
    public function getWebHookConfigAuthApiKeyValue()
    {
        return $this->container['webHookConfigAuthApiKeyValue'];
    }

    /**
     * Sets webHookConfigAuthApiKeyValue
     *
     * @param string|null $webHookConfigAuthApiKeyValue The value of the api key
     *
     * @return self
     */
    public function setWebHookConfigAuthApiKeyValue($webHookConfigAuthApiKeyValue)
    {
        $this->container['webHookConfigAuthApiKeyValue'] = $webHookConfigAuthApiKeyValue;

        return $this;
    }

    /**
     * Gets webHookConfigAuthBasicAuthPassword
     *
     * @return string|null
     */
    public function getWebHookConfigAuthBasicAuthPassword()
    {
        return $this->container['webHookConfigAuthBasicAuthPassword'];
    }

    /**
     * Sets webHookConfigAuthBasicAuthPassword
     *
     * @param string|null $webHookConfigAuthBasicAuthPassword The password to be sent in the HTTP Basic Auth Header
     *
     * @return self
     */
    public function setWebHookConfigAuthBasicAuthPassword($webHookConfigAuthBasicAuthPassword)
    {
        $this->container['webHookConfigAuthBasicAuthPassword'] = $webHookConfigAuthBasicAuthPassword;

        return $this;
    }

    /**
     * Gets webHookConfigAuthBasicAuthUser
     *
     * @return string|null
     */
    public function getWebHookConfigAuthBasicAuthUser()
    {
        return $this->container['webHookConfigAuthBasicAuthUser'];
    }

    /**
     * Sets webHookConfigAuthBasicAuthUser
     *
     * @param string|null $webHookConfigAuthBasicAuthUser The username to be sent in the HTTP Basic Auth Header
     *
     * @return self
     */
    public function setWebHookConfigAuthBasicAuthUser($webHookConfigAuthBasicAuthUser)
    {
        $this->container['webHookConfigAuthBasicAuthUser'] = $webHookConfigAuthBasicAuthUser;

        return $this;
    }

    /**
     * Gets webHookConfigAuthType
     *
     * @return string|null
     */
    public function getWebHookConfigAuthType()
    {
        return $this->container['webHookConfigAuthType'];
    }

    /**
     * Sets webHookConfigAuthType
     *
     * @param string|null $webHookConfigAuthType HTTP Auth Method to use for the Web-Hook
     *
     * @return self
     */
    public function setWebHookConfigAuthType($webHookConfigAuthType)
    {
        $this->container['webHookConfigAuthType'] = $webHookConfigAuthType;

        return $this;
    }

    /**
     * Gets webHookConfigBody
     *
     * @return string|null
     */
    public function getWebHookConfigBody()
    {
        return $this->container['webHookConfigBody'];
    }

    /**
     * Sets webHookConfigBody
     *
     * @param string|null $webHookConfigBody URI pointing to the JsonNet template used for Web-Hook payload generation. Only used for those HTTP methods, which support HTTP body payloads.
     *
     * @return self
     */
    public function setWebHookConfigBody($webHookConfigBody)
    {
        $this->container['webHookConfigBody'] = $webHookConfigBody;

        return $this;
    }

    /**
     * Gets webHookConfigCanInterrupt
     *
     * @return bool|null
     */
    public function getWebHookConfigCanInterrupt()
    {
        return $this->container['webHookConfigCanInterrupt'];
    }

    /**
     * Sets webHookConfigCanInterrupt
     *
     * @param bool|null $webHookConfigCanInterrupt If enabled allows the web hook to interrupt / abort the self-service flow. It only applies to certain flows (registration/verification/login/settings) and requires a valid response format.
     *
     * @return self
     */
    public function setWebHookConfigCanInterrupt($webHookConfigCanInterrupt)
    {
        $this->container['webHookConfigCanInterrupt'] = $webHookConfigCanInterrupt;

        return $this;
    }

    /**
     * Gets webHookConfigMethod
     *
     * @return string|null
     */
    public function getWebHookConfigMethod()
    {
        return $this->container['webHookConfigMethod'];
    }

    /**
     * Sets webHookConfigMethod
     *
     * @param string|null $webHookConfigMethod The HTTP method to use (GET, POST, etc) for the Web-Hook
     *
     * @return self
     */
    public function setWebHookConfigMethod($webHookConfigMethod)
    {
        $this->container['webHookConfigMethod'] = $webHookConfigMethod;

        return $this;
    }

    /**
     * Gets webHookConfigResponseIgnore
     *
     * @return bool|null
     */
    public function getWebHookConfigResponseIgnore()
    {
        return $this->container['webHookConfigResponseIgnore'];
    }

    /**
     * Sets webHookConfigResponseIgnore
     *
     * @param bool|null $webHookConfigResponseIgnore Whether to ignore the Web Hook response
     *
     * @return self
     */
    public function setWebHookConfigResponseIgnore($webHookConfigResponseIgnore)
    {
        $this->container['webHookConfigResponseIgnore'] = $webHookConfigResponseIgnore;

        return $this;
    }

    /**
     * Gets webHookConfigUrl
     *
     * @return string|null
     */
    public function getWebHookConfigUrl()
    {
        return $this->container['webHookConfigUrl'];
    }

    /**
     * Sets webHookConfigUrl
     *
     * @param string|null $webHookConfigUrl The URL the Web-Hook should call
     *
     * @return self
     */
    public function setWebHookConfigUrl($webHookConfigUrl)
    {
        $this->container['webHookConfigUrl'] = $webHookConfigUrl;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
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
    public function offsetSet($offset, $value)
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
    public function offsetUnset($offset)
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


