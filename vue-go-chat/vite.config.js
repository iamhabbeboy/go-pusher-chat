import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/

// export default ({ mode }) => {
//   process.env = {...process.env, ...loadEnv(mode)}
// }
export default defineConfig({
  plugins: [vue()]
})