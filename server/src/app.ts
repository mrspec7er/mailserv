import { Server } from "./smtp-listener";
import env from "dotenv";

env.config();

const smtpServer = Server;

const port = process.env.PORT;
const host = process.env.HOST;

smtpServer.listen(Number(port), host);
