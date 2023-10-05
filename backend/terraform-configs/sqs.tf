provider "aws" {
  region = "us-west-1" # Change this to your desired AWS region
}

resource "aws_sqs_queue" "main_queue" {
  name_prefix                = var.use_name_prefix ? var.name : ""
  name                       = var.name
  fifo_queue                 = var.fifo_queue
  content_based_deduplication = var.content_based_deduplication
  delay_seconds              = var.delay_seconds
  max_message_size           = var.max_message_size
  message_retention_seconds  = var.message_retention_seconds
  receive_wait_time_seconds  = var.receive_wait_time_seconds
  kms_master_key_id          = var.kms_master_key_id
  kms_data_key_reuse_period_seconds = var.kms_data_key_reuse_period_seconds
  visibility_timeout_seconds = var.visibility_timeout_seconds
  tags                       = var.tags
  redrive_policy             = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.dead_letter_queue.arn
    maxReceiveCount     = 5 # Adjust this value as needed
  })
}

resource "aws_sqs_queue" "dead_letter_queue" {
  count                      = var.create_dlq ? 1 : 0
  name_prefix                = var.use_name_prefix ? "${var.dlq_name}-dlq-" : ""
  name                       = var.dlq_name
  fifo_queue                 = var.fifo_queue
  content_based_deduplication = var.dlq_content_based_deduplication
  delay_seconds              = var.dlq_delay_seconds
  max_message_size           = var.dlq_max_message_size
  message_retention_seconds  = var.dlq_message_retention_seconds
  receive_wait_time_seconds  = var.dlq_receive_wait_time_seconds
  kms_master_key_id          = var.dlq_kms_master_key_id
  kms_data_key_reuse_period_seconds = var.dlq_kms_data_key_reuse_period_seconds
  visibility_timeout_seconds = var.dlq_visibility_timeout_seconds
  tags                       = var.dlq_tags
}

output "queue_arn" {
  value = aws_sqs_queue.main_queue.arn
}

output "queue_id" {
  value = aws_sqs_queue.main_queue.id
}

output "queue_name" {
  value = aws_sqs_queue.main_queue.name
}

output "queue_url" {
  value = aws_sqs_queue.main_queue.url
}

output "dead_letter_queue_arn" {
  value = aws_sqs_queue.dead_letter_queue[0].arn
}

output "dead_letter_queue_id" {
  value = aws_sqs_queue.dead_letter_queue[0].id
}

output "dead_letter_queue_name" {
  value = aws_sqs_queue.dead_letter_queue[0].name
}

output "dead_letter_queue_url" {
  value = aws_sqs_queue.dead_letter_queue[0].url
}
