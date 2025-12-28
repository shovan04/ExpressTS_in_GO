package shared

func GetGitIgnoreContent() []byte {
	return []byte(`// Node / Package Managers
node_modules/
npm-debug.log*
yarn-debug.log*
yarn-error.log*
pnpm-debug.log*

// Build output
dist/
build/
out/

// TypeScript
*.tsbuildinfo

// Environment variables
.env
.env.*
!.env.example

// Logs
logs/
*.log

// Coverage / Tests
coverage/
.nyc_output/

// Runtime data
pids/
*.pid
*.seed
*.pid.lock

// OS / Editor junk
.DS_Store
Thumbs.db

// IDEs
.vscode/
!.vscode/extensions.json
.idea/
*.swp
*.swo

// Misc
.cache/
tmp/
temp/
`)}