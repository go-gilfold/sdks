# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import openapi_client
from openapi_client.api.businesses_api import BusinessesApi  # noqa: E501
from openapi_client.rest import ApiException


class TestBusinessesApi(unittest.TestCase):
    """BusinessesApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.businesses_api.BusinessesApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_business_business_id_settings_get(self):
        """Test case for business_business_id_settings_get

        Get Settings for a business  # noqa: E501
        """
        pass

    def test_businesses_business_id_get(self):
        """Test case for businesses_business_id_get

        Get Business Details  # noqa: E501
        """
        pass

    def test_businesses_business_id_patch(self):
        """Test case for businesses_business_id_patch

        Update Business Details  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
