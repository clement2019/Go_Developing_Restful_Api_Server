# Go_Developing_Restful_Api_web service
This projects demonstrates the ability to develop restful Api web service in Go from scratch creating the struct, interfaces and building the RestApi web services all done using Go language


//use the below code to run data to add newalbum data to the the existing album records

curl http://localhost:8084/albums \
   --include \
    --header "Content-Type: application/json" \
   --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
