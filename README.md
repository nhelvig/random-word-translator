# random-word-translator
This is just a simple webapp that will translate from a given list of words to a language of your choosing.

## Running Locally

### Prerequisites
-Clone this repository onto your machine with `git clone https://github.com/nhelvig/random-word-translator.git`
-You must have [docker](https://docs.docker.com/engine/installation/) installed on your machine.

### Run the container
Run the `startContainer.sh` script. This will both build the images and start the container in the background.  
_Note: You may need to give it execute permissions with `chmod +x /path/to/script/startContainer.sh`_  

The script takes two options:  
`-p` to specify which port you would like (by default it uses 9001)  
`-i no` to declare if you do not want to build the images again (it does by default)  
By default, it will use port 9001. You can pass another port as a parameter to the script (i.e. `./startContainer.sh 12345`)

### Using the application
Hit the endpoint `http://localhost:9001/word/{desired language code}` (or with whatever port you passed in)  
-If you pass in a valid language code `http://localhost:9001/word/pt` (for Portuguese), you will see the translation printed in the broswer.  
-If you pass in a valid language code, you will be shown the valid codes you can use.

