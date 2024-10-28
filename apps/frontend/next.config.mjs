/** @type {import('next').NextConfig} */
const nextConfig = {
  output: "standalone",
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "cdn.litres.ru",
        port: "",
        pathname: "/**",
      },
    ],
  },
};

export default nextConfig;
