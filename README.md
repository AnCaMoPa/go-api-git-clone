# api-git-clone
An API in go with fiber that will clone any number of git repositories you want in parrallel with goroutines. You need to make a post request with the following json:

{
  "repositorios": [
      {
        "url" : "git@github.com:AnCaMoPa/go-crud-api-NoSQL.git"
      },
      {
        "url" : "git@github.com:AnCaMoPa/go-crud-api-NoSQL.git"
      }
    ]
}

To TEST the API you cand use Postman or Thunder Client in Visual Studio Code to make the post request.

   
    
