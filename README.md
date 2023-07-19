# API-GIT-CLONE

## INTRODUCTION
An API in go with fiber that will clone any number of git repositories you want in parrallel with the goroutines you choose and in the path you want. 

## PRE REQUISITES
To work, you need to have ```git``` and ```go``` install in your computer.

## JSON BODY STRUCTURE
Once you have the app running, you need to make a post request with the following json:

```json
{
  "goroutines": "2", //Number of operations you want in parrallel
  "path": "./", //Path in which you cloned repositories will be
  "time_out":"10s", //Max time the operations will be in executing before its get cancel
  "repositories": [ //List of the repositories you want to clone
      {
        "url" : "git@github.com:AnCaMoPa/go-api-git-clone.git"
      },
      {
        "url" : "git@github.com:AnCaMoPa/full-stack-test.git"
      }
    ]
}
```

To test the API you cand use Postman or Thunder Client in Visual Studio Code to make the post request.

   
    
