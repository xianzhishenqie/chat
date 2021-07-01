package s_conf

var Conf = conf{
	SERVING: ":8000",
	APPS: []string{
		"sv_base",
		"m_chat",
	},
}
