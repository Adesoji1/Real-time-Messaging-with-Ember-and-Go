variable "name" {
  description = "This is the human-readable name of the queue. If omitted, Terraform will assign a random name"
  type        = string
  default     = null
}

# ... (Declare other variables similarly)

variable "tags" {
  description = "A mapping of tags to assign to all resources"
  type        = map(string)
  default     = {}
}

# ... (Continue declaring other variables)
variable "name_prefix" {
  description = "A prefix used for naming resources."
  type        = string
}

variable "vpc_id" {
  description = "The VPC ID."
  type        = string
}

variable "vpc_security_group_ids" {
  description = "List of VPC security groups to associate to the cluster in addition to the SG we create in this module"
  type        = list(string)
  default     = []
}

variable "subnet_ids" {
  description = "A list of VPC subnet IDs."
  type        = list(string)
  default     = []
}

variable "username" {
  description = "Username for the master DB user."
  type        = string
}

variable "password" {
  description = "Password for the master DB user."
  type        = string
}

variable "database_name" {
  description = "Name for an automatically created database on cluster creation."
  type        = string
  default     = "main"
}

variable "port" {
  description = "The port on which the DB accepts connections."
  type        = number
  default     = 5439
}

variable "engine" {
  description = "The name of the database engine to be used for this DB cluster."
  type        = string
  default     = "aurora-postgresql"
}

variable "engine_version" {
  description = "The engine version to use."
  type        = string
  default     = ""
}

variable "engine_mode" {
  description = "The database engine mode. Valid values: global, parallelquery, provisioned, serverless."
  type        = string
  default     = "provisioned"
}

variable "enable_http_endpoint" {
  description = "Enable HTTP endpoint (data API). Only valid when engine_mode is set to serverless."
  type        = bool
  default     = false
}

variable "instance_type" {
  description = "The DB instance class to use."
  type        = string
  default     = "db.r5.large"
}

variable "instance_count" {
  description = "Number of DB instances to provision for the cluster."
  type        = number
  default     = 1
}

variable "global_cluster_identifier" {
  description = "The global cluster identifier specified on aws_rds_global_cluster."
  type        = string
  default     = ""
}

variable "replication_source_identifier" {
  description = "ARN of a source DB cluster or DB instance if this DB cluster is to be created as a Read Replica."
  default     = ""
}

variable "publicly_accessible" {
  description = "Bool to control if instances is publicly accessible."
  type        = bool
  default     = false
}

variable "storage_encrypted" {
  description = "Specifies whether the DB cluster is encrypted."
  type        = bool
  default     = true
}

variable "kms_key_arn" {
  description = "The ARN for the KMS encryption key. When specifying kms_key_id, storage_encrypted needs to be set to true."
  type        = string
  default     = ""
}

variable "skip_final_snapshot" {
  description = "Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created."
  type        = bool
  default     = false
}

variable "snapshot_identifier" {
  description = "Specifies whether or not to create this cluster from a snapshot."
  type        = string
  default     = ""
}

variable "final_snapshot_identifier_prefix" {
  description = "The prefix name to use when creating a final snapshot on cluster destroy, appends a random 8 digits to name to ensure it's unique too."
  type        = string
  default     = ""
}

variable "copy_tags_to_snapshot" {
  description = "Copy all Cluster tags to snapshots."
  type        = bool
  default     = false
}

variable "deletion_protection" {
  description = "If the DB instance should have deletion protection enabled"
  type        = bool
  default     = false
}

variable "tags" {
  description = "A map of tags (key-value pairs) passed to resources."
  type        = map(string)
  default     = {}
}

variable "backup_retention_period" {
  description = "The days to retain backups for. Default 7"
  type        = number
  default     = 7
}

variable "backtrack_window" {
  description = "The target backtrack window, in seconds. Only available for aurora engine currently. To disable backtracking, set this value to 0. Defaults to 0. Must be between 0 and 259200 (72 hours)"
  type        = number
  default     = 0
}

variable "preferred_backup_window" {
  description = "The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Default 02:00-03:00 (UTC)"
  type        = string
  default     = "02:00-03:00"
}

variable "preferred_maintenance_window" {
  description = "The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30"
  type        = string
  default     = "wed:04:00-wed:04:30"
}

