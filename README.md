# Go_Developing_Restful_Api_Server
This projects demonstrates the ability to develop restful Api in Go from scratch creating the struct, interfaces and building the RestApi server all don using Go language


//use the below code to run data to add newalbum data to the the existing album records

curl http://localhost:8080/albums \
   --include \
    --header "Content-Type: application/json" \
   --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
