package ddd

func GetMainAppContent() []byte {
	return []byte(`
import 'dotenv/config';
import express from 'express';
import { userRouter } from '../interfaces/http/routes/UserRoutes.js';

const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());
app.use('/api', userRouter);
app.listen(port, () => {
  console.log("Server running on port " + port);
});
`)
}
