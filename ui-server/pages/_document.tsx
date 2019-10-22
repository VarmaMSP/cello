import Document, { Head, Main, NextScript } from 'next/document'
import React from 'react'
import { GA_TRACKING_ID } from 'utils/gtag'

export default class extends Document {
  render() {
    return (
      <html>
        <Head>
          <link rel="icon" href="favicon.ico" type="image/x-icon" />
          <link rel="shortcut icon" href="favicon.ico" type="image/x-icon" />
          <meta
            name="viewport"
            content="width=device-width,minimum-scale=1,initial-scale=0,user-scalable=no"
          />
          <script
            async
            src={`https://www.googletagmanager.com/gtag/js?id=${GA_TRACKING_ID}`}
          />
          <script
            dangerouslySetInnerHTML={{
              __html: `window.dataLayer = window.dataLayer || []; function gtag(){dataLayer.push(arguments);} gtag('js', new Date()); gtag('config', '${GA_TRACKING_ID}');`,
            }}
          />
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </html>
    )
  }
}
