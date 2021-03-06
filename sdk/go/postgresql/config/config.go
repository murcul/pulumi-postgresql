// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"github.com/pulumi/pulumi/sdk/go/pulumi/config"
)

// Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
func GetConnectTimeout(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "postgresql:connectTimeout")
}

// The name of the database to connect to in order to conenct to (defaults to `postgres`).
func GetDatabase(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:database")
}

// Database username associated to the connected user (for user name maps)
func GetDatabaseUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:databaseUsername")
}

// Specify the expected version of PostgreSQL.
func GetExpectedVersion(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:expectedVersion")
}

// Name of PostgreSQL server address to connect to
func GetHost(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:host")
}

// Maximum number of connections to establish to the database. Zero means unlimited.
func GetMaxConnections(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "postgresql:maxConnections")
}

// Password to be used if the PostgreSQL server demands password authentication
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:password")
}

// The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
func GetPort(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "postgresql:port")
}

func GetSslMode(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:sslMode")
}

// This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
// PostgreSQL server
func GetSslmode(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:sslmode")
}

// PostgreSQL user name to connect as
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "postgresql:username")
}
