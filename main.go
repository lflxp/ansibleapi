package main

import (
	//"github.com/lflxp/ansibleapi/api"
)
import "strings"

func main() {
	//println(api.Apis.Version())
//	tmp := api.Apis.Setup("all")
//	for _,y := range tmp {
//		println(y.Host)
//		println(y.Status)
//		println(y.Origin)
//	}

	a := `10.6.200.121 | SUCCESS => {
    "ansible_facts": {
        "ansible_SHM.TWM": {
            "active": true,
            "device": "SHM.TWM",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "on [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "off [fixed]",
                "netns_local": "on [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "off [fixed]",
                "rx_vlan_filter": "off [fixed]",
                "rx_vlan_offload": "off [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "on [fixed]",
                "tx_checksum_ipv4": "off [fixed]",
                "tx_checksum_ipv6": "off [fixed]",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "off [fixed]",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "on [fixed]",
                "tx_scatter_gather": "on [fixed]",
                "tx_scatter_gather_fraglist": "on [fixed]",
                "tx_tcp6_segmentation": "on",
                "tx_tcp_ecn_segmentation": "on",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "off [fixed]",
                "udp_fragmentation_offload": "on [fixed]",
                "vlan_challenged": "off [fixed]"
            },
            "ipv4": {
                "address": "1.2.1.2",
                "broadcast": "",
                "netmask": "255.255.255.255",
                "network": "1.2.1.2"
            },
            "macaddress": "0a:06:c8:79",
            "mtu": 1476,
            "promisc": false,
            "type": "unknown"
        },
        "ansible_all_ipv4_addresses": [
            "10.6.200.121",
            "1.2.1.2"
        ],
        "ansible_all_ipv6_addresses": [
            "fe80::a00:27ff:fe99:75cc"
        ],
        "ansible_apparmor": {
            "status": "disabled"
        },
        "ansible_architecture": "x86_64",
        "ansible_bios_date": "12/01/2006",
        "ansible_bios_version": "VirtualBox",
        "ansible_cmdline": {
            "KEYBOARDTYPE": "pc",
            "KEYTABLE": "us",
            "LANG": "en_US.UTF-8",
            "SYSFONT": "latarcyrheb-sun16",
            "crashkernel": "129M@0M",
            "quiet": true,
            "rd_NO_DM": true,
            "rd_NO_LUKS": true,
            "rd_NO_LVM": true,
            "rd_NO_MD": true,
            "rhgb": true,
            "ro": true,
            "root": "UUID=e9c0395d-a06b-49a3-bede-6e2a0b605c19"
        },
        "ansible_date_time": {
            "date": "2017-06-22",
            "day": "22",
            "epoch": "1498097890",
            "hour": "10",
            "iso8601": "2017-06-22T02:18:10Z",
            "iso8601_basic": "20170622T101810458734",
            "iso8601_basic_short": "20170622T101810",
            "iso8601_micro": "2017-06-22T02:18:10.458841Z",
            "minute": "18",
            "month": "06",
            "second": "10",
            "time": "10:18:10",
            "tz": "CST",
            "tz_offset": "+0800",
            "weekday": "Thursday",
            "weekday_number": "4",
            "weeknumber": "25",
            "year": "2017"
        },
        "ansible_default_ipv4": {
            "address": "10.6.200.121",
            "alias": "eth0",
            "broadcast": "10.6.200.255",
            "gateway": "10.6.200.254",
            "interface": "eth0",
            "macaddress": "08:00:27:99:75:cc",
            "mtu": 1500,
            "netmask": "255.255.255.0",
            "network": "10.6.200.0",
            "type": "ether"
        },
        "ansible_default_ipv6": {},
        "ansible_devices": {
            "sda": {
                "holders": [],
                "host": "SATA controller: Intel Corporation 82801HM/HEM (ICH8M/ICH8M-E) SATA Controller [AHCI mode] (rev 02)",
                "model": "VBOX HARDDISK",
                "partitions": {
                    "sda1": {
                        "holders": [],
                        "sectors": "16777216",
                        "sectorsize": 512,
                        "size": "8.00 GB",
                        "start": "2048",
                        "uuid": "e9c0395d-a06b-49a3-bede-6e2a0b605c19"
                    },
                    "sda2": {
                        "holders": [],
                        "sectors": "4194304",
                        "sectorsize": 512,
                        "size": "2.00 GB",
                        "start": "16779264",
                        "uuid": "606e984a-1730-4a44-bcc2-81c31859ac41"
                    },
                    "sda3": {
                        "holders": [],
                        "sectors": "2097152",
                        "sectorsize": 512,
                        "size": "1.00 GB",
                        "start": "20973568",
                        "uuid": "7387c4cf-e2fb-4510-a8f3-71c51bde88d6"
                    }
                },
                "removable": "0",
                "rotational": "1",
                "sas_address": null,
                "sas_device_handle": null,
                "scheduler_mode": "cfq",
                "sectors": "39962880",
                "sectorsize": "512",
                "size": "19.06 GB",
                "support_discard": "0",
                "vendor": "ATA"
            },
            "sr0": {
                "holders": [],
                "host": "IDE interface: Intel Corporation 82371AB/EB/MB PIIX4 IDE (rev 01)",
                "model": "CD-ROM",
                "partitions": {},
                "removable": "1",
                "rotational": "1",
                "sas_address": null,
                "sas_device_handle": null,
                "scheduler_mode": "cfq",
                "sectors": "2097151",
                "sectorsize": "512",
                "size": "1024.00 MB",
                "support_discard": "0",
                "vendor": "VBOX"
            }
        },
        "ansible_distribution": "CentOS",
        "ansible_distribution_major_version": "6",
        "ansible_distribution_release": "Final",
        "ansible_distribution_version": "6.7",
        "ansible_dns": {
            "nameservers": [
                "10.23.70.97"
            ]
        },
        "ansible_domain": "localdomain",
        "ansible_effective_group_id": 0,
        "ansible_effective_user_id": 0,
        "ansible_env": {
            "G_BROKEN_FILENAMES": "1",
            "HOME": "/root",
            "LANG": "POSIX",
            "LC_CTYPE": "zh_CN.UTF-8",
            "LESSOPEN": "||/usr/bin/lesspipe.sh %s",
            "LOGNAME": "root",
            "MAIL": "/var/mail/root",
            "PATH": "/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin",
            "PWD": "/root",
            "SHELL": "/bin/bash",
            "SHLVL": "2",
            "SSH_CLIENT": "10.6.200.131 39207 22",
            "SSH_CONNECTION": "10.6.200.131 39207 10.6.200.121 22",
            "SSH_TTY": "/dev/pts/1",
            "TERM": "xterm-256color",
            "USER": "root",
            "_": "/usr/bin/python"
        },
        "ansible_eth0": {
            "active": true,
            "device": "eth0",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "off [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "off [fixed]",
                "netns_local": "off [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "on",
                "rx_vlan_filter": "on [fixed]",
                "rx_vlan_offload": "on [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "on",
                "tx_checksum_ipv4": "off",
                "tx_checksum_ipv6": "off",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "off",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "off [fixed]",
                "tx_scatter_gather": "on",
                "tx_scatter_gather_fraglist": "off [fixed]",
                "tx_tcp6_segmentation": "off",
                "tx_tcp_ecn_segmentation": "off",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "on [fixed]",
                "udp_fragmentation_offload": "off [fixed]",
                "vlan_challenged": "off [fixed]"
            },
            "ipv4": {
                "address": "10.6.200.121",
                "broadcast": "10.6.200.255",
                "netmask": "255.255.255.0",
                "network": "10.6.200.0"
            },
            "ipv6": [
                {
                    "address": "fe80::a00:27ff:fe99:75cc",
                    "prefix": "64",
                    "scope": "link"
                }
            ],
            "macaddress": "08:00:27:99:75:cc",
            "module": "e1000",
            "mtu": 1500,
            "pciid": "0000:00:03.0",
            "promisc": false,
            "speed": 1000,
            "type": "ether"
        },
        "ansible_fips": false,
        "ansible_form_factor": "Other",
        "ansible_fqdn": "localhost.localdomain",
        "ansible_gather_subset": [
            "hardware",
            "network",
            "virtual"
        ],
        "ansible_gre0": {
            "active": false,
            "device": "gre0",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "on [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "off [fixed]",
                "netns_local": "on [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "off [fixed]",
                "rx_vlan_filter": "off [fixed]",
                "rx_vlan_offload": "off [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "on [fixed]",
                "tx_checksum_ipv4": "off [fixed]",
                "tx_checksum_ipv6": "off [fixed]",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "off [fixed]",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "on [fixed]",
                "tx_scatter_gather": "on [fixed]",
                "tx_scatter_gather_fraglist": "on [fixed]",
                "tx_tcp6_segmentation": "on",
                "tx_tcp_ecn_segmentation": "on",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "off [fixed]",
                "udp_fragmentation_offload": "on [fixed]",
                "vlan_challenged": "off [fixed]"
            },
            "macaddress": "0a:06:c8:79",
            "mtu": 1476,
            "promisc": false,
            "type": "unknown"
        },
        "ansible_gretap0": {
            "active": false,
            "device": "gretap0",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "on [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "off [fixed]",
                "netns_local": "on [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "off [fixed]",
                "rx_vlan_filter": "off [fixed]",
                "rx_vlan_offload": "off [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "on [fixed]",
                "tx_checksum_ipv4": "off [fixed]",
                "tx_checksum_ipv6": "off [fixed]",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "off [fixed]",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "on [fixed]",
                "tx_scatter_gather": "on [fixed]",
                "tx_scatter_gather_fraglist": "on [fixed]",
                "tx_tcp6_segmentation": "on",
                "tx_tcp_ecn_segmentation": "on",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "off [fixed]",
                "udp_fragmentation_offload": "on [fixed]",
                "vlan_challenged": "off [fixed]"
            },
            "mtu": 1476,
            "promisc": false,
            "type": "ether"
        },
        "ansible_hostname": "localhost",
        "ansible_interfaces": [
            "lo",
            "gretap0",
            "eth0",
            "gre0",
            "SHM.TWM"
        ],
        "ansible_kernel": "2.6.32-573.el6.x86_64",
        "ansible_lo": {
            "active": true,
            "device": "lo",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "on [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "on [fixed]",
                "netns_local": "on [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "on [fixed]",
                "rx_vlan_filter": "off [fixed]",
                "rx_vlan_offload": "off [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "off [fixed]",
                "tx_checksum_ipv4": "off [fixed]",
                "tx_checksum_ipv6": "off [fixed]",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "on [fixed]",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "on [fixed]",
                "tx_scatter_gather": "on [fixed]",
                "tx_scatter_gather_fraglist": "on [fixed]",
                "tx_tcp6_segmentation": "on",
                "tx_tcp_ecn_segmentation": "on",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "off [fixed]",
                "udp_fragmentation_offload": "on",
                "vlan_challenged": "on [fixed]"
            },
            "ipv4": {
                "address": "127.0.0.1",
                "broadcast": "host",
                "netmask": "255.0.0.0",
                "network": "127.0.0.0"
            },
            "ipv6": [
                {
                    "address": "::1",
                    "prefix": "128",
                    "scope": "host"
                }
            ],
            "mtu": 65536,
            "promisc": false,
            "type": "loopback"
        },
        "ansible_lvm": {
            "lvs": {},
            "vgs": {}
        },
        "ansible_machine": "x86_64",
        "ansible_machine_id": "a30309c6d63673caa26316d40000000a",
        "ansible_memfree_mb": 1058,
        "ansible_memory_mb": {
            "nocache": {
                "free": 1750,
                "used": 127
            },
            "real": {
                "free": 1058,
                "total": 1877,
                "used": 819
            },
            "swap": {
                "cached": 0,
                "free": 1023,
                "total": 1023,
                "used": 0
            }
        },
        "ansible_memtotal_mb": 1877,
        "ansible_mounts": [
            {
                "device": "/dev/sda1",
                "fstype": "ext4",
                "mount": "/",
                "options": "rw",
                "size_available": 6597599232,
                "size_total": 8320901120,
                "uuid": "N/A"
            },
            {
                "device": "/dev/sda2",
                "fstype": "ext4",
                "mount": "/home",
                "options": "rw",
                "size_available": 1936121856,
                "size_total": 2046640128,
                "uuid": "N/A"
            }
        ],
        "ansible_nodename": "localhost.localdomain",
        "ansible_os_family": "RedHat",
        "ansible_pkg_mgr": "yum",
        "ansible_processor": [
            "GenuineIntel",
            "Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz"
        ],
        "ansible_processor_cores": 1,
        "ansible_processor_count": 1,
        "ansible_processor_threads_per_core": 1,
        "ansible_processor_vcpus": 1,
        "ansible_product_name": "VirtualBox",
        "ansible_product_serial": "0",
        "ansible_product_uuid": "E48C99B4-2B2D-416F-A58B-37287B7EB95B",
        "ansible_product_version": "1.2",
        "ansible_python": {
            "executable": "/usr/bin/python",
            "has_sslcontext": false,
            "type": "CPython",
            "version": {
                "major": 2,
                "micro": 6,
                "minor": 6,
                "releaselevel": "final",
                "serial": 0
            },
            "version_info": [
                2,
                6,
                6,
                "final",
                0
            ]
        },
        "ansible_python_version": "2.6.6",
        "ansible_real_group_id": 0,
        "ansible_real_user_id": 0,
        "ansible_selinux": false,
        "ansible_service_mgr": "upstart",
        "ansible_ssh_host_key_dsa_public": "AAAAB3NzaC1kc3MAAACBAMLsf75qwlmXJ4X6r8pDdPAP6VD1/DPn4KCaucy/AUaUaC3/giFXmwp+B4euwlTyHPOFYWUdeSe6fDrm9oXV5OMmgMQjTXIQzBj/Ai6HNoXV/bGxibFSRF1hXeVPU+5qb5usUWnHNuOLOlOJ236ze8PRPXsaXNK7glOyvsg0gBgVAAAAFQDIBDw5qg3eM/14WXOET+a9l3g2twAAAIEAuEVevatbNJ5ZFGS/mNwJvRNu9IgbvNLGDEmk685YSH8Mt602kKkWOcC+QM2HKZi7t9tcJqub3rgaPJHIIyIyUAlC9Qj294bq/HhvtYbx5dQK7EhwhVk+cuWt0aJkHcZG4uW2LABxNhNIAn4mNWAfmpXXnzGYHKMuP00HyY69np4AAACBAIDsRTKpJY7SIYIksa0mhtGnqnUQ1FAYd9m2OROtasMLILxE9+sghc58caTRLwVvNbgyd53l2p8GubeQrSGR7GjY/leO7/RbiEuuYIJM+qXuQCa1rq+xyvyh/46V7DqdXoV3q3L2BYvMTC/cib9BErUkI/TMg4UVlbPQTpzHDU24",
        "ansible_ssh_host_key_rsa_public": "AAAAB3NzaC1yc2EAAAABIwAAAQEApQTMHItcBpl73bA2AUeDQRk1jrZdOvq0n6GXThVPd0mv7ccP/ZF1m+1zJ/Wr8jqn7VOEErIxw1puPhMqnSsHSQLN9S7qEsrRTPyZzRcCXwEIl/A3Pws/5C4EV2+bR9t9OJkgVpg5nRRl5E3i4naNeUZB54DFsse92tI9Wde0hoMxx5ShwaJBbPKoXqJhc1EePKzLCmA54CB0D14xS2MCSyYtmoRHglS/OYJ/ptSASK7QmzpmNIQEyKj39bNMlqeqox/IzOgQ4oDX1fMOQLB8TvgaNzBmYDZsKNq+tzUsPxYFUy9Ip+fELJ5bJa2hmVd5tAxIlrqZbzNhHcRN4BhrsQ==",
        "ansible_swapfree_mb": 1023,
        "ansible_swaptotal_mb": 1023,
        "ansible_system": "Linux",
        "ansible_system_capabilities": [],
        "ansible_system_capabilities_enforced": "False",
        "ansible_system_vendor": "innotek GmbH",
        "ansible_uptime_seconds": 98131,
        "ansible_user_dir": "/root",
        "ansible_user_gecos": "root",
        "ansible_user_gid": 0,
        "ansible_user_id": "root",
        "ansible_user_shell": "/bin/bash",
        "ansible_user_uid": 0,
        "ansible_userspace_architecture": "x86_64",
        "ansible_userspace_bits": "64",
        "ansible_virtualization_role": "guest",
        "ansible_virtualization_type": "virtualbox",
        "module_setup": true
    },
    "changed": false
}
10.6.200.125 | SUCCESS => {
    "ansible_facts": {
        "ansible_TWM.CQB": {
            "active": true,
            "device": "TWM.CQB",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "on [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "off [fixed]",
                "netns_local": "on [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "off [fixed]",
                "rx_vlan_filter": "off [fixed]",
                "rx_vlan_offload": "off [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "on [fixed]",
                "tx_checksum_ipv4": "off [fixed]",
                "tx_checksum_ipv6": "off [fixed]",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "off [fixed]",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "on [fixed]",
                "tx_scatter_gather": "on [fixed]",
                "tx_scatter_gather_fraglist": "on [fixed]",
                "tx_tcp6_segmentation": "on",
                "tx_tcp_ecn_segmentation": "on",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "off [fixed]",
                "udp_fragmentation_offload": "on [fixed]",
                "vlan_challenged": "off [fixed]"
            },
            "ipv4": {
                "address": "1.3.5.1",
                "broadcast": "",
                "netmask": "255.255.255.255",
                "network": "1.3.5.1"
            },
            "macaddress": "0a:06:c8:7d",
            "mtu": 1476,
            "promisc": false,
            "type": "unknown"
        },
        "ansible_TWM.SHM": {
            "active": true,
            "device": "TWM.SHM",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "on [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "off [fixed]",
                "netns_local": "on [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "off [fixed]",
                "rx_vlan_filter": "off [fixed]",
                "rx_vlan_offload": "off [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "on [fixed]",
                "tx_checksum_ipv4": "off [fixed]",
                "tx_checksum_ipv6": "off [fixed]",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "off [fixed]",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "on [fixed]",
                "tx_scatter_gather": "on [fixed]",
                "tx_scatter_gather_fraglist": "on [fixed]",
                "tx_tcp6_segmentation": "on",
                "tx_tcp_ecn_segmentation": "on",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "off [fixed]",
                "udp_fragmentation_offload": "on [fixed]",
                "vlan_challenged": "off [fixed]"
            },
            "ipv4": {
                "address": "1.2.1.1",
                "broadcast": "",
                "netmask": "255.255.255.255",
                "network": "1.2.1.1"
            },
            "macaddress": "0a:06:c8:7d",
            "mtu": 1476,
            "promisc": false,
            "type": "unknown"
        },
        "ansible_all_ipv4_addresses": [
            "10.6.200.125",
            "1.3.5.1",
            "1.2.1.1"
        ],
        "ansible_all_ipv6_addresses": [
            "fe80::a00:27ff:fe4f:2037"
        ],
        "ansible_apparmor": {
            "status": "disabled"
        },
        "ansible_architecture": "x86_64",
        "ansible_bios_date": "12/01/2006",
        "ansible_bios_version": "VirtualBox",
        "ansible_cmdline": {
            "KEYBOARDTYPE": "pc",
            "KEYTABLE": "us",
            "LANG": "en_US.UTF-8",
            "SYSFONT": "latarcyrheb-sun16",
            "quiet": true,
            "rd_NO_DM": true,
            "rd_NO_LUKS": true,
            "rd_NO_LVM": true,
            "rd_NO_MD": true,
            "rhgb": true,
            "ro": true,
            "root": "UUID=7afe9693-dca3-456e-b130-c55eb8a017aa"
        },
        "ansible_date_time": {
            "date": "2017-06-22",
            "day": "22",
            "epoch": "1498097889",
            "hour": "10",
            "iso8601": "2017-06-22T02:18:09Z",
            "iso8601_basic": "20170622T101809916931",
            "iso8601_basic_short": "20170622T101809",
            "iso8601_micro": "2017-06-22T02:18:09.917027Z",
            "minute": "18",
            "month": "06",
            "second": "09",
            "time": "10:18:09",
            "tz": "CST",
            "tz_offset": "+0800",
            "weekday": "Thursday",
            "weekday_number": "4",
            "weeknumber": "25",
            "year": "2017"
        },
        "ansible_default_ipv4": {
            "address": "10.6.200.125",
            "alias": "eth0",
            "broadcast": "10.6.200.255",
            "gateway": "10.6.200.254",
            "interface": "eth0",
            "macaddress": "08:00:27:4f:20:37",
            "mtu": 1500,
            "netmask": "255.255.255.0",
            "network": "10.6.200.0",
            "type": "ether"
        },
        "ansible_default_ipv6": {},
        "ansible_devices": {
            "sda": {
                "holders": [],
                "host": "SATA controller: Intel Corporation 82801HM/HEM (ICH8M/ICH8M-E) SATA Controller [AHCI mode] (rev 02)",
                "model": "VBOX HARDDISK",
                "partitions": {
                    "sda1": {
                        "holders": [],
                        "sectors": "16777216",
                        "sectorsize": 512,
                        "size": "8.00 GB",
                        "start": "2048",
                        "uuid": "7afe9693-dca3-456e-b130-c55eb8a017aa"
                    },
                    "sda2": {
                        "holders": [],
                        "sectors": "4194304",
                        "sectorsize": 512,
                        "size": "2.00 GB",
                        "start": "16779264",
                        "uuid": "57f081cd-53b5-48b0-a4d0-72263645ee10"
                    },
                    "sda3": {
                        "holders": [],
                        "sectors": "2097152",
                        "sectorsize": 512,
                        "size": "1.00 GB",
                        "start": "20973568",
                        "uuid": "28bd1e0e-f238-411e-81cc-549765f62ef3"
                    }
                },
                "removable": "0",
                "rotational": "1",
                "sas_address": null,
                "sas_device_handle": null,
                "scheduler_mode": "cfq",
                "sectors": "34490464",
                "sectorsize": "512",
                "size": "16.45 GB",
                "support_discard": "0",
                "vendor": "ATA"
            },
            "sr0": {
                "holders": [],
                "host": "IDE interface: Intel Corporation 82371AB/EB/MB PIIX4 IDE (rev 01)",
                "model": "CD-ROM",
                "partitions": {},
                "removable": "1",
                "rotational": "1",
                "sas_address": null,
                "sas_device_handle": null,
                "scheduler_mode": "cfq",
                "sectors": "2097151",
                "sectorsize": "512",
                "size": "1024.00 MB",
                "support_discard": "0",
                "vendor": "VBOX"
            }
        },
        "ansible_distribution": "CentOS",
        "ansible_distribution_major_version": "6",
        "ansible_distribution_release": "Final",
        "ansible_distribution_version": "6.7",
        "ansible_dns": {
            "nameservers": [
                "127.0.0.1",
                "10.188.34.23"
            ]
        },
        "ansible_domain": "localdomain",
        "ansible_effective_group_id": 0,
        "ansible_effective_user_id": 0,
        "ansible_env": {
            "G_BROKEN_FILENAMES": "1",
            "HOME": "/root",
            "LANG": "POSIX",
            "LC_CTYPE": "zh_CN.UTF-8",
            "LESSOPEN": "||/usr/bin/lesspipe.sh %s",
            "LOGNAME": "root",
            "MAIL": "/var/mail/root",
            "PATH": "/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin",
            "PWD": "/root",
            "SHELL": "/bin/bash",
            "SHLVL": "2",
            "SSH_CLIENT": "10.6.200.131 35451 22",
            "SSH_CONNECTION": "10.6.200.131 35451 10.6.200.125 22",
            "SSH_TTY": "/dev/pts/1",
            "TERM": "xterm-256color",
            "USER": "root",
            "_": "/usr/bin/python"
        },
        "ansible_eth0": {
            "active": true,
            "device": "eth0",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "off [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "off [fixed]",
                "netns_local": "off [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "on",
                "rx_vlan_filter": "on [fixed]",
                "rx_vlan_offload": "on [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "on",
                "tx_checksum_ipv4": "off",
                "tx_checksum_ipv6": "off",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "off",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "off [fixed]",
                "tx_scatter_gather": "on",
                "tx_scatter_gather_fraglist": "off [fixed]",
                "tx_tcp6_segmentation": "off",
                "tx_tcp_ecn_segmentation": "off",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "on [fixed]",
                "udp_fragmentation_offload": "off [fixed]",
                "vlan_challenged": "off [fixed]"
            },
            "ipv4": {
                "address": "10.6.200.125",
                "broadcast": "10.6.200.255",
                "netmask": "255.255.255.0",
                "network": "10.6.200.0"
            },
            "ipv6": [
                {
                    "address": "fe80::a00:27ff:fe4f:2037",
                    "prefix": "64",
                    "scope": "link"
                }
            ],
            "macaddress": "08:00:27:4f:20:37",
            "module": "e1000",
            "mtu": 1500,
            "pciid": "0000:00:03.0",
            "promisc": false,
            "speed": 1000,
            "type": "ether"
        },
        "ansible_fips": false,
        "ansible_form_factor": "Other",
        "ansible_fqdn": "localhost.localdomain",
        "ansible_gather_subset": [
            "hardware",
            "network",
            "virtual"
        ],
        "ansible_gre0": {
            "active": false,
            "device": "gre0",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "on [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "off [fixed]",
                "netns_local": "on [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "off [fixed]",
                "rx_vlan_filter": "off [fixed]",
                "rx_vlan_offload": "off [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "on [fixed]",
                "tx_checksum_ipv4": "off [fixed]",
                "tx_checksum_ipv6": "off [fixed]",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "off [fixed]",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "on [fixed]",
                "tx_scatter_gather": "on [fixed]",
                "tx_scatter_gather_fraglist": "on [fixed]",
                "tx_tcp6_segmentation": "on",
                "tx_tcp_ecn_segmentation": "on",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "off [fixed]",
                "udp_fragmentation_offload": "on [fixed]",
                "vlan_challenged": "off [fixed]"
            },
            "macaddress": "0a:06:c8:7d",
            "mtu": 1476,
            "promisc": false,
            "type": "unknown"
        },
        "ansible_gretap0": {
            "active": false,
            "device": "gretap0",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "on [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "off [fixed]",
                "netns_local": "on [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "off [fixed]",
                "rx_vlan_filter": "off [fixed]",
                "rx_vlan_offload": "off [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "on [fixed]",
                "tx_checksum_ipv4": "off [fixed]",
                "tx_checksum_ipv6": "off [fixed]",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "off [fixed]",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "on [fixed]",
                "tx_scatter_gather": "on [fixed]",
                "tx_scatter_gather_fraglist": "on [fixed]",
                "tx_tcp6_segmentation": "on",
                "tx_tcp_ecn_segmentation": "on",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "off [fixed]",
                "udp_fragmentation_offload": "on [fixed]",
                "vlan_challenged": "off [fixed]"
            },
            "mtu": 1476,
            "promisc": false,
            "type": "ether"
        },
        "ansible_hostname": "localhost",
        "ansible_interfaces": [
            "gre0",
            "TWM.SHM",
            "lo",
            "gretap0",
            "TWM.CQB",
            "eth0"
        ],
        "ansible_kernel": "2.6.32-573.el6.x86_64",
        "ansible_lo": {
            "active": true,
            "device": "lo",
            "features": {
                "fcoe_mtu": "off [fixed]",
                "generic_receive_offload": "on",
                "generic_segmentation_offload": "on",
                "highdma": "on [fixed]",
                "large_receive_offload": "off [fixed]",
                "loopback": "on [fixed]",
                "netns_local": "on [fixed]",
                "ntuple_filters": "off [fixed]",
                "receive_hashing": "off [fixed]",
                "rx_checksumming": "on [fixed]",
                "rx_vlan_filter": "off [fixed]",
                "rx_vlan_offload": "off [fixed]",
                "scatter_gather": "on",
                "tcp_segmentation_offload": "on",
                "tx_checksum_fcoe_crc": "off [fixed]",
                "tx_checksum_ip_generic": "off [fixed]",
                "tx_checksum_ipv4": "off [fixed]",
                "tx_checksum_ipv6": "off [fixed]",
                "tx_checksum_sctp": "off [fixed]",
                "tx_checksum_unneeded": "on [fixed]",
                "tx_checksumming": "on",
                "tx_fcoe_segmentation": "off [fixed]",
                "tx_gre_segmentation": "off [fixed]",
                "tx_gso_robust": "off [fixed]",
                "tx_lockless": "on [fixed]",
                "tx_scatter_gather": "on [fixed]",
                "tx_scatter_gather_fraglist": "on [fixed]",
                "tx_tcp6_segmentation": "on",
                "tx_tcp_ecn_segmentation": "on",
                "tx_tcp_segmentation": "on",
                "tx_udp_tnl_segmentation": "off [fixed]",
                "tx_vlan_offload": "off [fixed]",
                "udp_fragmentation_offload": "on",
                "vlan_challenged": "on [fixed]"
            },
            "ipv4": {
                "address": "127.0.0.1",
                "broadcast": "host",
                "netmask": "255.0.0.0",
                "network": "127.0.0.0"
            },
            "ipv6": [
                {
                    "address": "::1",
                    "prefix": "128",
                    "scope": "host"
                }
            ],
            "mtu": 65536,
            "promisc": false,
            "type": "loopback"
        },
        "ansible_lvm": {
            "lvs": {},
            "vgs": {}
        },
        "ansible_machine": "x86_64",
        "ansible_machine_id": "f86bf844594ba927270bd6db0000000c",
        "ansible_memfree_mb": 732,
        "ansible_memory_mb": {
            "nocache": {
                "free": 896,
                "used": 100
            },
            "real": {
                "free": 732,
                "total": 996,
                "used": 264
            },
            "swap": {
                "cached": 0,
                "free": 1023,
                "total": 1023,
                "used": 0
            }
        },
        "ansible_memtotal_mb": 996,
        "ansible_mounts": [
            {
                "device": "/dev/sda1",
                "fstype": "ext4",
                "mount": "/",
                "options": "rw",
                "size_available": 6950420480,
                "size_total": 8320901120,
                "uuid": "N/A"
            },
            {
                "device": "/dev/sda2",
                "fstype": "ext4",
                "mount": "/home",
                "options": "rw",
                "size_available": 1936121856,
                "size_total": 2046640128,
                "uuid": "N/A"
            }
        ],
        "ansible_nodename": "localhost.localdomain",
        "ansible_os_family": "RedHat",
        "ansible_pkg_mgr": "yum",
        "ansible_processor": [
            "GenuineIntel",
            "Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz"
        ],
        "ansible_processor_cores": 1,
        "ansible_processor_count": 1,
        "ansible_processor_threads_per_core": 1,
        "ansible_processor_vcpus": 1,
        "ansible_product_name": "VirtualBox",
        "ansible_product_serial": "0",
        "ansible_product_uuid": "BC10C9D2-05AF-4E96-BA26-E1FCA1FAF09F",
        "ansible_product_version": "1.2",
        "ansible_python": {
            "executable": "/usr/bin/python",
            "has_sslcontext": false,
            "type": "CPython",
            "version": {
                "major": 2,
                "micro": 6,
                "minor": 6,
                "releaselevel": "final",
                "serial": 0
            },
            "version_info": [
                2,
                6,
                6,
                "final",
                0
            ]
        },
        "ansible_python_version": "2.6.6",
        "ansible_real_group_id": 0,
        "ansible_real_user_id": 0,
        "ansible_selinux": false,
        "ansible_service_mgr": "upstart",
        "ansible_ssh_host_key_dsa_public": "AAAAB3NzaC1kc3MAAACBAIhSrPf/f6w0qDFvcMPAg7ilTWaoAlhSYDy3jfHIaFbY4GpbVpTnkBVtX1IAaG15TZiDFd3pRHYSc63oTP/NWsvxU8xJ1UWTQGx3d8a0s8Y2pSza9MgA/JANOaQyywlJpfgc1WVGHMLMJbKGO/xcs7wh5pgrokamO39Q8Ll2JQb9AAAAFQC5nbdOv2aobSStnYymPNHHe6yt8QAAAIABMERIRugsuBN0I9OHJZkNi66Ily50x9k2z17LWxHyBLaeWmNWesz6v6c0Oy9gKLBlPVi0maxCw0zL/SxJcCMNjQhzesGL1xXbft9gYwwCNyVTvnoL7smUfZoY+yrq2RWw5kh9DwwMPlOYJrtvXXl3jKKsSRHeLrHc6+MbydoIwgAAAIByyony+zut4b/TjBcdrOdLdOqGLydblQVkIRJ3utVyldVeBfoQ2qWYXLyDCh4oChzu9gb/JClj0qWJ15i+0dHlEKuxr/R3HS0CJj3NKDdU5GKMKM/0gJ8YMnnuB0uHMhRjjeIL1g0WkCDK1TXDCi0rG0Z9q/0d/Vx1i1RhqHzxaA==",
        "ansible_ssh_host_key_rsa_public": "AAAAB3NzaC1yc2EAAAABIwAAAQEA6nGUvz/kJ48S1w3MnlPp0vOwkCyO7KbHLyESBs3koFiuEYIGs0rpglol2esMDHYMVTbJm3TldEPiixXYMqp1O4ovXGSRkGOJy2ERnDXzqtJKCTRbISuxIJwq/dAkR9YtOhDHc6ypx0jGbrP4D0BPJ2B6LGqAXtsPx0+96FjyQmXwPYduf2l0Pu+TtHM35QN4ECjMiMsWwz5KrPlVBqrdX0+49rLcUWv/b9pEdBik1fq/QSpkB8Lw5/AinMTJT6dn3kkd6BvSdxOJthEBn1lfxY5KUNNcgLYzoZQLtseInJvqnXgZbePd8ckVJeJNA4iad34DkvuzhOwz/ctlJKW/Kw==",
        "ansible_swapfree_mb": 1023,
        "ansible_swaptotal_mb": 1023,
        "ansible_system": "Linux",
        "ansible_system_capabilities": [],
        "ansible_system_capabilities_enforced": "False",
        "ansible_system_vendor": "innotek GmbH",
        "ansible_uptime_seconds": 98128,
        "ansible_user_dir": "/root",
        "ansible_user_gecos": "root",
        "ansible_user_gid": 0,
        "ansible_user_id": "root",
        "ansible_user_shell": "/bin/bash",
        "ansible_user_uid": 0,
        "ansible_userspace_architecture": "x86_64",
        "ansible_userspace_bits": "64",
        "ansible_virtualization_role": "guest",
        "ansible_virtualization_type": "virtualbox",
        "module_setup": true
    },
    "changed": false
}
10.6.200.123 | UNREACHABLE! => {
    "changed": false,
    "msg": "Failed to connect to the host via ssh: ssh: connect to host 10.6.200.123 port 22: No route to host\r\n",
    "unreachable": true
}
`
	for x,y := range strings.Split(a,"=>") {
		println(x,y)
	}
}
