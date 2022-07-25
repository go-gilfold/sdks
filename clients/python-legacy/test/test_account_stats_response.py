# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.3
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import openapi_client
from openapi_client.models.account_stats_response import AccountStatsResponse  # noqa: E501
from openapi_client.rest import ApiException

class TestAccountStatsResponse(unittest.TestCase):
    """AccountStatsResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test AccountStatsResponse
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.account_stats_response.AccountStatsResponse()  # noqa: E501
        if include_optional :
            return AccountStatsResponse(
                total = 1.337, 
                daily = [
                    openapi_client.models.account_stats_response_daily_inner.AccountStatsResponse_daily_inner(
                        date = '', 
                        amount = 1.337, )
                    ], 
                monthly = [
                    openapi_client.models.account_stats_response_daily_inner.AccountStatsResponse_daily_inner(
                        date = '', 
                        amount = 1.337, )
                    ]
            )
        else :
            return AccountStatsResponse(
                total = 1.337,
                daily = [
                    openapi_client.models.account_stats_response_daily_inner.AccountStatsResponse_daily_inner(
                        date = '', 
                        amount = 1.337, )
                    ],
                monthly = [
                    openapi_client.models.account_stats_response_daily_inner.AccountStatsResponse_daily_inner(
                        date = '', 
                        amount = 1.337, )
                    ],
        )

    def testAccountStatsResponse(self):
        """Test AccountStatsResponse"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
