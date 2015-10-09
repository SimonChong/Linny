# Linny (Work in progress)
**A light-weight and simple micro ad server**

Linny is a light weight and simple micro ad server that is designed to serve one web display ad or campaign. It can measure simple but important statistics, such as views and conversion via click-throughs.

## Features

* Serves one ad or ad campaign
* Measures ad views and click-throughs
* Measures any number of conversions on any number of pages
* Simple click-through link tagging and generation
* Simple link resolution templating


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
	- go get ....(TODO URL and command to download latest)
2. Run Linny
	- linny ....(TODO go command to run Linny)

### Quick Start

1. Create a new directory ```mkdir adServer```
2. Run the linny initialization command ```linny -init``` (this will create all the necessary files)
4. Run the linny server ```linny -serve```
3. Modify and add files until your creative is ready


### Ad Folder Structure

- Server Config
	- configLinny.json (configures the server and where the ad directory)


- Ad Directory
	- configAd.json (the ad configuration)
	- assets (directory for all creative code)
		- index.html (default html file)
	- header.frag (the header prepended to all HTML files)
	- footer.frag (the footer appended to all HTML files)


## Linny Commands

```
linny -serve
```

#### Packaging your Ad
Your can package your ad so that you can send it via email or sftp
```
linny -pack
```
To unpack your ad just use the following command
```
linny -unpack
```


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

## Ad Development

### Internal Resources References
Ad creatives usually rely on loading a number of assets and resources. To ensure they resolve to the correct address, absolute urls are generated using the following syntax.
- Single quotes and double quotes are supported.
- "ilk" stands for internal link

```js
{{ilk "someAsset.png" }}
```

Example usage inside your HTML ad:

```html
<script src="{{ilk 'someAsset.js'}}"></script>
```

### Click-through Links with Tracking
To track when users click on an ad you can use the following code to track the clicks. Appropriate conversion attribution can be tracked as per the conversion tagging section.

```js
{{mlk "http://www.testURL.com" tag="testTAG"}}
```

Example usage inside you HTML ad:
```html
<a href='{{mlk "http://www.testURL.com" tag="testTAG"}}'>Link</a>
```

### Conversion Tagging
This helps you track the effectiveness of your ad. After a user clicks on the ad, you can track if user bought something or did something by adding or loading the following script on that action.


```html
<script>
var s = document.createElement("script");
s.src = "//localhost:8000/m/c.js?t=YOURTAG";
document.body.appendChild(s);
</script>
```

- Replace "YOURTAG" with a label describing the conversion
- Replace "localhost:8000" with the domain of your ad server
- Using different tags will allow you to track multiple conversions

#### Example:

If you host an ad to sell something, putting this script on the thank you page after someone buys something will allow you to see how many people bought something after clicking on the ad i.e the conversion rate.


## Display Ad Development Tutorials

### Building a Single Medium Rectangle (300x250) Display Ad

1. Install Linny
2. Create a folder called "MedRecAdServer"
3. Run "linny -init" this will initialize and create a blank ad
4. Edit the "index.html" file in MedRecAdServer/newAdDir/assests with your favorite text editor.


### Building a Display Ad Campaign

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
