package shared

func GetPackageJsonContent(prjName string, prjDesc string) []byte {
	return []byte(`{
  "name": "` + prjName + `",
  "version": "1.0.0",
  "description": "` + prjDesc + `",
  "main": "dist/bin/app.js",
  "type": "module",
  "scripts": {
    "dev:build": "tsc --build -w",
    "dev:start": "nodemon dist/bin/app.js",
    "dev": "pnpm dev:build & pnpm dev:start"
  },
  "packageManager": "pnpm@10.19.0",
  "dependencies": {
    "class-transformer": "*",
    "class-validator": "*",
    "cookie-parser": "*",
    "cors": "*",
    "dotenv": "*",
    "express": "*"
  },
  "devDependencies": {
    "@types/cookie-parser": "*",
    "@types/cors": "*",
    "@types/express": "*",
    "nodemon": "*",
    "reflect-metadata": "*",
    "typescript": "*"
  }
}`)
}
