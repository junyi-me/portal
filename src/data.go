package main

type Link struct {
    Text string
    URL  string
    Favicon string
    Description string
}

var hosted = []Link{
    {"profile.junyi.me", "https://profile.junyi.me/", "https://profile.junyi.me/favicon.png", "Personal website"},
    {"blog.junyi.me", "https://blog.junyi.me/", "https://blog.junyi.me/favicon.ico", "My blog"},
}

var external = []Link{
    {"github.com/jywang99", "https://github.com/jywang99/", "https://github.githubassets.com/favicons/favicon.svg", "GitHub profile"},
    {"www.linkedin.com/in/junyi-wang-976a94199", "https://www.linkedin.com/in/junyi-wang-976a94199/", "https://static.licdn.com/aero-v1/sc/h/akt4ae504epesldzj74dzred8", "LinkedIn profile"},
    {"junyi.wang.007@gmail.com", "mailto:junyi.wang.007@gmail.com", "https://ssl.gstatic.com/ui/v1/icons/mail/rfr/gmail.ico", "Email me"},
}

