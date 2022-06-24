package aci

var resourceContractTest = map[string]interface{}{
	"weight": map[string]interface{}{
		"valid":   []interface{}{-97.98930471462685, 163.1833847711053, 28.38634022582533, 99.19214666776175, -3.7950847208787444},
		"invalid": []interface{}{"random", 10},
	},

	"ipv4_for": map[string]interface{}{
		"valid":   []interface{}{"140.213.70.124", "140.213.170.61", "140.213.23.42", "140.213.178.198", "140.213.126.255", "140.213.24.218", "140.213.254.226", "140.213.110.10", "140.213.166.188", "140.213.149.132", "140.213.167.179", "140.213.99.188", "140.213.205.126", "140.213.75.14", "140.213.136.25"},
		"invalid": []interface{}{"284.297.291.283"},
	},

	"port_number": map[string]interface{}{
		"valid":   []interface{}{1, 65535, 6127, 56107, 2327, 42801, 46371, 23929, 40817, 63397, 4899, 14962, 63035, 40263, 4075},
		"invalid": []interface{}{0, 65536},
	},

	"test_score": map[string]interface{}{
		"valid":   []interface{}{1, 100, 50, 7, 45, 42, 98, 42, 49, 47, 6, 94, 7, 54, 73},
		"invalid": []interface{}{0, 101},
	},

	"string_in_some_names": map[string]interface{}{
		"valid":   []interface{}{"parth", "aarsh", "arjun", "alfatah", "krunal"},
		"invalid": []interface{}{"3npllak04l"},
	},

	"valid_cidr": map[string]interface{}{
		"valid":   []interface{}{0, 32, 16, 6, 29, 20, 16, 13, 6, 7, 25, 22, 17, 22, 30},
		"invalid": []interface{}{-1, 33},
	},

	"percentage": map[string]interface{}{
		"valid":   []interface{}{0, 100, 50.0, 11.515536049074903, 17.433330898926098, 3.2804641632253593, 58.93643040105934, 13.17540857218283, 44.075584561467785, 9.590478663554332, 3.6907867278564823, 8.173093034377672, 6.414365094603028, 31.65650829681244, 0.0, 25.204629565742028},
		"invalid": []interface{}{-1, 101},
	},

	"filter": map[string]interface{}{
		"filter_name": map[string]interface{}{
			"valid":   []interface{}{"lkem2khaa3", "3u5xeo6v54", "iov58oh3k9", "fmf0r2ftyd", "e9fg2652j9"},
			"invalid": []interface{}{10, 12.43},
		},

		"id": map[string]interface{}{
			"valid":   []interface{}{"c77ifrm9o2", "p6v11hf9ye", "xalvvnvyop", "87byfc1wih", "pphoe0vsm9"},
			"invalid": []interface{}{10, 12.43},
		},

		"description": map[string]interface{}{
			"valid":   []interface{}{"fyqxihiygb", "75rm4bh6br", "noqamt4pfd", "smfb6rceos", "5fb1uw6dc7"},
			"invalid": []interface{}{10, 12.43},
		},

		"filter_entry": map[string]interface{}{
			"filter_entry_name": map[string]interface{}{
				"valid":   []interface{}{"n7aq0oxmqh", "elnujyfy01", "vvfwpkuqgk", "3k14w988uy", "lb08j39xdb"},
				"invalid": []interface{}{10, 12.43},
			},

			"ipv6": map[string]interface{}{
				"valid":   []interface{}{"2001:db8::34f4:0:0:f39d", "2001:db8::34f4:0:0:f3d2", "2001:db8::34f4:0:0:f361", "2001:db8::34f4:0:0:f36b", "2001:db8::34f4:0:0:f396", "2001:db8::34f4:0:0:f344", "2001:db8::34f4:0:0:f3b5", "2001:db8::34f4:0:0:f3df", "2001:db8::34f4:0:0:f339", "2001:db8::34f4:0:0:f387", "2001:db8::34f4:0:0:f319", "2001:db8::34f4:0:0:f3fd", "2001:db8::34f4:0:0:f397", "2001:db8::34f4:0:0:f34b", "2001:db8::34f4:0:0:f360"},
				"invalid": []interface{}{"invalidIPv6"},
			},

			"apply_to_frag": map[string]interface{}{
				"valid":   []interface{}{"yes", "no"},
				"invalid": []interface{}{"a6vuh3u485"},
			},
		},
	},
}
