# Veraison-Testing

Testing the functionallity of Veraison

# Prerequisites

- You wil need GO language installed in your system for more info on how to install look [here][1]
- Besides your ``.go`` file you will need the package handler ``go.mod`` in the same folder to verify the imports
- I highly suggest the usage of **Vs Code** and the extension of **GO by team Google** which autocompletes the import functions 

[1]: https://go.dev/doc/install
[2]: https://github.com/veraison/go-cose

# Importing the Veraison verifier
Inside your folder of project execute `` go get github.com/veraison/go-cose ``, consider trying ``sudo`` in case of access denied.
Code can be found in the repo or in the original [Veraison project][2] 

# Executing and verifying a message
To verify a message just open a terminal and run ``go build test.go`` then you will see the ``test.exe`` which you can simply run as ``./test``
Also, feel free to modify the message in the code to test the verification of other messages but "Hello World".

LICENSE by Verasion Project
