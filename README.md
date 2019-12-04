RevatureGo
========

RevatureGo will replace the portfolio system on RevaturePro by making a better interface and automating processes.


Look how easy it is to use:

- Ask slackbot to make a portfolio


Features
--------

- List of features


Installation
------------

Install RevatureGo by running:

- Get to the slackbot?
- Clone from git?
- Dockerhub?
- Azure?
- Kubernetes

testing 

RevatureGo-API
========

- Hosted on localhost:8080.
- Presents the user with a form to fill out their portfolio.
- Stores the portfolio information into structs.
- Saves said information as a json file localy under portfolios.
- Allows editing of existing portfolios.
- Sends portfolios to an AWS.


Producion notes
--------

## README for API-HTMLFORM-JSON
### 11/21
- Combined Nathan's API-HTMLFORM with Garner's HTMLFORM-JSON code. 
- User runs main.go , opens localhost:8080, goes to http://localhost:8080/html/portfolioform.html
- User inputs information, upon clicking "Submit Portfolio" button, portfolio is saved as a json file.
- Json file is named using the user's inputted fullname variable.

### 11/22
- Added a "Status" field to the portfolio struct, so that administrators can check to see if the portfolio
    has been accepted, denied, or unchecked(default value set).
- This "Status" field can be viewed and edited by an administrator using the Approve-Deny-Portfolio-Check code. 
- Program now replaces black spaces with underscore(_). This was to help with the programming of the Approve-Deny-Portfolio-Check code. 

RevatureGo-Client
========

- Hosted on localhost:7004.
- Downloads a portfolio using downloads.html.
- Creates portfolio template on page.html.
- Creates a blank file named after full path.
- Inputs portfolio information into the file.
- Allows you to approve or deny the portfolio.
