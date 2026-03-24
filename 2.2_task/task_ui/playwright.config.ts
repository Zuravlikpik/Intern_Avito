import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './tests',

  use: {
    baseURL: 'https://cerulean-praline-8e5aa6.netlify.app',
    headless: true,
    viewport: { width: 1280, height: 720 }
  },

  reporter: [['html']]
});