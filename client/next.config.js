/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    // swcMinify: true,
    images: {
      domains: [
        'us-fbcloud.net',
        'lh3.googleusercontent.com'
      ]
    }
}

module.exports = nextConfig
