# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Schema(pulumi.CustomResource):
    if_not_exists: pulumi.Output[bool]
    """
    When true, use the existing schema if it exists. (Default: true)
    """
    name: pulumi.Output[str]
    """
    The name of the schema. Must be unique in the PostgreSQL
    database instance where it is configured.
    """
    owner: pulumi.Output[str]
    """
    The ROLE who owns the schema.
    """
    policies: pulumi.Output[list]
    """
    Can be specified multiple times for each policy.  Each
    policy block supports fields documented below.
    """
    def __init__(__self__, resource_name, opts=None, if_not_exists=None, name=None, owner=None, policies=None, __name__=None, __opts__=None):
        """
        The ``postgresql_schema`` resource creates and manages [schema
        objects](https://www.postgresql.org/docs/current/static/ddl-schemas.html) within
        a PostgreSQL database.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] if_not_exists: When true, use the existing schema if it exists. (Default: true)
        :param pulumi.Input[str] name: The name of the schema. Must be unique in the PostgreSQL
               database instance where it is configured.
        :param pulumi.Input[str] owner: The ROLE who owns the schema.
        :param pulumi.Input[list] policies: Can be specified multiple times for each policy.  Each
               policy block supports fields documented below.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['if_not_exists'] = if_not_exists

        __props__['name'] = name

        __props__['owner'] = owner

        __props__['policies'] = policies

        super(Schema, __self__).__init__(
            'postgresql:schema/schema:Schema',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

