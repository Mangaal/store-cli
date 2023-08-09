# Store CLI
  Store is a CLI tool developed using Cobra and the Go programming language. It interacts with the API exposed by the project Text Store to perform various file-related operations.

# Commands
```
  store files: Update and create files on the server.
 
  store ls: List all the files uploaded to the server.
 
  store wc: Count the total number of words across all the uploaded files.
 
  store freq-words: Get the most frequently used words from the uploaded documents.
```

# File Upload and Management

  When a new file is uploaded, a key-value JSON data is stored in a defined path. The key is the name of the file, and the value is a string representing the calculated SHA-256 hash of the file 
  content. This hash allows for content comparison between files. If the content of two files is the same, their hash values will be the same.
```
* If a file's name and content are the same as an already uploaded file, it won't be processed.

* If a file's content is the same as an already uploaded file, but the name is different, the file will be renamed on the server.
```
  
##### For each upload operation, the JSON file will be checked and updated with the latest information.



# Installation

   Set up environment variables:
 
### Place where uploaded files will be stored
```
   export DATA_DIR="/tmp/da"
```

### API URL for the Text Store project

```
   export STORE_URL="localhost:8080"
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