variable "create_security_group" {
  description = "Whether to create security group for RDS cluster"
  type        = bool
  default     = true
}

variable "security_group_description" {
  description = "The description of the security group. If value is set to empty string it will contain cluster name in the description."
  type        = string
  default     = "Terraformed security group."
}

variable "security_groups" {
  description = "A list of Security Group ID's to allow access to."
  default     = []
}

variable "allowed_cidr_blocks" {
  description = "A list of CIDR blocks which are allowed to access the database"
  type        = list(string)
  default     = []
}

variable "enabled_cloudwatch_logs_exports" {
  description = "List of log types to export to cloudwatch"
  type        = list(string)
  default     = []
}


variable "monitoring_interval" {
  description = "The interval (seconds) between points when Enhanced Monitoring metrics are collected"
  type        = number
  default     = 0
}

variable "performance_insights_enabled" {
  description = "Specifies whether Performance Insights is enabled or not."
  type        = bool
  default     = false
}

variable "predefined_metric_type" {
  description = "The metric type to scale on. Valid values are RDSReaderAverageCPUUtilization and RDSReaderAverageDatabaseConnections."
  default     = "RDSReaderAverageCPUUtilization"
}

variable "db_parameter_group_name" {
  description = "The name of a DB parameter group to use"
  type        = string
  default     = null
}

variable "db_cluster_parameter_group_name" {
  description = "The name of a DB Cluster parameter group to use"
  type        = string
  default     = null
}

variable "db_subnet_group_name" {
  description = "The existing subnet group name to use"
  type        = string
  default     = ""
}

variable "availability_zones" {
  type        = list(string)
  description = "List of AZs to use"
  default     = []
}

variable "scaling_configuration" {
  description = "Map of nested attributes with scaling properties. Only valid when engine_mode is set to `serverless`"
  type        = map(string)
  default     = {}
}


variable "apply_immediately" {
  description = "Determines whether or not any DB modifications are applied immediately, or during the maintenance window"
  type        = bool
  default     = false
}


variable "iam_database_authentication_enabled" {
  description = "Specifies whether IAM Database authentication should be enabled or not. Not all versions and instances are supported. Refer to the AWS documentation to see which versions are supported."
  type        = bool
  default     = false
}

variable "iam_roles" {
  description = "A List of ARNs for the IAM roles to associate to the RDS Cluster."
  type        = list(string)
  default     = []
}


variable "enable_cloudwatch_alarms" {
  description = "Whether to enable CloudWatch alarms - requires `cloudwatch_sns_topic` is specified"
  type        = bool
  default     = false
}

variable "cloudwatch_sns_topic" {
  description = "An SNS topic to publish CloudWatch alarms to"
  type        = string
  default     = ""
}

variable "cloudwatch_max_conns" {
  type        = string
  default     = "500"
  description = "Connection count beyond which to trigger a CloudWatch alarm"
}

variable "cloudwatch_max_cpu" {
  description = "CPU threshold above which to alarm"

  type    = string
  default = "85"
}

variable "cloudwatch_eval_period_connections" {
  description = "Evaluation period for the DB connections alarms"
  type        = number
  default     = 1
}

variable "cloudwatch_eval_period_cpu" {
  description = "Evaluation period for the DB CPU alarms"
  type        = number
  default     = 2
}

variable "replica_scale_enabled" {
  type        = bool
  default     = false
  description = "Whether to enable autoscaling for RDS Aurora (MySQL) read replicas"
}

variable "replica_scale_connections" {
  description = "Average number of connections to trigger autoscaling at. Default value is 70% of db.r4.large's default max_connections"
  type        = number
  default     = 700
}

variable "replica_scale_max" {
  description = "Maximum number of replicas to allow scaling for"
  type        = number
  default     = 0
}

variable "replica_scale_min" {
  description = "Minimum number of replicas to allow scaling for"
  type        = number
  default     = 2
}

variable "replica_scale_cpu" {
  description = "CPU usage to trigger autoscaling at"
  type        = number
  default     = 70
}

variable "replica_scale_in_cooldown" {
  description = "Cooldown in seconds before allowing further scaling operations after a scale in"
  type        = number
  default     = 300
}

