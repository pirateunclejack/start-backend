package mq

import "start-backend/helper"

var kafka_address   string
var kafka_topic     string
var kafka_partition int
var rabbit_address  string
var rabbit_port     string
var rabbit_username string
var rabbit_password string
var rabbit_queue    string

func init() {
    vip := helper.GetConfig()
    kafka_address   = vip.GetString("KAFKA_ADDRESS")
    kafka_topic     = vip.GetString("KAFKA_TOPIC")
    kafka_partition = vip.GetInt("KAFKA_PARTITION")
    rabbit_address  = vip.GetString("RABBIT_ADDRESS")
    rabbit_port     = vip.GetString("RABBIT_PORT")
    rabbit_username = vip.GetString("RABBIT_USERNAME")
    rabbit_password = vip.GetString("RABBIT_PASSWORD")
    rabbit_queue    = vip.GetString("RABBIT_QUEUE")
}
