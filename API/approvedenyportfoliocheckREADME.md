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