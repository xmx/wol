package wol

import (
	"bytes"
	"net"
)

// Wake 向指定的 mac 地址发送唤醒魔包
func Wake(mac net.HardwareAddr, nums ...int) error {
	num := 3
	if len(nums) > 0 && nums[0] > 0 {
		num = nums[0]
	}

	buf := bytes.NewBuffer([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	for i := 0; i < 16; i++ {
		buf.Write(mac)
	}
	packet := buf.Bytes()

	ifs, err := net.InterfaceAddrs()
	if err != nil {
		return err
	}

	for _, addr := range ifs {
		if inet, ok := addr.(*net.IPNet); ok {
			_ = send(inet.IP, net.IPv4bcast, packet, num)
		}
	}

	return nil
}

func send(src, dest net.IP, packet []byte, n int) error {
	conn, err := net.DialUDP("udp", &net.UDPAddr{IP: src}, &net.UDPAddr{IP: dest})
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer conn.Close()

	for i := 0; i < n; i++ {
		_, _ = conn.Write(packet)
	}

	return nil
}
