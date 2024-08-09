# Ascii Art Web

## Description

This is a simple web application that converts text to ascii art. It is built using golang, html and css.

## Usage

To run the application, you need to have golang installed on your machine. You can download it from [here](https://golang.org/dl/).

After installing golang, you can run the following commands to run the application:

```bash
go run .
```

This will start the server on port 8080. You can access the application by visiting [http://localhost:8080](http://localhost:8080) in your browser.

## Implementation

The application is built using golang and html. The golang server serves the html files and handles the conversion of text to ascii art.

The html file contains a form that takes in the text input and the banner radio and sends it to the server for conversion. The converted ascii art is then displayed on /ascii_art route.

The ascii art is generated using the [asciiArtGen](./utils/asciiArtGen.go) function. This function takes in a string and a banner and returns the a string containing the ascii art.

![Screenshot 1](images/s2.png)
![Screenshot 2](images/s3.png)
![Screenshot 3](images/s4.png)
![Screenshot 3](images/s5.png)




## Authors

Khalid El Amrani [(kelamrani)](https://learn.zone01oujda.ma/git/kelamrani)<br>
Mustapha Lbahja [(mlbahja)](https://learn.zone01oujda.ma/git/mlbahja)<br>
Ali ES-SLIMANI [(aesslima)](https://learn.zone01oujda.ma/git/aesslima)<br>