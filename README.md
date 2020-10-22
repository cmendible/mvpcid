![mvpcid](https://github.com/cmendible/mvpcid/workflows/mvpcid/badge.svg)

# mvpcid

Small console application to append a MVP Creator ID to any valid url previously copied to the clipboard.

## Setup

* Download the the latest version from the [releases](https://github.com/cmendible/mvpcid/releases) page.
* Add a new environment variable with name: **MVP_CREATOR_ID** and the **Creator ID** (without the 'WT.mc_id=' prefix) as the value. (i.e. `export MVP_CREATOR_ID=AZ-MVP-5002618`)

## Usage

### Windows

* Run the command line program.
* Copy a valid url.
* Press `Ctrl+Shift+M` to add your MVP Creator ID to the url in the clipboard.
* Paste the new url with your MVP Creator ID wherever you need it.
* Press `Ctrl+Shift+Q` to close the program at any time.

### Linux & Mac OS

* Copy a valid url.
* Run the command line program.
* Paste the new url with your MVP Creator ID wherever you need it.

Note: On Linux 'xclip' or 'xsel' are required.