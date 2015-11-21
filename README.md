# Linny (pre alpha)
**A light-weight and simple micro ad server**

Linny is a light weight and simple micro ad server that is designed to serve one web display ad or campaign. It can measure simple but important statistics, such as views and conversion via click-throughs.

# Table of Contents
- [Features](## Features)
- [Principles](## Principles)
- [Prerequisites](## Prerequisites)
- [Getting Started](## Getting Started)
- [Linny Commands](## Linny Commands)
- [Server Configuration](## Server Configuration)
- [Ad Development](## Ad Development)
- [Accessing the DATA](## Accessing the DATA)
- [Display Ad Development Tutorials](## Display Ad Development Tutorials)
- [Deploying to Production](## Deploying to Production)
- [About](## About)
- [Copyright and Licence](## Copyright and Licence)

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

## Prerequisites

1. A server with a recent copy of Linux installed
2. GoLang installed and configured

## Getting Started

When you have the perquisites in place you can follow the steps below to get the ad server running.

1. Download the latest copy of Linny
	- go get github.com/SimonChong/Linny
2. Run Linny
	- linny -serve

### Quick Start

1. Create a new directory ```mkdir adServer```
2. Run the linny initialization command ```linny -init``` (this will create all the necessary files)
4. Run the linny server ```linny -serve```
5. Open your browser and go to ```localhost:8000```
6. Modify and add files until your creative is ready


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

#### Ad server
Start the linny ad server:
```
linny -serve
```
```
linny -serve -bind=":8000"
```
### Development
Create a new server directory with example ad directory
```
linny -init
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


## Server Configuration
There are two configurations required to serve ads, a server configuration and an ad configuration.

#### Server Configuration
The file name configLinny.json in the current working directory configures the server.

```json
{
    "ContentRoot": "./exampleCampaign"
}
```

- **ContentRoot** is the directory of the ad or campaign exists

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

**The following tools are available only within HTML files.**

### Internal Resources References
Ad creatives usually rely on loading a number of assets and resources. To ensure they resolve to the correct address, absolute urls are generated using the following syntax.
- Single quotes and double quotes are supported.
- "ilk" stands for Internal Link

```js
{{ilk "someAsset.png" }}
```

Example usage inside your HTML ad:

```html
<script src="{{ilk 'someAsset.js'}}"></script>
```

### Click-through Links with Tracking
To track when users click on an ad you can use the following code to track the clicks. Appropriate conversion attribution can be tracked as per the conversion tagging section.
- Single quotes and double quotes are supported.
- "mlk" stands for Metrics Link

```js
{{mlk "http://www.testURL.com" tag="testTAG"}}
```

Example usage inside you HTML ad:
```html
<a href='{{mlk "http://www.testURL.com" tag="testTAG"}}'>Link</a>
```

### Conversion Tagging
This helps you track the effectiveness of your ad. After a user clicks on the ad (using an MLK link), you can track if user bought something or did something by adding or triggering the following script on a page.


```html
<script>
var s = document.createElement("script");
s.src = "//localhost:8000/m/c.js?t=YOURTAG";
document.body.appendChild(s);
</script>
```

- Replace "YOURTAG" with a label describing the conversion
	- Tags are limited to 64 characters in length
- Replace "localhost:8000" with the domain of your ad server
- Using different tags will allow you to track multiple conversions

#### Example:

If you host an ad to sell something, putting this script on the thank you page after someone buys something will allow you to see how many people bought something after clicking on the ad i.e the conversion rate.


## Accessing the DATA

TODO

## Display Ad Development Tutorials

### Building a Single Medium Rectangle (300x250) Display Ad

1. Install Linny
2. Create a folder called "MedRecAdServer"
3. Run "linny -init" this will initialize and create a blank ad
4. Edit the "index.html" file in MedRecAdServer/newAdDir/assets with your favorite text editor.
5. Add code until it looks like the following

```html
<style type="text/css">
body {
    font-family: Arial, sans-serif;
}

#mr {
    display: block;
    width: 300px;
    height: 250px;
    color: white;
    background-color: black;
    margin: 0 auto;
}

.title {
    text-align: center;
    font-size: 40px;
    padding: 15px;
}
</style>
<div id="mr">
    <div class="title">
        Medium Rectangle 300x250
    </div>
</div>
```

6. Run ```linny -serve```
7. Open our browser and go to ```localhost:8000```

8. Embedding it into a publisher's site can be done using an iframe as per the below code. This also works with other Ad Servers such as Double Click for Publishers.

```html
<iframe src="//YOUR_DOMAIN/"></iframe>
```


### Building a Display Ad Campaign

1. Install Linny
2. Create a folder called "AdCampaign"
3. Run "linny -init" this will initialize and create a blank ad


#### Medium Rectangle Ad 300x250

Please refer to the the section above.

#### Half Page Ad 300x600

1. Create a file named ```half_page.html``` in the AdCampaign/newAdDir/assets directory.
2. Edit ```half_page.html``` with your favorite text editor and add code until it looks like the following:

```html
<style type="text/css">
body {
    font-family: Arial, sans-serif;
}

#halfpage {
    display: block;
    width: 300px;
    height: 600px;
    color: white;
    background-color: black;
    margin: 0 auto;
}

.title {
    text-align: center;
    font-size: 40px;
    padding: 15px;
}
</style>
<div id="halfpage">
    <div class="title">
        Half Page 300x600
    </div>
</div>
```

3. Embedding it into a publisher's site can be done using an iframe as per the below code. This also works with other Ad Servers such as Double Click for Publishers.

```html
<iframe src="//YOUR_DOMAIN/half_page.html"></iframe>
```

#### Banner Ad 728x90

1. Create a file named ```banner.html``` in the AdCampaign/newAdDir/assets directory.
2. Edit ```banner.html``` with your favorite text editor and add code until it looks like the following:

```html
<style type="text/css">
body {
    font-family: Arial, sans-serif;
}

#banner {
    display: block;
    width: 728px;
    height: 90px;
    color: white;
    background-color: black;
    margin: 0 auto;
}

.title {
    text-align: center;
    font-size: 40px;
    padding: 15px;
}
</style>
<div id="banner">
    <div class="title">
        Banner Ad 728x90
    </div>
</div>
```

3. Embedding it into a publisher's site can be done using an iframe as per the below code. This also works with other Ad Servers such as Double Click for Publishers.

```html
<iframe src="//YOUR_DOMAIN/banner.html"></iframe>
```

#### Mobile Banner Ad 320x50

1. Create a file named ```banner.html``` in the AdCampaign/newAdDir/assets directory.
2. Edit ```banner.html``` with your favorite text editor and add code until it looks like the following:

```html
<style type="text/css">
body {
    font-family: Arial, sans-serif;
}

#banner {
    display: block;
    width: 728px;
    height: 90px;
    color: white;
    background-color: black;
    margin: 0 auto;
}

.title {
    text-align: center;
    font-size: 40px;
    padding: 15px;
}
</style>
<div id="banner">
    <div class="title">
        Banner Ad 728x90
    </div>
</div>
```

3. Embedding it into a publisher's site can be done using an iframe as per the below code. This also works with other Ad Servers such as Double Click for Publishers.

```html
<iframe src="//YOUR_DOMAIN/banner.html"></iframe>
```

#### Wallpaper Ad (Advanced)

1. Before creating a wallpaper ad its best to ask the publisher how this can be best achieved using their styles. It's more than likely that there is a JavaScript function that you can use to "hook" into to set the wallpaper / background image.

2. Create a wallpaper image called ```wallpaper.jpg``` and put it in the AdCampaign/newAdDir/assets directory.
	- A jpeg ".jpg" image is usually best.
	- It should have a width of 1024 and wider depending on the publisher's site.

3. Create a JavaScript file called ```wallpaper.js``` in AdCampaign/newAdDir/assets directory.

2. Edit ```wallpaper.js``` with your favorite text editor and add code until it looks like the following:

```js
(function(){
	var top = window.top || window, doc = top.document, body = doc.getElementsByTagName("body")[0];

	body..style.backgroundImage = "//YOUR_DOMAIN/wallpaper.jpg"

})();
```



## Deploying to Production

### Upstart script
To keep the server up, even after restarting add the following upstart script.

- Create a file called "linny.conf" in the ```/etc/init/``` directory. Paste the following into your terminal to create it.
- Edit the paths to ensure that they are correct based upon your configuration.

```bash
cat >/etc/init/linny.conf <<EOL
description "Linny Ad Server"
author      "Simon Chong"

start on filesystem or runlevel [2345]
stop on runlevel [!2345]

# Automatically Respawn:
respawn
respawn limit 99 5

# Max open files are @ 1024 by default. Bit few.
limit nofile 32768 32768

script
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=/go
    export PATH=$PATH:/go/bin

    cd /PATH_TO_YOUR_LINNY_BASE_DIR
    exec linny -serve >> /var/log/linny.log 2>&1
end script
EOL

```


## About

Linny was built to make it easy and accessible for developers and businesses alike to build, host and measure the performance of their web display ads.

## Copyright and Licence

Copyright 2015, Simon Chong. Licensed under the Apache License, Version 2.0 .
