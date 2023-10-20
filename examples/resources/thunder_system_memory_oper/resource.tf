provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

data "thunder_system_memory_oper" "test"{}

output "get_throughput" {
  value = ["${data.thunder_system_memory_oper.test}"]
}