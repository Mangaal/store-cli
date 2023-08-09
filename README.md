# Store CLI
Store is a CLI tool developed using Cobra and the Go programming language. It interacts with the API exposed by the project Text Store to perform various file-related operations.

# Commands
```
  store files: Update and create files on the server.
 
  store ls: List all the files uploaded to the server.
 
  store wc: Count the total number of words across all the uploaded files.
 
  store freq-words: Get the most frequently used words from the uploaded documents.
```


#Installation

   Set up environment variables:
 
### Place where uploaded files will be stored
```
   export DATA_DIR="/tmp/da"
```

### API URL for the Text Store project

```
   export STORE_URL="http://localhost:8080"
```

Clone the repository:
```
   git clone https://github.com/your-username/store-cli.git

   cd store-cli
```
Run the installation script:
```
   ./install.sh
```

# Usage

### To update or create files on the server, RUN:
```
   store files filenames....
```
### To list all files uploaded to the server, RUN:

```
   store ls
````
### To list all files uploaded to the server, RUN:

```
   store rm filenames....
````
### To count the total number of words across all uploaded files, RUN:
```
   store wc
```
### To get the most frequently used words from the uploaded documents, RUN:
```
   store freq-words
```
### To display the version:
```
To test the CLI, run:

  store -v
