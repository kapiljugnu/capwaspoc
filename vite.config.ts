import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
  root: './src',
  build: {
    outDir: '../dist',
    minify: false,
    emptyOutDir: true,
  },
  plugins: [{
    name: 'watch-wasm',
    handleHotUpdate({ file, server }) {
      if (file.endsWith('.wasm')) {
        server.ws.send({ type: "full-reload", path: '*' })
      }
    }
  }],
});
