// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package schema

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``postgresql_schema`` resource creates and manages [schema
// objects](https://www.postgresql.org/docs/current/static/ddl-schemas.html) within
// a PostgreSQL database.
type Schema struct {
	s *pulumi.ResourceState
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOpt) (*Schema, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["ifNotExists"] = nil
		inputs["name"] = nil
		inputs["owner"] = nil
		inputs["policies"] = nil
	} else {
		inputs["ifNotExists"] = args.IfNotExists
		inputs["name"] = args.Name
		inputs["owner"] = args.Owner
		inputs["policies"] = args.Policies
	}
	s, err := ctx.RegisterResource("postgresql:schema/schema:Schema", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Schema{s: s}, nil
}

// GetSchema gets an existing Schema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SchemaState, opts ...pulumi.ResourceOpt) (*Schema, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["ifNotExists"] = state.IfNotExists
		inputs["name"] = state.Name
		inputs["owner"] = state.Owner
		inputs["policies"] = state.Policies
	}
	s, err := ctx.ReadResource("postgresql:schema/schema:Schema", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Schema{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Schema) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Schema) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// When true, use the existing schema if it exists. (Default: true)
func (r *Schema) IfNotExists() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["ifNotExists"])
}

// The name of the schema. Must be unique in the PostgreSQL
// database instance where it is configured.
func (r *Schema) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ROLE who owns the schema.
func (r *Schema) Owner() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["owner"])
}

// Can be specified multiple times for each policy.  Each
// policy block supports fields documented below.
func (r *Schema) Policies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["policies"])
}

// Input properties used for looking up and filtering Schema resources.
type SchemaState struct {
	// When true, use the existing schema if it exists. (Default: true)
	IfNotExists interface{}
	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	Name interface{}
	// The ROLE who owns the schema.
	Owner interface{}
	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	Policies interface{}
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	// When true, use the existing schema if it exists. (Default: true)
	IfNotExists interface{}
	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	Name interface{}
	// The ROLE who owns the schema.
	Owner interface{}
	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	Policies interface{}
}
