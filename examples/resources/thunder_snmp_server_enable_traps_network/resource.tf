

provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_network"  "SnMPServerEnableTrapsNetworks"  { 
       trunk_port_threshold = 1
} 
    

