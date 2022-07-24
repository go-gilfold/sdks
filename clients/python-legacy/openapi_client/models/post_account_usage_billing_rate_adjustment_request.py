# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by: https://openapi-generator.tech
"""


try:
    from inspect import getfullargspec
except ImportError:
    from inspect import getargspec as getfullargspec
import pprint
import re  # noqa: F401
import six

from openapi_client.configuration import Configuration


class PostAccountUsageBillingRateAdjustmentRequest(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'billing_cost': 'float'
    }

    attribute_map = {
        'billing_cost': 'billingCost'
    }

    def __init__(self, billing_cost=None, local_vars_configuration=None):  # noqa: E501
        """PostAccountUsageBillingRateAdjustmentRequest - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._billing_cost = None
        self.discriminator = None

        self.billing_cost = billing_cost

    @property
    def billing_cost(self):
        """Gets the billing_cost of this PostAccountUsageBillingRateAdjustmentRequest.  # noqa: E501


        :return: The billing_cost of this PostAccountUsageBillingRateAdjustmentRequest.  # noqa: E501
        :rtype: float
        """
        return self._billing_cost

    @billing_cost.setter
    def billing_cost(self, billing_cost):
        """Sets the billing_cost of this PostAccountUsageBillingRateAdjustmentRequest.


        :param billing_cost: The billing_cost of this PostAccountUsageBillingRateAdjustmentRequest.  # noqa: E501
        :type billing_cost: float
        """
        if self.local_vars_configuration.client_side_validation and billing_cost is None:  # noqa: E501
            raise ValueError("Invalid value for `billing_cost`, must not be `None`")  # noqa: E501

        self._billing_cost = billing_cost

    def to_dict(self, serialize=False):
        """Returns the model properties as a dict"""
        result = {}

        def convert(x):
            if hasattr(x, "to_dict"):
                args = getfullargspec(x.to_dict).args
                if len(args) == 1:
                    return x.to_dict()
                else:
                    return x.to_dict(serialize)
            else:
                return x

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            attr = self.attribute_map.get(attr, attr) if serialize else attr
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: convert(x),
                    value
                ))
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], convert(item[1])),
                    value.items()
                ))
            else:
                result[attr] = convert(value)

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, PostAccountUsageBillingRateAdjustmentRequest):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, PostAccountUsageBillingRateAdjustmentRequest):
            return True

        return self.to_dict() != other.to_dict()
