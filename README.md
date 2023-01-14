# API Document for Blog server 

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
  "data": {
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
  "data": {
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
  "data": 
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
  "data": {
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
  "data": "Blog deleted successfully"
}
```