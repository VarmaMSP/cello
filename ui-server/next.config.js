const glob = require('glob')
const css = require('@zeit/next-css')
const purgeCss = require('next-purgecss')
const withPlugins = require('next-compose-plugins')

const purgeCssConfig = {
  paths: () => [
    ...glob.sync(`${__dirname}/pages/**/*.ts?(x)`, { nodir: true }),
    ...glob.sync(`${__dirname}/components/**/*.ts?(x)`, { nodir: true }),
  ],
  extractors: [
    {
      extractor: class {
        static extract(content) {
          return content.match(/[\w-/:]+(?<!:)/g) || []
        }
      },
      extensions: ['tsx'],
    },
  ],
}

const nextConfig = {
  distDir: 'next',
  generateEtags: false,
  poweredByHeader: false,
}

module.exports = withPlugins(
  [
    [
      css,
      purgeCss({
        purgeCss: purgeCssConfig,
        purgeCssEnabled: ({ dev }) => !dev,
      }),
    ],
  ],
  nextConfig,
)
