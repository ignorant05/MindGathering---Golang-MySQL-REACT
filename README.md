# MindGathering

## Description

A simple DIY full-stack webapp built for learning purposes.
The web app is simply an imitation for blog post web app.

---

##   Stack

<details>
<summary><strong>Frontend</strong></summary>
<br>

- React v20.
- Tailwindcss.
- Nginx.

</details>

<details>
<summary><strong>Backend</strong></summary>
<br>

- Golang.
- Goose.
- MySQL.
- Docker + docker compose. 
- Supports user authentication using jwt appoach.

</details> 

> *Note:* You can check the .env file if you want to check or change the expiration date, the secret key (for both access and refresh tokens) & db credentials.

---

##   Project layout

<details>
<summary><strong>MindGathering</strong></summary>
<br>

```text.
.
│----backend
│   ├── cmd
│   │   ├── app
│   │   │   └── main.go
│   │   └── server
│   │       └── server.go
│   ├── config
│   │   └── db
│   │       └── conn.go
│   ├── db
│   │   ├── migration
│   │   │   ├── 20250618185346_refresh_tokens_schema.sql
│   │   │   ├── 20250618185352_users_schema.sql
│   │   │   ├── 20250620160301_blogs_schema.sql
│   │   │   ├── 20250620160310_comments_schema.sql
│   │   │   └── 20250627143823_images_schema.sql
│   │   ├── queries
│   │   │   ├── auth.queries.sql
│   │   │   └── user.queries.sql
│   │   └── sqlc.yaml
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── handler
│   │   ├── authentication.handler.go
│   │   ├── handler.go
│   │   └── user.handler.go
│   └── internal
│   ├── db
│   │   ├── auth.queries.sql.go
│   │   ├── db.go
│   │   ├── models.go
│   │   ├── querier.go
│   │   └── user.queries.sql.go
│   ├── helpers
│   │   ├── hash.pwd.go
│   │   └── load.env.go
│   ├── middleware
│   │   └── authentication.go
│   ├── routes
│   │   └── routes.go
│   └── utils
│       └── jwt.utils.go
├── docker-compose.yaml
├── frontend
│   ├── Dockerfile
│   ├── eslint.config.js
│   ├── index.html
│   ├── nginx
│   │   └── nginx.conf
│   ├── package.json
│   ├── package-lock.json
│   ├── postcss.config.js
│   ├── README.md
│   ├── src
│   │   ├── app
│   │   │   ├── api
│   │   │   │   └── apiSlice.jsx
│   │   │   └── store.jsx
│   │   ├── App.css
│   │   ├── App.jsx
│   │   ├── assets
│   │   │   ├── default_picture.png
│   │   │   └── logo.jpg
│   │   ├── components
│   │   │   ├── BlogCard.jsx
│   │   │   ├── CommentCard.jsx
│   │   │   ├── shared
│   │   │   │   └── Navbar.jsx
│   │   │   └── UserCard.jsx
│   │   ├── index.css
│   │   ├── main.jsx
│   │   ├── redux
│   │   │   └── slices
│   │   │       ├── authApiSlice.jsx
│   │   │       ├── authSlice.jsx
│   │   │       ├── commentApiSlice.jsx
│   │   │       ├── commentsSlice.jsx
│   │   │       ├── draftApiSlice.jsx
│   │   │       ├── draftSlice.jsx
│   │   │       ├── pageSlice.jsx
│   │   │       ├── readBlogSlice.jsx
│   │   │       └── usersApiSlice.jsx
│   │   └── screens
│   │       ├── BlogViewScreen.jsx
│   │       ├── CreateBlogScreen.jsx
│   │       ├── EditBlogScreen.jsx
│   │       ├── EditProfileScreen.jsx
│   │       ├── HomeScreen.jsx
│   │       ├── LoginScreen.jsx
│   │       ├── MyBlogsListScreen.jsx
│   │       ├── MyCommentsListScreen.jsx
│   │       ├── RegistrationScreen.jsx
│   │       ├── SeeOtherUserProfileScreen.jsx
│   │       └── SeeProfileScreen.jsx
│   ├── tailwind.config.js
│   └── vite.config.js
└── README.md

```

</details>


---

##   Installation 

```bash
# Clone the project or download the it's zip file
$ git clone https://github.com/ignorant05/MindGathering---Golang-MySQL-REACT.git
```

> *Note:* docker and docker compose are required here otherwise you need to install all the dependencies for both backend and frontend, so make sure docker is installed, enabled and started. 

---

##   Deployment

just run :

```bash
$ sudo docker-compose up --build
```

> *Note:* In-case you want to add/remove/change anything in the *db/queries* make sure that you have [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) to run the following command while you are in *db/* directory:

```bash
$ sqlc generate 
```
> *Note:* In-case you want change migrations, go check the [goose](https://github.com/pressly/goose) github repository.

Optionally you can build and clean it before you re-deploy it again: 

```bash
$ sudo docker-compose down 
$ sudo docker-compose build --no-cache
```

