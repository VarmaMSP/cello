const withCss = require('@zeit/next-css')
const withPurgeCss = require('next-purgecss')

module.exports = withCss(
  withPurgeCss(
    {
      purgeCssPaths: [
        'pages/**/*.tsx',
        'components/**/*.tsx',
      ],
      purgeCss: {
        extractors: [
          {
            extractor: class {
              static extract(content) {
                return content.match(/[\w-/:]+(?<!:)/g) || []
              }
            },
            extensions: ['tsx']
          }
        ]
      }
    }
  )
)
