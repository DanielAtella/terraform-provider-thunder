provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_rate_limit_log" "test_thunder_slb_rate_limit_log" {
  sampling_enable {
    counters1 = "all"
  }
}

