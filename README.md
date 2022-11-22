##### Msocials server with JWT authentication(RSA) built with GO 
#### Installation Instructions

* Edit config/local.yaml (*Change dbname or you can leave default* **_Should export mysql file for running local server_** )
* `sh run.sh` from root directory (*Generates RSA keys for JWT authentication*)
* `go run main.go` run server from root directory 
* Test URL http://localhost:8080/ping

##### API server will run at: http://localhost:8080

##### API Information


| Endpoint | Method | JSON Format | Note
| --- | --- | --- | --- |
| `/register` | POST | ```{ "email": "johndoe@gmail.com","name" : "John Doe","gender" : "Male","country" : "US","age" : 26,"password": "password"}``` | *Email should be unique*
| `/login` | POST | ```{"email" : "johndoe@gmail.com","password" : "password"}```| Gives Access-Token in header after valid credentials
| `/user` | GET | **_Not required_** | This requires authentication, pass access-token value provided after login as bearer
| `/request` | POST | ```{"email" : "abc@gmail.com"}```| Requires Authentication and send friend request if email found
| `/list` | GET |  **_Not required_**| Requires Authentication and lists friend request user have recieved only
| `/accept` | POST | ```{"uuid" : "c9fd7fa9-e471-4fc2-bbb8-8a6f3ab1fe95" }```|  Accepts friend request if uuid is validated of requested user (Requires Authentication)
| `/friend-list` | GET |  **_Not required_**| Requires Authentication and Loads friends list
| `/post` | POST | ```{"title" : "hello world", "description" : "Goodbye word -Just for description" } ```|  Post message and saves to DB (Requires Authentication)
| `/post` | GET |  **_Not required_**| List posts from friends and self(requied authentication)
| `/post/{post_id} ` | POST |  **_Not required_**| Like a post(requied authentication)
| `/like` | GET |  **_Not required_**| List posts user liked(requied authentication)
| `/post/share/{post_id} ` | POST |  **_Not required_**| Share a post of only friend(requied authentication)
| `/post/timeline` | GET |  **_Not required_**| List user created post and shared post(requied authentication)







