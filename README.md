# Linny (Work in progress)
**A light-weight and simple micro ad server**

Linny is a light weight and simple micro ad server that is designed to serve one web display ad or campaign. It can measure simple but important statistics, such as views and conversion via click-throughs.

## Features

* Serves one ad or ad campaign
* Measures ad views and click-throughs
* Measures any number of conversion landing pages
* Simple click-through link tagging and generation


## Principles

The following are guiding principles used to design and build this system.

1. It should be light weight (less than 10MB).
2. It should be fast (a response time of less than 200ms).
3. It must be able to serve 1,000 ads per second on a $5 [Linode](http://www.linode.com) or [Droplet](http://www.digitalocean.com).
4. It is designed to only serve one web display ad or ad campaign (a group of ads with one click-through url).
5. It should be quick to configure and deploy (~10 minuites)

## Prerequsites

1. A server with a recent copy of Linux installed
2. GoLang installed and configured

## Getting Started

When you have the prequsites in place you can follow the steps below to get the ad server running.

1. Download the latest copy of Linny
	- wget ....(TODO URL and command to download latest)
2. Run Linny
	- go run ....(TODO go command to run Linny)

## Display Ad Development

### Building a Single Web Display Ad

#### Medium Rectangle Ad 300x250

### Building a Web Display Ad Campaign

#### Medium Rectangle Ad 300x250
#### Half Page Ad 300x600
#### Banner Ad 728x90
#### Mobile Banner Ad 320x50
#### Wallpaper Ad (Advanced)


#### Internal Resources References
{{ilk "assets/testURL.com" }}

#### Click-through Links with Tracking
{{mlk "http://www.testURL.com" tag="testTAG"}}

## Deploying to Production


## About

Linny was built to make it easy and accessable for developers and businesses alike to build, host and measure the performance of their web display ads.

## Copyright and Licence

Copyright 2015, Simon Chong. Licensed under the Apache License, Version 2.0 .
