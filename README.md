# Linny (Work in progress)
**A light-weight and simple micro ad server**

Linny is a light weight and simple micro ad server that is designed to serve one web display ad or campaign. It can measure simple but important statistics, such as views and conversion via click-throughs.

## Features

* Serves one ad or ad campaign
* Measures ad views and click-throughs
* Measures any number of conversions on any number of pages
* Simple click-through link tagging and generation
* Simple link resolution


## Principles

The following are guiding principles used to design and build this system.

1. It should be light weight (less than 10MB).
2. It should be fast (a response time of less than 200ms).
3. It must be able to serve 1,000 ads per second on a server with 1GB of RAM and 1 Core.
4. It is designed to only serve one web display ad or ad campaign (a group of ads of related ads).
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

### Folder Structure

### Configuration
There are two configurations required to serve ads, a server configuration and an ad configuration.

#### Server Configuration
The file name configLinny.json in the current working directory configures the server.

```json
{
    "ContentRoot": "./exampleCampaign"
}
```

- **ContentRoot** the directory of the ad or campaign exists

#### Ad Configuration
This configuration file is specific to the ad or campaign itself.

```json
{
	"Id" : "ABC123",
	"Name": "Example Campaign",
	"HeaderFrag" : "header.frag",
	"FooterFrag" : "footer.frag"
}
```

- **Id** is a string that uniquely identifying the ad or campaign
- **Name** the name of the ad or campaign (should not be longer than 255 characters)
- **HeaderFrag** location of the header html fragment that wraps all HTML files
- **FooterFrag** location of the footer html fragment that wraps all HTML files


### Internal Resources References
Ad creatives usually rely on loading a number of assets and resources. To ensure they resolve to the correct address, absolute urls are generated using the following syntax.
- Single quotes and double quotes are supported.
- "ilk" stands for internal link

```js
{{ilk "someAsset.png" }}
```

Example usage inside your HTML:

```html
<a href="{{ilk 'someAsset.js'}}"> LINK</a>
```

### Click-through Links with Tracking
To track when users click on links to pages
{{mlk "http://www.testURL.com" tag="testTAG"}}

### Conversion Taging
```js
var s = document.createElement("script");
s.src = "//localhost:8000/m/c.js?t=YOURTAG";
document.body.appendChild(s);
```


## Display Ad Development

### Building a Single Web Display Ad

#### Medium Rectangle Ad 300x250

### Building a Web Display Ad Campaign

#### Medium Rectangle Ad 300x250
#### Half Page Ad 300x600
#### Banner Ad 728x90
#### Mobile Banner Ad 320x50
#### Wallpaper Ad (Advanced)



## Deploying to Production


## About

Linny was built to make it easy and accessable for developers and businesses alike to build, host and measure the performance of their web display ads.

## Copyright and Licence

Copyright 2015, Simon Chong. Licensed under the Apache License, Version 2.0 .
