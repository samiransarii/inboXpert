import { defineConfig, mergeConfig, UserConfig } from "vite";
import react from "@vitejs/plugin-react";
// import { viteStaticCopy } from "vite-plugin-static-copy";
import tailwindcss from "tailwindcss";

// Shared configuration
const sharedConfig: UserConfig = {
  plugins: [react()],
  css: {
    postcss: {
      plugins: [tailwindcss()],
    },
  },
  build: {
    sourcemap: true,
    emptyOutDir: false, // Prevent deleting the output directory
    outDir: "build",
    assetsDir: "assets", // All assets will be in build/assets
    rollupOptions: {
      output: {
        manualChunks: undefined, // Disable code splitting to prevent extra assets
        assetFileNames: "assets/[name].[ext]", // Place assets in the assets folder
      },
    },
  },
  define: {
    "process.env.NODE_ENV": JSON.stringify(
      process.env.NODE_ENV || "development"
    ),
  },
};

// Main configuration for the web application
const mainConfig: UserConfig = mergeConfig(sharedConfig, {
  build: {
    rollupOptions: {
      input: "./index.html",
      output: {
        entryFileNames: "[name].js",
        chunkFileNames: "[name].js",
        assetFileNames: "assets/[name].[ext]", // Ensure assets are in assets folder
      },
    },
  },
  optimizeDeps: {
    include: ["react", "react-dom"],
  },
});

// Content script configuration
const contentConfig: UserConfig = mergeConfig(sharedConfig, {
  build: {
    emptyOutDir: false, // Do not empty the output directory
    lib: {
      entry: "./src/content/content.tsx",
      name: "ContentScript",
      formats: ["iife"], // IIFE format for content script
      fileName: () => "content/content.js", // Place content script in content/ subfolder
    },
    rollupOptions: {
      output: {
        entryFileNames: "content/[name].js", // Ensure content script is in content/ folder
        chunkFileNames: "content/[name].js",
        assetFileNames: "assets/[name].[ext]", // Assets in the same assets folder
        manualChunks: undefined, // Disable code splitting
      },
    },
  },
});

// Background script configuration
const backgroundConfig: UserConfig = mergeConfig(sharedConfig, {
  build: {
    emptyOutDir: false, // Do not empty the output directory
    lib: {
      entry: "./src/background/background.ts",
      name: "BackgroundScript",
      formats: ["iife"], // IIFE format for background script
      fileName: () => "background/background.js", // Place background script in background/ subfolder
    },
    rollupOptions: {
      output: {
        entryFileNames: "background/[name].js", // Ensure background script is in background/ folder
        chunkFileNames: "background/[name].js",
        assetFileNames: "assets/[name].[ext]", // Assets in the same assets folder
        manualChunks: undefined, // Disable code splitting
      },
    },
  },
});

// Export the configuration
export default defineConfig(() => {
  // Determine the specific build target based on an environment variable
  if (process.env.BUILD_TARGET === "content") {
    return contentConfig;
  } else if (process.env.BUILD_TARGET === "background") {
    return backgroundConfig;
  } else {
    return mainConfig; // Fallback to the main configuration
  }
});
