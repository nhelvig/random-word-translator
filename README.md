# random-word-translator
This is just a simple webapp that will translate from a given list of words to a language of your choosing.

## Running Locally

### Prerequisites
-Clone this repository onto your machine with `git clone https://github.com/nhelvig/random-word-translator.git`
-You must have [docker](https://docs.docker.com/engine/installation/) installed on your machine.

### Build the image
1. Navigate to the directory that contains the Dockerfile
2. Run `docker build -t random_word_translator .`

### Run the container
Run the `startContainer.sh` script.  
_Note: You may need to give it execute permissions with `chmod +x /path/to/script/startContainer.sh`_  

By default, it will use port 9001. You can pass another port as a parameter to the script (i.e. `./startContainer.sh 12345`)

### Using the application
Hit the endpoint `http://localhost:9001/word/{desired language code}`  
-If you pass in a valid language code `http://localhost:9001/word/pt` (for Portuguese), you will see the translation printed in the broswer.  
-If you pass in a valid language code, you will be shown the valid codes you can use.