variable "replica_scale_out_cooldown" {
  description = "Cooldown in seconds before allowing further scaling operations after a scale out"
  type        = number
  default     = 300
}

variable "use_name_prefix" {
  description = "Flag to determine if a name prefix should be used for the queue"
  type        = bool
  default     = false # or true depending on your preference
}

variable "create_dlq" {
  description = "Flag to determine if a dead letter queue should be created"
  type        = bool
  default     = true # defaulting to true so a DLQ is created by default
}

variable "fifo_queue" {
  description = "Boolean designating a FIFO queue"
  type        = bool
  default     = false # or true, depending on your requirement
}

variable "content_based_deduplication" {
  description = "Enables content-based deduplication for published messages"
  type        = bool
  default     = false # or true, depending on your requirement
}
variable "var.dlq_content_based_deduplication" {
  description = "Enables variable dql-based deduplication for published messages"
  type        = bool
  default     = false # or true, depending on your requirement
}

variable "delay_seconds" {
  description = "The time in seconds that the delivery of all messages in the queue will be delayed"
  type        = number
  default     = 0 # or any appropriate default value
}


variable "max_message_size" {
  description = "The maximum size of messages in the queue"
  type        = number
  default     = 262144  # Set your default or adjust accordingly
}

variable "message_retention_seconds" {
  description = "The number of seconds to retain a message"
  type        = number
  default     = 345600  # Set your default or adjust accordingly
}

variable "receive_wait_time_seconds" {
  description = "The wait time for message processing"
  type        = number
  default     = 0       # Set your default or adjust accordingly
}

variable "kms_master_key_id" {
  description = "The ID of an AWS-managed customer master key for Amazon SQS or a custom CMK"
  type        = string
  default     = ""      # Set your default or adjust accordingly
}

variable "kms_data_key_reuse_period_seconds" {
  description = "The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again"
  type        = number
  default     = 300     # Set your default or adjust accordingly
}

variable "visibility_timeout_seconds" {
  description = "The visibility timeout for the queue in seconds"
  type        = number
  default     = 30      # Set your default or adjust accordingly
}

variable "dlg_name" {
  description = "The name of the dead letter queue"
  type        = string
  default     = "4"      # Set your default or adjust accordingly
}

variable "dlq_content_based_deduplication" {
  description = "Enables content-based deduplication for the Dead Letter Queue (DLQ)"
  type        = bool
  default     = false  # Adjust this default value if necessary
}

variable "dlq_name" {
  description = "Enables naming for the Dead Letter Queue (DLQ)"
  type        = bool
  default     = false  # Adjust this default value if necessary
}


variable "dlg_content_based_deduplication" {
  description = "Enables content-based deduplication for DLQ"
  type        = bool
  default     = false   # Set your default or adjust accordingly
}

variable "dlq_delay_seconds" {
  description = "Delay for DLQ in seconds"
  type        = number
  default     = 0       # Set your default or adjust accordingly
}

variable "dlq_max_message_size" {
  description = "Max message size for DLQ"
  type        = number
  default     = 262144  # Set your default or adjust accordingly
}

variable "dlq_message_retention_seconds" {
  description = "Message retention in seconds for DLQ"
  type        = number
  default     = 345600  # Set your default or adjust accordingly
}

variable "dlq_receive_wait_time_seconds" {
  description = "Wait time for message processing for DLQ"
  type        = number
  default     = 0       # Set your default or adjust accordingly
}

variable "dlq_kms_master_key_id" {
  description = "KMS key ID for DLQ"
  type        = string
  default     = ""      # Set your default or adjust accordingly
}

variable "dlq_kms_data_key_reuse_period_seconds" {
  description = "Data key reuse period in seconds for DLQ"
  type        = number
  default     = 300     # Set your default or adjust accordingly
}

variable "dlq_visibility_timeout_seconds" {  # There seems to be a typo in your error message (dlg_vi _timeout_seconds). I assume it should be dlq_visibility_timeout_seconds
  description = "Visibility timeout in seconds for DLQ"
  type        = number
  default     = 30      # Set your default or adjust accordingly
}

variable "dlq_tags" {
  description = "Tags for the dead letter queue"
  type        = map(string)
  default     = {}      # Set your default or adjust accordingly
}
