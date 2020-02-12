# GOnyx

[![Go Report Card](https://goreportcard.com/badge/github.com/tempor1s/gonyx)](https://goreportcard.com/report/github.com/tempor1s/gonyx)

A clone of Onyx Bot that was built using Go because the old version needed a lot of work and I need Golang practice!

## Getting Started

Install docker. Easiest way on mac is to `brew install docker`

Once you have docker installed and running, you need to rename `.env.example` to `.env` and fill out the blank variables.

When you are finished, it should look similar to below.
```bash
BOT_TOKEN=NTU3NDI1MDkzNDI1gw.If0maDHLCGdp7sMntGfNP54tW_c
LOG_CHANNEL=536328234556518033
WEEKLY_INFO_CHANNEL=510632234556588032
```

Then, all you need to do is run `docker-compose up --build` and you should in theory have a working version to do whatever you want with! :)