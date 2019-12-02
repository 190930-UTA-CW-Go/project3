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

## README for APPROVE-DENY-PORTFOLIO-CHECK

### 11/22
- This code is for administrators to check user profiles, and approve/deny them. By default, user
    profiles have their status set to "UNCHECKED". 
- This code creates a new identical file with the updated status, and discards the old file. 
- HOW TO USE:
    - go run main.go
    - Prompt will request for a file name (example.json)
    - Prompt will display name of user's profile as well as status(UNCHECKED, DENIED, or APPROVED)
    - Prompt will ask for action (Change status to APPROVED or DENIED. Or exit without any changes)
    - Program terminates. 

### 11/26
- User can now edit profiles. 
- Implemented HTML navigation functionality. 
- Implemented ability to automatically scp "push" the user portfolio to the amazon server. 
    - **THE "USER" IS HARDCODED RIGHT NOW as "localuser".**
        - Also makes sure the folder for "localuser" exists on the amazon server.
        - If it already exists, nothing happens.
    - In moving files: 
        - When "fileA" exists on amazon server, and a NEW "fileA" is pushed onto the amazon server,
            **The old "fileA" will cease to exist.**
        - A function to 'pull' the file from amazon for the user to read has not been implemented
            - May be an issue if an admin edits the file(such as updating its STATUS field). This 
                file on the amazon server will never be seen by the user, and will only be overwritten. ("I have no access or business with the amazon server code" - Garner)
    - **YOU MUST HAVE THE "rego.pem" KEY AND ALL JSON FILES IN THE SAME FOLDER AS main.go**










## (Misc Code)
- ssh -i rego.pem ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com
- Public IP: 18.188.174.65