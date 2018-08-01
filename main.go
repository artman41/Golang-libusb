package main

type libusb_version struct {
    major const uint16
    minor const uint16
    micro const uint16
    nano const uint16
    rc const string
    describe const string
}

type libusb_context struct {
    major const uint16
    minor const uint16
    micro const uint16
    nano const uint16
    rc const string
    describe const string
}


func Libusb_init(ctx libusb_context**) int {
    
}

func Libusb_set_debug(ctx libusb_context*. level int) {
    
}

func Libusb_exit(ctx libusb_context*) {
    
}
