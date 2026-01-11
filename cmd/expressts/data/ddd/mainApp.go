package ddd

func GetMainAppContent() []byte {
	return []byte(`import "reflect-metadata";
import express from "express";
import dotenv from "dotenv";
import mainRouter from "../interfaces/http/routes/index.js";
import cookieParser from "cookie-parser";
import cors from "cors";
import GlobalErrorHandler from "../interfaces/http/middlewares/GlobalErrorHandler.js";

dotenv.config();

const PORT = process.env.PORT || 8080;
const HOST = process.env.HOST || "localhost";
const BASE_API_PATH = process.env.BASE_API_PATH || "/api/v1";
const COOKIE_SECRET = process.env.COOKIE_SECRET || "default_cookie_secret";

const server = express();

server.use(cors())
server.use(express.json());
server.use(express.urlencoded({ extended: true }));
server.use(cookieParser(COOKIE_SECRET));

// Registering Main Router
server.use(BASE_API_PATH, mainRouter);
// Global Error Handler Middleware
server.use(GlobalErrorHandler);

server.listen(Number(PORT), HOST, () => {
    console.log(
        "Server is running at http://" +
        HOST + ":" + PORT + BASE_API_PATH
    );
});
`)
}
