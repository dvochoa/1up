// @ts-check

/**
 * @type {import('next').NextConfig}
 */
const nextConfig = {
  // TODO: What does this do?
  output: "standalone",
  webpack: (config, { dev }) => {
    if (dev) {
      config.watchOptions = {
        poll: 1000,
        aggregateTimeout: 300,
      };
    }
    return config;
  },
};

export default nextConfig;
