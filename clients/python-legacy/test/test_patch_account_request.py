# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import openapi_client
from openapi_client.models.patch_account_request import PatchAccountRequest  # noqa: E501
from openapi_client.rest import ApiException

class TestPatchAccountRequest(unittest.TestCase):
    """PatchAccountRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test PatchAccountRequest
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.patch_account_request.PatchAccountRequest()  # noqa: E501
        if include_optional :
            return PatchAccountRequest(
                account_name = ''
            )
        else :
            return PatchAccountRequest(
        )

    def testPatchAccountRequest(self):
        """Test PatchAccountRequest"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
