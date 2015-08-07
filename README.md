# Linny (Work in progress)
**A light-weight ad server**

Linny is a light weight micro ad server that is designed to serve one web display ad or campaign. It can measure simple but important statistics, such as views and conversion via click-throughs.

## Features

* Serves one ad or ad campaign
* Measures ad views and click-through conversions
* Simple click-through link taging and generation


## Principles

The following are guiding principles used to design and build this system.

1. It should be light weight (less than 10MB). 
2. It should be fast (a response time of less than 5ms).
3. It must be able to serve 1,000 ads per second on a $5 [Linode](http://www.linode.com) or [Droplet](http://www.digitalocean.com).
4. It is designed to only serve one web display ad or ad campaign (a group of ads with one click-through url).
5. It should be quick to configure and deploy (~10 minuites)

## Prerequsites

1. A server with a recent copy of Linux installed
2. GoLang installed and configured

## Getting Started

When you have the prequsites in place you can follow the steps below to get the ad server running.

1. Download the latest copy of Linny 
	- (TODO linux command to download latest)
2. Run Linny
	- (TODO go command to run Linny)

## Display Ad Development

### Building a single web display ad

#### Medium rectangle ad (Med Rec 300x200 TODO check size)

### Building a web display ad campaign

#### Medium rectangle ad (Med Rec 300x200 TODO check size)
#### Half Page Ad (300x600 TODO check size)
#### Banner Ad (TODO check size)
#### Wallpaper Ad

### Click-through tagging / generation

## Deploying to production


## About

Linny was built to make it easy and accessable for developers and businesses alike to build, host and measure the performance of their web display ads.

## Copyright and Licence

Copyright 2015, Simon Chong. Licensed under the Apache License, Version 2.0 .