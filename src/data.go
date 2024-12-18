package main

type Link struct {
    Text string
    URL  string
    Description string
}

var links = []Link{
    {"profile.junyi.me", "https://profile.junyi.me/", "My personal website"},
    {"blog.junyi.me", "https://blog.junyi.me/", "My blog"},
}

