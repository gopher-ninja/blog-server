# Golang Blog server 
The blog-server backend is developed using `gonic/gin` and `ent.io` ORM library. This backend server has 5 APIs as following:

1. Add blog API
2. Get single blog API
3. Get all blogs API
4. Update blog API
5. Delete Blog API

#### Assumptions
* Only 1 admin user is there, since we do not have any process for sign-in.
* Due to less functionalities only 1 service is developed. We can distribute the load using 2 service( 1 to handle API requests and response & another service to handle blog related operations or business logics).

## Installation steps

#### Pre-requisites
* Golang(1.19.5)
* Docker
* make

##### Manual steps
To run binary
```
git clone https://github.com/gopher-ninja/blog-server
make
# Please update following variable as per you postgres server details
export DB_HOST=192.168.29.85
export DB_PORT=54320
export DB_NAME=postgres
export DB_USER="user"
export DB_PASSWORD="admin"
export GIN_MODE=release
./blog-server
```

To run as a docker container
```
git clone https://github.com/gopher-ninja/blog-server
make docker
docker run -d -p 0.0.0.0:5432:5432 -e POSTGRES_USER='user' -e POSTGRES_PASSWORD='admin' postgres:latest
docker run -d -p 0.0.0.0:8080:8080 -e DB_HOST=192.168.29.85 -e DB_PORT=54320 -e DB_NAME=postgres -e DB_USER="user" -e DB_PASSWORD="admin" -e GIN_MODE=release gopher-ninja/blog-server```
```
To run without cloning
```json
docker run -d -p 0.0.0.0:5432:5432 -e POSTGRES_USER='user' -e POSTGRES_PASSWORD='admin' postgres:latest
docker run -d -p 0.0.0.0:8080:8080 -e DB_HOST=192.168.29.85 -e DB_PORT=54320 -e DB_NAME=postgres -e DB_USER="user" -e DB_PASSWORD="admin" -e GIN_MODE=release gopher-ninja/blog-server```
```

## API Specification
### Add Blog
  
**POST** http://127.0.0.1:8080/api/v1/blog

*Authentication* : Basic Auth
```json
username: admin
password: Admin@123
```
  
  *Request Body*
  ```
  {
    "blogTitle":"blog3",
    "blogContent":"Here is the content"
}
  ```
*Response*
```json
{
  "blogDetail": {
    "id": 3,
    "blogId": "70c40064-63a4-47cf-bd53-0d9a3e8b0df1",
    "blogTitle": "blog3",
    "blogContent": "Here is the content"
  }
}
```

### Get Single Blog

**GET** http://127.0.0.1:8080/api/v1/blog/{blogId}

*Response*
```json
{
  "blogDetail": {
    "id": 3,
    "blogId": "70c40064-63a4-47cf-bd53-0d9a3e8b0df1",
    "blogTitle": "blog3",
    "blogContent": "Here is the content"
  }
}
```

### Get All Blogs

**GET** http://127.0.0.1:8080/api/v1/blogs

*Response*
```json
{
  "blogs": 
  [
    {
    "id": 3,
    "blogId": "70c40064-63a4-47cf-bd53-0d9a3e8b0df1",
    "blogTitle": "blog3",
    "blogContent": "Here is the content"
  },
    {
      "id": 4,
      "blogId": "34kkjklj-63a4-47cf-bd53-0d9a3e8b0df1",
      "blogTitle": "blog5",
      "blogContent": "Here is the content for blog 5"
    }
  ]
}
```

### Update Blog

**POST** http://127.0.0.1:8080/api/v1/blog/{BlogId}

*Authentication* : Basic Auth
```json
username: admin
password: Admin@123
```

*Request Body*
  ```
  {
    "blogTitle":"blog3",
    "blogContent":"Here is the content"
}
  ```
*Response Body*
```json
{
  "blogDetail": {
    "id": 3,
    "blogId": "70c40064-63a4-47cf-bd53-0d9a3e8b0df1",
    "blogTitle": "blog3",
    "blogContent": "Here is the content"
  }
}
```


### Delete Blog
**DELETE** http://127.0.0.1:8080/api/v1/blog/{BlogId}

*Authentication* : Basic Auth
```json
username: admin
password: Admin@123
```

*Response Body*
```json
{
  "msg": "Blog deleted successfully"
}
```