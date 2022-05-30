# Privy Test
This repo for privy test

## How to install
- Clone this repo to your directory and then open in terminal
- Rename **.env.example** to **.env** and change settings with yours 
- Download required packages with command
    > **go get**
- Migrate database with command (change **user**, **password**, **dbname** with yours)
    > **goose -dir .\database\migration\ mysql "user:password@/dbname?parseTime=true" up**
- Start the app with command
    > **go run main.go** 
- Test Cake API, example :
    > localhost:8000/v1/admin/cakes
