import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

const AUTH_SERVER_URI = "http://localhost:3000";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5000,
    proxy: {
      "/api/v1": {
        target: AUTH_SERVER_URI,
        changeOrigin: true,
      },
    },
  },
});
