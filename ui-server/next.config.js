const glob = require('glob')
const withCss = require('@zeit/next-css')
const withPurgeCss = require('next-purgecss')

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
      extensions: ['ts', 'tsx'],
    },
  ],
}

module.exports = withCss(
  withPurgeCss({
    purgeCss: purgeCssConfig,
    purgeCssEnabled: ({ dev }) => !dev,
  }),
)
