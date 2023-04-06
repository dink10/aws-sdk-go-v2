// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modify the settings for an Amazon Aurora DB cluster or a Multi-AZ DB cluster.
// You can change one or more settings by specifying these parameters and the new
// values in the request. For more information on Amazon Aurora DB clusters, see
// What is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see  Multi-AZ DB cluster deployments
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) ModifyDBCluster(ctx context.Context, params *ModifyDBClusterInput, optFns ...func(*Options)) (*ModifyDBClusterOutput, error) {
	if params == nil {
		params = &ModifyDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBCluster", params, optFns, c.addOperationModifyDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBClusterInput struct {

	// The DB cluster identifier for the cluster being modified. This parameter isn't
	// case-sensitive. Constraints: This identifier must match the identifier of an
	// existing DB cluster. Valid for: Aurora DB clusters and Multi-AZ DB clusters
	//
	// This member is required.
	DBClusterIdentifier *string

	// The amount of storage in gibibytes (GiB) to allocate to each DB instance in the
	// Multi-AZ DB cluster. Type: Integer Valid for: Multi-AZ DB clusters only
	AllocatedStorage *int32

	// A value that indicates whether major version upgrades are allowed. Constraints:
	// You must allow major version upgrades when specifying a value for the
	// EngineVersion parameter that is a different major version than the DB cluster's
	// current version. Valid for: Aurora DB clusters only
	AllowMajorVersionUpgrade bool

	// A value that indicates whether the modifications in this request and any pending
	// modifications are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the DB cluster. If this parameter is
	// disabled, changes to the DB cluster are applied during the next maintenance
	// window. Most modifications can be applied immediately or during the next
	// scheduled maintenance window. Some modifications, such as turning on deletion
	// protection and changing the master password, are applied immediately—regardless
	// of when you choose to apply them. By default, this parameter is disabled. Valid
	// for: Aurora DB clusters and Multi-AZ DB clusters
	ApplyImmediately bool

	// A value that indicates whether minor engine upgrades are applied automatically
	// to the DB cluster during the maintenance window. By default, minor engine
	// upgrades are applied automatically. Valid for: Multi-AZ DB clusters only
	AutoMinorVersionUpgrade *bool

	// The target backtrack window, in seconds. To disable backtracking, set this value
	// to 0. Default: 0 Constraints:
	//
	// * If specified, this value must be set to a
	// number from 0 to 259,200 (72 hours).
	//
	// Valid for: Aurora MySQL DB clusters only
	BacktrackWindow *int64

	// The number of days for which automated backups are retained. Specify a minimum
	// value of 1. Default: 1 Constraints:
	//
	// * Must be a value from 1 to 35
	//
	// Valid for:
	// Aurora DB clusters and Multi-AZ DB clusters
	BackupRetentionPeriod *int32

	// The configuration setting for the log types to be enabled for export to
	// CloudWatch Logs for a specific DB cluster. The values in the list depend on the
	// DB engine being used. RDS for MySQL Possible values are error, general, and
	// slowquery. RDS for PostgreSQL Possible values are postgresql and upgrade. Aurora
	// MySQL Possible values are audit, error, general, and slowquery. Aurora
	// PostgreSQL Possible value is postgresql. For more information about exporting
	// CloudWatch Logs for Amazon RDS, see  Publishing Database Logs to Amazon
	// CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide. For more information about exporting CloudWatch
	// Logs for Amazon Aurora, see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	CloudwatchLogsExportConfiguration *types.CloudwatchLogsExportConfiguration

	// A value that indicates whether to copy all tags from the DB cluster to snapshots
	// of the DB cluster. The default is not to copy them. Valid for: Aurora DB
	// clusters and Multi-AZ DB clusters
	CopyTagsToSnapshot *bool

	// The compute and memory capacity of each DB instance in the Multi-AZ DB cluster,
	// for example db.m6gd.xlarge. Not all DB instance classes are available in all
	// Amazon Web Services Regions, or for all database engines. For the full list of
	// DB instance classes and availability for your engine, see  DB Instance Class
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. Valid for: Multi-AZ DB clusters only
	DBClusterInstanceClass *string

	// The name of the DB cluster parameter group to use for the DB cluster. Valid for:
	// Aurora DB clusters and Multi-AZ DB clusters
	DBClusterParameterGroupName *string

	// The name of the DB parameter group to apply to all instances of the DB cluster.
	// When you apply a parameter group using the DBInstanceParameterGroupName
	// parameter, the DB cluster isn't rebooted automatically. Also, parameter changes
	// are applied immediately rather than during the next maintenance window. Default:
	// The existing name setting Constraints:
	//
	// * The DB parameter group must be in the
	// same DB parameter group family as this DB cluster.
	//
	// * The
	// DBInstanceParameterGroupName parameter is valid in combination with the
	// AllowMajorVersionUpgrade parameter for a major version upgrade only.
	//
	// Valid for:
	// Aurora DB clusters only
	DBInstanceParameterGroupName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection isn't enabled. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	DeletionProtection *bool

	// The Active Directory directory ID to move the DB cluster to. Specify none to
	// remove the cluster from its current domain. The domain must be created prior to
	// this operation. For more information, see Kerberos Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service. Valid for: Aurora DB clusters only
	DomainIAMRoleName *string

	// A value that indicates whether to enable this DB cluster to forward write
	// operations to the primary cluster of an Aurora global database (GlobalCluster).
	// By default, write operations are not allowed on Aurora DB clusters that are
	// secondary clusters in an Aurora global database. You can set this value only on
	// Aurora DB clusters that are members of an Aurora global database. With this
	// parameter enabled, a secondary cluster can forward writes to the current primary
	// cluster and the resulting changes are replicated back to this cluster. For the
	// primary DB cluster of an Aurora global database, this value is used immediately
	// if the primary is demoted by the FailoverGlobalCluster API operation, but it
	// does nothing until then. Valid for: Aurora DB clusters only
	EnableGlobalWriteForwarding *bool

	// A value that indicates whether to enable the HTTP endpoint for an Aurora
	// Serverless v1 DB cluster. By default, the HTTP endpoint is disabled. When
	// enabled, the HTTP endpoint provides a connectionless web service API for running
	// SQL queries on the Aurora Serverless v1 DB cluster. You can also query your
	// database from inside the RDS console with the query editor. For more
	// information, see Using the Data API for Aurora Serverless v1
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in
	// the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	EnableHttpEndpoint *bool

	// A value that indicates whether to enable mapping of Amazon Web Services Identity
	// and Access Management (IAM) accounts to database accounts. By default, mapping
	// isn't enabled. For more information, see  IAM Database Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	EnableIAMDatabaseAuthentication *bool

	// A value that indicates whether to turn on Performance Insights for the DB
	// cluster. For more information, see  Using Amazon Performance Insights
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon RDS User Guide. Valid for: Multi-AZ DB clusters only
	EnablePerformanceInsights *bool

	// The version number of the database engine to which you want to upgrade. Changing
	// this parameter results in an outage. The change is applied during the next
	// maintenance window unless ApplyImmediately is enabled. If the cluster that
	// you're modifying has one or more read replicas, all replicas must be running an
	// engine version that's the same or later than the version you specify. To list
	// all of the available engine versions for Aurora MySQL version 2 (5.7-compatible)
	// and version 3 (MySQL 8.0-compatible), use the following command: aws rds
	// describe-db-engine-versions --engine aurora-mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for MySQL 5.6-compatible Aurora, use the following command: aws rds
	// describe-db-engine-versions --engine aurora --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for Aurora PostgreSQL, use the following command: aws rds
	// describe-db-engine-versions --engine aurora-postgresql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for RDS for MySQL, use the following command: aws rds
	// describe-db-engine-versions --engine mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for RDS for PostgreSQL, use the following command: aws rds
	// describe-db-engine-versions --engine postgres --query
	// "DBEngineVersions[].EngineVersion" Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	EngineVersion *string

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for each DB instance in the Multi-AZ DB cluster. For
	// information about valid IOPS values, see Amazon RDS Provisioned IOPS storage
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide. Constraints: Must be a multiple between .5 and 50
	// of the storage amount for the DB cluster. Valid for: Multi-AZ DB clusters only
	Iops *int32

	// A value that indicates whether to manage the master user password with Amazon
	// Web Services Secrets Manager. If the DB cluster doesn't manage the master user
	// password with Amazon Web Services Secrets Manager, you can turn on this
	// management. In this case, you can't specify MasterUserPassword. If the DB
	// cluster already manages the master user password with Amazon Web Services
	// Secrets Manager, and you specify that the master user password is not managed
	// with Amazon Web Services Secrets Manager, then you must specify
	// MasterUserPassword. In this case, RDS deletes the secret and uses the new
	// password for the master user specified by MasterUserPassword. For more
	// information, see Password management with Amazon Web Services Secrets Manager
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html)
	// in the Amazon RDS User Guide and Password management with Amazon Web Services
	// Secrets Manager
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	ManageMasterUserPassword *bool

	// The new password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@". Constraints:
	//
	// * Must contain
	// from 8 to 41 characters.
	//
	// * Can't be specified if ManageMasterUserPassword is
	// turned on.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	MasterUserPassword *string

	// The Amazon Web Services KMS key identifier to encrypt a secret that is
	// automatically generated and managed in Amazon Web Services Secrets Manager. This
	// setting is valid only if both of the following conditions are met:
	//
	// * The DB
	// cluster doesn't manage the master user password in Amazon Web Services Secrets
	// Manager. If the DB cluster already manages the master user password in Amazon
	// Web Services Secrets Manager, you can't change the KMS key that is used to
	// encrypt the secret.
	//
	// * You are turning on ManageMasterUserPassword to manage the
	// master user password in Amazon Web Services Secrets Manager. If you are turning
	// on ManageMasterUserPassword and don't specify MasterUserSecretKmsKeyId, then the
	// aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in a
	// different Amazon Web Services account, then you can't use the aws/secretsmanager
	// KMS key to encrypt the secret, and you must use a customer managed KMS key.
	//
	// The
	// Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or
	// alias name for the KMS key. To use a KMS key in a different Amazon Web Services
	// account, specify the key ARN or alias ARN. There is a default KMS key for your
	// Amazon Web Services account. Your Amazon Web Services account has a different
	// default KMS key for each Amazon Web Services Region. Valid for: Aurora DB
	// clusters and Multi-AZ DB clusters
	MasterUserSecretKmsKeyId *string

	// The interval, in seconds, between points when Enhanced Monitoring metrics are
	// collected for the DB cluster. To turn off collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0. If MonitoringRoleArn is specified, also
	// set MonitoringInterval to a value other than 0. Valid Values: 0, 1, 5, 10, 15,
	// 30, 60 Valid for: Multi-AZ DB clusters only
	MonitoringInterval *int32

	// The Amazon Resource Name (ARN) for the IAM role that permits RDS to send
	// Enhanced Monitoring metrics to Amazon CloudWatch Logs. An example is
	// arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring
	// role, see To create an IAM role for Amazon RDS Enhanced Monitoring
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html#USER_Monitoring.OS.IAMRole)
	// in the Amazon RDS User Guide. If MonitoringInterval is set to a value other than
	// 0, supply a MonitoringRoleArn value. Valid for: Multi-AZ DB clusters only
	MonitoringRoleArn *string

	// The network type of the DB cluster. Valid values:
	//
	// * IPV4
	//
	// * DUAL
	//
	// The network
	// type is determined by the DBSubnetGroup specified for the DB cluster. A
	// DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6
	// protocols (DUAL). For more information, see  Working with a DB instance in a VPC
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	NetworkType *string

	// The new DB cluster identifier for the DB cluster when renaming a DB cluster.
	// This value is stored as a lowercase string. Constraints:
	//
	// * Must contain from 1
	// to 63 letters, numbers, or hyphens
	//
	// * The first character must be a letter
	//
	// *
	// Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-cluster2
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	NewDBClusterIdentifier *string

	// A value that indicates that the DB cluster should be associated with the
	// specified option group. DB clusters are associated with a default option group
	// that can't be modified.
	OptionGroupName *string

	// The Amazon Web Services KMS key identifier for encryption of Performance
	// Insights data. The Amazon Web Services KMS key identifier is the key ARN, key
	// ID, alias ARN, or alias name for the KMS key. If you don't specify a value for
	// PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is
	// a default KMS key for your Amazon Web Services account. Your Amazon Web Services
	// account has a different default KMS key for each Amazon Web Services Region.
	// Valid for: Multi-AZ DB clusters only
	PerformanceInsightsKMSKeyId *string

	// The number of days to retain Performance Insights data. The default is 7 days.
	// The following values are valid:
	//
	// * 7
	//
	// * month * 31, where month is a number of
	// months from 1-23
	//
	// * 731
	//
	// For example, the following values are valid:
	//
	// * 93 (3
	// months * 31)
	//
	// * 341 (11 months * 31)
	//
	// * 589 (19 months * 31)
	//
	// * 731
	//
	// If you
	// specify a retention period such as 94, which isn't a valid value, RDS issues an
	// error. Valid for: Multi-AZ DB clusters only
	PerformanceInsightsRetentionPeriod *int32

	// The port number on which the DB cluster accepts connections. Constraints: Value
	// must be 1150-65535 Default: The same port as the original DB cluster. Valid for:
	// Aurora DB clusters only
	Port *int32

	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each Amazon
	// Web Services Region. To view the time blocks available, see  Backup window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow)
	// in the Amazon Aurora User Guide. Constraints:
	//
	// * Must be in the format
	// hh24:mi-hh24:mi.
	//
	// * Must be in Universal Coordinated Time (UTC).
	//
	// * Must not
	// conflict with the preferred maintenance window.
	//
	// * Must be at least 30
	// minutes.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each Amazon
	// Web Services Region, occurring on a random day of the week. To see the time
	// blocks available, see  Adjusting the Preferred DB Cluster Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	// Constraints: Minimum 30-minute window. Valid for: Aurora DB clusters and
	// Multi-AZ DB clusters
	PreferredMaintenanceWindow *string

	// A value that indicates whether to rotate the secret managed by Amazon Web
	// Services Secrets Manager for the master user password. This setting is valid
	// only if the master user password is managed by RDS in Amazon Web Services
	// Secrets Manager for the DB cluster. The secret value contains the updated
	// password. For more information, see Password management with Amazon Web Services
	// Secrets Manager
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html)
	// in the Amazon RDS User Guide and Password management with Amazon Web Services
	// Secrets Manager
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html)
	// in the Amazon Aurora User Guide. Constraints:
	//
	// * You must apply the change
	// immediately when rotating the master user password.
	//
	// Valid for: Aurora DB
	// clusters and Multi-AZ DB clusters
	RotateMasterUserPassword *bool

	// The scaling properties of the DB cluster. You can only modify scaling properties
	// for DB clusters in serverless DB engine mode. Valid for: Aurora DB clusters only
	ScalingConfiguration *types.ScalingConfiguration

	// Contains the scaling configuration of an Aurora Serverless v2 DB cluster. For
	// more information, see Using Amazon Aurora Serverless v2
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html)
	// in the Amazon Aurora User Guide.
	ServerlessV2ScalingConfiguration *types.ServerlessV2ScalingConfiguration

	// Specifies the storage type to be associated with the DB cluster. Valid values:
	// io1 When specified, a value for the Iops parameter is required. Default: io1
	// Valid for: Multi-AZ DB clusters only
	StorageType *string

	// A list of VPC security groups that the DB cluster will belong to. Valid for:
	// Aurora DB clusters and Multi-AZ DB clusters
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type ModifyDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster. For
	// an Amazon Aurora DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster, DeleteDBCluster, DescribeDBClusters,
	// FailoverDBCluster, ModifyDBCluster, PromoteReadReplicaDBCluster,
	// RestoreDBClusterFromS3, RestoreDBClusterFromSnapshot,
	// RestoreDBClusterToPointInTime, StartDBCluster, and StopDBCluster. For a Multi-AZ
	// DB cluster, this data type is used as a response element in the operations
	// CreateDBCluster, DeleteDBCluster, DescribeDBClusters, FailoverDBCluster,
	// ModifyDBCluster, RebootDBCluster, RestoreDBClusterFromSnapshot, and
	// RestoreDBClusterToPointInTime. For more information on Amazon Aurora DB
	// clusters, see  What is Amazon Aurora?
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
	// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
	// see  Multi-AZ deployments with two readable standby DB instances
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
	// in the Amazon RDS User Guide.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBCluster(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBCluster",
	}
}
