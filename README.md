
# 💬 Real-time Chat App with Go (Fiber + WebSocket)

A real-time chat application built using **Golang**, **Fiber**, and **WebSocket**.

## 🚀 Tech Stack

- **Golang** — Backend language
- **Fiber** — Fast HTTP web framework
- **WebSocket** — Real-time communication
- **MySql** — Database SQL

---

## ✨ Features

- ✅ Real-time Chat (WebSocket)
- ✅ Register & Login with JWT
- ✅ Add, Update, and Delete Friends
- ✅ Send & Delete Chat Messages
- ✅ Get User Information

---

## 🔐 AUTHENTICATION

### 🔸 Register

- **URL**: `POST http://localhost:3000/api/users/register`

#### 🟦 Request:
```json
{
  "username": "sacho",
  "password": "12345678"
}
```

#### 🟩 Response:
```json
{
  "data": {
    "username": "sacho",
    "token": "<JWT_TOKEN>"
  },
  "message": "Success Register",
  "success": true
}
```

---

### 🔸 Login

- **URL**: `POST http://localhost:3000/api/users/login`

#### 🟦 Request:
```json
{
  "username": "sacho",
  "password": "12345678"
}
```

#### 🟩 Response:
```json
{
  "data": {
    "username": "sacho",
    "token": "<JWT_TOKEN>"
  },
  "message": "Success Login",
  "success": true
}
```

---

## 👤 USER

### 🔸 Update User

- **URL**: `PATCH http://localhost:3000/api/auth/users/update`

#### 🟦 Request:
```json
{
  "name": "jadon",
  "username": "sacho"
}
```

#### 🟩 Response:
```json
{
  "data": true,
  "message": "Success UpdateUser",
  "success": true
}
```

---

### 🔸 Get Current User

- **URL**: `GET http://localhost:3000/api/auth/users/me`

#### 🟩 Response:
```json
{
  "data": {
    "id": 1,
    "unique_number": 50417333,
    "name": "jadon",
    "username": "sacho",
    "password": "",
    "last_login": null
  },
  "message": "Success Get User",
  "success": true
}
```

---

## 👥 FRIEND

### 🔸 Get Friend by ID

- **URL**: `GET http://localhost:3000/api/auth/friend/13`

#### 🟩 Response:
```json
{
  "data": {
    "id": 13,
    "name": "maryy konco plek",
    "friend": 2,
    "user_id": 1,
    "room_id": 9
  },
  "message": "success get friend data",
  "success": true
}
```

---

### 🔸 Add Friend

- **URL**: `POST http://localhost:3000/api/auth/friend/add`

#### 🟦 Request:
```json
{
  "name": "maryy konco plek",
  "unique_number": 99596897
}
```

#### 🟩 Response:
```json
{
  "data": 200,
  "message": "success add friend user",
  "success": true
}
```

---

### 🔸 Update Friend

- **URL**: `PATCH http://localhost:3000/api/auth/friend/update/1`

#### 🟦 Request:
```json
{
  "name": "mary sigma"
}
```

#### 🟩 Response:
```json
{
  "data": 200,
  "message": "success update friend user",
  "success": true
}
```

---

### 🔸 Delete Friend

- **URL**: `DELETE http://localhost:3000/api/auth/friend/delete`

#### 🟦 Request:
```json
{
  "id": 1
}
```

---

### 🔸 Get All Friends

- **URL**: `GET http://localhost:3000/api/auth/friends`

#### 🟩 Response:
```json
{
  "data": [
    {
      "id": 13,
      "room_id": 9,
      "name": "maryy konco plek",
      "unique_number": 99596897
    }
  ],
  "message": "success get friends user",
  "success": true
}
```

---

## 💬 CHAT

### 🔸 Get Chats by Room ID

- **URL**: `GET http://localhost:3000/api/auth/chats/4`

#### 🟩 Response:
```json
{
  "data": [
    {
      "id": 21,
      "chat": "test",
      "room_id": 9,
      "user_id": 1
    }
  ],
  "message": "success",
  "success": true
}
```

---

### 🔸 Send Chat Message

- **URL**: `POST http://localhost:3000/api/auth/chats/add`

#### 🟦 Request:
```json
{
  "room_id": 9,
  "chat": "test"
}
```

#### 🟩 Response:
```json
{
  "data": {
    "id": 21,
    "chat": "test",
    "room_id": 9,
    "user_id": 1
  },
  "message": "success",
  "success": true
}
```

---

### 🔸 Delete Chat

- **URL**: `DELETE http://localhost:3000/api/auth/chats/delete/1`

#### 🟩 Response:
```json
{
  "data": null,
  "message": "success",
  "success": true
}
```

---

## 🔌 WEBSOCKET

### Connect via WebSocket

- **URL**:
```
ws://localhost:3000/ws/chat?room_id=8
```

All users connected to the same `room_id` will receive real-time messages instantly.

---

## 🗃️ Database Structure

📥 SQL Schema:
```sql
CREATE TABLE `users`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `unique_number` SMALLINT NOT NULL,
    `name` TEXT NOT NULL,
    `username` TEXT NOT NULL,
    `password` TEXT NOT NULL,
    `last_login` DATETIME NOT NULL
);
ALTER TABLE
    `users` ADD UNIQUE `users_unique_number_unique`(`unique_number`);
CREATE TABLE `chats`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `text` TEXT NOT NULL,
    `room_id` BIGINT NOT NULL,
    `user_id` BIGINT NOT NULL
);
CREATE TABLE `room`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
);
CREATE TABLE `friends`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` BIGINT NOT NULL,
    `friend` BIGINT NOT NULL,
    `user_id` BIGINT NOT NULL,
    `room_id` BIGINT NOT NULL
);
ALTER TABLE
    `chats` ADD CONSTRAINT `chats_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `users`(`id`);
ALTER TABLE
    `friends` ADD CONSTRAINT `friends_room_id_foreign` FOREIGN KEY(`room_id`) REFERENCES `room`(`id`);
ALTER TABLE
    `chats` ADD CONSTRAINT `chats_room_id_foreign` FOREIGN KEY(`room_id`) REFERENCES `room`(`id`);
ALTER TABLE
    `friends` ADD CONSTRAINT `friends_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `users`(`id`);
ALTER TABLE
    `friends` ADD CONSTRAINT `friends_friend_foreign` FOREIGN KEY(`friend`) REFERENCES `users`(`id`);
```

🧭 ERD Visual Link:
[View on DrawSQL](https://drawsql.app/teams/devmare/diagrams/web-chat)
