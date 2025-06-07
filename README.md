# Stock Portfolio Tracker ArnavStocks

https://github.com/user-attachments/assets/8f407503-9370-4ea3-a442-9cd14777bc03

**Updated(RabbitMQ Implementation Server Side Logging)**

https://github.com/user-attachments/assets/fa2d4e0b-11f8-471d-a8ae-3df196c1818d



Updates() : 
Implemented JWT Auth at Login
React SignUp and Login Page Setup
Backend Rate Limiter implemented with max 4 bursts of requests with consume rate of 2

## Description

This project is a stock portfolio tracker application. It leverages React, Chakra UI, and MongoDB for the frontend, while using Golang with Gin and MySQL for the backend. The application provides a comprehensive platform for managing and tracking stock investments, allowing users to monitor trending stocks, add stocks to their watchlist, view investment performance, and access the latest financial news.

## Features

- **Trending Stocks**: View the top gaining and losing stocks.
- **Watchlist**: Add and remove stocks from a personal watchlist.
- **Investment Management**: Track buying and selling of stocks, with real-time balance updates.
- **Financial News**: Stay updated with the latest financial news.

## Technologies

### Frontend
- React
- Chakra UI
- Vite

### Backend
- Golang
- Gin Framework
- MySQL



## Prerequisites

- Node.js
- MongoDB
- Golang
- MySQL

## Installation

### Frontend

1. **Clone the repository**
    ```bash
    git clone https://github.com/arnav45656/GolangStocksPortfolio.git
    cd GolangStocksPortfolio
    ```

2. **Install dependencies**
    ```bash
    npm install
    ```

3. **Run the application**
    ```bash
    npm run dev
    ```

### Backend

1. **Navigate to the backend directory**
    ```bash
    cd backend
    ```

2. **Install dependencies**
    ```bash
    go get -u github.com/gin-gonic/gin
    go get -u github.com/jinzhu/gorm
    go get -u github.com/jinzhu/gorm/dialects/mysql
    ```

3. **Set up MySQL Database**
    Create a database in MySQL named `stock_tracker`.

4. **Create a `.env` file in the backend directory**
    ```env
    DB_USER=your_mysql_username
    DB_PASSWORD=your_mysql_password
    DB_NAME=stock_tracker
    DB_HOST=localhost
    DB_PORT=3306
    ```

5. **Run the backend server**
    ```bash
    go run main.go
    ```

## Project Structure

### Frontend

stock-portfolio-tracker/
├── public/
│ ├── favicon.ico
│ └── index.html
├── src/
│ ├── api/
│ │ ├── api.js
│ ├── assets/
│ ├── components/
│ │ ├── Layout/
│ │ │ ├── Layout.jsx
│ │ │ ├── LayoutComponents/
│ │ │ │ ├── Header.jsx
│ │ │ │ └── Sidebar.jsx
│ │ ├── Stocks/
│ │ │ ├── StocksComponents/
│ │ │ │ ├── StockList.jsx
│ │ │ │ └── StockItem.jsx
│ │ ├── Watchlist/
│ │ │ ├── Watchlist.jsx
│ │ │ ├── WatchlistItem.jsx
│ ├── hooks/
│ │ ├── useFetch.js
│ ├── pages/
│ │ ├── HomePage.jsx
│ │ ├── InvestmentPage.jsx
│ │ ├── NewsPage.jsx
│ ├── App.jsx
│ ├── index.jsx
├── package.json
└── vite.config.js




