# harpun
[![Build Status](https://travis-ci.com/80-am/harpun.svg?branch=master&status=started)](https://travis-ci.com/80-am/harpun)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/80-am/harpun)](https://golang.org/)
[![License: GPL-3.0](https://img.shields.io/github/license/80-am/harpun)](https://opensource.org/licenses/GPL-3.0)

Hunting whales on OMXS.

## Getting Started
These instructions will get you up and running on your local machine.

```sql
CREATE DATABASE harpun;
SOURCE db.sql;
```

Copy [config.yml.sample](config.yml.sample) into config.yml and fill in your secrets.

```yml
# db
user: "your db user"
password: "your db password"
schema: "/harpun"
```
