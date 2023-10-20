provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_passthrough" "test_thunder_slb_passthrough" {
  sampling_enable {
    counters1 = "all"
  }
}

