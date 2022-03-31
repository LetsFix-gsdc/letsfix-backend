## Requirements
Have Go installed on your system, following the instructions at https://go.dev/doc/install

## Running
1. Clone this repo and cd to root
2. Fill in config.env with your database connection details, or export those values to environment variables
3. When starting with a new database for the first time, run ```go run main.go seed``` to seed recycler data from file
4. Subsequently, run ```go run main.go``` in terminal