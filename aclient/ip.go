package aclient

import (
    "net"
    "strings"
)

func GetLocalIp() (ip string, err error) {
    conn, err := net.Dial("udp", "baidu.com:80")
    if err != nil {
        return "", err
    }
    defer conn.Close()
    ip = strings.Split(conn.LocalAddr().String(), ":")[0]
    return ip, nil
}

func GetLocalIps() (ips []string, err error) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return
    }
    for _, addr := range addrs {
        // 这个网络地址是IP地址: ipv4, ipv6
        //if ipNet, isIpNet := addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
        if ipNet, isIpNet := addr.(*net.IPNet); isIpNet {
            // 跳过IPV6
            if ipNet.IP.To4() != nil {
                ipv4 := ipNet.IP.String()
                ips = append(ips, ipv4)
            }
        }
    }
    return
}
