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
from openapi_client.models.transaction_stats_response_daily_inner import TransactionStatsResponseDailyInner  # noqa: E501
from openapi_client.rest import ApiException

class TestTransactionStatsResponseDailyInner(unittest.TestCase):
    """TransactionStatsResponseDailyInner unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test TransactionStatsResponseDailyInner
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.transaction_stats_response_daily_inner.TransactionStatsResponseDailyInner()  # noqa: E501
        if include_optional :
            return TransactionStatsResponseDailyInner(
                date = '', 
                incoming_amount = 1.337, 
                outgoing_amount = 1.337, 
                incoming_count = 1.337, 
                outgoing_count = 1.337
            )
        else :
            return TransactionStatsResponseDailyInner(
                date = '',
                incoming_amount = 1.337,
                outgoing_amount = 1.337,
                incoming_count = 1.337,
                outgoing_count = 1.337,
        )

    def testTransactionStatsResponseDailyInner(self):
        """Test TransactionStatsResponseDailyInner"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()