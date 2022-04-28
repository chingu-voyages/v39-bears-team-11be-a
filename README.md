This repository contains endpoints for user signup, login and authentication 
of the litetalk chat application.

Litetalk is a lightweight social media app for connecting with family and friends. You can chat, share media and make calls with family and friends for free. Litetalk uses less bandwidth so it supports 2G network areas pretty well.

     Getting Started
This is an example of how you can setup your project locally. To get a local copy up and running follow these simple example steps.
Prerequisites
Go 1.16 or lastest version already installed on your local machine.
MongoDB


Installation

1. Run an instance of MongoDB

    $mongod

2. Create and populate a .env file with its keys corresponding values as listed in 

    example.env

3. Run litetalk from project root directory

   $cd/path/to/litetalk

   $go run main.go


Testing

1. Lint checks are done with golangci-lint - an aggregator of linters

    $ cd /path/to/litetalk

    $ golangci-lint run