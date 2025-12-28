package shared

func GetTsConfigJsonContent() []byte {
	return []byte(`{
  "compilerOptions": {
    /* Output */
    "outDir": "dist",
    "rootDir": "src", // Helps structure input/output clearly

    /* Module system */
    "module": "nodenext",
    "target": "esnext",
    "moduleResolution": "nodenext",

    /* Code safety */
    "strict": true,
    "noUncheckedIndexedAccess": true,
    "exactOptionalPropertyTypes": true,
    // "verbatimModuleSyntax": true,
    "isolatedModules": true,
    "noUncheckedSideEffectImports": true,
    "moduleDetection": "force",
    "skipLibCheck": true,
    "experimentalDecorators": true,
    "emitDecoratorMetadata": true,

    /* Interop */
    "esModuleInterop": true, // Enables smooth import express from 'express'
    "allowSyntheticDefaultImports": true, // Allows default import of CJS modules
    "allowJs": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist"]
}`)
}
