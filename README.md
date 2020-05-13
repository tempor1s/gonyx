# GOnyx

[![Go Report Card](https://goreportcard.com/badge/github.com/tempor1s/gonyx)](https://goreportcard.com/report/github.com/tempor1s/gonyx)

A simple bot that is used to display weekly information for Destiny 2 in a channel of your choice. Also has a few other management features like autorole on join! Was custom built for the Destiny 2 clan Onyx and customed to their needs.

## Getting Started

Install docker.

Once you have docker installed and running, you need to rename `.env.example` to `.env` and fill out the blank variables.

When you are finished, it should look similar to below.
```bash
BOT_TOKEN=NTU3NDI1MDkzNDI1gw.If0maDHLCGdp7sMntGfNP54tW_c
LOG_CHANNEL=536328234556518033
WEEKLY_INFO_CHANNEL=51063223556588032
ROLE_ON_JOIN=6458533011868870
```

Then, all you need to do is run `docker-compose up --build` and you should in theory have a working version to do whatever you want with! :)
