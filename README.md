# USERS
### 1) userGroup.Post("/", controllers.CreateUser)
```console
POST http://localhost:3000/users
{
    "name": "Naushad",
    "email": "naushad@example.com",
    "balance": 100
}

```

### 2) userGroup.Get("/", controllers.GetAllUsers)
```console
GET http://localhost:8080/users
```

### 3) userGroup.Post("/transfer", controllers.TransferFunds)
```console
PUT http://localhost:3000/users/transfer
{
    "from_user_id": 1,
    "to_user_id": 2,
    "amount": 20
}
```

### 4) userGroup.Get("/with-posts", controllers.GetUsersWithPosts)
```console
GET http://localhost:8080/users/with-posts
```

### 5) userGroup.Get("/post-counts", controllers.GetUserPostCounts)
```console
GET http://localhost:8080/users/post-counts
```

### 6) userGroup.Get("/recent-posts", controllers.GetUsersWithRecentPosts)
```console
GET http://localhost:8080/users/recent-posts
```

# POSTS
### 1) postRoutes.POST("/", controllers.CreatePost)
```console
POST http://localhost:3000/posts
{
    "title": "My First Post",
    "content": "This is the content of the post",
    "user_id": 2
}
```

### 2) postRoutes.GET("/:id", controllers.GetPost) , postRoutes.GET("/", controllers.GetAllPosts)
```console
GET http://localhost:3000/posts
```


